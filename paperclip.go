package main

import (
	"flag"
	"os"

	"github.com/SIndujan28/paperclip/cmd"
)

func main() {
	// q := "Eric Elliott - Composing Software-Leanpub.pdf"
	// a := cmd.Run(fmt.Sprintf("/home/sindu/test/%s", q), "success")
	// kafka.Producer(fmt.Sprintf("/home/sindu/test/%s", q), a)
	// kafka.Producer()
	// parser.Test()
	textPtr := flag.String("p", "", "path for the book. (Required)")
	flag.Parse()

	if *textPtr == "" {
		flag.PrintDefaults()
		os.Exit(1)
	} else {
		cmd.Run(*textPtr)
	}

}
