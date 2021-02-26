package internal

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Routes []Route

type Route struct {
	Method  string
	Path    string
	Handler gin.HandlerFunc
}

var routeMap = make(map[string]Route)

func AddRoute(routes Routes) {
	for _, route := range routes {
		id := route.Method + " " + route.Path
		if _, ok := routeMap[id]; ok {
			panic("duplicate register router: " + id)
		}
		routeMap[route.Path] = route
	}
}

func Start(host string, port int) {
	r := gin.Default()
	r.Use(Cors()) // 跨域中间件必须放在 handler 注册之前
	for _, route := range routeMap {
		r.Handle(route.Method, route.Path, route.Handler)
	}
	if err := r.Run(fmt.Sprintf("%s:%d", host, port)); err != nil {
		panic(err)
	}
}

// 跨域中间件
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Access-Control-Allow-Origin,Authorization,Content-Type")
		c.Header("Access-Control-Max-Age", "172800")
		c.Header("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.JSON(http.StatusOK, nil)
		}
		c.Next()
	}
}

func CheckValue(c *gin.Context, checkValue interface{}, errMsg ...string) {
	switch val := checkValue.(type) {
	case error:
		if val != nil {
			switch internalErr := val.(type) {
			case WrapError:
				FailJsonRsp(c, internalErr.ExposeError())
			default:
				FailJsonRsp(c, internalErr.Error())
			}
			panic(val)
		}
	case bool:
		if !val {
			errMsg1 := strings.Join(errMsg, "\n")
			FailJsonRsp(c, errMsg1)
			panic(errMsg1)
		}
	}
}

func SuccessJsonRsp(c *gin.Context, result interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}

func FailJsonRsp(c *gin.Context, errMsg string) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"message": errMsg,
	})
}
