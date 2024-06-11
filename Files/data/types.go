package data

import "fmt"

// type integer = int

// var x integer

type Location string

func (location Location) DistanceTo(destination Location) distance {
	// TODO calculations
	fmt.Printf("Origin : %v , Destination : %v ", location, destination)
	return 10
}

func locationTest() {
	nyc := Location("33.4324 , 34.324")

	nyc.DistanceTo(Location("40.7128 , 74.0060"))

}

type distance float64   //miles
type distanceKm float64 //kilometers

// Method
func (miles distance) ToKm() distanceKm {
	return distanceKm(miles * 1.60934)

}
func (kilometres distanceKm) ToMiles() distance {
	return distance(kilometres / 1.60934)

}
func test() {
	d := distance(4.5)
	km := d.ToKm()
	againMiles := km.ToMiles()
	print(d)
	print(km)
	print(againMiles)
}
