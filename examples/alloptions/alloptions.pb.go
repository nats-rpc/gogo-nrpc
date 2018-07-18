// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: alloptions.proto

/*
Package main is a generated protocol buffer package.

It is generated from these files:
	alloptions.proto

It has these top-level messages:
	StringArg
	SimpleStringReply
*/
package main

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/nats-rpc/gogo-nrpc"
import _ "github.com/gogo/protobuf/gogoproto"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type StringArg struct {
	Arg1 string `protobuf:"bytes,1,opt,name=arg1,proto3" json:"arg1,omitempty"`
}

func (m *StringArg) Reset()                    { *m = StringArg{} }
func (*StringArg) ProtoMessage()               {}
func (*StringArg) Descriptor() ([]byte, []int) { return fileDescriptorAlloptions, []int{0} }

func (m *StringArg) GetArg1() string {
	if m != nil {
		return m.Arg1
	}
	return ""
}

type SimpleStringReply struct {
	Reply string `protobuf:"bytes,1,opt,name=reply,proto3" json:"reply,omitempty"`
}

func (m *SimpleStringReply) Reset()                    { *m = SimpleStringReply{} }
func (*SimpleStringReply) ProtoMessage()               {}
func (*SimpleStringReply) Descriptor() ([]byte, []int) { return fileDescriptorAlloptions, []int{1} }

func (m *SimpleStringReply) GetReply() string {
	if m != nil {
		return m.Reply
	}
	return ""
}

func init() {
	proto.RegisterType((*StringArg)(nil), "main.StringArg")
	proto.RegisterType((*SimpleStringReply)(nil), "main.SimpleStringReply")
}
func (this *StringArg) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*StringArg)
	if !ok {
		that2, ok := that.(StringArg)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *StringArg")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *StringArg but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *StringArg but is not nil && this == nil")
	}
	if this.Arg1 != that1.Arg1 {
		return fmt.Errorf("Arg1 this(%v) Not Equal that(%v)", this.Arg1, that1.Arg1)
	}
	return nil
}
func (this *StringArg) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*StringArg)
	if !ok {
		that2, ok := that.(StringArg)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Arg1 != that1.Arg1 {
		return false
	}
	return true
}
func (this *SimpleStringReply) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*SimpleStringReply)
	if !ok {
		that2, ok := that.(SimpleStringReply)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *SimpleStringReply")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *SimpleStringReply but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *SimpleStringReply but is not nil && this == nil")
	}
	if this.Reply != that1.Reply {
		return fmt.Errorf("Reply this(%v) Not Equal that(%v)", this.Reply, that1.Reply)
	}
	return nil
}
func (this *SimpleStringReply) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*SimpleStringReply)
	if !ok {
		that2, ok := that.(SimpleStringReply)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Reply != that1.Reply {
		return false
	}
	return true
}
func (this *StringArg) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&main.StringArg{")
	s = append(s, "Arg1: "+fmt.Sprintf("%#v", this.Arg1)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *SimpleStringReply) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&main.SimpleStringReply{")
	s = append(s, "Reply: "+fmt.Sprintf("%#v", this.Reply)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringAlloptions(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *StringArg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StringArg) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Arg1) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintAlloptions(dAtA, i, uint64(len(m.Arg1)))
		i += copy(dAtA[i:], m.Arg1)
	}
	return i, nil
}

func (m *SimpleStringReply) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SimpleStringReply) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Reply) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintAlloptions(dAtA, i, uint64(len(m.Reply)))
		i += copy(dAtA[i:], m.Reply)
	}
	return i, nil
}

func encodeFixed64Alloptions(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Alloptions(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintAlloptions(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedStringArg(r randyAlloptions, easy bool) *StringArg {
	this := &StringArg{}
	this.Arg1 = string(randStringAlloptions(r))
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedSimpleStringReply(r randyAlloptions, easy bool) *SimpleStringReply {
	this := &SimpleStringReply{}
	this.Reply = string(randStringAlloptions(r))
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyAlloptions interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneAlloptions(r randyAlloptions) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringAlloptions(r randyAlloptions) string {
	v1 := r.Intn(100)
	tmps := make([]rune, v1)
	for i := 0; i < v1; i++ {
		tmps[i] = randUTF8RuneAlloptions(r)
	}
	return string(tmps)
}
func randUnrecognizedAlloptions(r randyAlloptions, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldAlloptions(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldAlloptions(dAtA []byte, r randyAlloptions, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateAlloptions(dAtA, uint64(key))
		v2 := r.Int63()
		if r.Intn(2) == 0 {
			v2 *= -1
		}
		dAtA = encodeVarintPopulateAlloptions(dAtA, uint64(v2))
	case 1:
		dAtA = encodeVarintPopulateAlloptions(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateAlloptions(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateAlloptions(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateAlloptions(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateAlloptions(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *StringArg) Size() (n int) {
	var l int
	_ = l
	l = len(m.Arg1)
	if l > 0 {
		n += 1 + l + sovAlloptions(uint64(l))
	}
	return n
}

func (m *SimpleStringReply) Size() (n int) {
	var l int
	_ = l
	l = len(m.Reply)
	if l > 0 {
		n += 1 + l + sovAlloptions(uint64(l))
	}
	return n
}

func sovAlloptions(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozAlloptions(x uint64) (n int) {
	return sovAlloptions(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *StringArg) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&StringArg{`,
		`Arg1:` + fmt.Sprintf("%v", this.Arg1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *SimpleStringReply) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&SimpleStringReply{`,
		`Reply:` + fmt.Sprintf("%v", this.Reply) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringAlloptions(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *StringArg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAlloptions
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StringArg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StringArg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Arg1", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAlloptions
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAlloptions
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Arg1 = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAlloptions(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAlloptions
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SimpleStringReply) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAlloptions
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SimpleStringReply: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SimpleStringReply: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reply", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAlloptions
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAlloptions
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Reply = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAlloptions(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAlloptions
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipAlloptions(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAlloptions
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowAlloptions
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowAlloptions
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthAlloptions
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowAlloptions
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipAlloptions(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthAlloptions = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAlloptions   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("alloptions.proto", fileDescriptorAlloptions) }

var fileDescriptorAlloptions = []byte{
	// 499 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0x31, 0x6f, 0xd3, 0x40,
	0x14, 0xce, 0x41, 0x80, 0xe6, 0xa0, 0x4d, 0xb8, 0x22, 0x70, 0x32, 0x1c, 0x55, 0xa6, 0x22, 0x11,
	0x47, 0x0d, 0x13, 0x63, 0xe9, 0x80, 0x84, 0x70, 0x85, 0x62, 0x89, 0x0e, 0x0c, 0x91, 0xe3, 0x1c,
	0xee, 0x21, 0xdb, 0xe7, 0xde, 0x3d, 0x57, 0x62, 0x43, 0x1d, 0x3b, 0x75, 0xec, 0x4f, 0x60, 0x8e,
	0x14, 0x89, 0x81, 0x01, 0x31, 0x21, 0x26, 0x46, 0xc6, 0xc6, 0x0c, 0xac, 0x48, 0x5d, 0x18, 0x91,
	0xef, 0x4c, 0xa8, 0x23, 0x55, 0x04, 0xba, 0xd8, 0xef, 0xbd, 0xfb, 0xde, 0xf7, 0xbd, 0xfb, 0x9e,
	0x8d, 0x1b, 0x5e, 0x18, 0x8a, 0x04, 0xb8, 0x88, 0x95, 0x9d, 0x48, 0x01, 0x82, 0x54, 0x23, 0x8f,
	0xc7, 0xad, 0x4e, 0xc0, 0x61, 0x37, 0x1d, 0xda, 0xbe, 0x88, 0xba, 0xb1, 0x07, 0xaa, 0x23, 0x13,
	0xbf, 0x1b, 0x88, 0x40, 0x74, 0xe2, 0x3c, 0x9a, 0x3d, 0x4c, 0x53, 0x09, 0x9e, 0xa3, 0xba, 0xba,
	0x3c, 0x4c, 0x5f, 0xea, 0x4c, 0x27, 0x3a, 0x32, 0xf0, 0xf6, 0x5d, 0x5c, 0x73, 0x41, 0xf2, 0x38,
	0xd8, 0x94, 0x01, 0x21, 0xb8, 0xea, 0xc9, 0x60, 0xc3, 0x42, 0x6b, 0x68, 0xbd, 0xd6, 0xd7, 0x71,
	0xfb, 0x1e, 0xbe, 0xe9, 0xf2, 0x28, 0x09, 0x99, 0x81, 0xf5, 0x59, 0x12, 0xbe, 0x26, 0xb7, 0xf0,
	0x15, 0x99, 0x07, 0x05, 0xd2, 0x24, 0xbd, 0xef, 0x97, 0x70, 0xc3, 0xdd, 0xf7, 0xb7, 0x52, 0x05,
	0x22, 0x72, 0xd3, 0xe1, 0x2b, 0xe6, 0x03, 0xd9, 0xc6, 0xcb, 0x0e, 0x18, 0x06, 0xd3, 0x5b, 0xb7,
	0xf3, 0x6b, 0xd9, 0x33, 0xd5, 0xd6, 0x9d, 0xa2, 0x30, 0xaf, 0xd2, 0x5e, 0x3d, 0x18, 0x37, 0xeb,
	0x11, 0x0c, 0x94, 0x3e, 0x19, 0x68, 0x11, 0x72, 0x1f, 0x5f, 0x77, 0xe0, 0xb9, 0xe0, 0xa3, 0x73,
	0xd8, 0xb0, 0xad, 0xcd, 0xc8, 0x11, 0xed, 0x0a, 0x79, 0x98, 0xa3, 0xb7, 0x45, 0x9f, 0xed, 0xa5,
	0x4c, 0x01, 0xa9, 0x9b, 0xc3, 0x59, 0xe1, 0x7c, 0xed, 0x0a, 0xd9, 0xc4, 0x75, 0x07, 0x5c, 0x90,
	0xcc, 0x8b, 0xd8, 0xe8, 0x5f, 0x47, 0xaf, 0x1e, 0x8d, 0x9b, 0x88, 0x6c, 0xe1, 0xdb, 0xbf, 0x67,
	0xdd, 0x2b, 0x33, 0x9d, 0x99, 0xf2, 0x2f, 0x24, 0x2d, 0xf2, 0xf1, 0xd4, 0x5a, 0xf1, 0xb5, 0xa7,
	0x03, 0x65, 0x4c, 0xed, 0xbd, 0x37, 0x4e, 0x17, 0x1e, 0x3f, 0xf3, 0xa4, 0x17, 0x29, 0xf2, 0x04,
	0xaf, 0x3a, 0xb0, 0xc3, 0x61, 0xb7, 0x5c, 0x5e, 0x48, 0x6a, 0xe5, 0x70, 0xdc, 0xbc, 0x1c, 0x25,
	0x1b, 0xe6, 0xd5, 0x23, 0x2f, 0xf0, 0xda, 0xdc, 0xe5, 0xff, 0x93, 0x98, 0x94, 0x89, 0xb5, 0x2d,
	0xeb, 0xb8, 0x66, 0x96, 0x32, 0xef, 0xc4, 0xf2, 0x9f, 0xf5, 0x98, 0x1d, 0x3c, 0xc6, 0xe4, 0xcc,
	0xfa, 0x76, 0x0a, 0xe1, 0xc5, 0xb7, 0x78, 0xad, 0x50, 0x6f, 0xdd, 0xf8, 0x7c, 0x6a, 0x2d, 0xf9,
	0x21, 0x67, 0x31, 0xf0, 0x51, 0xcf, 0xc1, 0x8d, 0x59, 0xb3, 0xcb, 0xe4, 0x3e, 0xf7, 0xd9, 0x05,
	0xbe, 0x94, 0x47, 0x4f, 0x0f, 0x26, 0x56, 0x55, 0x0a, 0x01, 0x87, 0x13, 0x6b, 0x89, 0xc7, 0x0a,
	0xbc, 0xd8, 0x67, 0x47, 0x13, 0x0b, 0x1d, 0x4f, 0x2c, 0xf4, 0x75, 0x4a, 0x2b, 0x27, 0x53, 0x8a,
	0x7e, 0x4c, 0x29, 0xfa, 0x39, 0xa5, 0xe8, 0x4d, 0x46, 0xd1, 0xdb, 0x8c, 0xa2, 0x77, 0x19, 0x45,
	0x1f, 0x32, 0x8a, 0x3e, 0x65, 0x14, 0x7d, 0xc9, 0x28, 0x3a, 0xc9, 0x28, 0x3a, 0xfe, 0x46, 0x2b,
	0xc3, 0xab, 0xfa, 0xc7, 0x7c, 0xf0, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x02, 0xac, 0xaf, 0x45, 0x10,
	0x04, 0x00, 0x00,
}
