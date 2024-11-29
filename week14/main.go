package main

import "fmt"

type visitor struct {
	age  int
	cost int
}

func claculateCost(visitors []visitor) int {
	totalCost := 0
	for _, value := range visitors {
		totalCost += value.cost
	}
	return totalCost
}

func main() {
	var numVisitor int

	fmt.Print("방문 인원은 몇명인가? : ")
	fmt.Scanln(&numVisitor)

	var visitors []visitor
	visitors = make([]visitor, numVisitor)

	for i := 0; i < numVisitor; i++ {
		var age int
		fmt.Print("나이를 입력하세요. : ")
		fmt.Scanln(&age)
		switch {
		case age < 12:
			visitors[i] = visitor{age: age, cost: 5000}
		case age >= 12 && age < 65:
			visitors[i] = visitor{age: age, cost: 10000}
		default:
			visitors[i] = visitor{age: age, cost: 7000}
		}
	}

	fmt.Printf("Total price is %dwon.\n", claculateCost(visitors))
}
