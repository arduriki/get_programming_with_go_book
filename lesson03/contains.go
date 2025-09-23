package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("You find yourself in a dimly lit cavern.")

	command := "walk outside"
	// Asks if the command variable contains the text "outside"
	// it returns a boolean.
	exit := strings.Contains(command, "outside")

	fmt.Println("You leave the cave:", exit)
	fmt.Println()

	// Comparisons.
	fmt.Println("There is a sign near the entrance that reads 'No Minors'.")

	age := 41
	minor := age < 18

	fmt.Printf("At age %v, am I a minor? %v\n", age, minor)
}
