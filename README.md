# go-mall
本项目用于本人的go语言的项目学习，主要从web端请求，到最终入库的过程

项目参考作者[CocaineCong](https://github.com/CocaineCong)的[gin-mall](https://github.com/CocaineCong/gin-mall)
只做了简单的开发，主要是为了将java中的mvc开发模式如何迁移到go语言中，适合初学者。

项目中的框架应用到：gorm + mysql+redis


### 项目运行
```
cd ./cmd
go run main.go
```

### 项目结构
````
gin-mall
├── api             # 用于定义接口函数，也就是controller的作用
├── cmd             # 程序入口
├── conf            # 配置文件
├── doc             # 文档
├── middleware      # 中间件
├── model           # 数据库模型
├── pkg             # 项目的一些公共函数
│  ├── e            # 错误码
│  └── util         # 工具函数
├── repository      # 项目的一些公共函数
│  ├── cache        # Redis缓存
│  ├── db           # 持久层的mysql
│  │  ├── dao       # dao层，对db进行操作
│  │  └── model     # 定义mysql的模型
│  ├── es           # ElasticSearch，形成elk体系
│  └── mq           # 放置各种mq，kafka，rabbitmq等等
├── routes          # 路由逻辑处理
├── serializer      # 将数据序列化为 json 的函数，便于返回给前端
├── service         # 接口函数的实现
└── static          # 存放静态文件
````

### 配置文件
配置文件放置在 /conf/config.yaml
````
system:
  domain : mall
  version : 1.0
  env : "dev"
  HttpPort: ":5001"
  Host: "localhost"
  UploadModel: "local"

mysql:
  default:
    dialect: "mysql"
    dbHost: "127.0.0.1"
    dbPort: "3306"
    dbName: "user_db"
    userName: "user"
    password: "user"
    charset: "utf8mb4"
    MaxIdleConns : 10
    MaxOpenConns : 100
    ConnMaxLifetime : 20
redis:
  redisDbName: 0
  redisHost: r-3ns7huz6q0teloxquw.redis.rds.aliyuncs.com
  redisPort: 6379
  redisUsername: front_risk_rw
  redisPassword: E4gb2ji580
  redisNetwork: "tcp"


encryptSecret:
  jwtSecret: "FanOne666Secret"
  emailSecret: "EmailSecret"
  phoneSecret: "PhoneSecret"
 
````

### 项目规划
1. 引入微服务架构grpc
2. 引入ddd领域模型
3. docker部署运行该项目
4. 集成企业应用的中间件，如：rocketmq、es、kafka、链路追踪、elk
5. 优化日志输出，统一用日志对象
