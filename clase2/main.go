package main

import (
	"fmt";
)

func main() {
	var input int
	fmt.Println("Input some number")
	fmt.Scanln(&input)

	even := generateEvenNumbers(input)

	for k, v := range even {
		fmt.Printf("%d : %d \n", k, v)
	}
}

func generateEvenNumbers (input int) []int {
	var even []int
	for i:=0; i <= input; i++ {
		if i % 2 == 0{
			even = append(even, i)
		}
	}
	return even
}
