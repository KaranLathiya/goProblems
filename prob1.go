package main

import "fmt"

func main() {
	var arraySize int
	// arr := [5]int{11, 22, 33, 44,55}

	fmt.Println("Enter the size of array: ")
	fmt.Scanln(&arraySize)
	arr := []int{}
	for i := 0; i < arraySize; i++ {
		val := 0
		fmt.Printf("Enter the element number %v: ", i+1)
		fmt.Scanln(&val)
		arr = append(arr, val)
	}
}

func divisible(arr []int) {
	n := len(arr)
	num := 0
	for i := 0; i < n; i++ {
		num *= 10
		digit := arr[i]
		for digit >= 10 {
			digit = digit / 10
		}
		num += digit
	}
	if num%n == 0 {
		fmt.Printf("Therefore the number formed is %v which is divisible by %v.", num, n)
	} else {
		fmt.Printf("Therefore the number formed is %v which is not divisible by %v.", num, n)
	}
	return
}
