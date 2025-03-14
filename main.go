package main

import (
	"fmt"
)

func main () {
	names := [3]string{"a", "b", "c"}
	fmt.Println(names)
	fmt.Println(names[2])

	selectedNames := names[0:len(names)] // slice, second index is not included
	fmt.Println(selectedNames)

	// map
	websites := map[string]string{
		"google": "https://google.com",
		"facebook": "https://facebook.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["google"])

	// add new key-value pair
	websites["twitter"] = "https://twitter.com"
	fmt.Println(websites)

	// delete key-value pair
	delete(websites, "google")

	// update value
	websites["facebook"] = "https://facebook.com/updated"
}