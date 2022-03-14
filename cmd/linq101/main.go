package main

import (
	"fmt"
)

type collection[T comparable] []T

func (input collection[T]) where(pred func(T) bool) collection[T] {
	result := collection[T]{}
	for _, j := range input {
		if pred(j) {
			result = append(result, j)
		}
	}
	return result
}

func select2[T comparable, U comparable](input collection[T], transform func(T) U) collection[U] {
	result := collection[U]{}
	for _, j := range input {
		result = append(result, transform(j))
	}
	return result
}

type collectionPlusSecondType[T comparable, U comparable] []T

func (input collectionPlusSecondType[T, U]) select3(transform func(T) U) collection[U] {
	result := collection[U]{}
	for _, j := range input {
		result = append(result, transform(j))
	}
	return result
}

type product struct {
	ProductName  string
	UnitsInStock int
}

func linq1() {
	numbers := collection[int]{5, 4, 1, 3, 9, 8, 6, 7, 2, 0}
	lowNums := numbers.where(func(i int) bool { return i < 5 })
	fmt.Println("Numbers < 5:")
	fmt.Println(lowNums)
}

func linq2() {
	products := collection[product]{
		{ProductName: "milk", UnitsInStock: 0},
		{ProductName: "water", UnitsInStock: 4},
		{ProductName: "beer", UnitsInStock: 0}}
	soldOutProducts := collectionPlusSecondType[product, string](products.where(func(p product) bool { return p.UnitsInStock == 0 })).
		select3(func(p product) string { return p.ProductName })
	fmt.Println("Sold out products:")
	fmt.Println(soldOutProducts)
}

func main() {
	linq2()
}
