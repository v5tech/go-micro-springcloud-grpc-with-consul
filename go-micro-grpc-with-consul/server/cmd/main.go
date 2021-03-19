package main

import (
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
	"github.com/asim/go-micro/v3/util/log"
	"grpc-lib/pkg/proto/grpc.hello/proto"
	"server/pkg/srv"
)

func main() {
	// 初始化 consul 服务注册中心
	reg := consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			"172.24.107.63:8500",
		}
	})
	service := micro.NewService(
		// 把consul注册到micro
		micro.Registry(reg),
		// name很重要 服务发现就靠它
		micro.Name("golang_grpc_server"),
	)
	service.Init()
	// 注册服务
	er := proto.RegisterHelloHandler(service.Server(), new(srv.App))
	if er != nil {
		log.Errorf("RegisterHelloHandler err: %s", er.Error())
	}
	// run server
	if err := service.Run(); err != nil { // 启动
		panic(err)
	}
}
