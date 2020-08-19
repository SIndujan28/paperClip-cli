package main

import (
	"fmt"

	"github.com/SIndujan28/paperclip/parser"
)

func main() {
	filename := "toc.txt"
	s := parser.Fetch(filename)
	fmt.Println(s)
	// kafka.Producer()
	// parser.Test()
}
