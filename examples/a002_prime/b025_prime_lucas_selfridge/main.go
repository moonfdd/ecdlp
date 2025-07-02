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
		rightLimit.SetString("10000", 10)
		for ; num.Cmp(rightLimit) <= 0; num.Add(num, big.NewInt(1)) {

			r := LucasSelfridge(num)
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

func LucasSelfridge(n *big.Int) bool {
	if n.Cmp(big.NewInt(2)) < 0 {
		return false
	}
	if n.Cmp(big.NewInt(2)) == 0 {
		return true
	}
	if n.Bit(0) == 0 {
		return false
	}

	// 判断是否是平方数
	sq := big.NewInt(0).Sqrt(n)
	sq.Mul(sq, sq)
	if sq.Cmp(n) == 0 {
		return false
	}

	//确定d。5,-7,9,-11,13,-15
	d := big.NewInt(5) //5,-7,9,-11,13,-15
	p := big.NewInt(1)
	q := big.NewInt(0)
	jacobi := big.NewInt(0)
	iSign := big.NewInt(1) //1,-1,1,-1,1,-1
	zD := big.NewInt(5)    //5,7,9,11,13,15
	ulGCD := big.NewInt(0)
	for {
		d.Mul(iSign, zD)
		iSign.Neg(iSign)
		ulGCD.GCD(nil, nil, n, zD)
		if ulGCD.Cmp(big.NewInt(1)) > 0 && n.Cmp(ulGCD) > 0 {
			return false
		}
		jacobi = ecdlp.Jacobi(d, n)
		if jacobi.Cmp(big.NewInt(-1)) == 0 {
			break
		}
		zD.Add(zD, big.NewInt(2))
	}
	q.Sub(p, d).Rsh(q, 2) //q=(1-d)/4

	//因为Jacobi是-1，所以n+1是偶数
	nPlus1 := big.NewInt(0)
	nPlus1.Add(n, big.NewInt(1))

	cc := ecdlp.LucasParam{p, q}
	u, _ := cc.GetUnAndVnMod(nPlus1, n)
	if u.Cmp(big.NewInt(0)) == 0 {
		return true
	} else {
		return false
	}

}
