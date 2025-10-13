package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	wordCount := make(map[string]int)
	words := strings.Fields(s)

	for _,word := range words {
		wordCount[word] = wordCount[word] + 1
	}
	return wordCount
	
}

func func23() {
	fmt.Println("func 23: ")
	wc.Test(WordCount)
}