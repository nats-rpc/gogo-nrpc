package helloworld

//go:generate protoc -I. -I../../.. -I../../../../../.. --gogo_out=Mgoogle/protobuf/descriptor.proto=github.com/gogo/protobuf/protoc-gen-gogo/descriptor:. --gogo-nrpc_out . helloworld.proto
