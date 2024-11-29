package main

import "fmt"

type student struct {
	id   int
	name string
	gpa  float32
	age  int
}

func main() {
	var student1 student
	student1.id = 20241100
	student1.name = "Kim inha"
	student1.gpa = 4.5
	student1.age = 20
	fmt.Printf("%#v\n", student1.id)

	var student2 student
	student2.id = 20242200
	student2.name = "Kim inha"
	student2.gpa = 4.5
	student2.age = 20
	fmt.Printf("%#v\n", student2.name)
	// var school struct { 구조체는 타입이 아니다. 클래스도 아니다.
	// 	student stu
	// }
}
