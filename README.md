# Go programming language - gRPC

## VScode extension

`` vscode-proto3 ``

## Protocol Buffers Documentation

[Language Guide (proto 3)](https://protobuf.dev/programming-guides/proto3/)

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