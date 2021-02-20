package test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
)

var (
	agent *HttpAgent
	host  = "127.0.0.1:8101"
)

func init() {
	agent = NewHttpAgent()
}

func readRsp(rsp *http.Response) string {
	buf, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		panic(err)
	}
	defer rsp.Body.Close()
	return string(buf)
}

func Test_EtcdAdd(t *testing.T) {
	url1 := fmt.Sprintf("http://%s/etcd/add", host)
	params := url.Values{
		"key": {"key1"},
		"val": {"val2"},
	}
	rsp, err := agent.PostForm(url1, nil, params)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(readRsp(rsp))
}

func Test_EtcdGet(t *testing.T) {
	url := fmt.Sprintf("http://%s/etcd/get", host)
	params := map[string]string{
		"key": "key1",
	}
	rsp, err := agent.Get(url, nil, params)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(readRsp(rsp))
}

func Test_EtcdDelete(t *testing.T) {
	url1 := fmt.Sprintf("http://%s/etcd/delete", host)
	params := url.Values{
		"key": {"key1"},
	}
	rsp, err := agent.PostForm(url1, nil, params)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(readRsp(rsp))
}
