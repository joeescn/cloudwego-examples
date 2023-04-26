package main

import (
	"context"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/joeescn/hello-svc/kitex_gen/api"
	"github.com/joeescn/hello-svc/kitex_gen/api/hello"
)

func main() {
	c, err := hello.NewClient(
		"hello-svc",
		client.WithHostPorts("127.0.0.1:8888"),
	)
	if err != nil {
		klog.Fatal("new hello client faild", err)
	}

	// 异常调用
	call(c, "")
	// 正常调用
	call(c, "rpc connect")
}

func call(c hello.Client, message string) {
	ctx := context.Background()
	r, err := c.Echo(ctx, &api.Request{Message: message})
	if err != nil {
		klog.Error("rcp call faild", err)
		return
	}
	klog.Info("rpc call success", r)
}
