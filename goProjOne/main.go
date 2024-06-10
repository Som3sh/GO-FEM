package main

import (
	"fmt"
)

// func calculateTax(price float32) (float32, float32) {
// 	return price * 0.09, price * 0.02
// }
// func calculateTaxTwo(price float32) (float32, float32) {
// 	return price * 0.07, price * 0.05
// }

// func init() {
// 	fmt.Println("This is init function")

// }
func birthday(pointerAge *int) {
	if *pointerAge > 50 {
		panic("You are too old")
	}
}
func main() {

	// fmt.Println("fuck off")
	// fmt.Println(data.MaxSpeed)
	// printData()
	// // names := []string{"Mary", "John", "Casper"}
	// // // names := append(names , "Carol")
	// // println(len(names))
	// stateTax, cityTax := calculateTax(100)
	// fmt.Println(stateTax, cityTax)

	// countryTax, _ := calculateTaxTwo(2000)
	// fmt.Println(countryTax)
	defer fmt.Println("Bye")
	age := 23
	birthday(&age)
	fmt.Println(age)

}
