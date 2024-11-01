package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("정수 입력 : ")
	r := bufio.NewReader(os.Stdin)
	i, err := r.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	i = strings.TrimSpace(i)
	n, err := strconv.Atoi(i)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%f, int = %d\n", math.Sqrt(53.0), int(math.Sqrt(53.0))) //실수를 정수로 캐스팅하면 소수점 날라감.

	isPrime := true
	if n < 2 {
		isPrime = false
	} else if n > 2 && n%2 == 0 {
		isPrime = false
	} else {
		for j := 3; j*j <= n; j += 2 {
			if n%j == 0 {
				isPrime = false
				break
			}
		}
	}

	if isPrime { //2보다 작은, 1, 0, -1 ... 은 소수가 아니다.
		fmt.Printf("%d는(은) 소수입니다.", n)
	} else {
		fmt.Printf("%d는(은) 소수가 아닙니다.", n)
	}

	// fmt.Printf("%d is %t", n, IsPrimeSqrt(n))
}

func IsPrimeSqrt(n int) bool {
	if n == 1 {
		return false
	}
	sqrtN := int(math.Sqrt(float64(n))) + 1
	for j := 2; j < sqrtN; j++ {
		if n%j == 0 {
			return false
		}
	}
	return true
}
