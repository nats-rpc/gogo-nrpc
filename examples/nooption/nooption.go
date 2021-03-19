package nooption

//go:generate protoc -I. -I../.. -I../../protobuf/protobuf -I../../protobuf/gogoproto --gogo_out . --gogo-nrpc_out . nooption.proto
