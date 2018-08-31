package main

import (
	"fmt"
//	"os"
//	"strconv"
)

// 0 = white
// 1 = black
// pattern of keys. Mapped to white/black

func main() {
	number := 2017
	pattern := [16]int{0,1,2,1,2,1,2,1,2,1,2,1,2,1,2,3}
	var keys []int  

	keyMap := map[int][]int{
		0 : []int{0,1,0},
		1 : []int{0,1,0,1,0},
		2 : []int{0,1,0,1,0,1,0},
		3 : []int{0},
	} // keyMap
	
	// make the keys slice
	for _, item := range pattern { 
		keys = append(keys, keyMap[item]...)
	} // for
	
	//strArg := os.Args[1:]
	//number, err := strconv.Atoi(strArg)

	//if err == nil {
	//	fmt.Println("number conversion:  ")
	//	fmt.Println(number)
	//} // if
	
	//original := os.Args[1]

	var original int = number
	
	for number > 88 {
		number -= 88
	} // for

	colorIndex := keys[number-1]
	color := ""
	
	switch colorIndex {
	case 0:
		color = "white"
	case 1:
		color = "black"
	default:
		color = "FAIL"
	} // switch

	fmt.Println(original)
	fmt.Println(color)
	
} // main()
