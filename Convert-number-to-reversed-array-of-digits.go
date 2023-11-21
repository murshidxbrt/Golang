package main

import (
	"fmt"
)

func Digitize(n int) []int {
	var reversedDigits []int

	// Special case: If the number is 0, return [0]
	if n == 0 {
		return []int{0}
	}

	// Extract digits and append to the reversedDigits slice
	for n > 0 {
		digit := n % 10
		reversedDigits = append(reversedDigits, digit)
		n /= 10
	}

	return reversedDigits
}

func main() {
	number := 35231
	reversedDigits := Digitize(number)
	fmt.Println(reversedDigits)
}
