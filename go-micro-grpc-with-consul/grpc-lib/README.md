# grpc-lib

根据proto文件生成pb文件

```bash
protoc --proto_path=$GOPATH/src:. --micro_out=. --go_out=. hello.proto
```