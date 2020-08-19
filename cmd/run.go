package main

import (
	"fmt"
	"os/exec"
)

func main() {
	app := "mutool"

	arg0 := "info"

	cmd := exec.Command(app, arg0, "../toc.txt")

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + string(output))
		return
	}
	fmt.Println(string(output))
}
