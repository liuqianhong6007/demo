package main

import (
	"flag"
	"os"

	"go.uber.org/zap"
)

var (
	gGenerate = flag.Bool("generate", false, "generate rest frame")
	gMetadata = flag.String("metadata", "metadata", "api metadata filename")
	addr      = flag.String("addr", ":8600", "server address")
	workDir   = flag.String("work_dir", ".", "dolt command work dir")
)

func main() {
	flag.Parse()

	if *gGenerate {
		err := Gen(*gMetadata)
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
	gServer.Init(*addr, *workDir, logger)

	// server listen and serve
	gServer.Serve()
}
