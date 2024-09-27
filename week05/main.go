package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

func imports() {
	fmt.Println(math.Floor(2.75))
	fmt.Println(math.Ceil(2.1))
	fmt.Printf("%s\n", strings.Title("head fist go"))
}

func I_O() {
	var f float32 = 4.3
	i := 55   //int
	j := 55.0 //float64 == double
	//var i int = 55
	// var i int
	// i = 55

	fmt.Printf("value i : %d\n", i)
	fmt.Print("value i : ", i, "\n")
	fmt.Println("value i :", i)
	fmt.Println(reflect.TypeOf(f))
	fmt.Println(reflect.TypeOf(j))
}

func main() {
	I_O()
	// imports()
}
