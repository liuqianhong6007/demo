package route

import (
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

func RegisterRoute(engine *gin.Engine) {
	// 跨域中间件必须放在 handler 注册之前
	engine.Use(Cors())
	for _, route := range routeMap {
		engine.Handle(route.Method, route.Path, route.Handler)
	}
	// 静态资源服务器
	engine.StaticFS("/html", http.Dir("./html"))
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

func checkValue(c *gin.Context, checkValue interface{}, errMsg ...string) {
	switch val := checkValue.(type) {
	case error:
		if val != nil {
			errMsg1 := strings.Join(errMsg, "\n") + "\n" + val.Error()
			FailJsonRsp(c, errMsg1)
			panic(errMsg1)
		}
	case bool:
		if !val {
			errMsg1 := strings.Join(errMsg, "\n")
			FailJsonRsp(c, errMsg1)
			panic(errMsg1)
		}
	}
}

type Code int

const (
	OK      = 1000
	UNKNOWN = 9999
)

func SuccessJsonRsp(c *gin.Context, result interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code":   OK,
		"result": result,
	})
}

func FailJsonRsp(c *gin.Context, errMsg string) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"code":    UNKNOWN,
		"message": errMsg,
	})
}
