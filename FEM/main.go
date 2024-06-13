package main

import (
	"fmt"
	// "time"

	"learninggo.com/go/FEM/data"
)

func main() {

	// var som data.Instructor
	// or
	som := data.Instructor{Id: 3, LastName: "Mohanty", Score: 100}
	som.FirstName = "Somesh"

	// kyle := data.NewInstructor("Kyle", "Simpson")
	// fmt.Printf("%v\n", goCourse)
	// print(som.Print())
	goCourse := data.Course{Id: 2, Name: "Go Programming", Instructor: som}

	goWS := data.NewWorkshop("Go Programming Workshop", som)
	// fmt.Printf("%v", goWS)

	var courses [2]data.Signable
	courses[0] = goCourse
	courses[1] = goWS

	for _, course := range courses {
		fmt.Println(course)
	}
}
