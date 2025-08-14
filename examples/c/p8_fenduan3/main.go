package main

import (
	"fmt"
	"math"
)

// To store the prime numbers
var allPrimes map[int]bool

// Function that find prime numbers
// till limit
func simpleSieve(limit int) []int {
	mark := make([]bool, limit+1)
	prime := make([]int, 0)

	// Find primes using Sieve of Eratosthenes
	for i := 2; i <= limit; i++ {
		if !mark[i] {
			prime = append(prime, i)
			for j := i; j <= limit; j += i {
				mark[j] = true
			}
		}
	}

	return prime
}

// Function that finds all prime
// numbers in given range using Segmented Sieve
func primesInRange(low, high int) {
	// Find the limit
	limit := int(math.Sqrt(float64(high))) + 1

	// To store the prime numbers
	prime := simpleSieve(limit)

	// Count the elements in the range [low, high]
	n := high - low + 1

	// Declaration and initialization of mark
	mark := make([]bool, n+1)

	// Traverse the prime numbers till limit
	for _, p := range prime {
		loLim := (low / p) * p

		// Find the minimum number in [low..high] that is a multiple of p
		if loLim < low {
			loLim += p
		}

		if loLim == p {
			loLim += p
		}

		// Mark the multiples of p in [low, high] as true
		for j := loLim; j <= high; j += p {
			mark[j-low] = true
		}
	}

	// Elements that are not marked in range are Prime
	allPrimes = make(map[int]bool)
	for i := low; i <= high; i++ {
		if !mark[i-low] {
			allPrimes[i] = true
		}
	}
}

// Function that finds longest subarray of prime numbers
func maxPrimeSubarray(arr []int) int {
	currentMax := 0
	maxSoFar := 0

	for _, num := range arr {
		// If element is Non-prime then update currentMax to 0
		if !allPrimes[num] {
			currentMax = 0
		} else {
			// If element is prime, then update currentMax and maxSoFar
			currentMax++
			maxSoFar = int(math.Max(float64(currentMax), float64(maxSoFar)))
		}
	}

	// Return the count of longest subarray
	return maxSoFar
}

// https://www.geeksforgeeks.org/longest-sub-array-of-prime-numbers-using-segmented-sieve/?ref=lbp
// Driver Code
func main() {
	arr := []int{1, 2, 4, 3, 29, 11, 7, 8, 9}
	// n := len(arr)

	// Find minimum and maximum element
	maxEl := math.Inf(-1)
	minEl := math.Inf(1)
	for _, num := range arr {
		if float64(num) > maxEl {
			maxEl = float64(num)
		}
		if float64(num) < minEl {
			minEl = float64(num)
		}
	}

	// Find prime in the range [minEl, maxEl]
	primesInRange(int(minEl), int(maxEl))

	// Function call
	fmt.Println(maxPrimeSubarray(arr))
}
