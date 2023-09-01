package main

import "fmt"

func main() {
	// N:=5
	var N int 
	// arr := [5]int{11, 22, 33, 44,55}
	// arr[4]=23
	
    fmt.Println("Enter the size of array: ")
	fmt.Scanln(&N)
	arr := []int{}
	for i := 0; i < N; i++ {
	val:=0
	fmt.Printf("Enter the element number %v: ",i+1)
	fmt.Scanln(&val)
	arr = append(arr, val)
	}
	result:=divisible(arr)
	fmt.Printf("\n%v",result)
}

func divisible(arr []int)  (flag bool) {
	flag=false
	n := len(arr)
	num := 0
	for i := 0; i < n; i++ {
		num *= 10
		digit:=arr[i]
		for digit>=10{
			digit=digit/10
		}
		// fmt.Printf("%v", digit)
		num += digit
	}
	// fmt.Printf("%v", num)
	if num%n==0{
		fmt.Printf("Therefore the number formed is %v which is divisible by %v.",num,n)
		flag=true
		return 
	} else{
		fmt.Printf("Therefore the number formed is %v which is not divisible by %v.",num,n)
		return
	}
}