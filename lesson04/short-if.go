package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Use short declaration to declare a new variable in an if statement
	if num := rand.Intn(3); num == 0 {
		fmt.Println("Space Adventures")
	} else if num == 1 {
		fmt.Println("SpaceX")
	} else {
		fmt.Println("Virgin Galactic")
	}
}
