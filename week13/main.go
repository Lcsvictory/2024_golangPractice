package main

import (
	"fmt"
)

func main() {

	//MAPS
	var myMaps map[string]int = make(map[string]int)

	myMap := map[string]int{
		"na":  1,
		"lee": 2,
	}

	myMaps["kim"] = 10
	fmt.Println(myMaps, myMap)
}
