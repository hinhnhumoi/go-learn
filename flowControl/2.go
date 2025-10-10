package main

import "fmt"

func func2() {
	sum := 1
	for ; sum < 1000 ; {
		sum += sum
	}
	fmt.Println("Sum:", sum)
}