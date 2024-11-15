package main

import (
	"fmt"

	"github.com/headfirstgo/keyboard"
)

func main() {
	var gpa [3]float64
	for i := 0; i < len(gpa); i++ {
		fmt.Print("Input float number(your grade):")
		gpa[i], _ = keyboard.GetFloat() // go get github.com/headfirstgo/keyboard
	}

	for i, v := range gpa {
		fmt.Printf("%d번학생의 학점 : %f\n", i, v)
	}

}
