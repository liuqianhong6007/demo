package route

import (
	"github.com/liuqianhong6007/demo/etcd/etcd_backend/etcd"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	AddRoute(Routes{
		{
			Method:  http.MethodGet,
			Path:    "/etcd/get",
			Handler: getEtcd,
		},
		{
			Method:  http.MethodPost,
			Path:    "/etcd/add",
			Handler: putEtcd,
		},
		{
			Method:  http.MethodPost,
			Path:    "/etcd/delete",
			Handler: deleteEtcd,
		},
	})
}

func getEtcd(c *gin.Context) {
	key := c.Query("key")
	checkValue(c, key != "", "param[key] is null")

	val, err := etcd.Get(key)
	checkValue(c, err, "get etcd value error")

	SuccessJsonRsp(c, val)
}

func putEtcd(c *gin.Context) {
	key := c.PostForm("key")
	checkValue(c, key != "", "param[key] is null")

	val := c.PostForm("val")
	checkValue(c, val != "", "param[val] is null")

	err := etcd.Put(key, val)
	checkValue(c, err, "put etcd value error")

	SuccessJsonRsp(c, nil)
}

func deleteEtcd(c *gin.Context) {
	key := c.PostForm("key")
	checkValue(c, key != "", "param[key] is null")

	err := etcd.Delete(key)
	checkValue(c, err, "delete etcd value error")

	SuccessJsonRsp(c, nil)
}
