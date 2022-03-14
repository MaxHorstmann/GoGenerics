package main

import (
	"fmt"
)

func where[T comparable](input []T, pred func(T) bool) []T {
	result := []T{}
	for _, j := range input {
		if pred(j) {
			result = append(result, j)
		}
	}
	return result
}

func main() {
	numbers := []int{5, 4, 1, 3, 9, 8, 6, 7, 2, 0}
	lowNums := where(numbers, func(i int) bool { return i < 5 })
	fmt.Println("Numbers < 5:")
	fmt.Println(lowNums)
}
