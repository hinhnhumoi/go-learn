package main

import "fmt"

func func12() {

	fmt.Println("func 12: ")
	var s []int
	fmt.Println(s, len(s), cap(s))
	// if s == nil {
	// 	fmt.Println("slice is nil")
	// }
}