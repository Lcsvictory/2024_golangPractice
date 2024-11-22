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
	fmt.Println(gpas, gpaSlice, reflect.TypeOf(gpaSlice))

}
