# FLY — A CLI tool for building go application.
FLY是一个基于Golang的应用脚手架,
该项目是站在巨人的肩膀上，
它是由Golang生态中各种非常流行的库整合而成的，
它们的组合可以帮助你快速构建一个高效、可靠的应用程序,
目前该项目还在实验个人学到的知识点，不建议放在商业项目中

## 第三方库
- **Gin**: https://github.com/gin-gonic/gin
- **Gorm**: https://github.com/go-gorm/gorm
- **Wire**: https://github.com/google/wire
- **Viper**: https://github.com/spf13/viper
- **Zap**: https://github.com/uber-go/zap
- **Golang-jwt**: https://github.com/golang-jwt/jwt
- **Go-redis**: https://github.com/go-redis/redis
- **Testify**: https://github.com/stretchr/testify
- More...

## 简洁分层架构
FLY采用了经典的分层架构。同时，为了更好地实现模块化和解耦，采用了依赖注入框架`Wire`。
尝试使用DDD领域架构

## 安装
```shell
go install github.com/renfy96/fly@latest
```