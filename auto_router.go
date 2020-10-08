package gin_route

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"reflect"
	"strings"
)

func AutoRouter(gRouteGroup *gin.RouterGroup, controller interface{}) {
	t := reflect.TypeOf(controller)
	structName := t.Elem().Name()
	if t.Kind() != reflect.Ptr {
		fmt.Println(structName + " is not struct")
		return
	}
	v := reflect.ValueOf(controller)
	for i := 0; i < v.NumMethod(); i++ {
		methodName := t.Method(i).Name
		fun := v.MethodByName(methodName)
		gRouteGroup.Any(strings.ToLower("/"+structName+"/"+methodName), func(context *gin.Context) {
			fun.Call([]reflect.Value{reflect.ValueOf(context)})
		})
	}
}
