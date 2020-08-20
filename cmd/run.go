package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"

	"github.com/SIndujan28/paperclip/parser"
)

//Run fn executes mutool cmd
func Run(bpath string) []string {
	app := "mutool"
	arg0 := "show"
	arg1 := bpath
	arg2 := "outline"
	// Open the output file
	outfile, err := os.Create(fmt.Sprintf("%s.txt", path.Base(bpath)))
	if err != nil {
		log.Fatal(err)
	}
	defer outfile.Close()
	cmd := exec.Command(app, arg0, arg1, arg2)
	cmd.Stdout = outfile
	if err = cmd.Start(); err != nil {
		log.Fatal(err)
	}

	log.Print("Waiting for command to finish...")

	// Wait for the command to finish.
	if err = cmd.Wait(); err != nil {
		log.Fatal(err)
	}
	s := parser.Fetch(fmt.Sprintf("%s.txt", f))
	return s
}
