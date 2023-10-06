package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	//  used when to know particular file exist or not and it returns file information like length ... if it exists
	fileInfo, error := os.Stat("create.go")

	if error != nil {
		log.Fatal(error)
	} else {
		fmt.Println(fileInfo)
	}
}
