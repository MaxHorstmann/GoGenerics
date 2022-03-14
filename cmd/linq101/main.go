package main

import (
	"fmt"
)

type collection[T comparable] []T

func (input collection[T]) where(pred func(T) bool) []T {
	result := []T{}
	for _, j := range input {
		if pred(j) {
			result = append(result, j)
		}
	}
	return result
}

func linq1() {
	numbers := collection[int]{5, 4, 1, 3, 9, 8, 6, 7, 2, 0}
	lowNums := numbers.where(func(i int) bool { return i < 5 })
	fmt.Println("Numbers < 5:")
	fmt.Println(lowNums)
}

type product struct {
	ProductName  string
	UnitsInStock int
}

func linq2() {
	products := collection[product]{
		{ProductName: "milk", UnitsInStock: 0},
		{ProductName: "water", UnitsInStock: 4},
		{ProductName: "beer", UnitsInStock: 0}}
	soldOutProducts := products.where(func(p product) bool { return p.UnitsInStock == 0 })
	fmt.Println("Sold out products:")
	for _, p := range soldOutProducts {
		fmt.Printf("%s is sold out!\n", p.ProductName)
	}
}

func main() {
	linq2()
}
