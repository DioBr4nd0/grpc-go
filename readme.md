# gRPC Go Example

This repository contains a simple gRPC example implementation in Go with a client-server architecture.

## Prerequisites

- Go 1.15 or higher
- Protocol Buffers compiler (protoc)
- Go plugins for Protocol Buffers

### Installing Protocol Buffers

Follow the official guide to install Protocol Buffers:
[Protocol Buffers Installation Guide](https://grpc.io/docs/protoc-installation/)

## Generate Proto Files

Run the following command to generate the protocol buffer files:

```bash
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    helloworld.proto
```

## Running the Application

1. Start the server:
```bash
go run server/main.go
```

2. In a new terminal, run the client:
```bash
go run client/main.go
```

### Expected Output

You should see something like:
```
2025/03/23 21:22:03 response received: Hello World!
```

## Project Structure

- `helloworld.proto`: Protocol Buffers service definition
- `server/`: Server implementation
- `client/`: Client implementation