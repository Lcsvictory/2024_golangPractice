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
	// emptySlice = make([]int, 5)

	if len(emptySlice) == 0 { // 슬라이스의 값이 없을때.
		fmt.Printf("emptySlice : %#v\n", emptySlice)
		emptySlice = append(emptySlice, 77)
	}
	fmt.Printf("emptySlice : %#v\n", emptySlice)

}
