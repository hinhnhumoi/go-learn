package main

import "fmt"

func func22() {
	fmt.Println("func 22: ")
	m22 := make(map[string]int)

	m22["Answer"] = 42
	fmt.Println("The value: ", m22["Answer"])

	m22["Answer"] = 48
	fmt.Println("The value:", m22["Answer"])

	delete(m22, "Answer")
	fmt.Println("The value:", m22["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}