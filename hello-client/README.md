# hello-client

## 创建服务

```shell
# 在 kitex-examples 文件夹下
md hello-client
cd hello-client
go mod init github.com/joeescn/hello-client
```

## 添加 `hello-svc` 依赖

这里使用golang提供的 work 功能，添加依赖

```shell
# 在 kitex-emamples 文件夹下执行
go work init hello-svc hello-client
```


## 调用示例

### HTTP 调用

通过 `hertz-gateway` 服务调用 `hello-svc`

### RPC 调用

通过 `hello-svc` 服务端口，直接调用。