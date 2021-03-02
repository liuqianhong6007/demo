package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	AddRoute(Routes{
		{
			Method:  http.MethodGet,
			Path:    "/etcd/ping",
			Handler: ping,
		},
	})
}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
