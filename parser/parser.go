package parser

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
)

//Test export function
func Test() {
	f, err := os.Open("toc.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	w := bufio.NewScanner(f)
	w.Split(func(data []byte, atEOF bool) (int, []byte, error) {
		trimspace := func(b []byte) []byte {
			return bytes.TrimSpace(b)
		}
		if len(data) == 0 {
			return 0, nil, nil
		}
		for i := 0; i < len(data); i++ {
			if data[i] == '#' {
				return i + 1, trimspace(data[:i]), nil
			}
		}
		if atEOF {
			return len(data), trimspace(data), nil
		}
		return 0, nil, nil
	})
	var wo []string
	for w.Scan() {
		wo = append(wo, w.Text())
	}
	fmt.Println(wo)
}
