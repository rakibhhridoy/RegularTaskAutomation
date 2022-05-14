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
	code := `package main

import (
	"fmt"
	"log"
)

func main(){
	fmt.Println("")
}
`
	fileName := os.Args[1]
	isextention := strings.HasSuffix(fileName, ".go")
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
		fmt.Println("Golang File should have extension of .go")
	}

}
