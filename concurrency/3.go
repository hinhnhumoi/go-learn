package main

import "fmt"

func func3() {
	fmt.Println("func 3: ")
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}