package main

import "fmt"
func push(array *[]int){
	val:=0
	fmt.Printf("Please enter the value you want to push:")
	fmt.Scanln(&val)
	*array =append(*array,val)
}
func pop(array *[]int){
	length:=len(*array)
	if length>0{
	fmt.Printf("Popped out value is %v \n",(*array)[length-1])
	*array =(*array)[:length-1]
	} else {
		fmt.Print("Stack is already empty.\n")
	}
}
func peek(array *[]int){
	length:=len(*array)
	if length>0{
		fmt.Printf("Peek value is %v \n",(*array)[length-1])
	} else {
		fmt.Print("Stack is already empty\n.")
	}
}

func main(){
	var casee string
	array := []int{}
	fmt.Printf("To perform PUSH operation please enter 1\n")
	fmt.Printf("To perform POP operation please enter 2\n")
	fmt.Printf("To perform PEEK operation please enter 3\n")
	fmt.Printf("To perform PRINT operation please enter 4\n")
	fmt.Printf("To exit please enter 5\n")
	
	flag:=true
	for flag{
		fmt.Printf("Please enter the number\n")
		fmt.Scanln(&casee)
		switch casee{
		case "1":
			push(&array)
		case "2":
			pop(&array)
		case "3":
			peek(&array)	
		case "4":
			fmt.Print(array)
		case "5":
			flag=false
		default:
			fmt.Println("Enter the number in between 1-5")
	}
}
}