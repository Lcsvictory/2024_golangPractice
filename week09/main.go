package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	second := time.Now().Unix()
	rand.Seed(second)
	answer := rand.Intn(6) + 1 //0 ~ n-1
	// fmt.Printf("%d\n", answer)
	win := false
	for count := 3; count > 0; count-- {
		fmt.Printf("남은 기회 : %d회, 정답은? : ", count)
		reader := bufio.NewReader(os.Stdin)
		i, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		i = strings.TrimSpace(i)
		guess, err := strconv.Atoi(i)
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Println(guess)

		if answer == guess {
			fmt.Println("정답입니다.")
			win = true
			break
		} else if answer > guess {
			fmt.Println("숫자가 작습니다. LOW")
		} else {
			fmt.Println("숫자가 큽니다. HIGH")
		}
	}
	if win {
		fmt.Printf("You Win! 정답 : %d\n", answer)
	} else {
		fmt.Printf("You Lose.. 정답 : %d\n", answer)
	}

}
