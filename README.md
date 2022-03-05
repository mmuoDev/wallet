# Wallet Service

The wallet service basically implements the WalletServer interface from the wallet [proto file](https://github.com/mmuoDev/core-proto/blob/master/gen/wallet/wallet_grpc.pb.go#L21).

## Requirements
1. Postgres
2. [Set up grpc](https://grpc.io/docs/languages/go/quickstart/#prerequisites)

## Usage
The implemeation of the WalletServer is done in the `internal/server/server.go` file.

### Starting gPRC Server
To start the server, run 
```bash 
make run
```

### Testing
To run the tests, 
```bash
make test
```