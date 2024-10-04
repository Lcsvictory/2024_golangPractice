package main

import (
	"fmt"
)

func main() {
	var f float64
	var i int
	var b bool
	var s string
	fmt.Println(f, s, i, b)
	fmt.Printf("%f %s %d %t\n", f, s, i, b) //zero value
	f = 2.7
	i = 3
	fmt.Print(f > float64(i), "\n")
}
