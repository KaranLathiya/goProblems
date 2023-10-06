package main

import "fmt"

func steps(N int) int {
	start := 2
	sum := 3
	temp := 0
	if N < 4 {
		return N
	} else {
		for i := 3; i < N; i++ {
			temp = sum
			sum += start
			start = temp
		}
		return sum
	}
}
func main() {
	var N int
	for {
		fmt.Printf("Please enter the number of staircase steps:")
		fmt.Scanln(&N)
		if 0 < N && N < 51 {
			count := steps(N)
			fmt.Printf("There are %v ways to climb to the top.", count)
			break
		} else {
			fmt.Printf("Number must be in between 1-50\n")
		}
	}
}
