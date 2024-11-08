package main

import (
	"fmt"
	"log"
	"week11/greeting"
	"week11/keyboard"
)

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

func main() {
	greeting.Hi("Hong")
	greeting.Hello("Kim")

	// 넘버1
	fmt.Print("시작 정수 입력 : ")
	n1, err := keyboard.GetInteger()
	if err != nil {
		log.Fatal(err)
	}
	// 넘버 2
	fmt.Print("끝 정수 입력 : ")
	n2, err := keyboard.GetInteger()
	if err != nil {
		log.Fatal(err)
	}
	for i := n1; i <= n2; i++ {
		if isPrime(i) { //2보다 작은, 1, 0, -1 ... 은 소수가 아니다.
			fmt.Printf("%d ", i)
		}
	}

}
