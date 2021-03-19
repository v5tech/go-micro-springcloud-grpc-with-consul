package srv

import (
	"context"
	"fmt"
	"grpc-lib/pkg/proto/grpc.hello/proto"
)

func (s *App) SayHello(ctx context.Context, req *proto.HelloRequest, reply *proto.HelloReply) error {
	fmt.Println("received:", req.Name)
	reply.Message = fmt.Sprintf("Hello,%s! this message come form Golang GRPC Server Application!", req.Name)
	return nil
}
