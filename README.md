# gin-api-bootstrap

gin开发http api项目的启动脚手架

安装依赖包

```
go get -u github.com/gin-gonic/gin
go get -u github.com/swaggo/gin-swagger
go mod download github.com/swaggo/files
go get -u github.com/swaggo/swag/cmd/swag
go install github.com/swaggo/swag/cmd/swag@latest
```

依赖包版本

- gin v1.9.0
- gin-swagger v1.3.0

# 目录结构说明

```
├─api           # http接口
├─config        # 项目配置
├─docs          # 由 swag 生成的swagger接口文档
├─middleware    # 中间件
├─model         # ORM模型
├─serializer    # 序列化对象
├─server        # URL路由
├─service       # 业务服务
└─util          # 辅助工具
```

生成swag文档

```
swag init
```

启动服务

```
go run main.go
```

访问下面地址，即可访问接口文档

http://127.0.0.1:3000/swagger/index.html


参考资料

- https://github.com/gourouting/singo
