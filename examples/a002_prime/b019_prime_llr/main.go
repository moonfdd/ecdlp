package main

import (
	"fmt"
	"math/big"

	"github.com/moonfdd/ecdlp"
)

func Llr(k, n int64) bool {
	N := new(big.Int).Sub(new(big.Int).Mul(big.NewInt(k), new(big.Int).Exp(big.NewInt(2), big.NewInt(n), nil)), big.NewInt(1))
	// N = k * 2^n - 1
	//新增的部分
	if true {
		kk := k % 6
		if (kk == 1 && n&1 == 0) || (kk == 5 && n&1 == 1) {
			if N.Cmp(big.NewInt(3)) == 0 {
				return true
			}
			return false
		}
	}

	s := big.NewInt(4)

	//新增的部分
	if k != 1 {
		P := big.NewInt(0)
		for {
			if ecdlp.Jacobi(big.NewInt(0).Sub(P, big.NewInt(2)), N).Cmp(big.NewInt(1)) == 0 && ecdlp.Jacobi(big.NewInt(0).Add(P, big.NewInt(2)), N).Cmp(big.NewInt(-1)) == 0 {
				break
			}
			P.Add(P, big.NewInt(1))
		}

		ll := &ecdlp.LucasParam{P, big.NewInt(1)}
		_, s = ll.GetUnAndVnMod(big.NewInt(k), N)
	}

	// Repeat for n-2 iterations.
	for i := int64(0); i < n-2; i++ {
		// s = (s^2 - 2) % N
		s.Exp(s, big.NewInt(2), N).Sub(s, big.NewInt(2)).Mod(s, N)
	}

	// If we reach here, N passed the test - it might be prime.
	return s.Cmp(big.NewInt(0)) == 0
}

func main() {
	if true {
		k := int64(3)
		n := int64(20)
		primeCount := 0
		for n = 1; n < 7; n++ {
			for k = 1; k < 1<<n; k += 2 {
				isPrime := Llr(k, n)
				aa := new(big.Int).Sub(new(big.Int).Mul(big.NewInt(k), new(big.Int).Exp(big.NewInt(2), big.NewInt(n), nil)), big.NewInt(1))
				rr := aa.ProbablyPrime(0)
				if rr == isPrime {
					//fmt.Printf("正确The number N = %d * 2^%d - 1 is prime: %v %v %v\n", k, n, isPrime, rr, aa)
				} else {
					fmt.Printf("错误The number N = %d * 2^%d - 1 is prime: %v %v %v\n", k, n, isPrime, rr, aa)
					return
				}
				if isPrime {
					primeCount++
					fmt.Printf("The number N = %d * 2^%d - 1 is prime: %v %v\n", k, n, isPrime, aa)
				}
			}
		}
		fmt.Println("完全正确", primeCount)
	}
}
