package api

import (
	"github.com/liuqianhong6007/demo/k8s/internal"
	"net/http"

	"github.com/gin-gonic/gin"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/liuqianhong6007/demo/k8s/com"
)

func init() {
	com.AddRoute(com.Routes{
		{
			Method:  http.MethodGet,
			Path:    "/namespace/list",
			Handler: ListNamespace,
		},
	})
}

func ListNamespace(c *gin.Context) {
	namespaces, err := internal.K8sClientset().CoreV1().Namespaces().List(c, v1.ListOptions{})
	com.CheckValue(c, err, "list namespace error")

	c.JSON(http.StatusOK, namespaces)
}
