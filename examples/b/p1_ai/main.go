package main

import (
	"fmt"
)

func sieveOfEratosthenes(n int) {
	// Create a boolean array "prime[0..n]" and initialize
	// all entries it as true. A value in prime[i] will
	// finally be false if i is Not a prime, else true.
	prime := make([]bool, n+1)
	for i := range prime {
		prime[i] = true
	}

	for p := 2; p*p <= n; p++ {
		// If prime[p] is not changed, then it is a prime
		if prime[p] {
			// Update all multiples of p greater than or
			// equal to the square of it numbers which are
			// multiple of p and are less than p^2 are
			// already been marked.
			for i := p * p; i <= n; i += p {
				prime[i] = false
			}
		}
	}

	// Print all prime numbers
	for p := 2; p <= n; p++ {
		if prime[p] {
			fmt.Printf("%d ", p)
		}
	}
}

// https://www.geeksforgeeks.org/sieve-of-eratosthenes/?ref=lbp
func main() {
	n := 30
	fmt.Printf("Following are the prime numbers smaller than or equal to %d\n", n)
	sieveOfEratosthenes(n)
}
