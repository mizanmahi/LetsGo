package main

import "fmt"

func main() {

	// anonymous function
	func() {
		println("Hello World")
	}() 
	// function with parameter
	// func(a int, b int) {
	// 	println(a + b)
	// }(1, 2) 

	// fmt.Println(super(5))


	prices := []int{10,50,60,45,78}

	taxes := [2]float64{5.0, 7.5}

	priceWithTaxedPrice := map[int]int{}


	for _, price := range prices {
		fmt.Println(price)
		var taxedPrice float64 = float64(price)
		for _, tax := range taxes {
			fmt.Println(tax, taxedPrice)
			taxedPrice += taxedPrice * tax
		}

		fmt.Println(taxedPrice)
		priceWithTaxedPrice[price] = int(taxedPrice)


	}

	fmt.Println(priceWithTaxedPrice)




	
	
	
}


func super(n int) int {
	if n == 1 {
		return 1
	}
	return n * super(n-1)
}

// variadic function
func sum(num ...int) int {
	total := 0
	for _, v := range num {
		total += v
	}
	return total
}

