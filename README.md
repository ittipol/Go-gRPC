# Go programming language - gRPC

## Packages

- protobuf [https://pkg.go.dev/google.golang.org/protobuf#section-directories](https://pkg.go.dev/google.golang.org/protobuf#section-directories)
    - Directories section > cmd > click protoc-gen-go [1]
- [1] protoc-gen-go [https://pkg.go.dev/google.golang.org/protobuf@v1.30.0/cmd/protoc-gen-go](https://pkg.go.dev/google.golang.org/protobuf@v1.30.0/cmd/protoc-gen-go)
- protoc-gen-go-grpc [https://pkg.go.dev/google.golang.org/grpc/cmd/protoc-gen-go-grpc](https://pkg.go.dev/google.golang.org/grpc/cmd/protoc-gen-go-grpc)

``` bash
# Install protoc-gen-go package
go get google.golang.org/protobuf/cmd/protoc-gen-go

# Install protoc-gen-go-grpc package
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

## Install gRPC tool

``` bash
go install google.golang.org/protobuf/cmd/protoc-gen-go

go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

# Check tool installed
cd cd $(go env GOPATH)/bin

find . -name "protoc-*"
```

## Add the workspace's bin subdirectory to your PATH
``` bash
$ export PATH=$PATH:$(go env GOPATH)/bin

# check
printenv
```

## Protocol Buffers

Protocol Buffers Release - [https://github.com/protocolbuffers/protobuf/releases](https://github.com/protocolbuffers/protobuf/releases)

## Install protoc command
[https://github.com/protocolbuffers/protobuf/releases](https://github.com/protocolbuffers/protobuf/releases)
``` bash
#mac os
brew install protobuf

# Check version
protoc --version
```

## Fix brew error

Solution 1. Update hosts file
- open /etc/hosts file
- Add a new line after the host's file and enter 185.199.108.133 raw.githubusercontent.com
- Save the modified host file

Solution 2. Add Public DNS IP addresses
[https://developers.google.com/speed/public-dns/docs/using](https://developers.google.com/speed/public-dns/docs/using)
- 8.8.8.8
- 8.8.4.4

Solution 3. Disable IPv6
- Go to System Preferences > Network  > TCP/IP > Change Configure IPv6 to Link-local only

## VScode extension

- vscode-proto3

## Protocol Buffers Documentation

[Language Guide (proto 3)](https://protobuf.dev/programming-guides/proto3/)