package data

import "fmt"

var Country [10]string
var Slice []int
var Codes map[int]string

const MaxSpeed = 6000000

func init() {
	Country[0] = "India"
	Country[1] = "USA"
	Country[2] = "UK"
	Country[6] = "France"
	Country[7] = "Italy"
	Country[8] = "Japan"
	Country[9] = "China"
	qty := len(Country)
	fmt.Println("Countries :", qty)
}
