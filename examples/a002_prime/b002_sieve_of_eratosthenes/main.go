package main

import "fmt"

func main() {
	if true {
		fmt.Println("埃氏筛")
		sieveOfEratosthenes1(100)
		return
	}
	if true {
		fmt.Println("奇数的埃氏筛。不是重点，可以不看。")
		n := 100
		primes := sieveOfEratosthenes2(n)
		for i := 1; i <= n; i++ {
			if i == 2 {
				fmt.Print(i, " ")
			} else if i%2 == 1 && !primes[i/2] {
				fmt.Print(i, " ")
			}
		}
	}
	fmt.Println("")
}

// https://www.geeksforgeeks.org/sieve-of-eratosthenes/?ref=lbp
// 第1个程序
func sieveOfEratosthenes1(n int) {
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
// 第2个程序
func sieveOfEratosthenes2(n int) []bool {
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
