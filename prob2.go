package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var Customer []User
var id int

func main() {
	fileData, error := os.ReadFile("prob2.json")
	if error != nil {
		fmt.Printf("%v", error)
		return
	}
	// fmt.Print(string(fileData))
	err := json.Unmarshal(fileData, &Customer)
	if err != nil {
		fmt.Println(err)
		return
	}
	var operation string

	var idFound bool
	fmt.Printf("Please enter you ID number\n")
	fmt.Scanln(&id)
	for i := range Customer {
		if Customer[i].Userid == id {
			idFound = true
			id = i
			fmt.Printf("Hello, %v\n", Customer[id].Name)
			break
		}
	}
	if !idFound {
		fmt.Println("Invalid ID")
		return
	}
	fmt.Printf("To perform Check balance operation please enter 1\n")
	fmt.Printf("To perform Withdraw operation please enter 2\n")
	fmt.Printf("To perform Deposit operation please enter 3\n")
	fmt.Printf("To exit please enter 4\n")
	complete := true
	for complete {
		fmt.Printf("Please enter the number\n")
		fmt.Scanln(&operation)

		switch operation {
		case "1":
			balance()
		case "2":
			withdraw()
		case "3":
			deposit()
		case "4":
			complete = false
		default:
			fmt.Println("Enter the number in between 1-4")
		}
		updataFileData()
	}

	fmt.Printf("Thank you, %v", Customer[id].Name)
}

type User struct {
	Userid  int     `json:"userId"`
	Name    string  `json:"name"`
	Balance float64 `json:"balance"`
}

func balance() {
	bal := Customer[id].Balance
	fmt.Printf("Your Balance is %v\n", bal)
}

func withdraw() {
	fmt.Println("Please the enter the amount of withdraw")
	var withdraw float64
	fmt.Scanln(&withdraw)
	bal := Customer[id].Balance
	if withdraw > bal {
		fmt.Println("Your balance is insufficient")
	} else if withdraw < 0 {
		fmt.Println("Withdraw amount must be greater than 0")
	} else {
		Customer[id].Balance -= withdraw
		fmt.Printf("Withdrawn amount is %v\n Now, Your balance is %v\n", withdraw, Customer[id].Balance)
	}
}
func deposit() {
	fmt.Println("Please the enter the amount of deposit")
	var deposit float64
	fmt.Scanln(&deposit)
	if deposit > 0 {
		Customer[id].Balance += deposit
		fmt.Printf("Deposited amount is %v\n Now, Your balance is %v\n", deposit, Customer[id].Balance)
	} else {
		fmt.Println("Deposit amount must be greater than 0")
	}
}
func updataFileData() {
	// fmt.Println(Customer)
	updated_data, _ := json.MarshalIndent(Customer, "", "  ")
	// fmt.Print(string(newval))
	newFile, errs := os.Create("prob2.json")
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
