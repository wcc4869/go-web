# go-web-edu

### 使用 gin + gorm 完成第二期课后练习，gorm 完成数据库的增删改查，viper读取配置文件


### 项目结构

```text
├── common
│   └── function.go
├── config
│   └── app.yml
├── controller
│   ├── article.go
│   └── category.go
├── go.mod
├── go.sum
├── main.go
├── model
│   ├── article.go
│   ├── category.go
│   └── time.go
├── response
│   └── response.go
├── router.go
├── routers
│   ├── Article.go
│   └── category.go
├── service
│   ├── article.go
│   ├── category.go
│   └── database.go
└── vendor

```

- controller: 控制器
- model: 模型文件
- configs: 配置文件
- routers: 路由文件
- response: 针对返回函数封装
- service: 服务逻辑处理层
- common: 通用的文件
- vendor: 三方包

### 工行创新项目目录文件约定

```text
├── common
│   └── function.go 通用的方法文件
├── config
│   └── app.yml    配置文件
├── controller
│   ├── article.go  文章控制器
│   └── category.go 分类控制器
├── go.mod
├── go.sum
├── main.go  入口文件
├── model
│   ├── article.go  文章模型
│   ├── category.go 分类模型
│   └── time.go     时间封装，处理时间格式
├── response
│   └── response.go  返回函数封装
├── router.go
├── routers
│   ├── Article.go  文章路由
│   └── category.go 分类路由
├── service
│   ├── article.go  文章操作数据库
│   ├── category.go 分类操作数据库
│   └── database.go 初始化数据库
```

