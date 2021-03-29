package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/liuqianhong6007/demo/k8s/internal"
)

func init() {
	internal.AddRoute(internal.Routes{
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
