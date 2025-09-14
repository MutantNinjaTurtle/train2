package main

import "fmt"

func main() {
	fmt.Println(MaxInt(2, 7))
}

func MaxInt(a, b int) int {
	if a >= b {
		return a
	}

	return b
}
