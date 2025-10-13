package main

import "fmt"

func func17()  {
	fmt.Println("func 17: ")
	pow := make([]int, 10)

	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}

	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}