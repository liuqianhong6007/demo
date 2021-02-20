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
			Path:    "/pod/list",
			Handler: ListPod,
		},
		{
			Method:  http.MethodGet,
			Path:    "/pod",
			Handler: GetPod,
		},
	})
}

func ListPod(c *gin.Context) {
	namespace := c.Query("namespace")
	checkValue(c, namespace != "", "param[namespace] is null")

	list, err := k8s.CoreV1().Pods(namespace).List(c, v1.ListOptions{})
	checkValue(c, err, "list pod error")

	c.JSON(http.StatusOK, list)
}

func GetPod(c *gin.Context) {
	namespace := c.Query("namespace")
	checkValue(c, namespace != "", "param[namespace] is null")

	podName := c.Query("podName")
	checkValue(c, podName != "", "param[podName] is null")

	pod, err := k8s.CoreV1().Pods(namespace).Get(c, podName, v1.GetOptions{})
	checkValue(c, err, "get pod error")

	c.JSON(http.StatusOK, pod)
}
