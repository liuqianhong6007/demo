package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
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
			cmdArg := arg
			arg = normalize(arg)
			cmdTemplate.Args = append(cmdTemplate.Args, [2]string{strings.ToUpper(string(arg[0])) + arg[1:], cmdArg})
		}
		for _, option := range cmd.Options {
			cmdOption := option
			option = normalize(option)
			cmdTemplate.Options = append(cmdTemplate.Options, [2]string{strings.ToUpper(string(option[0])) + option[1:], cmdOption})
		}
		for _, f := range cmd.Flags {
			cmdFlags := f
			f = normalize(f)
			cmdTemplate.Flags = append(cmdTemplate.Flags, [2]string{strings.ToUpper(string(f[0])) + f[1:], cmdFlags})
		}

		err = genFile(cmdTemplate.StructName+".auto.go", cmdTemplate)
		if err != nil {
			log.Println(err)
			return err
		}
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

	"github.com/gin-gonic/gin"
)

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
		return
	}

	var args []string
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

	// get command string
	cmdField, _ := reflect.TypeOf(param).FieldByName("Cmd")
	commandStr := cmdField.Tag.Get("cmd")

	// call dolt command
	result,err := Execute(exec.Command(commandStr, args...))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": "Internal Server error",
		})
		return
	}
	
	c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": result,
	})
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
	return strings.Replace(strings.Replace(s, " ", "_", -1), "-", "_", -1)
}

func genFile(filename string, data interface{}) error {
	tpl := template.Must(
		template.New("api").Funcs(template.FuncMap{}).Parse(apiTemplate),
	)
	fmt.Println(apiTemplate)
	var bf bytes.Buffer
	err := tpl.Execute(&bf, data)
	if err != nil {
		log.Println(err)
		return err
	}

	err = ioutil.WriteFile(filename, bf.Bytes(), os.ModePerm)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
