package main

import (
	"fmt"
)

func sieveOfSundaram(n int) {
	k := (n - 2) / 2
	integersList := make([]bool, k+1)

	for i := 1; i <= k; i++ {
		j := i
		for i+j+2*i*j <= k {
			integersList[i+j+2*i*j] = true
			j++
		}
	}

	if n > 2 {
		fmt.Print(2, " ")
		for i := 1; i <= k; i++ {
			if !integersList[i] {
				fmt.Print(2*i+1, " ")
			}
		}
	}
}

func main() {
	n := 1000
	sieveOfSundaram(n)
}
