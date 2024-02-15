package main

import (
	"fmt"
	"math"
)

var prime []int

func simpleSieve(limit int) {
	mark := make([]bool, limit+1)

	p := 2
	for p*p <= limit {
		if mark[p] == false {
			for i := p * p; i <= limit; i += p {
				mark[i] = true
			}
		}
		p++
	}

	for p := 2; p < limit; p++ {
		if mark[p] == false {
			prime = append(prime, p)
			fmt.Printf("%d ", p)
		}
	}
}

func segmentedSieve(n int) {
	limit := int(math.Floor(math.Sqrt(float64(n))) + 1)
	simpleSieve(limit)

	low := limit
	high := limit * 2

	for low < n {
		if high >= n {
			high = n
		}

		mark := make([]bool, limit+1)

		for _, p := range prime {
			loLim := int(math.Floor(float64(low)/float64(p)) * float64(p))
			if loLim < low {
				loLim += p
			}

			for j := loLim; j < high; j += p {
				mark[j-low] = true
			}
		}

		for i := low; i < high; i++ {
			if mark[i-low] == false {
				fmt.Printf("%d ", i)
			}
		}

		low += limit
		high += limit
	}
}

// https://www.geeksforgeeks.org/segmented-sieve/?ref=lbp 分段筛
func main() {
	n := 100
	fmt.Printf("Primes smaller than %d: ", n)
	segmentedSieve(100)
}
