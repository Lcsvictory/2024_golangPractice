package main

import (
	"fmt"
	"os"
	"reflect"
)

// func test(names string) { //단일매개변수 arguments
// fmt.Println(names)
// }

func test(names ...string) { //가변매개변수 variable arguments
	fmt.Println(names, reflect.TypeOf(names))
}

func main() {
	singer := os.Args[1:] //$0 실행파일경로 $1~... 공백으로 분리된 매개변수들
	// 실행시킬때 커맨드 창에서 입력받은 매개변수들.
	fmt.Println(singer)
	fmt.Printf("%T\n", singer[1]) // TypeOf()

	singer = append(singer, "IU", "BTS")
	fmt.Println(singer, len(singer))

	test("you", "are")
	test()
}
