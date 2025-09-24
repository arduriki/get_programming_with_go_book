package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Scope inside the main function.
	count := 0

	for count < 10 {
		// Scope inside the for loop.
		num := rand.Intn(10) + 1
		fmt.Println(num)

		count++
	}
}
