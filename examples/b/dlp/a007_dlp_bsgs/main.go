package main

import (
	"fmt"
	"math/big"
)

// https://zhuanlan.zhihu.com/p/132603308
// https://codeleading.com/article/14084397800/
func main() {
	r := Bsgs(big.NewInt(2), big.NewInt(4), big.NewInt(29))
	fmt.Println(r)
	fmt.Println("")
}

func Bsgs(a, b, n *big.Int) (ans *big.Int) {
	ans = big.NewInt(-1)
	m := big.NewInt(1)
	v := big.NewInt(1)
	e := big.NewInt(1)
	m.Sqrt(n)
	m.Add(m, big.NewInt(1))
	v.Exp(a, m, n)
	v.Exp(v, big.NewInt(0).Add(n, big.NewInt(-2)), n)
	x := make(map[string]*big.Int)
	x["1"] = big.NewInt(0)
	for i := big.NewInt(1); i.Cmp(m) < 0; i.Add(i, big.NewInt(1)) {
		e.Mul(e, a)
		e.Mod(e, n)
		if _, ok := x[e.Text(16)]; !ok {
			x[e.Text(16)] = big.NewInt(0).Add(i, big.NewInt(0))
		}
	}
	for i := big.NewInt(0); i.Cmp(m) < 0; i.Add(i, big.NewInt(1)) {
		if _, ok := x[b.Text(16)]; ok {
			i.Mul(i, m)
			ans.Add(i, x[b.Text(16)])
			return
		}
		b.Mul(b, v)
		b.Mod(b, n)
	}
	return
}
