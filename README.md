# gin-route 
gin增强路由模块，提供自动路由功能
## 安装
```go
go get github.com/zjlsliupei/gin-route
```

## 使用
```go
import (
    "github.com/gin-gonic/gin"
    "github.com/zjlsliupei/gin-route"
)
g := gin.Default()
r := g.Group("/admin")
gin-route.AutoRouter(r, &Test{})
g.Run(":8083")

