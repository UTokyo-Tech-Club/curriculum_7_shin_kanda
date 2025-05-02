package main

import (
	"fmt"
)

func Add(a, b int) int {
	return a + b
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	result := FizzBuzz(numbers)
	fmt.Printf("FizzBuzz result: %v\n", result)
}
