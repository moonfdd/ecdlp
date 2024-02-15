// https://en.wikipedia.org/wiki/P%C3%A9pin%27s_test
package main

import (
	"fmt"
	"math/big"
)

// n>=0
func PePin(n *big.Int) bool {
	if n.Cmp(big.NewInt(0)) == 0 {
		return true
	}
	fn := big.NewInt(0).Exp(big.NewInt(2), n, nil)
	fn.Exp(big.NewInt(2), fn, nil)
	fn.Add(fn, big.NewInt(1))
	fn_1 := big.NewInt(0).Sub(fn, big.NewInt(1))
	fn_1_r := big.NewInt(0).Rsh(fn_1, 1)
	r1 := big.NewInt(0).Exp(big.NewInt(3), fn_1_r, fn)
	if r1.Cmp(fn_1) == 0 {
		fmt.Println(n, fn)
	}
	return r1.Cmp(fn_1) == 0
}

func main() {
	if true {
		i := 1
		for n := big.NewInt(4); ; n.Lsh(n, 1) {
			i++
			if big.NewInt(0).Add(n, big.NewInt(-9)).ProbablyPrime(0) {
				fmt.Println(i, "是素数")
			}
		}
	}
	return
	for n := big.NewInt(1); n.Cmp(big.NewInt(100)) <= 0; n.Add(n, big.NewInt(1)) {
		isPrime := PePin(n)
		if isPrime {
			fmt.Printf("Fermat number with p=%d is prime\n", n)
		} else {
			fmt.Printf("Fermat number with p=%d is not prime\n", n)
		}
	}
}
