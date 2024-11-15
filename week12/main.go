package main

import (
	"fmt"
	"time"
)

func main() {
	// var dates [2]time.Time
	// dates[1] = time.Unix(14420000, 0)
	// fmt.Println(dates[1])

	// var dates [3]time.Time = [3]time.Time{
	// 	time.Unix(0, 0),
	// 	time.Unix(1, 0),
	// 	time.Unix(1244444401, 0),
	// }

	dates := [3]time.Time{
		time.Unix(0, 0),
		time.Unix(1, 0),
		time.Unix(1244444401, 0), //줄바꿈이 있다면 마지막 요소에 ,는 필수
	}

	// dates1 := [3]time.Time{
	// 	time.Unix(0, 0),
	// 	time.Unix(1, 0),
	// 	time.Unix(1947920001, 0)} //줄바꿈 없으면 없어도 됨.

	// fmt.Println(dates[0], dates[1], dates[2])
	// fmt.Printf("%#v\n", dates1) //literal 형태 그대로 출력.

	for i := 0; i < len(dates); i++ {
		fmt.Printf("%d, %v\n", i, dates[i])
	}
}
