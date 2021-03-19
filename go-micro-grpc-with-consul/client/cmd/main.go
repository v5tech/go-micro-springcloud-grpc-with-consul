package main

import (
	"context"
	"fmt"
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
	"grpc-lib/pkg/proto/grpc.hello/proto"
	"time"
)

func main() {
	reg := consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			"172.24.107.63:8500",
		}
	})
	service := micro.NewService(
		micro.Registry(reg),
	)
	service.Init()
	// callJava(service)
	callGolang(service)
}

// 调用Golang Grpc Service
func callGolang(service micro.Service) {
	// golang_grpc_server为注册在consul中的grpc服务名
	helloService := proto.NewHelloService("golang_grpc_server", service.Client())
	for {
		rsp, err := helloService.SayHello(context.Background(), &proto.HelloRequest{
			Name: "Microsoft",
		})
		if err != nil {
			panic(err)
		}
		now := time.Now().Format("2006-01-02 15:04:05")
		fmt.Println(fmt.Sprintf(now+" %v", rsp))
		time.Sleep(3 * time.Second)
	}
}

// 调用Java Grpc Service
func callJava(service micro.Service) {
	// grpc-server-springcloud-with-consul为注册在consul中的grpc服务名
	helloService := proto.NewHelloService("grpc-server-springcloud-with-consul", service.Client())
	for {
		rsp, err := helloService.SayHello(context.Background(), &proto.HelloRequest{
			Name: "Microsoft",
		})
		if err != nil {
			panic(err)
		}
		now := time.Now().Format("2006-01-02 15:04:05")
		fmt.Println(fmt.Sprintf(now+" %v", rsp))
		time.Sleep(3 * time.Second)
	}
}
