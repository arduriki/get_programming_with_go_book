/*
This program calculates how long it would take to travel to the Andromeda Galaxy
at the speed of light, using extremely large numbers that require special handling.
*/

package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Create a big integer from a regular number.
	lightSpeed := big.NewInt(299792)
	secondsPerDay := big.NewInt(86400)

	// Create an empty big integer.
	distance := new(big.Int)
	// base-10 number.
	distance.SetString("24000000000000000000", 10)
	fmt.Println("Andromeda Galaxy is", distance, "km away.")

	seconds := new(big.Int)
	seconds.Div(distance, lightSpeed)

	days := new(big.Int)
	days.Div(seconds, secondsPerDay)

	fmt.Println("That is", days, "days of travel at light speed.")
}
