package main

import (
	"flag"
	"os"

	"github.com/SIndujan28/paperclip/cmd"
)

func main() {

	textPtr := flag.String("p", "", "path for the book. (Required)")
	flag.Parse()

	if *textPtr == "" {
		flag.PrintDefaults()
		os.Exit(1)
	} else {
		cmd.Run(*textPtr)
	}

}
