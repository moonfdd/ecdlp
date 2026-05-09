package main

import (
	"fmt"
)

const MAX_SIZE = 1000001

func manipulatedSieve(N int) []bool {
	isprime := make([]bool, MAX_SIZE)
	prime := []int{}
	SPF := make([]int, MAX_SIZE)

	for i := 0; i < MAX_SIZE; i++ {
		isprime[i] = true
	}

	isprime[0] = false
	isprime[1] = false

	for i := 2; i < N; i++ {
		if isprime[i] {
			prime = append(prime, i)
			SPF[i] = i
		}

		for j := 0; j < len(prime) && i*prime[j] < N && prime[j] <= SPF[i]; j++ {
			isprime[i*prime[j]] = false
			SPF[i*prime[j]] = prime[j]
		}
	}

	return isprime
}

// https://www.geeksforgeeks.org/sieve-eratosthenes-0n-time-complexity/?ref=lbp
func main() {
	N := 100

	isprime := manipulatedSieve(N)

	for i := 0; i < len(isprime); i++ {
		if isprime[i] && i <= N {
			fmt.Printf("%d ", i)
		}
	}
}
