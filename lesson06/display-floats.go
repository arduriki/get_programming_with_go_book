// Formatted print for floating-point
package main

import "fmt"

func main() {
	third := 1.0 / 3
	fmt.Println(third)
	fmt.Printf("%v\n", third)
	// %f to specify the number of digits.
	fmt.Printf("%f\n", third)
	fmt.Printf("%.3f\n", third)
	// Specify the width to show %05 (fill with 0s if smaller)
	// and the number of decimals .2f
	fmt.Printf("%05.2f\n", third)
}
