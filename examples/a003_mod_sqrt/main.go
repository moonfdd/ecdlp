package main

import (
	"fmt"
	"math/big"
)

func main() {
	if true {
		p := big.NewInt(0)
		p.SetString("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEFFFFFC2F", 16)
		r3 := ModSqrt(big.NewInt(2), p)
		fmt.Println(r3)
		r3 = ModSqrt(big.NewInt(55), big.NewInt(103))
		fmt.Println(r3)
		r3 = ModSqrt(big.NewInt(186), big.NewInt(401))
		fmt.Println(r3)

	}
	if true {
		//系统自带的ModSqrt
		r3 := big.NewInt(0)
		p := big.NewInt(0)
		p.SetString("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEFFFFFC2F", 16)
		r3.ModSqrt(big.NewInt(2), p)
		fmt.Println(r3)
		r3.ModSqrt(big.NewInt(55), big.NewInt(103))
		fmt.Println(r3)
		r3.ModSqrt(big.NewInt(186), big.NewInt(401))
		fmt.Println(r3)
	}
	fmt.Println("")
}

// 判断模平方根是否存在
func IsModSqrt(c, p *big.Int) bool {
	t := big.NewInt(0).Add(p, big.NewInt(-1))
	t.Rsh(t, 1)
	return big.NewInt(0).Exp(c, t, p).Cmp(big.NewInt(1)) == 0
}

// 求模平方根
func ModSqrt(c, p *big.Int) (ans *big.Int) {
	ans = big.NewInt(0)
	if !IsModSqrt(c, p) {
		return
	}
	//存在
	t := big.NewInt(0)
	s := big.NewInt(0).Add(p, big.NewInt(-1))
	for big.NewInt(0).And(s, big.NewInt(1)).Cmp(big.NewInt(0)) == 0 {
		s.Rsh(s, 1)
		t.Add(t, big.NewInt(1))
	}
	if t.Cmp(big.NewInt(1)) == 0 {
		s.Add(s, big.NewInt(1))
		s.Rsh(s, 1)
		ans.Exp(c, s, p)
	} else if t.Cmp(big.NewInt(2)) >= 0 {
		x_ := big.NewInt(0).Exp(c, big.NewInt(0).Add(p, big.NewInt(-2)), p)
		n := big.NewInt(1)
		for IsModSqrt(n, p) {
			n.Add(n, big.NewInt(1))
		}
		b := big.NewInt(0).Exp(n, s, p)
		s.Add(s, big.NewInt(1))
		s.Rsh(s, 1)
		ans.Exp(c, s, p)
		t_ := big.NewInt(0)
		for t.Cmp(big.NewInt(1)) > 0 {
			aa := big.NewInt(0).Mul(x_, ans)
			aa.Mod(aa, p)
			aa.Mul(aa, ans)
			aa.Mod(aa, p)
			bb := big.NewInt(0).Exp(big.NewInt(2), big.NewInt(0).Add(t, big.NewInt(-2)), p)
			if big.NewInt(0).Exp(aa, bb, p).Cmp(big.NewInt(1)) != 0 {
				tt := big.NewInt(0).Exp(big.NewInt(2), t_, p)
				tt.Exp(b, tt, p)
				ans.Mul(ans, tt)
				ans.Mod(ans, p)
			}
			t.Add(t, big.NewInt(-1))
			t_.Add(t_, big.NewInt(1))
		}
	}

	return
}
