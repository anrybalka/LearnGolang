package main

import "fmt"

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func main() {
	//простые string, integer, float, bool

	//[u]int[n]

	//type rune = int32 //символы unicode
	//type byte = uint8 //байты

	//float32 ~6 знаков // быстро копятся ошибки
	//float64 ~15 знаков

	//bool
	var isBool bool = true
	var numBool int = boolToInt(isBool)
	fmt.Println(numBool)

	//string
	var str1 = "\tHello\n"
	fmt.Println(str1)

	//составные array, map, struct
}
