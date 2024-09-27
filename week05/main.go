package main

import (
	"fmt"
	"math"
	"strings"
)

func imports() {
	// var f float32
	fmt.Println(math.Floor(2.75))
	fmt.Println(math.Ceil(2.1))
	fmt.Printf("%s\n", strings.Title("head fist go"))
}

func I_O() {
	i := 55
	//var i int = 55
	// var i int
	// i = 55

	fmt.Printf("value i : %d\n", i)
	fmt.Print("value i : ", i, "\n")
	fmt.Println("value i :", i)
}

func main() {
	// I_O()
	imports()
}
