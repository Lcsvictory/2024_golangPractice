package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	second := time.Now().Unix()
	fmt.Printf("%d\n", second)
	rand.Seed(second)
	// rand.NewSource(14)
	target := rand.Intn(100) + 1 //0 ~ n-1
	fmt.Printf("%d\n", target)
}
