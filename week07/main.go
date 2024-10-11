package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("your name: ")
	r := bufio.NewReader(os.Stdin)
	// i, err := r.ReadString('\n')
	name, _ := r.ReadString('\n') //\n직전까지 읽어옴.
	fmt.Printf("Welcome %s.", name)
}
