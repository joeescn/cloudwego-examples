package main

import (
	"log"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	api "github.com/joeescn/hello-svc/kitex_gen/api/hello"
	consul "github.com/kitex-contrib/registry-consul"
)

func main() {

	// address 参数为 consul 地址
	r, err := consul.NewConsulRegister("127.0.0.1:8500")
	if err != nil {
		log.Fatal(err)
	}

	svr := api.NewServer(new(HelloImpl), server.WithRegistry(r), server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: "hello-svc", // 和初始化命令的 -service [ServiceName] 参数保持一致
	}))

	if err := svr.Run(); err != nil {
		log.Println(err.Error())
	}
}
