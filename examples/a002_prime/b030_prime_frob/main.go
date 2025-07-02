package main

import (
	"fmt"
	"math/big"

	"github.com/moonfdd/ecdlp"
)

// https://eprint.iacr.org/2008/124 Frobenius 伪素数检验的简单推导

// https://en.wikipedia.org/wiki/Frobenius_pseudoprime
func Frob(n, p, q *big.Int) bool {
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
	d.Sub(d, big.NewInt(0).Mul(q, big.NewInt(4))) //d = p^2 - 4q

	if d.Cmp(big.NewInt(0)) == 0 {
		panic("Does not produce a proper Lucas sequence")
	}
	t := big.NewInt(0)
	t.Mul(q, d).Lsh(t, 1) //t = 2qd
	t.GCD(nil, nil, n, big.NewInt(0).Abs(t))
	//1
	if t.Cmp(big.NewInt(1)) != 0 {
		if n.Cmp(t) > 0 {
			return false
		}
		if t.Cmp(n) >= 0 {
			fmt.Println(n, "可能", n.ProbablyPrime(0), p, q, d)
			return n.ProbablyPrime(0)
		}
		return false
	}

	index := big.NewInt(0).Set(n)
	ret := ecdlp.Jacobi(d, n)
	if ret.Cmp(big.NewInt(-1)) == 0 {
		index.Add(index, big.NewInt(1))
	} else if ret.Cmp(big.NewInt(1)) == 0 {
		index.Sub(index, big.NewInt(1))
	} else {
		panic("ret error")
	}

	ll := ecdlp.LucasParam{p, q}
	u, v := ll.GetUnAndVnMod(big.NewInt(0).Rsh(index, 0), n)

	//2.
	if u.Cmp(big.NewInt(0)) != 0 {
		return false
	}

	//3
	if true {
		vv := big.NewInt(0).Sub(big.NewInt(1), ret)
		vv.Rsh(vv, 1)
		vv.Exp(q, vv, n)
		vv.Lsh(vv, 1)
		vv.Neg(vv)
		vv.Add(vv, v)
		vv.Mod(vv, n)
		if vv.Cmp(big.NewInt(0)) != 0 {
			return false
		}
	}

	//3
	if false {
		u, v = ll.GetUnAndVnMod(n, n)
		vv := big.NewInt(0)
		vv.Sub(v, p).Mod(vv, n)
		if vv.Cmp(big.NewInt(0)) != 0 {
			return false
		}
	}

	return true
}

func main() {
	if true {
		if true {
			num := big.NewInt(0)
			num.SetString("2", 10)
			count := 0
			rightLimit := big.NewInt(0)
			rightLimit.SetString("2000000", 10)
			for ; num.Cmp(rightLimit) <= 0; num.Add(num, big.NewInt(1)) {
				// r := Frob(num, big.NewInt(1), big.NewInt(-1))
				r := Frob(num, big.NewInt(3), big.NewInt(-1))
				// r := Frob(num, big.NewInt(3), big.NewInt(-5))

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
