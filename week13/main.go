package main

import (
	"fmt"
	"reflect"
)

func main() {
	//슬라이스는 복사가 아닌 참조.

	//Create a slice by slicing an existing array
	gpas := [5]float64{3.5, 4.1, 4.5, 3.9, 4.23} // array := array literal
	gpaSlice := gpas[1:4]                        // slice := slicing array
	// gpaSlice[1] = 2.71
	gpas[2] = 2.71
	gpaSlice = append(gpaSlice, 4.3)
	fmt.Println(gpas, gpaSlice, reflect.TypeOf(gpaSlice))

	var emptySlice []int
	fmt.Printf("intSlice: %#v\n", emptySlice)
	emptySlice = make([]int, 10)
	//emptySlice = append(emptySlice, 27, 90, 101)
	fmt.Printf("intSlice: %#v\n", emptySlice)
	// sclice zero value == nil

}
