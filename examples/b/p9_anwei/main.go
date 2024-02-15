package main

import (
	"fmt"
)

// Checks whether x is prime or composite
func ifnotPrime(prime []int, x int) bool {
	// checking whether the value of element
	// is set or not. Using prime[x/64], we find
	// the slot in prime array. To find the bit
	// number, we divide x by 2 and take its mod
	// with 32.
	return (prime[x/64] & (1 << ((x >> 1) & 31))) != 0
}

// Marks x composite in prime[]
func makeComposite(prime []int, x int) {
	// Set a bit corresponding to given element.
	// Using prime[x/64], we find the slot in prime
	// array. To find the bit number, we divide x
	// by 2 and take its mod with 32.
	prime[x/64] |= (1 << ((x >> 1) & 31))
}

// Prints all prime numbers smaller than n.
func bitWiseSieve(n int) {
	// Assuming that n takes 32 bits, we reduce
	// size to n/64 from n/2.
	prime := make([]int, n/64+1)

	// Initializing values to 0 .
	for i := range prime {
		prime[i] = 0
	}

	// 2 is the only even prime so we can ignore that
	// loop starts from 3 as we have used in sieve of
	// Eratosthenes .
	for i := 3; i*i <= n; i += 2 {

		// If i is prime, mark all its multiples as
		// composite
		if !ifnotPrime(prime, i) {
			for j, k := i*i, i<<1; j < n; j += k {
				makeComposite(prime, j)
			}
		}
	}

	// writing 2 separately
	fmt.Print("2 ")

	// Printing other primes
	for i := 3; i <= n; i += 2 {
		if !ifnotPrime(prime, i) {
			fmt.Printf("%d ", i)
		}
	}
}

// https://www.geeksforgeeks.org/bitwise-sieve/?ref=lbp
func main() {
	n := 30
	bitWiseSieve(n)
}
