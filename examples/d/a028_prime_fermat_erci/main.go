package main

import (
	"fmt"
	"math/big"

	"github.com/moonfdd/ecdlp"
)

func main() {

	if false {
		//(1+d)(1+d)
		an, bn := exp(big.NewInt(1), big.NewInt(1), big.NewInt(2), big.NewInt(2), big.NewInt(5))
		fmt.Println(an, bn)
		an, bn = mul(big.NewInt(1), big.NewInt(1), big.NewInt(1), big.NewInt(1), big.NewInt(2), big.NewInt(5))
		fmt.Println(an, bn)
		return
	}
	if true {
		num := big.NewInt(0)
		num.SetString("2", 10)
		count := 0
		rightLimit := big.NewInt(0)
		rightLimit.SetString("100000000", 10)
		for ; num.Cmp(rightLimit) <= 0; num.Add(num, big.NewInt(1)) {

			r := Fermat(num)
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
}

// https://zhuanlan.zhihu.com/p/710321156 定理2.1 二次域上的费马小定理
// quadratic field
// (a+b*sqrt(d))^n = a+(d/n)*b*sqrt(d) mod n
func Fermat(n *big.Int) bool {
	if n.Cmp(big.NewInt(2)) < 0 {
		return false
	}
	if n.Cmp(big.NewInt(2)) == 0 {
		return true
	}
	if n.Bit(0) == 0 {
		return false
	}
	sq := big.NewInt(0).Sqrt(n)
	sq.Mul(sq, sq)
	if sq.Cmp(n) == 0 {
		return false
	}
	a := big.NewInt(2)
	b := big.NewInt(3)
	d := big.NewInt(1)
	jacobi := ecdlp.Jacobi(d, n)
	for {
		jacobi = ecdlp.Jacobi(d, n)
		if jacobi.Cmp(big.NewInt(-1)) == 0 {
			sq = big.NewInt(0).Sqrt(d)
			sq.Mul(sq, sq)
			if sq.Cmp(d) != 0 {
				break
			}
		}
		d.Add(d, big.NewInt(1))
	}
	an, bn := exp(a, b, n, d, n)
	an2 := big.NewInt(0).Set(a)
	an2.Mod(an2, n)
	bn2 := big.NewInt(0).Set(b)
	bn2.Mul(bn2, jacobi)
	bn2.Mod(bn2, n)
	if an.Cmp(an2) == 0 && bn.Cmp(bn2) == 0 {
		return true
	} else {
		return false
	}
}

// (a+b*sqrt(d))^k mod n
func exp(a, b, k, d, n *big.Int) (ak, bk *big.Int) {
	ak = big.NewInt(1)
	bk = big.NewInt(0)
	if k.Cmp(big.NewInt(0)) == 0 {
		return
	}
	at := big.NewInt(0).Set(a)
	bt := big.NewInt(0).Set(b)
	for i := 0; i < k.BitLen(); i++ {
		if k.Bit(i) != 0 {
			// fmt.Println("+", i, ak, bk, at, bt)
			ak, bk = mul(ak, bk, at, bt, d, n)
		}

		at, bt = mul(at, bt, at, bt, d, n)
		// fmt.Println("at bt = ", at, bt, a, b)
	}
	return
}

// (a1+b1*sqrt(d))*(a2+b2*sqrt(d)) mod n
func mul(a1, b1, a2, b2, d, n *big.Int) (a1a2, b1b2 *big.Int) {
	a1a2 = big.NewInt(0)
	temp := big.NewInt(0)
	temp.Mul(b1, b2)
	temp.Mul(temp, d)
	a1a2.Mul(a1, a2).Add(a1a2, temp) //a1*a2+b1*b2*d
	a1a2.Mod(a1a2, n)

	b1b2 = big.NewInt(0)
	b1b2.Mul(a1, b2)
	temp.Mul(a2, b1)
	b1b2.Add(b1b2, temp) //a1*b2+a2*b1
	b1b2.Mod(b1b2, n)

	return
}
