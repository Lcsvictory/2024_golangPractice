// print prime number between two input number
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

	// 넘버1
	fmt.Print("시작 정수 입력 : ")
	n1, err := getInteger()
	if err != nil {
		log.Fatal(err)
	}

	// 넘버 2
	fmt.Print("끝 정수 입력 : ")
	n2, err := getInteger()
	if err != nil {
		log.Fatal(err)
	}

	for i := n1; i <= n2; i++ {
		if isPrime(i) { //2보다 작은, 1, 0, -1 ... 은 소수가 아니다.
			fmt.Printf("%d ", i)
		}
	}

}

func getInteger() (int, error) {
	r := bufio.NewReader(os.Stdin)
	a, err := r.ReadString('\n')
	if err != nil {
		return 0, err
	}
	a = strings.TrimSpace(a)
	n, err := strconv.Atoi(a)
	if err != nil {
		return 0, err
	}
	return n, nil
}

func isPrime(n int) bool {
	if n < 2 { //1, 0, -1 ... 은 소수가 아님.
		return false
	} else if n > 2 && n%2 == 0 { //2보다 큰 수중에 짝수라면 소수가 아님.
		return false
	} else {
		for j := 3; j*j <= n; j += 2 { //3부터 2씩 증가해서 홀수만 가지고 판단함.
			if n%j == 0 {
				return false
			}
		}
	}
	return true
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
