package main

import (
	"fmt"
	"reflect"
)

func printType[T any](data T) {
	fmt.Println(reflect.TypeOf(data))
}

func main() {
	printType(10) // func printType(data int)
	printType("Hello") // func printType(data string)
	printType(true) // func printType(data bool)
}
