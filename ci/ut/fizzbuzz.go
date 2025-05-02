package main

import (
	"strconv"
)

const fizz, buzz = "Fizz", "Buzz"
const fizzbuzz = fizz + buzz

func FizzBuzz(numbers []int) []string {
	result := make([]string, 0, len(numbers))

	for _, n := range numbers {
		switch {
		case n > 0:
			switch {
			case isFizzBuzz(n):
				result = append(result, fizzbuzz)
			case isFizz(n):
				result = append(result, fizz)
			case isBuzz(n):
				result = append(result, buzz)
			default:
				result = appendNum(result, n)
			}
		case n <= 0:
			result = appendNum(result, n)
		}
	}

	return result
}

func isFizz(n int) bool                    { return n%3 == 0 }
func isBuzz(n int) bool                    { return n%5 == 0 }
func appendNum(s []string, n int) []string { return append(s, strconv.Itoa(n)) }
func isFizzBuzz(n int) bool                { return n%3 == 0 && n%5 == 0 }
