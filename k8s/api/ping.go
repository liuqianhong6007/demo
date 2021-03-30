package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/liuqianhong6007/demo/k8s/com"
)

func init() {
	com.AddRoute(com.Routes{
		{
			Method:  http.MethodGet,
			Path:    "/ping",
			Handler: ping,
		},
	})
}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
