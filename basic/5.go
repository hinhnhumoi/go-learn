package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func addShort(x, y int) int {
	return x + y
}

func func5() {
	fmt.Println(add(42,13))
	fmt.Println(addShort(15,33))
}