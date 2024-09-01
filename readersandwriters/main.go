package main

import (
	"fmt"
	"io"
	"strings"
)

func scanFromReader(reader io.Reader, template string,
	vals ...interface{}) (int, error) {
	return fmt.Fscanf(reader, template, vals...)
}

func scanSingle(reader io.Reader, val interface{}) (int, error) {
	return fmt.Fscan(reader, val)
}

func main() {
	reader := strings.NewReader("Kayak Watersports $279.00")
	for {
		var str string
		_, err := scanSingle(reader, &str)
		if err != nil {
			if err != io.EOF {
				Printfln("Error: %v", err.Error())
			}
			break
		}
		Printfln("Value: %v", str)
	}
}
