package main

import (
	"fmt"
	"reflect"
)

func main() {
	i := 13
	t := 12
	f := 12.9 //int(12.9) == 12
	fmt.Printf("value i :%d,  value F : %f\n", i, f)
	fmt.Printf("%d * %f = %.02f\n", i, f, float64(i)*f) //형변환
	fmt.Printf("%d * %f = %d\n", i, f, i*int(f))        //형변환 float -> int시 Floor적용.
	// fmt.Printf("%d * %f = %.02f\n", i, f, i*f) # int * float = error
	fmt.Println((reflect.TypeOf(i)))
	fmt.Println(t == int(f))
}
