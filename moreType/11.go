package main

import "fmt"

func func11()  {
	fmt.Println("func 11: ")
	slice := []int{2, 3, 5, 7, 11, 13}

	printSlice(slice)

	slice = slice[:0]
	printSlice(slice)

	slice = slice[:4]
	printSlice(slice)

	slice = slice[2:]
	printSlice(slice)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}