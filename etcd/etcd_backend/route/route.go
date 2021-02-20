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
	for _, route := range routeMap {
		engine.Handle(route.Method, route.Path, route.Handler)
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
