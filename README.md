# trades


Install Protobuf dependencies according to https://grpc.io/docs/languages/go/quickstart.

```
brew install protobuf
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

Generate the Go code according to the Protobuf messages.

```
protoc --go_out=typ/orders --go-grpc_out=typ/orders pbf/orders/orders.proto
protoc --go_out=typ/trades --go-grpc_out=typ/trades pbf/trades/trades.proto
```
