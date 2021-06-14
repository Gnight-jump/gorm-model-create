# gorm_model_create

#### Description
Connect to the database to automatically generate the corresponding form model

#### Software Architecture
Modify conf.yaml according to the comments
```yaml
# modify the input data as required, currently only support MySQL database
Database:
# database configuration
Host: 127.0.0.1
Port: 3306
User: the test
Password: 123321
DBName: gotest
GModel:
# store the path, the last also need a divider
StorePath: ". / dao/models/"
# do not overwrite the original file, false is not overwrite
ModelCover: true,
# The name of the stored Package
PackageName: "models"
# create table name from model
TableName:
- "screenshot"
- "omgphone"
```

# Command line instructions

一、initialize `conf.yaml`
```go
// Default initialization mode
go run main.go -isInit true
// Use a non-default path (generate conf.yaml in the ./config directory)
go run main.go --isInit -initPath "./config"
```

二、 generate `model`
```go
// Default mode (with default initialization mode)
go run main.go
// Use a non-default path (generate conf.yaml in the./config directory)
go run main.go -confPath "./config"
```