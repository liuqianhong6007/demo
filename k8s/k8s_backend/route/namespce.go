package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/liuqianhong6007/demo/k8s/k8s_backend/k8s"
)

func init() {
	AddRoute(Routes{
		{
			Method:  http.MethodGet,
			Path:    "/namespace/list",
			Handler: ListNamespace,
		},
	})
}

func ListNamespace(c *gin.Context) {
	namespaces, err := k8s.CoreV1().Namespaces().List(c, v1.ListOptions{})
	checkValue(c, err, "list namespace error")

	c.JSON(http.StatusOK, namespaces)
}
