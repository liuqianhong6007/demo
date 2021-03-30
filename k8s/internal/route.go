package internal

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
		if _, ok := routeMap[route.Path]; ok {
			panic("duplicate register router: " + route.Path)
		}
		routeMap[route.Path] = route
	}
}

func RegisterRoute(engine *gin.Engine) {
	for _, route := range routeMap {
		engine.Handle(route.Method, route.Path, route.Handler)
	}
}

func CheckValue(c *gin.Context, checkValue interface{}, errMsg ...string) {
	msg := strings.Join(errMsg, "\n")
	switch val := checkValue.(type) {
	case error:
		if val != nil {
			if msg == "" {
				msg = val.Error()
			}
			c.JSON(http.StatusInternalServerError, msg)
			panic(msg + "\n" + val.Error())
		}
	case bool:
		if !val {
			c.JSON(http.StatusInternalServerError, msg)
			panic(msg)
		}
	}
}
