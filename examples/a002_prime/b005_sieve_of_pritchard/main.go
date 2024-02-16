package main

import (
	"fmt"
	"math"
)

func pritchard(limit int) []int {
	members := make([]bool, limit+1)
	members[1] = true
	steplength, prime, sq, nlimit := 1, 2, int(math.Sqrt(float64(limit))), 2
	primes := []int{}
	for prime <= sq {
		if ISPEINT {
			fmt.Println("------------")
			fmt.Printf("prime = %d, steplength = %d, nlimit = %d\r\n", prime, steplength, nlimit)
		}
		if steplength < limit {
			if ISPEINT {
				fmt.Print("m1 = ")
				printSlice(members)
			}
			for w := 1; w < len(members); w++ {
				if members[w] {
					n := w + steplength
					for n <= nlimit {
						members[n] = true
						n += steplength
					}
				}
			}
			if ISPEINT {
				fmt.Print("m2 = ")
			}
			printSlice(members)
			steplength = nlimit
		}

		np := 5
		mcpy := make([]bool, len(members))
		copy(mcpy, members)
		if ISPEINT {
			fmt.Print("delete：")
		}
		for w := 1; w < len(members); w++ {
			if mcpy[w] {
				if np == 5 && w > prime {
					np = w
				}
				n := prime * w
				if n > nlimit {
					break
				}
				if members[n] {
					if ISPEINT {
						fmt.Print(n, " ")
					}
				}
				members[n] = false
			}
		}
		if ISPEINT {
			fmt.Println()
			fmt.Print("m3 = ")
		}
		printSlice(members)

		if np < prime {
			break
		}
		primes = append(primes, prime)
		if prime == 2 {
			prime = 3
		} else {
			prime = np
		}
		nlimit = int(math.Min(float64(steplength*prime), float64(limit)))
	}
	newprimes := []int{}
	for i := 2; i < len(members); i++ {
		if members[i] {
			newprimes = append(newprimes, i)
		}
	}
	if ISPEINT {
		fmt.Println("------------")
	}
	return append(primes, newprimes...)
}

const ISPEINT = false

func printSlice(s []bool) {

	if !ISPEINT {
		return
	}
	for i := 0; i < len(s); i++ {
		if s[i] {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()
}

// https://rosettacode.org/wiki/Sieve_of_Pritchard#Python
func main() {
	fmt.Println(pritchard(401))
	// fmt.Println("Number of primes up to 1,000,000:", len(pritchard(1000000)))
}
