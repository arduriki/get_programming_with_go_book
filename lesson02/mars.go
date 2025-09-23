// My weight loss program.
package main

import "fmt"

// main is the function where it all begins.
func main() {
	fmt.Print("My weight on the surface of Mars is ")
	fmt.Print(149.0 * 0.3783)
	fmt.Print(" lbs, and I would be ")
	fmt.Print(41 * 365 / 687)
	fmt.Print(" years old.")

	fmt.Printf("My weight on the surface of Mars is %v lbs,", 149.0*0.378)
	fmt.Printf(" and I would be %v years old.\n", 41*365/687)

	// Multiple values.
	fmt.Printf("My weight on the surface of %v is %v lbs.\n", "Earth", 149.0)

	// Align text.
	fmt.Printf("%-15v $%4v\n", "SpaceX", 94)
	fmt.Printf("%-15v $%4v\n", "Virgin Galactic", 100)
}
