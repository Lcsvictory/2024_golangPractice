package main

import (
	"fmt"
	"reflect"
)

func main() {
	// var notes [5]string = [5]string{"do", "re"} //정석버전
	notes := [5]string{ //Array literal
		"do",
		"re"}

	//방법1
	//var vars []string
	//vars = make([]string, 5)
	var vars []string = make([]string, 5)

	//방법2
	primes := make([]int, 5)

	//방법3
	names := []string{ //Slice literal
		"kim",
		"park",
	}
	fmt.Println(notes)                 //[do re   ]
	fmt.Println(reflect.TypeOf(names)) //[kim park]

	fmt.Println(reflect.TypeOf(vars))
	fmt.Println(reflect.TypeOf(primes))

	// Create a slice by make function
	var gpaSlice []float64
	gpaSlice = make([]float64, 3)
	gpaSlice[0] = 4.1
	gpaSlice[1] = 4.5
	gpaSlice[2] = 3.9
	gpaSlice = append(gpaSlice, 35.0)
	fmt.Println(gpaSlice, reflect.TypeOf(gpaSlice))

	// Create a slice by slice literal
	// gpaSlice := []float64{4.1, 4.5, 3.9} // slice literal
	// fmt.Println(gpaSlice, reflect.TypeOf(gpaSlice))

	// Create a slice by slicing an existing array
	// gpas := [5]float64{3.5, 4.1, 4.5, 3.9, 4.23} // array := array literal
	// fmt.Println(gpas, reflect.TypeOf(gpas))
	// gpaSlice := gpas[1:4] // slice := slicing array
	// fmt.Println(gpaSlice, reflect.TypeOf(gpaSlice))

}
