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
			cmdTemplate.Args = append(cmdTemplate.Args, [2]string{normalize(arg), arg})
		}
		for _, option := range cmd.Options {
			cmdTemplate.Options = append(cmdTemplate.Options, [2]string{normalize(option), option})
		}
		for _, f := range cmd.Flags {
			cmdTemplate.Flags = append(cmdTemplate.Flags, [2]string{normalize(f), f})
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
	}

	err = genTestUtilFile("dolt_test_util.go", nil)
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

type {{.StructName}}_Args struct{
	{{range .Args}} 
	{{index . 0}} string $(ARG_TAG)
	{{end}}
}

type {{.StructName}}_Opts struct{
	{{range .Options}}
	{{index . 0}} string $(OPT_TAG)
	{{end}}
}

type {{.StructName}}_Flags struct{
	{{range .Flags}}
	{{index . 0}} bool $(FLAG_TAG)
	{{end}}
}

type {{.StructName}} struct{
	Cmd *struct{} $(CMD_TAG)
	Args {{.StructName}}_Args
	Opts {{.StructName}}_Opts
	Flags {{.StructName}}_Flags
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
			"message": "Internal Server error",
		})
		s.logger.Error("Internal Server error",zap.Error(err))
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
			Args: {{.StructName}}_Args{
				{{range .Args}} 
				{{index . 0}}: "",
				{{end}}
			},
			Opts: {{.StructName}}_Opts{
				{{range .Options}} 
				{{index . 0}}: "",
				{{end}}		
			},
			Flags: {{.StructName}}_Flags{
				{{range .Flags}} 
				{{index . 0}}: false,
				{{end}}		
			},
		},
		nil,
		&rsp,
	)
	if err != nil{
		t.Fatal(err)
	}
	t.Log(rsp)
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

`

func init() {
	apiTemplate = strings.Replace(apiTemplate, "$(ARG_TAG)", "`cmd:\"{{index . 1}}\"`", -1)
	apiTemplate = strings.Replace(apiTemplate, "$(OPT_TAG)", "`cmd:\"{{index . 1}}\"`", -1)
	apiTemplate = strings.Replace(apiTemplate, "$(FLAG_TAG)", "`cmd:\"{{index . 1}}\"`", -1)
	apiTemplate = strings.Replace(apiTemplate, "$(CMD_TAG)", "`cmd:\"{{.Cmd}}\",json:\"-\"`", -1)
}

// template struct
type CmdTemplate struct {
	StructName string
	Args       [][2]string
	Options    [][2]string
	Flags      [][2]string
	Cmd        string // used for CMD
}

// metadata struct
type Cmd struct {
	Command string   `json:"command"`
	Args    []string `json:"args"`
	Options []string `json:"options"`
	Flags   []string `json:"flags"`
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

	err = ioutil.WriteFile(filename, bf.Bytes(), 0644)
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

	err = ioutil.WriteFile(filename, bf.Bytes(), 0644)
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

	err = ioutil.WriteFile(filename, bf.Bytes(), 0644)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
