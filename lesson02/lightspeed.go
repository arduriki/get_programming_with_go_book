// How long does it take to get to Mars?
package main

import "fmt"

func main() {
	// Constant.
	const lightSpeed = 299792
	// Variables, in this case declaring multiple ones.
	distance := 56000000

	fmt.Println(distance/lightSpeed, "seconds")

	// Changing the value in the variable.
	distance = 401000000
	fmt.Println(distance/lightSpeed, "seconds")
}
