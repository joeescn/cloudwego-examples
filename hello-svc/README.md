# hello-svc

使用 `go mod` 管理依赖，使用 `go work` 环境开发，所以代码可以放置在任意目录，不强制。

## 环境

- go1.19.8
- docker20.10.14
- hashicorp/consul@latest
    ```shell
    # 启动脚本
    docker run -d -p 8500:8500 --restart=always --name=consul consul:latest agent -server -bootstrap -ui -node=1 -client='0.0.0.0'
    ```
- 

## 安装代码生成工具 

- 安装 `kitex`
  ```shell
  go install github.com/cloudwego/kitex/tool/cmd/kitex@latest
  kitex -version
  #v0.5.2
  ```
- 安装 `thriftgo`
  ```shell
  go install github.com/cloudwego/thriftgo@latest
  thriftgo --version
  #thriftgo 0.2.9
  ```


## 创建 `workspace`

```shell
md kitex-examples
cd kitex-examples
```

## 创建服务端 `hello-svc`

```shell
md hello-svc
cd hello-svc
```


## 编辑 IDL

idl/hello-svc.thrifht

```thrift
namespace go api

struct Request {
    1: string message,
}

struct Response {
    1: string message,
}

struct AddRequest {
    1: i64 first,
    2: i64 second,
}

struct AddResponse {
    1: i64 sum,
}

service Hello {
    Response echo(1: Request req),
    AddResponse add(1: AddRequest req),
}
```

## 生成代码

```shell
kitex -module "github.com/joeescn/hello-svc" -service hello-svc idl/hello-svc.thrift
# 加载依赖
go mod tidy
```

## 服务注册

编辑 `main.go` , 将 `hello-svc` 服务注册到 `consul` 注册中心。

[cloudwego文档](https://www.cloudwego.io/zh/docs/kitex/tutorials/service-governance/discovery/)

### 安装依赖
```shell
go get github.com/kitex-contrib/registry-consul
```

### 修改代码 `main.go`

修改的目的是为了把 `hello-svc` 服务注册到注册中心，这样子后续的客户端就可以通过注册中心的服务发现调用该服务。

```go
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
```

## 编写业务代码

在 `handler.go` 修改代码，添加自己的业务逻辑。


## 启动服务

启动服务后，进入 WEB 控制台[consul-ui](http://localhost:8500) 查看服务注册情况，查看 `hello-svc` 是否注册成功。

```shell
go run .
```