package main

import "fmt"

// 0 = white
// 1 = black
// pattern of keys. Mapped to white/black

func main() {
 
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
	
	fmt.Println("Printing out the keys:")
	fmt.Println(keys)

	keyboard, sep := "", ""
	for _, k := range keys {
		color := ""
		
		switch k {
		case 0:
			color = "white"
		case 1:
			color = "black"
		default:
			color = "FAIL"
		}
		
		keyboard += sep + color
		sep = ", "

	} // for

	fmt.Println(keyboard)
	
} // main()
