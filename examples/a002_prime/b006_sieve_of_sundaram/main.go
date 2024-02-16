package main

import (
	"fmt"
)

// Function to print primes smaller than n using Sieve of Sundaram
func SieveOfSundaram(n int) {
	// In general Sieve of Sundaram, produces primes smaller
	// than (2*x + 2) for a number given number x.
	// Since we want primes smaller than n, we reduce n to half
	nNew := (n - 1) / 2

	// This array is used to separate numbers of the form i+j+2ij
	// from others where 1 <= i <= j
	marked := make([]bool, nNew+1)

	// Initialize all elements as not marked
	for i := 0; i <= nNew; i++ {
		marked[i] = false
	}
	if ISPEINT {
		fmt.Print("m1:")
	}
	printSlice(marked)
	// Main logic of Sundaram. Mark all numbers of the
	// form i + j + 2ij as true where 1 <= i <= j
	for i := 1; i <= nNew; i++ {
		for j := i; i+j+2*i*j <= nNew; j++ {
			marked[i+j+2*i*j] = true
		}
	}
	if ISPEINT {
		fmt.Print("m2:")
	}
	printSlice(marked)

	// Since 2 is a prime number
	if n > 2 {
		fmt.Print("2 ")
	}

	// Print other primes. Remaining primes are of the form
	// 2*i + 1 such that marked[i] is false.
	for i := 1; i <= nNew; i++ {
		if marked[i] == false {
			fmt.Print(2*i+1, " ")
		}
	}
}

const ISPEINT = false

func printSlice(s []bool) {

	if !ISPEINT {
		return
	}
	for i := 0; i < len(s); i++ {
		if !s[i] {
			fmt.Printf("%d ", 2*i+1)
		}
	}
	fmt.Println()
}

// https://www.geeksforgeeks.org/sieve-sundaram-print-primes-smaller-n/?ref=lbp
// Driver program to test above function
func main() {
	n := 100
	SieveOfSundaram(n)
}
