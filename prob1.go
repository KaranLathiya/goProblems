package main

import "fmt"

func main() {
	var arraySize int
	// arr := [5]int{11, 22, 33, 44,55}
	fmt.Println("Enter the size of array: ")
	_, err := fmt.Scanln(&arraySize)
	if err != nil {
		fmt.Println(err)
		return
	}
	arr := []int{}
	for i := 0; i < arraySize; i++ {
		val := 0
		fmt.Printf("Enter the element number %v: ", i+1)
		_, err := fmt.Scanln(&val)
		if err != nil {
			fmt.Println(err)
			return
		}
		arr = append(arr, val)
	}
	divisible(arr)
}

func divisible(arr []int) {
	arraySize := len(arr)
	num := 0
	for i := 0; i < arraySize; i++ {
		num *= 10
		element := arr[i]
		for element >= 10 {
			element = element / 10
		}
		num += element
	}
	if num%arraySize == 0 {
		fmt.Printf("Therefore the number formed is %v which is divisible by %v.", num, arraySize)
	} else {
		fmt.Printf("Therefore the number formed is %v which is not divisible by %v.", num, arraySize)
	}
	return
}
