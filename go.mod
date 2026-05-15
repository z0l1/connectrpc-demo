module github.com/z0l1/connectrpc-demo

go 1.26.2

tool (
	connectrpc.com/connect/cmd/protoc-gen-connect-go
	google.golang.org/protobuf/cmd/protoc-gen-go
)

require (
	buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go v1.36.11-20260415201107-50325440f8f2.1
	connectrpc.com/connect v1.19.2
	google.golang.org/protobuf v1.36.11
)
