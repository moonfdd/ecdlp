package main

import (
	"fmt"
)

func sieveOfEratosthenes(n int) []bool {
	primes := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		primes[i] = true
	}

	for p := 2; p*p <= n; p++ {
		if primes[p] == true {
			for i := p * p; i <= n; i += p {
				primes[i] = false
			}
		}
	}

	return primes
}

func main() {
	n := 100
	primes := sieveOfEratosthenes(n)
	for i := 2; i <= n; i++ {
		if primes[i] == true {
			fmt.Printf("%d ", i)
		}
	}
}
