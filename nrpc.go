package nrpc

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"runtime/debug"
	"strings"
	"sync"
	"time"

	"github.com/golang/protobuf/proto"
	nats "github.com/nats-io/go-nats"
)

// ContextKey type for storing values into context.Context
type ContextKey int

// ErrStreamInvalidMsgCount is when a stream reply gets a wrong number of messages
var ErrStreamInvalidMsgCount = errors.New("Stream reply received an incorrect number of messages")

//go:generate protoc -I. -I../../.. --gogo_out=Mgoogle/protobuf/descriptor.proto=github.com/gogo/protobuf/protoc-gen-gogo/descriptor:../../.. nrpc/nrpc.proto

type NatsConn interface {
	Publish(subj string, data []byte) error
	PublishRequest(subj, reply string, data []byte) error
	Request(subj string, data []byte, timeout time.Duration) (*nats.Msg, error)

	ChanSubscribe(subj string, ch chan *nats.Msg) (*nats.Subscription, error)
	Subscribe(subj string, handler nats.MsgHandler) (*nats.Subscription, error)
	SubscribeSync(subj string) (*nats.Subscription, error)
}

// ReplyInboxMaker returns a new inbox subject for a given nats connection.
type ReplyInboxMaker func(NatsConn) string

// GetReplyInbox is used by StreamCall to get a inbox subject
// It can be changed by a client lib that needs custom inbox subjects
var GetReplyInbox ReplyInboxMaker = func(NatsConn) string {
	return nats.NewInbox()
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s error: %s", Error_Type_name[int32(e.Type)], e.Message)
}

func Unmarshal(encoding string, data []byte, msg proto.Message) error {
	switch encoding {
	case "protobuf":
		return proto.Unmarshal(data, msg)
	case "json":
		return json.Unmarshal(data, msg)
	default:
		return errors.New("Invalid encoding: " + encoding)
	}
}

func UnmarshalResponse(encoding string, data []byte, msg proto.Message) error {
	switch encoding {
	case "protobuf":
		if len(data) > 0 && data[0] == 0 {
			var repErr Error
			if err := proto.Unmarshal(data[1:], &repErr); err != nil {
				return err
			}
			return &repErr
		}
		return proto.Unmarshal(data, msg)
	case "json":
		if len(data) > 13 && bytes.Equal(data[:13], []byte("{\"__error__\":")) {
			var rep map[string]*Error
			if err := json.Unmarshal(data, &rep); err != nil {
				return err
			}
			return rep["__error__"]
		}
		return json.Unmarshal(data, msg)
	default:
		return errors.New("Invalid encoding: " + encoding)
	}
}

func Marshal(encoding string, msg proto.Message) ([]byte, error) {
	switch encoding {
	case "protobuf":
		return proto.Marshal(msg)
	case "json":
		return json.Marshal(msg)
	default:
		return nil, errors.New("Invalid encoding: " + encoding)
	}
}

func MarshalErrorResponse(encoding string, repErr *Error) ([]byte, error) {
	switch encoding {
	case "protobuf":
		var (
			buf  = []byte{0}
			pBuf = proto.NewBuffer(buf)
		)
		if err := pBuf.Marshal(repErr); err != nil {
			return nil, err
		}
		return pBuf.Bytes(), nil
	case "json":
		return json.Marshal(map[string]*Error{
			"__error__": repErr,
		})
	default:
		return nil, errors.New("Invalid encoding: " + encoding)
	}
}

func ParseSubject(
	packageSubject string, packageParamsCount int,
	serviceSubject string, serviceParamsCount int,
	subject string,
) (packageParams []string, serviceParams []string,
	name string, tail []string, err error,
) {
	packageSubjectDepth := 0
	if packageSubject != "" {
		packageSubjectDepth = strings.Count(packageSubject, ".") + 1
	}
	serviceSubjectDepth := strings.Count(serviceSubject, ".") + 1
	subjectMinSize := packageSubjectDepth + packageParamsCount + serviceSubjectDepth + serviceParamsCount + 1

	tokens := strings.Split(subject, ".")
	if len(tokens) < subjectMinSize {
		err = fmt.Errorf(
			"Invalid subject len. Expects number of parts >= %d, got %d",
			subjectMinSize, len(tokens))
		return
	}
	if packageSubject != "" {
		for i, packageSubjectPart := range strings.Split(packageSubject, ".") {
			if tokens[i] != packageSubjectPart {
				err = fmt.Errorf(
					"Invalid subject prefix. Expected '%s', got '%s'",
					packageSubjectPart, tokens[i])
				return
			}
		}
		tokens = tokens[packageSubjectDepth:]
	}

	packageParams = tokens[0:packageParamsCount]
	tokens = tokens[packageParamsCount:]

	for i, serviceSubjectPart := range strings.Split(serviceSubject, ".") {
		if tokens[i] != serviceSubjectPart {
			err = fmt.Errorf(
				"Invalid subject. Service should be '%s', got '%s'",
				serviceSubjectPart, tokens[i])
			return
		}
	}
	tokens = tokens[serviceSubjectDepth:]

	serviceParams = tokens[0:serviceParamsCount]
	tokens = tokens[serviceParamsCount:]

	name = tokens[0]
	tokens = tokens[1:]

	tail = tokens
	return
}

func ParseSubjectTail(
	methodParamsCount int,
	tail []string,
) (
	methodParams []string, encoding string, err error,
) {
	if len(tail) < methodParamsCount || len(tail) > methodParamsCount+1 {
		err = fmt.Errorf(
			"Invalid subject tail length. Expects %d or %d parts, got %d",
			methodParamsCount, methodParamsCount+1, len(tail),
		)
		return
	}
	methodParams = tail[:methodParamsCount]
	tail = tail[methodParamsCount:]
	switch len(tail) {
	case 0:
		encoding = "protobuf"
	case 1:
		encoding = tail[0]
	default:
		panic("Got extra tokens, which should be impossible at this point")
	}
	return
}

func Call(req proto.Message, rep proto.Message, nc NatsConn, subject string, encoding string, timeout time.Duration) error {
	// encode request
	rawRequest, err := Marshal(encoding, req)
	if err != nil {
		log.Printf("nrpc: inner request marshal failed: %v", err)
		return err
	}

	if encoding != "protobuf" {
		subject += "." + encoding
	}

	// call
	if _, noreply := rep.(*NoReply); noreply {
		err := nc.Publish(subject, rawRequest)
		if err != nil {
			log.Printf("nrpc: nats publish failed: %v", err)
		}
		return err
	}
	msg, err := nc.Request(subject, rawRequest, timeout)
	if err != nil {
		log.Printf("nrpc: nats request failed: %v", err)
		return err
	}

	data := msg.Data

	if err := UnmarshalResponse(encoding, data, rep); err != nil {
		if _, isError := err.(*Error); !isError {
			log.Printf("nrpc: response unmarshal failed: %v", err)
		}
		return err
	}

	return nil
}

const (
	// RequestContextKey is the key for string the request into the context
	RequestContextKey ContextKey = iota
)

// NewRequest creates a Request instance
func NewRequest(ctx context.Context, conn NatsConn, subject string, replySubject string) *Request {
	return &Request{
		Context:      ctx,
		Conn:         conn,
		Subject:      subject,
		ReplySubject: replySubject,
		CreatedAt:    time.Now(),
	}
}

// GetRequest returns the Request associated with a context, or nil if absent
func GetRequest(ctx context.Context) *Request {
	request, _ := ctx.Value(RequestContextKey).(*Request)
	return request
}

// Request is a server-side incoming request
type Request struct {
	Context context.Context
	Conn    NatsConn

	KeepStreamAlive *KeepStreamAlive
	StreamContext   context.Context
	StreamCancel    func()
	StreamMsgCount  uint32

	Subject     string
	MethodName  string
	SubjectTail []string

	CreatedAt time.Time
	StartedAt time.Time

	Encoding     string
	NoReply      bool
	ReplySubject string

	PackageParams map[string]string
	ServiceParams map[string]string

	AfterReply func(r *Request, success bool, replySuccess bool)

	Handler func(context.Context) (proto.Message, error)
}

// Elapsed duration since request was started
func (r Request) Elapsed() time.Duration {
	return time.Since(r.CreatedAt)
}

// Run the handler and capture any error. Returns the response or the error
// that should be returned to the caller
func (r Request) Run() (msg proto.Message, replyError *Error) {
	r.StartedAt = time.Now()
	ctx := r.Context
	if r.StreamedReply() {
		ctx = r.StreamContext
	}
	ctx = context.WithValue(ctx, RequestContextKey, &r)
	msg, replyError = CaptureErrors(
		func() (proto.Message, error) {
			return r.Handler(ctx)
		})
	return
}

// RunAndReply calls Run() and send the reply back to the caller
func (r *Request) RunAndReply() {
	var failed, replyFailed bool
	resp, replyError := r.Run()
	if replyError != nil {
		failed = true
		log.Printf("%s handler failed: %s", r.MethodName, replyError)
	}
	if !r.NoReply {
		if err := r.SendReply(resp, replyError); err != nil {
			replyFailed = true
			log.Printf("%s failed to publish the response: %s", r.MethodName, err)
		}
	}
	if r.AfterReply != nil {
		r.AfterReply(r, !failed, !replyFailed)
	}
}

// PackageParam returns a package parameter value, or "" if absent
func (r *Request) PackageParam(key string) string {
	if r == nil || r.PackageParams == nil {
		return ""
	}
	return r.PackageParams[key]
}

// ServiceParam returns a package parameter value, or "" if absent
func (r *Request) ServiceParam(key string) string {
	if r == nil || r.ServiceParams == nil {
		return ""
	}
	return r.ServiceParams[key]
}

// SetPackageParam sets a package param value
func (r *Request) SetPackageParam(key, value string) {
	if r.PackageParams == nil {
		r.PackageParams = make(map[string]string)
	}
	r.PackageParams[key] = value
}

// SetServiceParam sets a service param value
func (r *Request) SetServiceParam(key, value string) {
	if r.ServiceParams == nil {
		r.ServiceParams = make(map[string]string)
	}
	r.ServiceParams[key] = value
}

// SetupStreamedReply initializes the reply stream
func (r *Request) SetupStreamedReply() {
	r.StreamContext, r.StreamCancel = context.WithCancel(r.Context)
	r.KeepStreamAlive = NewKeepStreamAlive(
		r.Conn, r.ReplySubject, r.Encoding, r.StreamCancel)
}

// StreamedReply returns true if the request reply is streamed
func (r Request) StreamedReply() bool {
	return r.KeepStreamAlive != nil
}

// SendStreamReply send a reply a part of a stream
func (r *Request) SendStreamReply(msg proto.Message) {
	log.Printf("nrpc: SendStreamReply")
	if err := r.sendReply(msg, nil); err != nil {
		log.Printf("nrpc: error publishing response")
		r.StreamCancel()
		return
	}
	r.StreamMsgCount++
}

// SendReply sends a reply to the caller
func (r *Request) SendReply(resp proto.Message, withError *Error) error {
	if r.StreamedReply() {
		r.KeepStreamAlive.Stop()
		if withError == nil {
			return r.sendReply(
				nil, &Error{Type: Error_EOS, MsgCount: r.StreamMsgCount},
			)
		}
	}
	return r.sendReply(resp, withError)
}

// sendReply sends a reply to the caller
func (r *Request) sendReply(resp proto.Message, withError *Error) error {
	return Publish(resp, withError, r.Conn, r.ReplySubject, r.Encoding)
}

var ErrEOS = errors.New("End of stream")
var ErrCanceled = errors.New("Call canceled")

func NewStreamCallSubscription(
	ctx context.Context, nc NatsConn, encoding string, subject string,
	timeout time.Duration,
) (*StreamCallSubscription, error) {
	sub := StreamCallSubscription{
		ctx:      ctx,
		nc:       nc,
		encoding: encoding,
		subject:  subject,
		timeout:  timeout,
		timeoutT: time.NewTimer(timeout),
		subCh:    make(chan *nats.Msg, 256),
		nextCh:   make(chan *nats.Msg),
		quit:     make(chan struct{}),
		errCh:    make(chan error),
	}
	ssub, err := nc.ChanSubscribe(subject, sub.subCh)
	if err != nil {
		return nil, err
	}
	sub.sub = ssub
	go sub.loop()
	return &sub, nil
}

type StreamCallSubscription struct {
	ctx      context.Context
	nc       NatsConn
	encoding string
	subject  string
	timeout  time.Duration
	timeoutT *time.Timer
	sub      *nats.Subscription
	subCh    chan *nats.Msg
	nextCh   chan *nats.Msg
	quit     chan struct{}
	errCh    chan error
	msgCount uint32
}

func (sub *StreamCallSubscription) stop() {
	close(sub.quit)
}

func (sub *StreamCallSubscription) loop() {
	hbSubject := sub.subject + ".heartbeat"
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	defer sub.sub.Unsubscribe()
	for {
		select {
		case msg := <-sub.subCh:
			sub.timeoutT.Reset(sub.timeout)

			if len(msg.Data) == 1 && msg.Data[0] == 0 {
				break
			}

			sub.nextCh <- msg
		case <-sub.timeoutT.C:
			sub.errCh <- nats.ErrTimeout
			return
		case <-sub.ctx.Done():
			// send a 'lastbeat' and quit
			b, err := Marshal(sub.encoding, &HeartBeat{Lastbeat: true})
			if err != nil {
				err = fmt.Errorf("Error marshaling heartbeat: %s", err)
				sub.errCh <- err
				return
			}
			if err := sub.nc.Publish(hbSubject, b); err != nil {
				err = fmt.Errorf("Error sending heartbeat: %s", err)
				sub.errCh <- err
				return
			}
			sub.errCh <- ErrCanceled
			return
		case <-ticker.C:
			msg, err := Marshal(sub.encoding, &HeartBeat{})
			if err != nil {
				err = fmt.Errorf("Error marshaling heartbeat: %s", err)
				sub.errCh <- err
				return
			}
			if err := sub.nc.Publish(hbSubject, msg); err != nil {
				err = fmt.Errorf("Error sending heartbeat: %s", err)
				sub.errCh <- err
				return
			}
		case <-sub.quit:
			return
		}
	}
}

func (sub *StreamCallSubscription) Next(rep proto.Message) error {
	if sub.sub == nil {
		return nats.ErrBadSubscription
	}
	select {
	case err := <-sub.errCh:
		sub.sub = nil
		return err
	case msg := <-sub.nextCh:
		if err := UnmarshalResponse(sub.encoding, msg.Data, rep); err != nil {
			sub.stop()
			sub.sub = nil
			if nrpcErr, ok := err.(*Error); ok {
				if nrpcErr.GetMsgCount() != sub.msgCount {
					log.Printf(
						"nrpc: received invalid number of messages. Expected %d, got %d",
						nrpcErr.GetMsgCount(), sub.msgCount)
				}
				if nrpcErr.GetType() == Error_EOS {
					if nrpcErr.GetMsgCount() != sub.msgCount {
						return ErrStreamInvalidMsgCount
					}
					return ErrEOS
				}
			} else {
				log.Printf("nrpc: response unmarshal failed: %v", err)
			}
			return err
		}
		sub.msgCount++
	}

	return nil
}

func StreamCall(ctx context.Context, nc NatsConn, subject string, req proto.Message, encoding string, timeout time.Duration) (*StreamCallSubscription, error) {
	rawRequest, err := Marshal(encoding, req)
	if err != nil {
		log.Printf("nrpc: inner request marshal failed: %v", err)
		return nil, err
	}

	if encoding != "protobuf" {
		subject += "." + encoding
	}

	reply := GetReplyInbox(nc)

	streamSub, err := NewStreamCallSubscription(ctx, nc, encoding, reply, timeout)
	if err != nil {
		return nil, err
	}

	if err := nc.PublishRequest(subject, reply, rawRequest); err != nil {
		streamSub.stop()
		return nil, err
	}
	return streamSub, nil
}

func Publish(resp proto.Message, withError *Error, nc NatsConn, subject string, encoding string) error {
	var rawResponse []byte
	var err error

	if withError != nil {
		rawResponse, err = MarshalErrorResponse(encoding, withError)
	} else {
		rawResponse, err = Marshal(encoding, resp)
	}

	if err != nil {
		log.Printf("nrpc: rpc response marshal failed: %v", err)
		return err
	}

	// send response
	if err := nc.Publish(subject, rawResponse); err != nil {
		log.Printf("nrpc: response publish failed: %v", err)
		return err
	}

	return nil
}

// CaptureErrors runs a handler and convert error and panics into proper Error
func CaptureErrors(fn func() (proto.Message, error)) (msg proto.Message, replyError *Error) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Caught panic: %s\n%s", r, debug.Stack())
			replyError = &Error{
				Type:    Error_SERVER,
				Message: fmt.Sprint(r),
			}
		}
	}()
	var err error
	msg, err = fn()
	if err != nil {
		var ok bool
		if replyError, ok = err.(*Error); !ok {
			replyError = &Error{
				Type:    Error_CLIENT,
				Message: err.Error(),
			}
		}
	}
	return
}

func NewKeepStreamAlive(nc NatsConn, subject string, encoding string, onError func()) *KeepStreamAlive {
	k := KeepStreamAlive{
		nc:       nc,
		subject:  subject,
		encoding: encoding,
		c:        make(chan struct{}),
		onError:  onError,
	}
	go k.loop()
	return &k
}

type KeepStreamAlive struct {
	nc       NatsConn
	subject  string
	encoding string
	c        chan struct{}
	onError  func()
}

func (k *KeepStreamAlive) Stop() {
	close(k.c)
}

func (k *KeepStreamAlive) loop() {
	hbChan := make(chan *nats.Msg, 256)
	hbSub, err := k.nc.ChanSubscribe(k.subject+".heartbeat", hbChan)
	if err != nil {
		log.Printf("nrpc: could not subscribe to heartbeat: %s", err)
		k.onError()
	}
	defer func() {
		if err := hbSub.Unsubscribe(); err != nil {
			log.Printf("nrpc: error unsubscribing from heartbeat: %s", err)
		}
	}()
	hbDelay := 0
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case msg := <-hbChan:
			var hb HeartBeat
			if err := Unmarshal(k.encoding, msg.Data, &hb); err != nil {
				log.Printf("nrpc: error unmarshaling heartbeat: %s", err)
				ticker.Stop()
				k.onError()
				return
			}
			if hb.Lastbeat {
				log.Printf("nrpc: client canceled the streamed reply. (%s)", k.subject)
				ticker.Stop()
				k.onError()
				return
			}
			hbDelay = 0
		case <-ticker.C:
			hbDelay++
			if hbDelay >= 5 {
				log.Printf("nrpc: No heartbeat received in 5 seconds. Canceling.")
				ticker.Stop()
				k.onError()
				return
			}
			if err := k.nc.Publish(k.subject, []byte{0}); err != nil {
				log.Printf("nrpc: error publishing response: %s", err)
				ticker.Stop()
				k.onError()
				return
			}
		case <-k.c:
			ticker.Stop()
			return
		}
	}
}

// ErrTooManyPendingRequests is returned if the pending queue is full
var ErrTooManyPendingRequests = errors.New("Too many pending requests")

// WorkerPool is a poof of workers
type WorkerPool struct {
	Context       context.Context
	contextCancel context.CancelFunc

	queue     chan *Request
	schedule  chan *Request
	waitGroup sync.WaitGroup
	m         sync.Mutex

	size               uint
	maxPending         uint
	maxPendingDuration time.Duration
}

// NewWorkerPool creates a pool of workers
func NewWorkerPool(
	ctx context.Context,
	size uint,
	maxPending uint,
	maxPendingDuration time.Duration,
) *WorkerPool {
	nCtx, cancel := context.WithCancel(ctx)
	pool := WorkerPool{
		Context:            nCtx,
		contextCancel:      cancel,
		queue:              make(chan *Request, maxPending),
		schedule:           make(chan *Request),
		maxPending:         maxPending,
		maxPendingDuration: maxPendingDuration,
	}
	pool.waitGroup.Add(1)
	go pool.scheduler()
	pool.SetSize(size)
	return &pool
}

func (pool *WorkerPool) getQueue() (queue chan *Request) {
	pool.m.Lock()
	queue = pool.queue
	pool.m.Unlock()
	return
}

func (pool *WorkerPool) scheduler() {
	defer pool.waitGroup.Done()

	for {
		queue := pool.getQueue()
		if queue == nil {
			return
		}
	queueLoop:
		for request := range queue {
			now := time.Now()
			deadline := request.CreatedAt.Add(pool.maxPendingDuration)

			log.Printf("scheduler: got a request. Deadline in=%s", deadline.Sub(now))
			if deadline.After(now) {
				log.Printf("scheduler: try scheduling")
				select {
				case pool.schedule <- request:
					log.Printf("scheduler: scheduled the request")
					continue queueLoop
				case <-time.After(deadline.Sub(now)):
					// Too late
					log.Printf("scheduler: could not schedule the request")
				}
			} else {
				log.Printf("scheduler: already too late, skip scheduling")
			}

			log.Printf("Sending SERVERTOOBUSY to the caller")
			request.SendReply(nil, &Error{
				Type:    Error_SERVERTOOBUSY,
				Message: "No worker available",
			})
		}
	}
}

func (pool *WorkerPool) worker() {
	defer pool.waitGroup.Done()
	for request := range pool.schedule {
		if request == nil {
			log.Printf("worker: got nil, exiting")
			return
		}
		log.Printf("worker: got a request, running it")
		request.RunAndReply()
	}
}

// SetMaxPending changes the queue size
func (pool *WorkerPool) SetMaxPending(value uint) {
	if pool.maxPending == value {
		return
	}
	pool.m.Lock()
	defer pool.m.Unlock()

	oldQueue := pool.queue
	pool.queue = make(chan *Request, value)
	pool.maxPending = value

	for request := range oldQueue {
		pool.queue <- request
	}
	close(oldQueue)
}

func (pool *WorkerPool) SetMaxPendingDuration(value time.Duration) {
	pool.maxPendingDuration = value
}

// SetSize changes the number of workers
func (pool *WorkerPool) SetSize(size uint) {
	if size == pool.size {
		return
	}
	for size < pool.size {
		pool.schedule <- nil
		pool.size--
	}
	for size > pool.size {
		pool.waitGroup.Add(1)
		go pool.worker()
		pool.size++
	}
}

// QueueRequest adds a request to the queue
// returns ErrTooManyPendingRequests if the queue is full
func (pool *WorkerPool) QueueRequest(request *Request) error {
	select {
	case pool.getQueue() <- request:
		return nil
	default:
		return ErrTooManyPendingRequests
	}
}

// Close stops all the workers and wait for their completion
// If the workers do not stop before the timeout, their context is canceled
// Will never return if a request ignore the context
func (pool *WorkerPool) Close(timeout time.Duration) {
	// Stops all the workers so nothing more gets scheduled
	pool.SetSize(0)

	pool.m.Lock()
	oldQueue := pool.queue
	pool.queue = nil
	pool.m.Unlock()

	close(oldQueue)
	for range oldQueue {
	}

	// Now wait for the workers to stop and cancel the context if they don't
	close(pool.schedule)
	timer := time.AfterFunc(timeout, pool.contextCancel)
	pool.waitGroup.Wait()
	timer.Stop()
}
