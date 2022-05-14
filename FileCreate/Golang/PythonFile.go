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
	code := `#!/bin/env python

def main():
	print("")

if __name__ == "__main__":
	main()
`
	fileName := os.Args[1]
	isextention := strings.HasSuffix(fileName, ".py")
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
		fmt.Println("Python File should have extension of .py")
	}

}
