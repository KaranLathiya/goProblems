package main

import (
	"encoding/json"
	"fmt"
	"os"
	"github.com/xuri/excelize/v2"
	"strconv"
)

type Customer struct{
	Id 				int 		`json:"id"`
	FirstName   	string 		`json:"firstName"`
	LastName 	    string 		`json:"lastName"`
	FullName   		string 		`json:"fullName"`
	Email    		string 		`json:"email"`
	Gender    		string 		`json:"gender"`
	IpAddress       string 		`json:"ipAddress"`
}
func Id(Cus []Customer) int{
	var id int
	number:=-1
	fmt.Printf("Please enter the ID number on you want to perform operation\n")
	fmt.Scanln(&id)
	for i := range Cus {
		if Cus[i].Id == id {
			number=i
			fmt.Printf("Hello, %v\n",Cus[i].FullName)
			break
		}	
	}
	if number==-1{
		fmt.Println("Invalid ID")
		return 0
	}
	return number
}
func Read(Cus []Customer){
	number:=Id(Cus)
	fmt.Print(Cus[number])
}

func main() {
	var Cus []Customer
	f, err := excelize.OpenFile("data.xlsx")
	 if err != nil {
        fmt.Println(err)
        return
    }
	fmt.Print(f)
	rows, err := f.GetRows("in")
    if err != nil {
        fmt.Println(err)
        return
    }
    for i, row := range rows {
		if i==0{
				continue
			}
			Cus1:=Customer{}
			Cus = append(Cus,Cus1)
        for j, colCell := range row {
			
            fmt.Print(colCell, "\t")
			switch j{
			case 0:
				Cus[i-1].Id,_ =strconv.Atoi(colCell)
		
			case 1:
				Cus[i-1].FirstName=colCell

			case 2:
				Cus[i-1].LastName =colCell

			case 3:
				Cus[i-1].Email=colCell

			case 4:
				Cus[i-1].Gender =colCell

			case 5:
				Cus[i-1].IpAddress =colCell
				fallthrough
			
			case 6:
				Cus[i-1].FullName =Cus[i-1].FirstName+" "+Cus[i-1].LastName 
			
			}
        }
        fmt.Println()
}
newval, err := json.MarshalIndent(Cus, "", "  ")
filee, errs := os.Create("dataofexcel.json")
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

   var casee string
	// array := []int{}
	
	fmt.Printf("To Read user deatails 1\n")
	fmt.Printf("To Update user deatails 2\n")
	fmt.Printf("To Delete user deatails 3\n")
	fmt.Printf("To Add user deatails 4\n")
	fmt.Printf("To get all user details in ascending order of fullname 3\n")
	fmt.Printf("To exit please enter 6\n")
	flag:=true
	for flag {
		fmt.Printf("Please enter the number\n")
		fmt.Scanln(&casee)
		
		switch casee {
		case "1":
			Read(Cus[:])
		
		case "4":
			flag=false
		default:
			fmt.Println("Enter the number in between 1-4")
		}
	}
}
