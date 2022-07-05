# trades


Install Protobuf dependencies according to https://grpc.io/docs/languages/go/quickstart.

```
brew install protobuf
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

Generate the Go code according to the Protobuf message.

```
protoc --go_out=typ/trades --go-grpc_out=typ/trades pbf/trades/trades.proto
```
