package main

import (
	"fmt"
)

func main() {
	// var notes [5]string = [5]string{"do", "re"} //정석버전
	notes := [5]string{ //Array literal
		"do",
		"re"}

	names := []string{ //Slice literal
		"kim",
		"park",
	}
	fmt.Println(notes) //[do re   ]
	fmt.Println(names) //[kim park]

}
