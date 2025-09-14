package main

import "fmt"

func main() {
	fmt.Println("ppc")
}

func MaxInt(a, b int) int {
	if a >= b {
		return a
	}

	return b
}
