package main

import (
	"flag"
	"os"

	"go.uber.org/zap"

	"github.com/liuqianhong6007/dolt/generate"
	"github.com/liuqianhong6007/dolt/server"
)

var (
	gGenerate = flag.Bool("generate", false, "generate rest frame")
	genOutDir = flag.String("gen_out_dir", ".", "generate out dir")
	gMetadata = flag.String("metadata", "metadata", "api metadata filename")
	addr      = flag.String("addr", ":8600", "server address")
	wd        = flag.String("wd", ".", "dolt command work dir")
)

func main() {
	flag.Parse()

	if *gGenerate {
		err := generate.Gen(*gMetadata, *genOutDir)
		if err != nil {
			os.Exit(1)
		}
		return
	}

	// init zap logger
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	// init server
	server.Init(*addr, *wd, logger)

	// server listen and serve
	server.Serve()
}
