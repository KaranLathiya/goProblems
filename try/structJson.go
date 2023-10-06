package main

import (
	"fmt"
)

func main() {

	type Customer struct {
		Id        int    `json:"id"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		FullName  string `json:"fullName"`
		Email     string `json:"email"`
		Gender    string `json:"gender"`
		IpAddress string `json:"ipAddress"`
	}

	// var Cus []Customer

	Cus := []Customer{}
	// Cus[0].Id=121
	// fmt.Print(Cus[0].Id)

	// To Perform Above Operation
	cus1 := Customer{}
	Cus = append(Cus, cus1)

	Cus[0].Id = 121
	fmt.Print(Cus[0].Id)
}
