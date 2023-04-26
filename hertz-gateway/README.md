# hertz-gateway

## doc

- https://www.cloudwego.io/zh/docs/hertz/getting-started/
- https://juejin.cn/post/7216917459009732668
- https://github.com/cloudwego/biz-demo


## 执行命令

```shell
md hertz-gateway
cd hertz-gateway

go mod init github.com/joeescn/hertz-gateway

cd ..
go work use hertz-gateway
cd hertz-gateway

go install github.com/cloudwego/hertz/cmd/hz@latest

hz -v
#hz version v0.6.2

# 创建代码
hz new -mod github.com/joeescn/hertz-gateway

```