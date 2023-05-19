

## go-zero-resource




## 前言
zero练手项目
- master分支 微服务版本 
- mono分支 单体版本，其中db未使用zero 自带的orm，改为gorm 兼容zero db缓存

## 接口地址
https://www.apifox.cn/apidoc/shared-6c1c58fe-fc04-45f9-abac-a4b7b71bbc62/api-7148139

## 其他
- go-stress-testing -c 1 -n 1 -p ossdetail.txt

## 模块
- api 1001
- sys 1002
- resource 1003
- message 1004
- mgtw 1005 21005
## 编译错误码
protoc errcode.proto --go_out=.
