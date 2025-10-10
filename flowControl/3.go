package main

import "fmt"

func func3() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println("Sum3: ", sum)
}