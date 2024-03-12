package main

import (
	"fmt"
	"math/big"
)

func lucasLehmer(p int) bool {
	m := big.NewInt(0).Exp(big.NewInt(2), big.NewInt(int64(p)), nil)
	m.Sub(m, big.NewInt(1))
	s := big.NewInt(4)

	for i := 0; i < p-2; i++ {
		s.Mul(s, s)
		s.Sub(s, big.NewInt(2))
		s.Mod(s, m)
		if s.Cmp(big.NewInt(0)) == 0 {
			return true
		}
	}

	return s.Cmp(big.NewInt(0)) == 0
}

func main() {
	if true {
		for p := 3; p < 10000; p++ {
			isPrime := lucasLehmer(int(p))
			aa := big.NewInt(0).Exp(big.NewInt(2), big.NewInt(int64(p)), nil)
			aa.Sub(aa, big.NewInt(1))
			rr := aa.ProbablyPrime(0)
			if rr == isPrime {
				if isPrime {
					fmt.Printf("数字 N = 2^%d - 1 是素数: %v\n", p, isPrime)
				}
			} else {
				fmt.Println("错误", p, isPrime)
				return
			}

		}
	}
}
