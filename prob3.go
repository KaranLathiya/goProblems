package main

import "fmt"

func push(stack []string) []string {
	top := len(stack) - 1
	var topVal string
	fmt.Printf("Please enter the topValue you want to push:")
	fmt.Scanln(&topVal)
	stack[top] = topVal
	return append(stack[:top])
}
func pop(stack []string) []string {
	top := len(stack)
	if top > 0 {
		fmt.Printf("Popped out topValue is %v \n", (stack)[top-1])
		return append(stack[:top-1])
	} else {
		fmt.Print("Stack is already empty.\n")
		return stack
	}
}
func peek(stack []string) {
	top := len(stack)
	if top > 0 {
		fmt.Printf("Peek topValue is %v \n", (stack)[top-1])
	} else {
		fmt.Print("Stack is already empty\n.")
	}
}

func main() {
	var operation string
	stack := []string{}
	fmt.Printf("To perform PUSH operation please enter 1\n")
	fmt.Printf("To perform POP operation please enter 2\n")
	fmt.Printf("To perform PEEK operation please enter 3\n")
	fmt.Printf("To perform PRINT operation please enter 4\n")
	fmt.Printf("To exit please enter 5\n")

	complete := false
	for !complete {
		fmt.Printf("Please enter the number\n")
		fmt.Scanln(&operation)
		switch operation {
		case "1":
			stack = append(stack, "0")
			push(stack)
		case "2":
			stack = pop(stack)
		case "3":
			peek(stack)
		case "4":
			fmt.Print(stack)
		case "5":
			complete = true
		default:
			fmt.Println("Enter the number in between 1-5")
		}
	}
}
