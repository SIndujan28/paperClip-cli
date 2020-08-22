package parser

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

//Fetch function to scrap values from txt file
func Fetch(path string) []string {
	f, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	var wo []string

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {

		rawLine := strings.Split(scanner.Text(), "#")

		reg, err := regexp.Compile(`\d+(\.\d+)*`)
		if err != nil {
			log.Fatal(err)
		}

		cleanLine := strings.TrimSpace(reg.ReplaceAllString(rawLine[0], ""))
		wo = append(wo, cleanLine)
	}

	fmt.Println(wo[5], len(wo))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return wo
}
