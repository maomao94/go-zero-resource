// 默认模板
goctl api go -api resource.api -dir .
// 指定模板
goctl api go -api resource.api -dir . -home ../../../../template
-f service/resource/cmd/api/etc/resource-api.yaml