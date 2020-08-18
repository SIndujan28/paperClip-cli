package parser

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

//Test export function
func Test() {
	w := bufio.NewScanner(os.Stdin)
	w.Split(func(data []byte, atEOF bool) (int, []byte, error) {
		trimspace := func(b []byte) []byte {
			return bytes.TrimSpace(b)
		}
		if len(data) == 0 {
			return 0, nil, nil
		}
		for i := 0; i < len(data); i++ {
			if isNumeric(data[i]) {
				return i + 1, data[:i], nil
			}
		}
		if atEOF {
			return len(data), trimspace(data), nil
		}
		return 0, nil, nil
	})
	for w.Scan() {
		fmt.Println(w.Text())
	}
}

func isNumeric(s byte) bool {

	return true
}
