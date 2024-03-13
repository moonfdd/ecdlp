package main

import (
	"fmt"
	"math/big"
)

func Proth(k, n int64) bool {
	num := new(big.Int).Add(new(big.Int).Mul(big.NewInt(k), new(big.Int).Exp(big.NewInt(2), big.NewInt(n), nil)), big.NewInt(1)) //k*2^n+1
	num_1 := big.NewInt(0).Sub(num, big.NewInt(1))                                                                               //k*2^n
	num_1_div_2 := big.NewInt(0).Div(num_1, big.NewInt(2))                                                                       //k*2^(n-1)
	t := big.NewInt(0).Exp(big.NewInt(3), num_1_div_2, num)
	t.Add(t, big.NewInt(1)).Mod(t, num)
	if t.Cmp(big.NewInt(0)) == 0 {
		return true
	}
	return false
}

func main() {

	if true {
		primeCount := 0
		for n := int64(2); n < 7; n++ {
			for k := int64(1); k <= 1<<n+1; k += 1 {
				if k%3 == 0 {
					continue
				}
				// fmt.Println("开始", k, n)
				r1 := Proth(k, n)
				aa := new(big.Int).Add(new(big.Int).Mul(big.NewInt(k), new(big.Int).Exp(big.NewInt(2), big.NewInt(n), nil)), big.NewInt(1))
				r2 := aa.ProbablyPrime(0)
				if r2 == r1 {
					//fmt.Printf("正确The number N = %d * 2^%d + 1 is prime: %v %v %v\n", k, n, isPrime, rr, aa)
				} else {
					fmt.Printf("错误The number N = %d * 2^%d + 1 is prime: %v %v %v\n", k, n, r1, r2, aa)
					return
				}

				if r1 {
					primeCount++
					fmt.Printf("The number N = %d * 2^%d + 1 is prime: %v %v\n", k, n, r1, aa)
				}
			}
		}
		fmt.Println("完全正确", primeCount)
	}
	fmt.Println("")
}
