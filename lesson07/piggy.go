package main

import (
	"fmt"
	"math/rand"
)

func main() {
	piggyBank := 0

	// Less than 2000 cents ($20.00).
	for piggyBank < 2000 {
		// Random coin selection.
		switch rand.Intn(3) {
		case 0:
			piggyBank += 5
		case 1:
			piggyBank += 10
		case 2:
			piggyBank += 25
		}

		// Convert cents to dollars.
		dollars := piggyBank / 100
		// Remaining cents.
		cents := piggyBank % 100
		// Print as a 2-digit number with leading zeros.
		fmt.Printf("$%d.%02d\n", dollars, cents)
	}
}
