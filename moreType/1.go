package main

import "fmt"

func func1() {
	i, j := 42, 27001
	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)
}