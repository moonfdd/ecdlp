package main

import (
	"fmt"
	"math/big"

	"github.com/moonfdd/ecdlp"
)

func mpz_lucas_prp1(n, p, q *big.Int) bool {
	index := big.NewInt(0)
	j := big.NewInt(0)
	index.Set(n)
	// //    long int d = p*p - 4*q;
	d := big.NewInt(0)
	d.Mul(p, p)
	d.Sub(d, big.NewInt(0).Mul(q, big.NewInt(4)))

	if d.Cmp(big.NewInt(0)) == 0 {
		panic("Does not produce a proper Lucas sequence")
	}
	if n.Cmp(big.NewInt(2)) < 0 {
		return false
	}
	if n.Cmp(big.NewInt(2)) == 0 {
		return true
	}
	if n.Bit(0) == 0 {
		return false
	}

	// if n.Cmp(big.NewInt(5)) > 0 && big.NewInt(0).Mod(n, big.NewInt(5)).Cmp(big.NewInt(0)) == 0 {
	// 	return false
	// }

	j.Set(ecdlp.Jacobi(d, n)).Neg(j)
	// j.Neg(j)
	index.Add(index, j)

	ll := &ecdlp.LucasParam{p, q}
	u, _ := ll.GetUnAndVnMod(index, n)
	if u.Cmp(big.NewInt(0)) == 0 {
		return true
	} else {
		return false
	}

}

func mpz_lucas_prp(n, p, q *big.Int) bool {
	if n.Cmp(big.NewInt(2)) < 0 {
		return false
	}
	if n.Cmp(big.NewInt(2)) == 0 {
		return true
	}
	if n.Bit(0) == 0 {
		return false
	}
	ll := &ecdlp.LucasParam{p, q}
	u, _ := ll.GetUnAndVnMod(n, n)
	j := ecdlp.Jacobi(big.NewInt(2), n)
	j.Neg(j)
	u.Add(u, j).Mod(u, n)
	if u.Cmp(big.NewInt(0)) == 0 {
		return true
	} else {
		return false
	}
}
func mpz_lucas_prp5(n, p, q *big.Int) bool {
	// index := big.NewInt(0)
	// index.Set(n)
	// //    long int d = p*p - 4*q;
	// d := big.NewInt(0)
	// d.Mul(p, p)
	// d.Sub(d, big.NewInt(0).Mul(q, big.NewInt(4)))

	// if d.Cmp(big.NewInt(0)) == 0 {
	// 	panic("Does not produce a proper Lucas sequence")
	// }
	if n.Cmp(big.NewInt(2)) < 0 {
		return false
	}
	if n.Cmp(big.NewInt(2)) == 0 {
		return true
	}
	if n.Bit(0) == 0 {
		return false
	}

	ll := &ecdlp.LucasParam{p, q}
	_, v := ll.GetUnAndVnMod(n, n)
	pmodn := big.NewInt(0).Mod(p, n)
	if v.Cmp(pmodn) == 0 {
		return true
	} else {
		return false
	}

}
func main() {
	if true {
		errCount := 0
		for n := big.NewInt(2); n.Cmp(big.NewInt(100000)) <= 0; n.Add(n, big.NewInt(1)) {
			r := mpz_lucas_prp1(n, big.NewInt(2), big.NewInt(-1))
			// r := mpz_lucas_prp(n, big.NewInt(2), big.NewInt(-1))
			// r := mpz_lucas_prp5(n, big.NewInt(2), big.NewInt(-1))
			r2 := n.ProbablyPrime(0)
			if r != r2 {
				errCount++
				fmt.Println("错误", n, r, r2)
			} else {
				if r {
					//fmt.Println("素数", n)
				}
			}
		}
		fmt.Println("错误次数", errCount)
	}
	fmt.Println("")
}
