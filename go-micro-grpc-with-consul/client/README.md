# client

grpc-go 原生方式调用

```bash
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
go get -u google.golang.org/grpc
protoc --go_out=plugins=grpc:. *.proto
```

调用示例：

```
package main

import (
	"client/pkg/proto/grpc.hello/proto"
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	address = "127.0.0.1:9898"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := proto.NewHelloClient(conn)
	for {
		r, err := c.SayHello(context.Background(), &proto.HelloRequest{Name: "Microsoft"})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Greeting: %s", r.GetMessage())
		time.Sleep(3 * time.Second)
	}
}
```

参考资料：

https://github.com/grpc/grpc-go