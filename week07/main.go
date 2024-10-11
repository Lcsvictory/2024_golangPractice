package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("your score : ")
	reader := bufio.NewReader(os.Stdin)
	i, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	i = strings.TrimSpace(i)
	score, _ := strconv.ParseInt(i, 16, 32) // 0~36진수 까지 가능. 32비트이냐 64비트이냐.
	// score, _ := strconv.ParseFloat(i, 32)  //32비트냐 64비트냐.
	fmt.Printf("%d", score)

	if score >= 60 {
		fmt.Println("A")
	} else {
		fmt.Println("BCDF")
	}

}

// 16 * 4 + 1
