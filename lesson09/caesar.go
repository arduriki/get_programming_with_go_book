package main

import "fmt"

func main() {
	// The encrypted message.
	message := "L fdph, L vdz, L frqtxhuhg."

	// Process each character.
	for i := 0; i < len(message); i++ {
		// Each character is stored as a byte value (ASCII code).
		c := message[i]
		// Decoding lowercase letters.
		if c >= 'a' && c <= 'z' {
			c -= 3
			// Is past 'a', we need to wrap to the end of the alphabet.
			if c < 'a' {
				c += 26
			}
			// Decoding uppercase letters.
		} else if c >= 'A' && c <= 'Z' {
			c -= 3
			// Wrap around.
			if c < 'A' {
				c += 26
			}
		}
		fmt.Printf("%c", c)
	}
	fmt.Println()
}
