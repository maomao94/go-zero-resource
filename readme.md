## 默认模板
goctl api go -api resource.api -dir .
## 指定模板
goctl api go -api resource.api -dir . -home ../../../../template
## goland配置
-f service/resource/cmd/api/etc/resource-api.yaml
## 更新goctl版本
go get -u github.com/tal-tech/go-zero/tools/goctl