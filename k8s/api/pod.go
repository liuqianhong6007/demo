package api

import (
	"github.com/liuqianhong6007/demo/k8s/internal"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/liuqianhong6007/demo/k8s/com"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func init() {
	com.AddRoute(com.Routes{
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
	com.CheckValue(c, namespace != "", "param[namespace] is null")

	list, err := internal.K8sClientset().CoreV1().Pods(namespace).List(c, metav1.ListOptions{})
	com.CheckValue(c, err, "list pod error")

	c.JSON(http.StatusOK, list)
}

func GetPod(c *gin.Context) {
	namespace := c.Query("namespace")
	com.CheckValue(c, namespace != "", "param[namespace] is null")

	podName := c.Query("podName")
	com.CheckValue(c, podName != "", "param[podName] is null")

	pod, err := internal.K8sClientset().CoreV1().Pods(namespace).Get(c, podName, metav1.GetOptions{})
	com.CheckValue(c, err, "get pod error")

	c.JSON(http.StatusOK, pod)
}
