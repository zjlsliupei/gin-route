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
    ginRouter "github.com/zjlsliupei/gin-route"
)

type Test struct {
}

func (t *Test) Hello(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"success": true,
	})
}

func main() {
    g := gin.Default()
    r := g.Group("/admin")
    ginRouter.AutoRouter(r, &Test{})
    g.Run(":8083")
}




