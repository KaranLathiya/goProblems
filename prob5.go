package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strconv"
	"github.com/xuri/excelize/v2"
)

type Customer struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	FullName  string `json:"fullName"`
	Email     string `json:"email"`
	Gender    string `json:"gender"`
	IpAddress string `json:"ipAddress"`
}

func Id() int {
	var id int
	number := -1
	fmt.Printf("Please enter the ID number on you want to perform operation\n")
	fmt.Scanln(&id)
	for i := range Cus {
		if Cus[i].Id == id {
			number = i
			fmt.Printf("Hello, %v\n", Cus[i].FullName)
			break
		}
	}
	if number == -1 {
		fmt.Println("Invalid ID")
		os.Exit(0)
	}
	return number
}
func Read() {
	id := Id()
	fmt.Println(Cus[id])

}
func Update() {
	id := Id()
	fmt.Println(Cus[id])
	Add(id)
}

func Remove() {
	id := Id()
	fmt.Println(Cus[id])
	Cus = append(Cus[:id], Cus[id+1:len(Cus)]...)
	fmt.Printf("Details successfully deleted")
}

func Add(code int) {
	var id int
	if code == -1 {
		id = len(Cus)
	} else {
		id = code + 1
	}
	Cus[id-1].Id = id
	var firstname, lastname, email, gender, ipaddress string
	fmt.Printf("Please enter the first name\n")
	fmt.Scanln(&firstname)
	Cus[id-1].FirstName = firstname
	fmt.Printf("Please enter the last name\n")
	fmt.Scanln(&lastname)
	Cus[id-1].LastName = lastname
	Cus[id-1].FullName = firstname + " " + lastname
	fmt.Printf("Please enter the email\n")
	fmt.Scanln(&email)
	Cus[id-1].Email = email
	fmt.Printf("Please enter the gender\n")
	fmt.Scanln(&gender)
	Cus[id-1].Gender = gender
	fmt.Printf("Please enter the ipaddress\n")
	fmt.Scanln(&ipaddress)
	Cus[id-1].IpAddress = ipaddress
	if code == -1 {
		fmt.Println("New details successfully added.")
	} else {
		fmt.Println("Details successfully updated.")
	}

}
func SortByFullName() {
	// objs.sort((a,b) => a.last_nom - b.last_nom)
	// sort.Strings(Cus[3].FullName)
	sort.Slice(Cus, func(i, j int) bool {
		return Cus[i].FullName < Cus[j].FullName
	})
	fmt.Println("Data is sorted in ascending order of FullName")
	for i := range Cus {
		fmt.Println(Cus[i], "\t")
	}
	sort.Slice(Cus, func(i, j int) bool {
		return Cus[i].Id < Cus[j].Id
	})
}
func updataFileData() {
	data, _ := json.MarshalIndent(Cus, "", "  ")
	dataFile, errs := os.Create("prob5.json")
	if errs != nil {
		fmt.Println("Failed to create file:", errs)
		return
	}
	defer dataFile.Close()

	// Write the string "Hello, World!" to the file
	_, errs = dataFile.Write(data)
	if errs != nil {
		fmt.Println("Failed to write to file:", errs) //print the failed message
		return
	}
}

var Cus []Customer

func main() {
	_, error := os.Stat("prob5.json")
	if error != nil {
		f, err := excelize.OpenFile("prob5.xlsx")
		if err != nil {
			fmt.Println(err)
			return
		}
		// fmt.Println(f)
		rows, err := f.GetRows("in")
		if err != nil {
			fmt.Println(err)
			return
		}
		for i, row := range rows {
			if i == 0 {
				continue
			}
			Cus1 := Customer{}
			Cus = append(Cus, Cus1)
			for j, colCell := range row {

				// fmt.Print(colCell, "\t")
				switch j {
				case 0:
					Cus[i-1].Id, _ = strconv.Atoi(colCell)

				case 1:
					Cus[i-1].FirstName = colCell

				case 2:
					Cus[i-1].LastName = colCell

				case 3:
					Cus[i-1].Email = colCell

				case 4:
					Cus[i-1].Gender = colCell

				case 5:
					Cus[i-1].IpAddress = colCell
					fallthrough

				case 6:
					Cus[i-1].FullName = Cus[i-1].FirstName + " " + Cus[i-1].LastName

				}
			}
		}
	} else {
		fileData, _ := os.ReadFile("prob5.json")
		err := json.Unmarshal(fileData, &Cus)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	var operation string
	fmt.Printf("To Read user deatails 1\n")
	fmt.Printf("To Update user deatails 2\n")
	fmt.Printf("To Delete user deatails 3\n")
	fmt.Printf("To Add user deatails 4\n")
	fmt.Printf("To get all user details in ascending order of fullname 5\n")
	fmt.Printf("To exit please enter 6\n")
	complete := false
	for !complete {
		fmt.Printf("\nPlease enter the number\n")
		fmt.Scanln(&operation)

		switch operation {
		case "1":
			Read()
		case "2":
			Update()
		case "3":
			Remove()
		case "4":
			Add(-1) // -1 for new customer add or id for update
		case "5":
			SortByFullName()
		case "6":
			complete = true
		default:
			fmt.Println("Enter the number in between 1-6")
		}
		updataFileData()
	}

}
