package grpc

import (
	"client/pkg/grpc/proto/grpc.hello/proto"
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
)

// grpc原生方式调用grpc服务
func Invoke(address string) {
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
