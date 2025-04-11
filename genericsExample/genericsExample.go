package main

import (
	"fmt"
	"reflect"
	"strings"
)

func printType[T int | string | bool](data T) {
	fmt.Println(reflect.TypeOf(data))
}


func Filter[T any](items []T, callback func(T) bool) []T {
	newItems := []T{}
	for _, item := range items {
		if callback(item) {
			newItems = append(newItems, item)
		}
		fmt.Println(item)
	}

	return newItems
}



func main() {
	printType(10) // func printType(data int)
	printType("Hello") // func printType(data string)
	printType(true) // func printType(data bool)


	fmt.Println("Initialized 🚀")

	numbers := []int{1, 2, 3, 4, 5, 6}
	fruits := []string{"apple", "banana", "cherry", "blue berry"}

	filteredNumbers := Filter(numbers, func(item int) bool {
		 return  item % 2 == 0
	})

	filteredStrings := Filter(fruits, func(item string) bool {
		 return strings.HasPrefix(item, "b") // string that starts with 'b'
	})

	fmt.Println(filteredNumbers)
	fmt.Println(filteredStrings)
}
