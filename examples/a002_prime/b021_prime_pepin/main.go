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
	fn.Add(fn, big.NewInt(1)) //fn是费马数
	fn_1 := big.NewInt(0).Sub(fn, big.NewInt(1))
	fn_1_r := big.NewInt(0).Rsh(fn_1, 1)
	r1 := big.NewInt(0).Exp(big.NewInt(3), fn_1_r, fn)
	r1.Add(r1, big.NewInt(1)).Mod(r1, fn)
	if r1.Cmp(big.NewInt(0)) == 0 {
		fmt.Println(n, fn)
	}
	return r1.Cmp(big.NewInt(0)) == 0
}
func main() {
	for n := big.NewInt(1); n.Cmp(big.NewInt(10000)) <= 0; n.Add(n, big.NewInt(1)) {
		r1 := PePin(n)
		fn := big.NewInt(0).Exp(big.NewInt(2), n, nil)
		fn.Exp(big.NewInt(2), fn, nil)
		fn.Add(fn, big.NewInt(1)) //fn是费马数
		r2 := fn.ProbablyPrime(0)
		if r1 == r2 {
			if r1 {
				fmt.Printf("Fermat number with p=%d is prime\n", n)
			} else {
				// fmt.Printf("Fermat number with p=%d is not prime\n", n)
			}
		} else {
			fmt.Println("error")
			return
		}
	}
}
