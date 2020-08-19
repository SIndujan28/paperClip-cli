package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

//Run fn executes mutool cmd
func Run(path string, f string) {
	app := "mutool"
	arg0 := "show"
	arg1 := path
	arg2 := "outline"
	// Open the output file
	outfile, err := os.Create(fmt.Sprintf("%s.txt", f))
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
}
