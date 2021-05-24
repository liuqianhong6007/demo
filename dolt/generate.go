package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"
)

func Gen(metadataFile string) error {
	buffArray, err := read(metadataFile)
	if err != nil {
		log.Println(err)
		return err
	}

	if len(buffArray) == 0 {
		return nil
	}

	var docBuf []byte
	docBuf = append(docBuf, []byte(docAPITitle)...)

	for _, buff := range buffArray {
		var cmd Cmd
		err = json.Unmarshal(buff, &cmd)
		if err != nil {
			log.Println(err)
			return err
		}

		var cmdTemplate CmdTemplate
		cmdTemplate.Cmd = cmd.Command
		cmdTemplate.StructName = normalize(cmd.Command)
		for _, arg := range cmd.Args {
			cmdTemplate.Args = append(cmdTemplate.Args, CmdTemplateItem{
				Value: normalize(arg.Value),
				Cmd:   arg.Value,
				Desc:  arg.Desc,
			})
		}
		for _, option := range cmd.Options {
			cmdTemplate.Options = append(cmdTemplate.Options, CmdTemplateItem{
				Value: normalize(option.Value),
				Cmd:   option.Value,
				Desc:  option.Desc,
			})
		}
		for _, f := range cmd.Flags {
			cmdTemplate.Flags = append(cmdTemplate.Flags, CmdTemplateItem{
				Value: normalize(f.Value),
				Cmd:   f.Value,
				Desc:  f.Desc,
			})
		}

		err = genFile(strings.ToLower(cmdTemplate.StructName)+".auto.go", cmdTemplate)
		if err != nil {
			log.Println(err)
			return err
		}

		err = genTestFile(strings.ToLower(cmdTemplate.StructName)+".auto_test.go", cmdTemplate)
		if err != nil {
			log.Println(err)
			return err
		}

		buf, err := genAPIDoc(cmdTemplate)
		if err != nil {
			log.Println(err)
			return err
		}
		docBuf = append(docBuf, buf...)
	}

	err = genTestUtilFile("dolt_test_util.go", nil)
	if err != nil {
		log.Println(err)
		return err
	}

	err = write("dolt_api_doc.md", docBuf)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

var apiTemplate = `
package main

import(
	"fmt"
	"net/http"
	"os/exec"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func init(){
	gServer.RegRoute(http.MethodPost, "/{{.StructName}}", gServer.{{.StructName}})
}

type {{.StructName}}_args struct{
	{{- range .Args}} 
	{{.Value}} string $(CMD_TAG)
	{{- end}}
}

type {{.StructName}}_opts struct{
	{{- range .Options}}
	{{.Value}} string $(CMD_TAG)
	{{- end}}
}

type {{.StructName}}_flags struct{
	{{- range .Flags}}
	{{.Value}} bool $(CMD_TAG)
	{{- end}}
}

type {{.StructName}} struct{
	Cmd *struct{} $(CMD_IGNORE_JSON_TAG)
	Args {{.StructName}}_args $(ARG_JSON_TAG)
	Opts {{.StructName}}_opts $(OPT_JSON_TAG)
	Flags {{.StructName}}_flags $(FLAG_JSON_TAG)
}

func (s *Server) {{.StructName}}(c *gin.Context){
	var param {{.StructName}}
	err := c.BindJSON(&param)
	if err != nil{
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": "Request param error",
		})
		s.logger.Error("Request param error",zap.Error(err))
		return
	}

	var args []string
	// get command string
	cmdField, _ := reflect.TypeOf(param).FieldByName("Cmd")
	sub := strings.Split(cmdField.Tag.Get("cmd")," ")
	commandStr := sub[0]
	args = append(args,sub[1:]...)

	// parse cmd opt
	{
		typeObj := reflect.TypeOf(param.Opts)
		valObj := reflect.ValueOf(param.Opts)
		for i := 0; i < typeObj.NumField(); i++ {
			if valObj.Field(i).String() != "" {
				args = append(args, fmt.Sprintf("--%s=%s", typeObj.Field(i).Tag.Get("cmd"), valObj.Field(i).String()))
			}
		}
	}

	// parse cmd flag
	{
		typeObj := reflect.TypeOf(param.Flags)
		valObj := reflect.ValueOf(param.Flags)
		for i := 0; i < typeObj.NumField(); i++ {
			if valObj.Field(i).Bool() {
				args = append(args, fmt.Sprintf("--%s", typeObj.Field(i).Tag.Get("cmd")))
			}
		}
	}

	// parse cmd arg
	{
		typeObj := reflect.TypeOf(param.Args)
		valObj := reflect.ValueOf(param.Args)
		for i := 0; i < typeObj.NumField(); i++ {
			if valObj.Field(i).String() != "" {
				args = append(args, valObj.Field(i).String())
			}
		}
	}

	// call dolt command
	result,err := Execute(exec.Command(commandStr, args...))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		s.logger.Error(err.Error())
		return
	}
	
	c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "OK",
		"ret":     result,
	})
}

`

var apiTestTemplate = `
package main

import(
	"net/http"
	"testing"
)

func Test_{{.StructName}}(t *testing.T){
	var rsp JsonResponse
	err := JsonRequest(
		http.MethodPost, 
		reqPath("/{{.StructName}}"), {{.StructName}}{
			Args: {{.StructName}}_args{
				{{- range .Args}} 
				{{.Value}}: "",
				{{- end}}
			},
			Opts: {{.StructName}}_opts{
				{{- range .Options}} 
				{{.Value}}: "",
				{{- end}}		
			},
			Flags: {{.StructName}}_flags{
				{{- range .Flags}} 
				{{.Value}}: false,
				{{- end}}		
			},
		},
		nil,
		&rsp,
	)
	if err != nil{
		t.Fatal(err)
	}
	t.Log(marshalJson(rsp))
	t.Log(rsp.Ret)
}

`

var apiTestUtilTemplate = `
package main

import(
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"	
	"net/http"
	"reflect"
	"strings"
)

type SignatureHook func(req *http.Request, msg []byte) error

func JsonRequest(method, path string, params interface{}, signature SignatureHook, ret interface{}) error {
	if ret != nil {
		if reflect.TypeOf(ret).Kind() != reflect.Ptr {
			return errors.New("ret must be a ptr type")
		}
	}

	client := http.Client{}
	req, err := http.NewRequest(method, path, nil)
	if err != nil {
		return err
	}
	req.Header.Set("content-type", "application/json")

	var msg []byte
	if params != nil {
		buf, err := json.Marshal(params)
		if err != nil {
			return err
		}
		req.ContentLength = int64(len(buf))
		req.Body = ioutil.NopCloser(bytes.NewReader(buf))
		msg = buf
	}

	if signature != nil {
		err = signature(req, msg)
		if err != nil {
			return err
		}
	}

	rsp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer rsp.Body.Close()

	switch rsp.StatusCode {
	case http.StatusOK:
		if ret != nil {
			buf, err := ioutil.ReadAll(rsp.Body)
			if err != nil {
				return err
			}
			err = json.Unmarshal(buf, ret)
			if err != nil {
				return err
			}
		}

	default:
		return errors.New("http return unexpected status code: " + rsp.Status)
	}

	return nil
}

func reqPath(rPath string) string {
	if strings.HasPrefix(rPath, "/") {
		rPath = strings.TrimPrefix(rPath, "/")
	}
	return fmt.Sprintf("http://127.0.0.1:8600/%s", rPath)
}

type JsonResponse struct{
	Code int
	Message string
	Ret interface{}
}

func marshalJson(obj interface{}) string {
	buf,_ := json.Marshal(obj)
	return string(buf)
}

`

var docAPITitle = `
# Dolt API 接口文档

## 统一规范

请求结果返回三个字段

* Code

返回状态码，200 正常返回，400 请求参数错误, 500 内部错误

* Message

当 Code 返回非 200 时，Message 显示具体错误

* Ret

当 Code 返回 200 时, Ret 返回结果详情

`

var docAPITemplate = `
## /{{.StructName}}

### Args

{{range .Args}} 
* {{.Value}}

{{transformDesc .Desc}}
{{end}}

### Opts

{{range .Options}} 
* {{.Value}}

{{transformDesc .Desc}}
{{end}}

### Flags

{{range .Flags}} 
* {{.Value}}

{{transformDesc .Desc}}
{{end}}

`

func init() {
	apiTemplate = strings.Replace(apiTemplate, "$(CMD_TAG)", "`cmd:\"{{.Cmd}}\"`", -1)
	apiTemplate = strings.Replace(apiTemplate, "$(ARG_JSON_TAG)", "`json:\"Args\"`", -1)
	apiTemplate = strings.Replace(apiTemplate, "$(OPT_JSON_TAG)", "`json:\"Opts\"`", -1)
	apiTemplate = strings.Replace(apiTemplate, "$(FLAG_JSON_TAG)", "`json:\"Flags\"`", -1)
	apiTemplate = strings.Replace(apiTemplate, "$(CMD_IGNORE_JSON_TAG)", "`cmd:\"{{.Cmd}}\",json:\"-\"`", -1)
}

// template struct
type CmdTemplate struct {
	StructName string
	Args       []CmdTemplateItem
	Options    []CmdTemplateItem
	Flags      []CmdTemplateItem
	Cmd        string // used for CMD
}

type CmdTemplateItem struct {
	Value string
	Desc  string
	Cmd   string
}

// metadata struct
type Cmd struct {
	Command string `json:"command"`
	Args    []Item `json:"args"`
	Options []Item `json:"options"`
	Flags   []Item `json:"flags"`
}

type Item struct {
	Value string `json:"value"`
	Desc  string `json:"desc"`
}

func read(filename string) ([][]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer file.Close()

	var buffArray [][]byte
	r := bufio.NewReader(file)
	for {
		buff, err := r.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		buffArray = append(buffArray, buff)
	}
	return buffArray, nil
}

func write(filename string, buf []byte) error {
	err := ioutil.WriteFile(filename, buf, 0644)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func normalize(s string) string {
	s = strings.Replace(strings.Replace(s, " ", "_", -1), "-", "_", -1)
	return strings.ToUpper(string(s[0])) + s[1:]
}

func genFile(filename string, data interface{}) error {
	tpl := template.Must(
		template.New("api").Funcs(template.FuncMap{}).Parse(apiTemplate),
	)
	var bf bytes.Buffer
	err := tpl.Execute(&bf, data)
	if err != nil {
		log.Println(err)
		return err
	}

	err = write(filename, bf.Bytes())
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func genTestFile(filename string, data interface{}) error {
	tpl := template.Must(
		template.New("api_test").Funcs(template.FuncMap{}).Parse(apiTestTemplate),
	)
	var bf bytes.Buffer
	err := tpl.Execute(&bf, data)
	if err != nil {
		log.Println(err)
		return err
	}

	err = write(filename, bf.Bytes())
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func genTestUtilFile(filename string, data interface{}) error {
	tpl := template.Must(
		template.New("api_test_util").Funcs(template.FuncMap{}).Parse(apiTestUtilTemplate),
	)
	var bf bytes.Buffer
	err := tpl.Execute(&bf, data)
	if err != nil {
		log.Println(err)
		return err
	}

	err = write(filename, bf.Bytes())
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func genAPIDoc(data interface{}) ([]byte, error) {
	tpl := template.Must(
		template.New("api_doc").Funcs(template.FuncMap{
			"transformDesc": func(raw string) string {
				raw = strings.Replace(raw, "{{.EmphasisLeft}}", "**", -1)
				raw = strings.Replace(raw, "{{.EmphasisRight}}", "**", -1)
				raw = strings.Replace(raw, "{{.LessThan}}", "<", -1)
				raw = strings.Replace(raw, "{{.GreaterThan}}", ">", -1)
				return raw
			},
		}).Parse(docAPITemplate),
	)
	var bf bytes.Buffer
	err := tpl.Execute(&bf, data)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return bf.Bytes(), nil
}
