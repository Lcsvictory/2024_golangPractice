package main

import (
	"fmt"
	"math/rand"
)

func main() {
	target := rand.Intn(10) + 1 //0 ~ n-1
	fmt.Println(target)
}
