package main

import "fmt"

func main() {
	var result string = fmt.Sprintf("%d / %d = %0.3f\n", 1, 3, 1.0/3.0)
	// var result string = fmt.Sprintf("%v / %v = %0.3f\n", 1, 3, 1.0/3.0)
	// var result string = fmt.Sprintf("%#v / %#v = %0.3f\n", 1, 3, 1.0/3.0)
	fmt.Print(result)
	i := 1
	for i <= 10 {
		fmt.Printf("%%%2d\n", i)
		i++
	}
}
