package main

import "fmt"

func func12() {
	defer fmt.Println("world")

	fmt.Println("hello")
}