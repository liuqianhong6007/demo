package test

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

var (
	agent *HttpAgent
	host  = "127.0.0.1:8081"
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

func Test_Register(t *testing.T) {
	url := fmt.Sprintf("http://%s/auth/register", host)
	hash := md5.New()
	password := hex.EncodeToString(hash.Sum([]byte("password")))
	buf, _ := json.Marshal(map[string]interface{}{
		"account":     "lqh",
		"password":    password,
		"invite_code": "123456",
	})
	rsp, err := agent.Post(url, nil, buf)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(readRsp(rsp))
}

func Test_Login(t *testing.T) {
	url := fmt.Sprintf("http://%s/auth/login", host)
	hash := md5.New()
	password := hex.EncodeToString(hash.Sum([]byte("password")))
	buff, _ := json.Marshal(map[string]interface{}{
		"account":  "lqh",
		"password": password,
	})
	rsp, err := agent.Post(url, nil, buff)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(readRsp(rsp))
}

func Test_CheckToken(t *testing.T) {
	url := fmt.Sprintf("http://%s/auth/checkToken", host)
	header := http.Header{}
	header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MDAwLCJpc3MiOiJscWgiLCJhY2NvdW50IjoibHFoIn0.b_3hQx2aIiSzt9SeFirahzFeD13qUzSjOpMZ-4zK68g")
	rsp, err := agent.Get(url, header, nil)
	if err != nil {
		t.Fatal(err)
	}
	if rsp.StatusCode == http.StatusUnauthorized {
		t.Fatal("unauthorized")
	} else if rsp.StatusCode == http.StatusOK {
		t.Log("check token success")
	}
	t.Log(readRsp(rsp))
}
