# connToDatabase

#### 介绍
本项目介绍了go语言连接数据库的实例，在阅读本项目之前请确保您安装了mysql数据库，更多相关内容可以参考[这里](https://pkg.go.dev/github.com/go-sql-driver/mysql)

#### 快速开始

安装数据库驱动

```go
go get -u github.com/go-sql-driver/mysql
```

初始化模块

```go
go mod init connToDatabase
```

安装gorm，关于gorm可以参考[这里](https://gorm.io/)

```go
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
```

执行go mod tidy，它可以引用项目需要的依赖增加到go.mod文件，也可以去掉go.mod文件中项目不需要的依赖。

```go
go mod tidy
```



