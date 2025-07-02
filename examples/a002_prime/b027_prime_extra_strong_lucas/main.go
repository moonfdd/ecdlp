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
		rightLimit.SetString("100000", 10)
		for ; num.Cmp(rightLimit) <= 0; num.Add(num, big.NewInt(1)) {

			r := ExtraStrongLucas(num, big.NewInt(3))
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

func ExtraStrongLucas(n, pp *big.Int) bool {
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

	d := big.NewInt(0) // 计算得到
	p := big.NewInt(0) // 计算得到
	q := big.NewInt(1) // 固定值
	if pp.Cmp(big.NewInt(3)) < 0 {
		p.SetInt64(3)
	} else {
		p.Set(pp)
	}
	ulGCD := big.NewInt(0)
	for {
		d.Mul(p, p).Sub(d, big.NewInt(4))
		ulGCD.GCD(nil, nil, n, big.NewInt(0).Abs(d))
		if ulGCD.Cmp(big.NewInt(1)) == 0 {
			break
		}
		if n.Cmp(ulGCD) > 0 {
			return false
		}
		p.Add(p, big.NewInt(1))
	}

	//n+1或者n-1
	jacobi := ecdlp.Jacobi(d, n)
	nPlusOrMinus1 := big.NewInt(0)
	if jacobi.Cmp(big.NewInt(-1)) == 0 {
		nPlusOrMinus1.Add(n, big.NewInt(1))
	} else if jacobi.Cmp(big.NewInt(1)) == 0 {
		nPlusOrMinus1.Sub(n, big.NewInt(1))
	} else {
		panic("正常来说不应该执行到这里")
	}

	//拆分n+1
	t := 0
	s := big.NewInt(0).Set(nPlusOrMinus1)
	for big.NewInt(0).And(s, big.NewInt(1)).Cmp(big.NewInt(0)) == 0 {
		s.Rsh(s, 1)
		t++
	}

	cc := ecdlp.LucasParam{p, q}
	u, v := cc.GetUnAndVnMod(s, n)
	//  Vs==0
	if v.Cmp(big.NewInt(0)) == 0 {
		return true
	}
	// Us==0 && Vs==±2
	if u.Cmp(big.NewInt(0)) == 0 {
		if v.Cmp(big.NewInt(2)) == 0 {
			return true
		}
		if v.Cmp(big.NewInt(0).Sub(n, big.NewInt(2))) == 0 {
			return true
		}
	}

	// V2s==0
	for i := 1; i < t; i++ {
		u, v = cc.TwoUAndTwoVMod(u, v, n)
		if v.Cmp(big.NewInt(0)) == 0 {
			return true
		}
	}

	return false

}
