package main

import (
	"fmt"
	"math/big"

	"github.com/moonfdd/ecdlp"
)

func main() {
	if true {
		num := big.NewInt(0)
		num.SetString("2", 10)
		count := 0
		rightLimit := big.NewInt(0)
		rightLimit.SetString("100000000", 10)
		for ; num.Cmp(rightLimit) <= 0; num.Add(num, big.NewInt(1)) {
			if big.NewInt(0).Mod(num, big.NewInt(10000)).Cmp(big.NewInt(0)) == 0 {
				fmt.Println(num, ":", count)
			}
			r := QuadraticFieldFermat(num)
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

func QuadraticFieldFermat(n *big.Int) bool {
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
			break
		}
		d.Add(d, big.NewInt(1))
	}
	return quadraticFieldFermat(n, a, b, d)
}

func quadraticFieldFermat(n, a, b, d *big.Int) bool {
	f := ecdlp.FieldParam{M: 2, N: n, D: d}
	anbn := f.ExpMod([]*big.Int{a, b}, n)
	jacobi := ecdlp.Jacobi(d, n)
	anbn = f.SubMod(anbn, []*big.Int{a, big.NewInt(0).Mul(b, jacobi)})
	if anbn[0].Cmp(big.NewInt(0)) == 0 && anbn[1].Cmp(big.NewInt(0)) == 0 {
		return true
	} else {
		return false
	}

}
