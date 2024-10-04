package main

import (
	"fmt"
	"reflect"
)

func main() {
	i := 13
	var t float64 = 12.8
	f := 12.9 //int(12.9) == 12

	//Rune타입, 자바의 chr(캐릭터)와 동일한 개념.
	c1 := 'A' //65 int32 == 4byte
	c2 := '김' //44608 0xAE40
	c3 := '0' //48
	c4 := 'a' //97

	fmt.Printf("value i :%d,  value F : %f\n", i, f)
	fmt.Printf("%d * %f = %.02f\n", i, f, float64(i)*f) //형변환
	fmt.Printf("%d * %f = %d\n", i, f, i*int(f))        //형변환 float -> int시 Floor적용.
	// fmt.Printf("%d * %f = %.02f\n", i, f, i*f) # int * float = error
	fmt.Println((reflect.TypeOf(i)))
	fmt.Println(reflect.TypeOf(i), reflect.TypeOf(t), reflect.TypeOf(c1), reflect.TypeOf(c2))
	fmt.Println(c1, c2, c3, c4)
}
