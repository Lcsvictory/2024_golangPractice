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
	fmt.Printf("%#v, %#v\n", myMaps, myMap)
	// map[kim:10] map[lee:2 na:1]
	// map[string]int{"kim":10}, map[string]int{"lee":2, "na":1}

	fmt.Printf("%T, %T, %T, %T\n", 10, "김문자", 'A', 3.14)
	// int, string, int32, float64

	a := '0' // Rune Type, // 0 = 48, A = 65, a = 97
	fmt.Println(a)
	// int32

	var mapLiteral map[int]int = map[int]int{10: 1, 20: 2}
	fmt.Printf("%#v, %T\n", mapLiteral, mapLiteral)
	//map[int]int{10:1, 20:2}, map[int]int

	var mapMake map[string]int = make(map[string]int, 10) //make는 슬라이스 or 맵 가능. array 불가능
	//var mapMake map[string]int = make(map[string]int)
	fmt.Printf("%#v, %T\n", mapMake, mapMake)
	//map[string]int{}, map[string]int

	var mapZeroValue map[string]int
	fmt.Printf("%#v, %T\n", mapZeroValue, mapZeroValue)
	//map[string]int(nil), map[string]int

	mapShort := map[float64]string{3.14: "PI", 2.718: "자연상수"}
	fmt.Printf("%#v, %T\n", mapShort, mapShort)
	//map[float64]string{2.718:"자연상수", 3.14:"PI"}, map[float64]string

	var array [5]int = [5]int{}
	fmt.Printf("%#v, %v %T\n", array, array, array)
	//[5]int{0, 0, 0, 0, 0}, [0 0 0 0 0] [5]int

	var arrayLiteral [5]int = [5]int{1, 2, 3, 4}
	fmt.Printf("%#v, %v %T\n", arrayLiteral, arrayLiteral, arrayLiteral)
	//[5]int{1, 2, 3, 4, 0}, [1 2 3 4 0] [5]int

	var arrayLiteralString [5]string = [5]string{"do", "re"}
	fmt.Printf("%#v, %v %T\n", arrayLiteralString, arrayLiteralString, arrayLiteralString)
	//[5]string{"do", "re", "", "", ""}, [do re   ] [5]string

	arrayShort := [5]int{}
	fmt.Printf("%#v, %v %T\n", arrayShort, arrayShort, arrayShort)
	//[5]int{0, 0, 0, 0, 0}, [0 0 0 0 0] [5]int

	var arrayZeroValue [1]int
	fmt.Printf("%#v, %v %T\n", arrayZeroValue, arrayZeroValue, arrayZeroValue)
	//[1]int{0}, [0] [1]int

	sliceLiteral := []string{ //Slice literal
		"kim",
		"park",
	}
	fmt.Printf("%#v, %v %T\n", sliceLiteral, sliceLiteral, sliceLiteral)
	//[]string{"kim", "park"}, [kim park] []string

	var sliceLiteralZeroValue []string
	fmt.Printf("%#v, %v %T\n", sliceLiteralZeroValue, sliceLiteralZeroValue, sliceLiteralZeroValue)
	//[]string(nil), [] []string

	var sliceLiteral2 []string = make([]string, 0)
	fmt.Printf("%#v, %v %T\n", sliceLiteral2, sliceLiteral2, sliceLiteral2)
	//[]string{}, [] []string

	var sliceLiteral3 []string = make([]string, 1)
	fmt.Printf("%#v, %v %T\n", sliceLiteral3, sliceLiteral3, sliceLiteral3)
	//[]string{""}, [] []string

	sliceAppend := make([]int, 2)
	sliceAppend = append(sliceAppend, 1, 2, 3, 4)
	fmt.Printf("%#v, %v %T\n", sliceAppend, sliceAppend, sliceAppend)
	//[]int{0, 0, 1, 2, 3, 4}, [0 0 1 2 3 4] []int

	arrayAppend := [3]int{1, 2, 3}
	arrayAppend = append(arrayAppend, 1) //array에는 append불가능
	//first argument to append must be a slice; have arrayAppend (variable of type [3]int)

	oneToFive := [2]int{1, 2}
	for i, v := range oneToFive {
		fmt.Printf("index : %d, value : %d\n", i, v)
	}
	//index : 0, value : 1
	//index : 1, value : 2

	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}
	//0 1 2 3 4 5 6 7 8 9
}
