package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
	// "strconv"
)

var pic []Photo
var newpic []Photo

type Photo struct {
	PhotoID   int    `json:"photoID"`
	PhotoName string `json:"photoName"`
	PhotoSize string `json:"photoSize"`
}

func main() {
	fileData, error := os.ReadFile("prob8.json")
	if error != nil {
		fmt.Printf("%v", error)
		return
	}
	// fmt.Print(string(fileData))
	err := json.Unmarshal(fileData, &pic)
	if err != nil {
		fmt.Println(err)
		return
	}
	ch := make(chan Photo)
	updated_ch := make(chan Photo)
	for i := 0; i < len(pic); i++ {

		go readData(ch, i)
		go updateData(ch, i, updated_ch)
		go storeData(updated_ch, i)

		fmt.Println("Waiting for goroutines to finish...")
		fmt.Println("Done!")

		time.Sleep(1 * time.Microsecond)
	}
	close(ch)
	close(updated_ch)
}

func readData(ch chan Photo, index int) {
	fmt.Println("Reading the data of ", index+1)
	ch <- pic[index]
}

func updateData(ch chan Photo, index int, updated_ch chan Photo) {
	fmt.Println("updating the data of ", index+1)
	createdPic := <-ch
	createdPic.PhotoSize = "2.8MB"
	updated_ch <- createdPic
}

func storeData(updated_ch chan Photo, index int) {
	fmt.Println("Storing the data of ", index+1)
	updatedPic := <-updated_ch
	newpic = append(newpic, updatedPic)
	updated_data, _ := json.MarshalIndent(newpic, "", "  ")
	// fmt.Print(string(newval))
	newFile, errs := os.Create("updated_prob8.json")
	if errs != nil {
		fmt.Println("Failed to create file:", errs)
		return
	}
	defer newFile.Close()

	// Write the string "Hello, World!" to the file
	_, errs = newFile.Write(updated_data)
	if errs != nil {
		fmt.Println("Failed to write to file:", errs) //print the failed message
		return
	}
}
