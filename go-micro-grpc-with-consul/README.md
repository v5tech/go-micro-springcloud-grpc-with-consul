# go-micro-grpc-with-consul

### 1、安装`go-micro`

```bash
go get -u github.com/asim/go-micro/v3
```

### 2、安装`protoc`

```bash
go get -u github.com/google/protobuf/protoc
```

### 3、安装`protoc-gen-go`

```bash
go get -u github.com/google/protobuf/protoc-gen-go
```

### 4、安装`protoc-gen-micro`

```bash
go get -u github.com/asim/go-micro/cmd/protoc-gen-micro/v3
```

### 5、安装`consul registry`

```bash
go get -u github.com/asim/go-micro/plugins/registry/consul/v3
```

### 6、根据proto文件生成pb文件

```bash
protoc --proto_path=$GOPATH/src:. --micro_out=. --go_out=. hello.proto
```

# 参考文档

https://github.com/asim/go-micro

https://github.com/fullstorydev/grpcui

https://developers.google.com/protocol-buffers/docs/proto3

https://github.com/grpc/grpc-go