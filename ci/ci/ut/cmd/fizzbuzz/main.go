package main

import (
	"fmt"

	"github.com/UTokyo-Tech-Club/curriculum_7_shin_kanda/ci/ut/pkg/fizzbuzz"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	result := fizzbuzz.FizzBuzz(numbers)
	fmt.Printf("FizzBuzz result: %v\n", result)

	sum := fizzbuzz.Add(5, 3)
	fmt.Printf("5 + 3 = %d\n", sum)
}
