package main

import (
	"flag"
)

var (
	gGenerate = flag.Bool("generate", false, "generate rest frame")
	gMetadata = flag.String("metadata", "", "api metadata filename")
)

func main() {
	flag.Parse()

	if *gGenerate {
		err := Gen(*gMetadata)
		if err != nil {
			panic(err)
		}
	}
}
