package main

import (
	"fmt"
	"math/rand"
)

// Package scope: visible to all the functions declared in the package.
// In this case, short declaration can't be used.
var era = "AD"

func main() {
	// Variable visible only in the main function.
	year := 2018

	// month is only visible in the switch statement.
	switch month := rand.Intn(12) + 1; month {
	case 2:
		// Every 'day' variable is declared within each case block.
		day := rand.Intn(28) + 1
		fmt.Println(era, year, month, day)
	case 4, 6, 9, 11:
		day := rand.Intn(30) + 1
		fmt.Println(era, year, month, day)
	default:
		day := rand.Intn(31) + 1
		fmt.Println(era, year, month, day)
	}
}
