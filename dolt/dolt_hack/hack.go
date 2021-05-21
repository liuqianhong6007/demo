package dolt_hack

import (
	"encoding/json"
	"log"
	"os"

	"github.com/dolthub/dolt/go/cmd/dolt/cli"
	"github.com/dolthub/dolt/go/libraries/utils/argparser"
)

var apiJsonFile *os.File

func init() {
	var err error
	if apiJsonFile, err = os.OpenFile("metadata", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644); err != nil {
		panic(err)
	}
}

type Item struct {
	Value string `json:"value"`
	Desc  string `json:"desc"`
}

func dump(file *os.File, cmdDoc cli.CommandDocumentation) {
	var args []Item
	for _, argListHelp := range cmdDoc.ArgParser.ArgListHelp {
		args = append(args, Item{
			Value: argListHelp[0],
			Desc:  argListHelp[1],
		})
	}

	var options, flags []Item
	for _, opt := range cmdDoc.ArgParser.Supported {
		ietm := Item{
			Value: opt.Name,
			Desc:  opt.Desc,
		}
		if opt.OptType == argparser.OptionalFlag {
			flags = append(flags, ietm)
		} else {
			options = append(options, ietm)
		}
	}

	buf, err := json.Marshal(map[string]interface{}{
		"command": cmdDoc.CommandStr,
		"args":    args,
		"options": options,
		"flags":   flags,
	})
	if err != nil {
		log.Println(err)
		return
	}

	if _, err = file.Write(append(buf, '\n')); err != nil {
		log.Println(err)
		return
	}
}
