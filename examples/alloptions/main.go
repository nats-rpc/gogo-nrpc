package main

//go:generate protoc -I. -I../.. -I../../protobuf/protobuf -I../../protobuf/gogoproto --gogo_out=Mgoogle/protobuf/descriptor.proto=github.com/gogo/protobuf/protoc-gen-gogo/descriptor:. --gogo-nrpc_out . alloptions.proto

func main() {

}
