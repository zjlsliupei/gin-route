package gin_route

import (
	"github.com/gin-gonic/gin"
	"testing"
)

type Test struct {
}

func (t *Test) Hello(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"success": true,
	})
}

func TestAutoRouter(t *testing.T) {
	g := gin.Default()
	r := g.Group("/admin")
	AutoRouter(r, &Test{})
	g.Run(":8083")
}
