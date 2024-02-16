package main

import (
	"fmt"
	"math"
)

// Function to check if a given
// number x is prime or not
func isPrime(N int) {
	isPrime := true
	// The Wheel for checking
	// prime number
	arr := []int{7, 11, 13, 17, 19, 23, 29, 31}

	// Base Case
	if N < 2 {
		isPrime = false
	}
	if N == 2 || N == 3 || N == 5 {
		isPrime = true
	}

	// Check for the number taken
	// as basis
	if N%2 == 0 || N%3 == 0 || N%5 == 0 {
		isPrime = false
	}
	sq := int(math.Sqrt(float64(N)))
	// Check for Wheel
	// Here i, acts as the layer
	// of the wheel
	for i := 0; i < sq; i += 30 {

		// Check for the list of
		// Sieve in arr[]
		for _, c := range arr {

			// If number is greater
			// than sqrt(N) break
			if c > sq {
				break
			}

			// Check if N is a multiple
			// of prime number in the
			// wheel
			if N%(c+i) == 0 {
				isPrime = false
				break
			}

			// If at any iteration
			// isPrime is false,
			// break from the loop
			if !isPrime {
				break
			}
		}
	}

	if isPrime {
		fmt.Println("Prime Number")
	} else {
		fmt.Println("Not a Prime Number")
	}
}

// https://www.geeksforgeeks.org/wheel-factorization-algorithm/?ref=lbp

// Driver's Code
func main() {
	N := 9973

	// Function call for primality
	// check
	isPrime(N)
}
