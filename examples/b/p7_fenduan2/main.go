package main

import "fmt"

const (
	maxN       = 10000
	filterSize = 101
)

var compositeFilter = make([]bool, filterSize)
var primesList = make([]int, 0)

func prepareSieve() {
	compositeFilter[0] = true
	compositeFilter[1] = true
	for p := 2; p*p < filterSize; p++ {
		if !compositeFilter[p] {
			for i := p * p; i < filterSize; i += p {
				compositeFilter[i] = true
			}
		}
	}

	for i, v := range compositeFilter {
		if !v {
			primesList = append(primesList, i)
		}
	}

}

func primesInRange(m, n int) {
	// we all know first prime starts at 2
	if m <= 1 {
		m = 2
	}
	d := (n - m) + 1

	composite := make([]bool, d)
	for _, p := range primesList {
		if p*p > n {
			break
		}

		i := (m / p) * p

		if i < m {
			i += p
		}

		// making sure we are not marking prime `p` as composite
		if i == p {
			i += p
		}
		// marking all the multiples of p as composite (except p itself)
		for ; i <= n; i += p {
			composite[i-m] = true
		}
	}

	for i := m; i <= n; i++ {
		if !composite[i-m] {
			fmt.Print(i, " ")
		}
	}
}

// https://www.geeksforgeeks.org/segmented-sieve-print-primes-in-a-range/?ref=lbp
func main() {
	prepareSieve()
	primesInRange(2, 100)
}
