package helloworld

//go:generate protoc --go_out . --gogo-nrpc_out plugins=prometheus:. helloworld.proto
