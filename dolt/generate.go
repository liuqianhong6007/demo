package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

func Gen(metadataFile string) error {
	buffArray, err := read(metadataFile)
	if err != nil {
		return err
	}

	if len(buffArray) == 0 {
		return nil
	}

	for _, buff := range buffArray {
		var cmd Cmd
		err = json.Unmarshal(buff, &cmd)
		if err != nil {
			return err
		}
		var temp struct {
			StructName string
			Args       []string
			Options    []string
		}
		temp.StructName = strings.Replace(cmd.Command, " ", "_", -1)
		temp.Args = cmd.Args
		temp.Options = cmd.Options
		err = genFile(temp.StructName+".auto.go", temp)
		if err != nil {
			return err
		}
	}

	return nil
}

func genFile(filename string, data interface{}) error {
	tpl := template.Must(
		template.New("api").Funcs(template.FuncMap{}).Parse(apiTemplate),
	)
	var bf bytes.Buffer
	err := tpl.Execute(&bf, data)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, bf.Bytes(), os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

var apiTemplate = `
package main

import(
	"net/http"
	"os/exec"
	"reflect"

	"github.com/gin-gonic/gin"
)

type {{.StructName}} struct{
	{{range .Args}}
	{{.}} string
	{{end}}
	{{range .Options}}
	{{.}} string
	{{end}}
}

func (s *Server) {{.StructName}}(c *gin.Context){
	var param {{.StructName}}
	err := c.BindJSON(&param)
	if err != nil{
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": "Request param error",
		})
		return
	}

	var args []string
	typeObj := reflect.TypeOf(param)
	valObj := reflect.ValueOf(param)
	for i := 0; i < typeObj.NumField(); i++ {
		args = append(args, valObj.Field(i).String())
	}

	result,err := Execute(exec.Command("dolt", args...))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": "Internal Server error",
		})
		return
	}
	
	c.JSON(http.StatusInternalServerError, map[string]interface{}{
		"code":    http.StatusOK,
		"message": result,
	})
}
`

type Cmd struct {
	Command string   `json:"command"`
	Args    []string `json:"args"`
	Options []string `json:"options"`
}

func read(filename string) ([][]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var buffArray [][]byte
	r := bufio.NewReader(file)
	for {
		buff, _, c := r.ReadLine()
		if c == io.EOF {
			break
		}
		buffArray = append(buffArray, buff)
	}
	return buffArray, nil
}
