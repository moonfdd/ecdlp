package main

import (
	"fmt"
	"math/big"

	"github.com/moonfdd/ecdlp"
)

func Qft(n, p, q *big.Int) bool {
	if n.Cmp(big.NewInt(2)) < 0 {
		return false
	}
	if n.Cmp(big.NewInt(2)) == 0 {
		return true
	}
	if n.Bit(0) == 0 {
		return false
	}
	d := big.NewInt(0)
	d.Mul(p, p)
	d.Sub(d, big.NewInt(0).Mul(q, big.NewInt(4)))
	sq:=big.NewInt(0)
	sq.Sqrt(d)
	sq.Mul(sq, sq)
	if sq.Cmp(d) == 0 {
		panic("d can't square number")
	}

	t := big.NewInt(0)
	t.Mul(q, d).Lsh(t, 1) //t = 2qd
	t2 := big.NewInt(0).Set(t)
	t.GCD(nil, nil, n, t)
	//1
	if t.Cmp(big.NewInt(1)) != 0 {
		if n.Cmp(t) > 0 {
			return false
		}
		if t.Cmp(n) >= 0 {
			fmt.Println(n, "可能", n.ProbablyPrime(0), p, q, d, t2)
			return n.ProbablyPrime(0)
		}
		return false
	}

	j := ecdlp.Jacobi(d, n)
	polynomial := ecdlp.PolynomialExpMod([]*big.Int{big.NewInt(1), big.NewInt(0)}, n, []*big.Int{big.NewInt(1), big.NewInt(0).Neg(p), big.NewInt(0).Set(q)}, n) //x^n mod x^2-px+q,n
	if j.Cmp(big.NewInt(1)) == 0 {
		polynomial = ecdlp.PolynomialAdd(polynomial, []*big.Int{big.NewInt(-1), big.NewInt(0)}, n) //x^n-x
	} else if j.Cmp(big.NewInt(-1)) == 0 {
		polynomial = ecdlp.PolynomialAdd(polynomial, []*big.Int{big.NewInt(1), big.NewInt(0).Neg(p)}, n) //x^n+x-p
	} else {
		res := n.ProbablyPrime(0)
		fmt.Println(n, "可能2", res)
		return res
	}
	if len(polynomial) == 1 && polynomial[0].Cmp(big.NewInt(0)) == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	if true {
		if true {
			num := big.NewInt(0)
			num.SetString("2", 10)
			count := 0
			rightLimit := big.NewInt(0)
			rightLimit.SetString("1000000", 10)
			for ; num.Cmp(rightLimit) <= 0; num.Add(num, big.NewInt(1)) {
				// r := Qft(num, big.NewInt(1), big.NewInt(-1))
				// r := Qft(num, big.NewInt(3), big.NewInt(-1))
				r := Qft(num, big.NewInt(3), big.NewInt(-5))
				r2 := num.ProbablyPrime(0)

				if r == r2 {
					if r {
						// fmt.Println(num, "是素数")
					}
				} else {
					fmt.Println("测试失败", r, r2, num)
					count++
					// return
				}

			}
			fmt.Println("失败次数", count)
			return
		}
		return
	}
}
