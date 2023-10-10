package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"
	// "strconv"
)

var pic []Photo
var photoId int

type Photo struct {
	PhotoID   int    `json:"photoID"`
	PhotoName string `json:"photoName"`
	PhotoSize string `json:"photoSize"`
}

func main() {
	var wg sync.WaitGroup
	// i:=1
	// photoname:="image1.jpg"
	for {
		wg.Add(1)
		go readData(&wg)
		wg.Wait()
		// time.Sleep(30 * time.Second)
		pic1 := Photo{PhotoID: 1111, PhotoName: "photoname", PhotoSize: "1.8MB"}
		pic = append(pic, pic1)
		wg.Add(1)

		go updateData(&wg)

		fmt.Println("Waiting for goroutines to finish...")
		wg.Wait()
		fmt.Println("Done!")
		time.Sleep(30 * time.Second)
		// i+=1
		// string_i := strconv.Itoa(i)
		// photoname="image"+string_i+".jpg"

	}
}

func readData(wg *sync.WaitGroup) {
	defer wg.Done()

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
	fmt.Print(pic)
}

func updateData(wg *sync.WaitGroup) {
	defer wg.Done()

	updated_data, _ := json.MarshalIndent(pic, "", "  ")
	// fmt.Print(string(newval))
	newFile, errs := os.Create("prob8.json")
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
