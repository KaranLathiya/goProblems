package main

import (
	"fmt"
	"os"
)

// create main function to execute the program
func main() {
	// Open the file for writing
	file, errs := os.Create("myfile.txt")
	if errs != nil {
		fmt.Println("Failed to create file:", errs)
		return
	}
	defer file.Close()

	// Write the string "Hello, World!" to the file
	_, errs = file.WriteString("Hello, World!")
	if errs != nil {
		fmt.Println("Failed to write to file:", errs) //print the failed message
		return
	}
	fmt.Println("Wrote to file 'myfile.txt'.") //print the success message
}
