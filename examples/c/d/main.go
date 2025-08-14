// https://github.com/dvberkel/pocklington/blob/master/pocklington/criteria.py
package main

import (
	"fmt"
	"math/big"
)

type PrimalityTester struct{}

func (pt *PrimalityTester) isGermainPrime(q int) bool {
	p := 2*q + 1           // p satisfies q | p-1
	a := pt.candidateTo(p) // s satisfies a**(p-1) % p == 1 mod p

	return pt.gcd(a*a-1, p) == 1
}

func (pt *PrimalityTester) candidateTo(p int) int {
	a := 2
	for pt.powmod(a, p-1, p) != 1 {
		a++
	}
	return a
}

func (pt *PrimalityTester) powmod(a, exponent, modulus int) int {
	result, power := 1, a
	for exponent > 0 {
		if exponent%2 != 0 {
			result = (result * power) % modulus
		}
		exponent, power = exponent/2, (power*power)%modulus
	}
	return result
}

func (pt *PrimalityTester) gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	count := 0
	for q := 2; q < 100; q++ {
		pt := PrimalityTester{}
		isGermainPrime := pt.isGermainPrime(q)
		// isGermainPrime = !isGermainPrime
		r := big.NewInt(int64(q)).ProbablyPrime(0)
		if r != isGermainPrime {
			fmt.Println(q, isGermainPrime, r)
			count++
		}

	}
	fmt.Println("错误次数：", count)
}
