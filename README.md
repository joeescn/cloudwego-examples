# kitex-examples

## 运行

```shell
# consul
docker run -d -p 8500:8500 --restart=always --name=consul consul:latest agent -server -bootstrap -ui -node=1 -client='0.0.0.0'

# 运行 http 网关
cd hello-gateway

# hello-svc 服务
cd hello-svc
go run .

# hello-client 服务
cd hello-client
# 直接调用
go run cmd/rpc/main.go 
# 服务发现调用
go run cmd/resolver/main.go
# http-gateway 网关调用
go run cmd/http-gateway/main.go
```

## hello-client 输出

```shell
➜  hello-client go run cmd/rpc/main.go
2023/04/27 00:08:00.949244 main.go:31: [Error] rcp call faildremote or network error[remote]: biz error: biz error: message cannot be empty
2023/04/27 00:08:00.949816 main.go:34: [Info] rpc call successResponse({Message:input message:rpc connect})
➜  hello-client go run cmd/resolver/main.go
2023/04/27 00:09:04.125617 main.go:35: [Error] resolver call errorremote or network error[remote]: biz error: biz error: message cannot be empty
2023/04/27 00:09:04.126155 main.go:38: [Info] resolver call successResponse({Message:input message:resolver})
➜  hello-client go run cmd/http-gateway/main.go
2023/04/27 00:16:56.175335 main.go:44: [Info] {}
2023/04/27 00:16:56.176537 main.go:44: [Info] {"ErrCode":0,"ErrMessage":"ok","message":"input message:http-gateway"}
```