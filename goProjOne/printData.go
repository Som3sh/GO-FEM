package main

import "fmt"

var name = "THis is Somesh"

func printData() {

	slice := []float32{1.1, 2.2, 3.3, 4.4, 5.5}
	slice = append(slice, 6.6)
	q := len(slice)
	fmt.Println(slice)
	println("Hello, World!")
	println("This is a new line")
	println(name)
	println(q)
	println("hell yeah!")
}
