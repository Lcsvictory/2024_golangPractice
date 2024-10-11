package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Print("your name: ")
	r := bufio.NewReader(os.Stdin)
	// i, err := r.ReadString('\n')
	name, err := r.ReadString('\n') //\n포함하여 읽어옴. \n 포함하지 않고 읽고 싶다.
	if err != nil {                 //(nil = 에러가 없다.) 즉 에러가 있다면.
		log.Fatal(err)
	} else {
		fmt.Printf("Welcome %s", name)
	}
}
