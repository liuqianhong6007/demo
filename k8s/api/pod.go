package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/liuqianhong6007/demo/k8s/internal"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func init() {
	internal.AddRoute(internal.Routes{
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
	internal.CheckValue(c, namespace != "", "param[namespace] is null")

	list, err := internal.CoreV1().Pods(namespace).List(c, metav1.ListOptions{})
	internal.CheckValue(c, err, "list pod error")

	c.JSON(http.StatusOK, list)
}

func GetPod(c *gin.Context) {
	namespace := c.Query("namespace")
	internal.CheckValue(c, namespace != "", "param[namespace] is null")

	podName := c.Query("podName")
	internal.CheckValue(c, podName != "", "param[podName] is null")

	pod, err := internal.CoreV1().Pods(namespace).Get(c, podName, metav1.GetOptions{})
	internal.CheckValue(c, err, "get pod error")

	c.JSON(http.StatusOK, pod)
}
