package main

import (
	"fmt"
	"math/big"
)

// https://zhuanlan.zhihu.com/p/132603308
// https://codeleading.com/article/14084397800/
func main() {
	if false {
		a := big.NewInt(2)
		b := big.NewInt(0)
		p := big.NewInt(8)

		r := Bsgs(a, b, p)
		fmt.Println(b, r)
		r = ExBsgs(a, b, p)
		fmt.Println(b, r)
		fmt.Println("----------")

	}
	if true {
		a := big.NewInt(2)
		p := big.NewInt(29)
		for b := big.NewInt(0); b.Cmp(p) < 0; b.Add(b, big.NewInt(1)) {
			r := Bsgs(a, b, p)
			fmt.Println(b, r)
			r2 := ExBsgs(a, b, p)
			fmt.Println(b, r)
			if r.Cmp(r2) != 0 {
				fmt.Println("error")
				return
			}
			fmt.Println("----------")
		}
	}
}

// https://en.oi-wiki.org/math/bsgs/
func Bsgs(a, b, p *big.Int) (ans *big.Int) { //a^(im)=b*a^j mod p
	m := big.NewInt(0)
	m.Sqrt(p).Add(m, big.NewInt(1))
	aExpM := big.NewInt(0).Exp(a, m, p)
	left := big.NewInt(1)
	mapRightJ := make(map[string]*big.Int)
	right := big.NewInt(0).Set(b)
	for j := big.NewInt(1); j.Cmp(m) <= 0; j.Add(j, big.NewInt(1)) { //1<=j<=m
		right.Mul(right, a).Mod(right, p)
		mapRightJ[right.Text(10)] = big.NewInt(0).Set(j)
	}

	for i := big.NewInt(1); i.Cmp(m) <= 0; i.Add(i, big.NewInt(1)) { //1<=i<=m
		left.Mul(left, aExpM).Mod(left, p)
		if j, ok := mapRightJ[left.Text(10)]; ok {
			ans = big.NewInt(0)
			ans.Mul(i, m).Sub(ans, j).Mod(ans, p)
			return
		}
	}

	return
}

// https://www.bilibili.com/video/BV1GR4y1X7Mc
// 08:32
func ExBsgs(a, b, m *big.Int) (ans *big.Int) {
	if b.Cmp(big.NewInt(1)) == 0 {
		ans = big.NewInt(0)
		// fmt.Println("abm ", a, b, m)
		return
	}
	x := big.NewInt(0)
	y := big.NewInt(0)
	d := big.NewInt(0).GCD(x, y, a, m)
	if d.Cmp(big.NewInt(1)) == 0 {
		ans = Bsgs(a, b, m)
		return
	}
	if big.NewInt(0).Mod(b, d).Cmp(big.NewInt(0)) != 0 {
		return
	}
	mm := big.NewInt(0)
	mm.Div(m, d)
	bb := big.NewInt(0)
	bb.Div(b, d).Mul(bb, x).Mod(bb, mm)
	ans = ExBsgs(a, bb, mm)
	if ans != nil {
		ans.Add(ans, big.NewInt(1))
	}
	return
}
