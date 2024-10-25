package main

import (
	"fmt"
)

func main() {
	u := 0
	for {
		u++
		fmt.Println(u)
		if u == 3 {
			break
		}
	}

	t := 0
	for t < 3 {
		t++
		fmt.Println(t)
	}

}
