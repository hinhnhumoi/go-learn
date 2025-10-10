package main

import (
	"fmt"
	"runtime"
)

func func9() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Print(" MacOS\n")
	case "linux":
		fmt.Print(" Linux \n")
	default:
		fmt.Printf(" %s\n", os)
	}
}