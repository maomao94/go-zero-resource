## 介绍
go-zero练手项目，包装了多种资源，整合成统一接口对外提供服务
## 默认模板
goctl api go -api resource.api -dir .
## 指定模板
goctl api go -api resource.api -dir . -home ../../../../template
## goland配置
-f service/resource/cmd/api/etc/resource-api.yaml
## 更新goctl版本
go get -u github.com/tal-tech/go-zero/tools/goctl
## 生成swagger
goctl api plugin -plugin goctl-swagger="swagger" -api resource.api -dir .
goctl api plugin -plugin goctl-swagger="swagger -filename resource.json -host 127.0.0.1:8888 -basepath /" -api resource.api -dir .
docker run --rm -p 8083:8080 -e SWAGGER_JSON=/foo/resource.json -v $PWD:/foo swaggerapi/swagger-ui
访问swagger地址 http://localhost:18006/doc.html#/home
## docker 镜像制作
goctl docker -go ./service/resource/cmd/api/resource.go
docker build -t resource:v1 .
docker tag resource:v1 hehanpeng/public/resource:v1