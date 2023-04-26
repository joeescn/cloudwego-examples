package main

import (
	"context"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/joeescn/hello-svc/kitex_gen/api"
	"github.com/joeescn/hello-svc/kitex_gen/api/hello"
	consul "github.com/kitex-contrib/registry-consul"
)

func main() {
	r, err := consul.NewConsulResolver("127.0.0.1:8500")
	if err != nil {
		klog.Fatal("Consul Resolver Error", err)
	}

	c, err := hello.NewClient(
		"hello-svc",
		client.WithResolver(r),
	)
	if err != nil {
		klog.Fatal("new hello client faild", err)
	}

	call(c, "")
	call(c, "resolver")

}

func call(c hello.Client, message string) {
	r, err := c.Echo(context.Background(), &api.Request{Message: message})
	if err != nil {
		klog.Error("resolver call error", err)
		return
	}
	klog.Info("resolver call success", r)
}
