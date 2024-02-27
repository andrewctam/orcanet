# Market Server

gRPC reference: https://grpc.io/docs/languages/go/quickstart

To recompile protobuf:

```Shell
protoc --go_out=. --go_opt=paths=source_relative \
  --go-grpc_out=. --go-grpc_opt=paths=source_relative \
  market/market.proto
```

## Running

To run the market server:

```Shell
go run server/main.go
```

To run the test client:

```Shell
go run test_client/main.go
```

## API

gRPC endpoints are in `market/market.proto`
