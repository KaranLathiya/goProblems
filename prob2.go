package main

import (
	"encoding/json"
	"fmt"
	"os"
	
)

type User struct {
	Userid  int	`json:"userId"`
	Name    string `json:"name"`
	Balance float64 `json:"balance"`
}
func balance(id int, Cus1 []User) {
			bal := Cus1[id].Balance
			fmt.Printf("Your Balance is %v\n", bal)
		}
	
func withdraw(id int, Cus1 []User) {
	fmt.Println("Please the enter the amount of withdraw")
	var withdraw float64 
	fmt.Scanln(&withdraw)
	bal := Cus1[id].Balance
	if withdraw>bal{
		fmt.Println("Your balance is insufficient")
	} else if withdraw<0{
		fmt.Println("Withdraw amount must be greater than 0")
	} else {
		Cus1[id].Balance-=withdraw
		fmt.Printf("Withdrawn amount is %v\n Now, Your balance is %v\n",withdraw,Cus1[id].Balance)
	}
}
func deposit(id int, Cus1 []User) {
	fmt.Println("Please the enter the amount of deposit")
	var deposit float64
	fmt.Scanln(&deposit)
	if deposit>0{
		Cus1[id].Balance+=deposit
		fmt.Printf("Deposited amount is %v\n Now, Your balance is %v\n",deposit,Cus1[id].Balance)
		} else{
			fmt.Println("Deposit amount must be greater than 0")
		}
}

func main() {
	file, _ := os.ReadFile("prob2.json")
	// fmt.Print(string(file))
	var Cus1 []User
	err := json.Unmarshal(file, &Cus1)
	if err != nil {
		fmt.Println(err)
	} else{
	// fmt.Println(Cus1)
	}
	var casee string
	var id,number int
	// array := []int{}
	fmt.Printf("Please enter you ID number\n")
	fmt.Scanln(&id)
	for i := range Cus1 {
		if Cus1[i].Userid == id {
			number=i
			fmt.Printf("Hello, %v\n",Cus1[i].Name)
			break
		}	
	}
	if number==0{
		fmt.Println("Invalid ID")
		return
	}
	fmt.Printf("To perform Check balance operation please enter 1\n")
	fmt.Printf("To perform Withdraw operation please enter 2\n")
	fmt.Printf("To perform Deposit operation please enter 3\n")
	fmt.Printf("To exit please enter 4\n")
	flag:=true
	for flag {
		fmt.Printf("Please enter the number\n")
		fmt.Scanln(&casee)
		
		switch casee {
		case "1":
			balance(number, Cus1)
		case "2":
			withdraw(number, Cus1)
		case "3":
			deposit(number, Cus1)
		case "4":
			flag=false
		default:
			fmt.Println("Enter the number in between 1-4")
		}
	}
	// fmt.Println(Cus1)
	newval, err := json.MarshalIndent(Cus1, "", "  ")
	// fmt.Print(string(newval))
	filee, errs := os.Create("prob2.json")
   if errs != nil {
      fmt.Println("Failed to create file:", errs)
      return
   }
   defer filee.Close()

   // Write the string "Hello, World!" to the file
   _, errs = filee.Write(newval)
   if errs != nil {
      fmt.Println("Failed to write to file:", errs) //print the failed message
      return
   }
   fmt.Printf("Thank you, %v",Cus1[number].Name)
}
