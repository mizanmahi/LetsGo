package main

import (
	"fmt"
)

func main () {
	// names := [3]string{"a", "b", "c"}
	// fmt.Println(names)
	// fmt.Println(names[2])

	// selectedNames := names[0:len(names)] // slice, second index is not included
	// fmt.Println(selectedNames)

	// // map
	// websites := map[string]string{
	// 	"google": "https://google.com",
	// 	"facebook": "https://facebook.com",
	// }
	// fmt.Println(websites)
	// fmt.Println(websites["google"])

	// // add new key-value pair
	// websites["twitter"] = "https://twitter.com"
	// fmt.Println(websites)

	// // delete key-value pair
	// delete(websites, "google")

	// // update value
	// websites["facebook"] = "https://facebook.com/updated"



	// example of make() function
	// make(map[key-type]value-type, initial size)
	websites2 := make(map[string]string, 10) // max 10 key-value pairs can be stored
	websites2["google"] = "https://google.com"
	websites2["facebook"] = "https://facebook.com"
	fmt.Println(websites2)

	// example of make() function with slice
	names := make([]string, 3, 10)
	names[0] = "a"
	names[1] = "b"
	names[2] = "c"
	names = append(names, "d")
	fmt.Println(names)

	// for loop on array
	for i := range names {
		fmt.Println(names[i])
	}

	// for loop on slice
	for i := range names {
		fmt.Println(names[i])
	}

	// for loop on map
	for key, value := range websites2 {
		fmt.Println(key, value)
	}


}