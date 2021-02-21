package route

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/liuqianhong6007/demo/etcd/etcd_backend/etcd"
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
	val, err := etcd.Get(key)
	checkValue(c, err, "get etcd value error")

	SuccessJsonRsp(c, val)
}

type PutEtcdParam struct {
	Key string `json:"key"`
	Val string `json:"val"`
}

func putEtcd(c *gin.Context) {
	var param PutEtcdParam
	err := c.BindJSON(&param)
	checkValue(c, err, "param format error")
	checkValue(c, param.Key != "", "param[key] is null")
	checkValue(c, param.Val != "", "param[val] is null")

	err = etcd.Put(param.Key, param.Val)
	checkValue(c, err, "put etcd value error")

	SuccessJsonRsp(c, nil)
}

type DelEtcdParam struct {
	Keys []string `json:"keys"`
}

func deleteEtcd(c *gin.Context) {
	var param DelEtcdParam
	err := c.BindJSON(&param)
	checkValue(c, err, "param format error")
	checkValue(c, len(param.Keys) != 0, "param[keys] is null")

	for _, key := range param.Keys {
		err := etcd.Delete(key)
		checkValue(c, err, "delete etcd value error")
	}

	SuccessJsonRsp(c, nil)
}
