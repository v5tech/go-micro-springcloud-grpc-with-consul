package grpc

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/consul/api"
	"log"
	"testing"
)

var (
	services = []string{"golang_grpc_server", "grpc-server-springcloud-with-consul"}
)

func TestInvokeService(t *testing.T) {
	// client, err := api.NewClient(api.DefaultConfig())
	client, err := api.NewClient(&api.Config{
		Scheme:  "http",
		Address: "172.24.107.63:8500",
	})
	if err != nil {
		panic(err)
	}

	for _, service := range services {
		_, infos, err := client.Agent().AgentHealthServiceByName(service)
		if err != nil {
			log.Fatal("error: ", err)
		}
		for _, info := range infos {
			address := info.Service.Address
			port := info.Service.Port
			Invoke(fmt.Sprintf("%s:%d", address, port))
			fmt.Println(address, port)
			marshal, _ := json.MarshalIndent(info, "", " ")
			fmt.Println(string(marshal))
		}
	}

	fmt.Println("...")
}
