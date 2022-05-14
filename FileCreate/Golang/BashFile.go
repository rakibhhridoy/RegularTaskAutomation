package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	FileCreate()
}

func FileCreate() {
	code := `#!/bin/bash

echo ""
`
	fileName := os.Args[1]
	isextention := strings.HasSuffix(fileName, ".sh")
	if isextention {
		file, err := os.Create(fileName)

		if err != nil {
			log.Fatal(err)
		}

		defer file.Close()

		_, err2 := file.WriteString(code)

		if err2 != nil {
			log.Fatal(err2)
		}

		fmt.Println("Done!")

	} else {
		fmt.Println("Bash File should have extension of .sh")
	}

}
