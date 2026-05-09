package main

import (
	"fmt"
)

func sieveOfEratosthenes(n int) []bool {
	primes := make([]bool, n+1)
	primes[0] = true
	for i := 3; i*i <= n; i += 2 {
		if !primes[i/2] {
			for j := 3 * i; j <= n; j += 2 * i {
				primes[j/2] = true
			}
		}
	}
	return primes
}

// https://www.geeksforgeeks.org/sieve-of-eratosthenes/?ref=lbp  第二个程序，奇数
func main() {
	n := 100
	primes := sieveOfEratosthenes(n)
	for i := 1; i <= n; i++ {
		if i == 2 {
			fmt.Print(i, " ")
		} else if i%2 == 1 && !primes[i/2] {
			fmt.Print(i, " ")
		}
	}
}
