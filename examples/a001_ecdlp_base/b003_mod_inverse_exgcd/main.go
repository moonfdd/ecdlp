package main

import (
	"fmt"
	"math/big"
)

func main() {
	if true {
		r := ModInverseRecursion(big.NewInt(89), big.NewInt(99)) //求逆 1/89%99
		fmt.Println(r)
	}
	if true {
		r := big.NewInt(0)
		r.ModInverse(big.NewInt(89), big.NewInt(99)) //求逆 1/89%99
		fmt.Println(r)
	}
	fmt.Println("")
}

// 求1/a%m
// a和m都是是非零整数
func ModInverseRecursion(a, m *big.Int) (ans *big.Int) {
	ans = big.NewInt(0)
	x := big.NewInt(0)
	y := big.NewInt(0)
	g := ExGCD(x, y, a, m)
	if g.Cmp(big.NewInt(1)) != 0 {
		return
	}
	ans.Mod(x, m)
	return
}

// 扩展欧几里得算法
func ExGCD(x, y, a, b *big.Int) (ans *big.Int) {
	ans = big.NewInt(0)
	if b.Cmp(big.NewInt(0)) == 0 {
		x.Set(big.NewInt(1))
		y.Set(big.NewInt(0))
		ans.Set(a)
	} else {
		ans.Set(ExGCD(x, y, b, big.NewInt(0).Mod(a, b)))
		t := big.NewInt(0)
		t.Div(a, b).Mul(t, y).Neg(t).Add(t, x)
		x.Set(y)
		y.Set(t)
	}
	return
}
