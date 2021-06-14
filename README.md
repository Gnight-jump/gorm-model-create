# gorm_model_create

#### 介绍
连接数据库，自动生成go的model

#### 使用方法
依照注释修改conf.yaml
```yaml
#配置文件，按照要求修改输入数据，当前仅支持mysql数据库
Database:
  #数据库配置
  Host: 127.0.0.1
  Port: 3306
  User: test
  Password: 123321
  DBName: gotest
GModel:
  #存储路径，最后面也需要分割符
  StorePath: "./dao/models/"
  #是否覆盖原本的文件，false为不覆盖
  ModelCover: true
  #存储的package名
  PackageName: "models"
  #需要生成model的表名，支持多个
  TableName:
    - "screenshot"
    - "omgphone"
```

# 命令行 使用说明

一、初始化`conf.yaml`
```go
// 默认初始化方式
go run main.go -isInit true
// 使用非默认路径(在conf目录下生成conf.yaml)
go run main.go --isInit -initPath "./config"
```

二、生成`model`
```go
// 默认方式(配合默认初始化方式)
go run main.go
// 使用非默认路径(在conf目录下生成conf.yaml)
go run main.go -confPath "./config"
```
