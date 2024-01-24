package main

import (
	"fmt"
	"math/big"
)

// https://www.docin.com/p-32645122.html

func main() {
	// 测试ModSqrt
	if true {

		p := big.NewInt(0)
		p.SetString("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEFFFFFC2F", 16)
		var r []*big.Int
		r = ModSqrt(big.NewInt(2), p)
		fmt.Println(r)
		r = ModSqrt(big.NewInt(55), big.NewInt(103))
		fmt.Println(r)
		r = ModSqrt(big.NewInt(186), big.NewInt(401))
		fmt.Println(r)

	}
	// 测试系统自带的ModSqrt
	if true {
		r := big.NewInt(0)
		p := big.NewInt(0)
		p.SetString("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEFFFFFC2F", 16)
		r.ModSqrt(big.NewInt(2), p)
		fmt.Println(r)
		r.ModSqrt(big.NewInt(55), big.NewInt(103))
		fmt.Println(r)
		r.ModSqrt(big.NewInt(186), big.NewInt(401))
		fmt.Println(r)
	}
	fmt.Println("")
}

// 求模平方根的个数
func ModSqrtCount(a, p *big.Int) int {
	t := big.NewInt(0).Add(p, big.NewInt(-1)) //t=(p-1)/2
	t.Rsh(t, 1)
	if big.NewInt(0).Exp(a, t, p).Cmp(big.NewInt(1)) == 0 {
		return 2
	} else {
		return 0
	}
}

// Tonelli–Shanks算法
// 求模平方根(lnp)^2
func ModSqrt(a, p *big.Int) (ans []*big.Int) {
	ans = make([]*big.Int, 0)
	x := big.NewInt(0)

	//a==0
	if a.Cmp(big.NewInt(0)) == 0 {
		ans = append(ans, x)
		return
	}

	//欧拉判别法
	if ModSqrtCount(a, p) == 0 {
		return
	}

	//存在模平方根，拆解成s和t
	//p-1=s*(2^t)  s是奇数
	t := 0
	s := big.NewInt(0).Add(p, big.NewInt(-1))
	for big.NewInt(0).And(s, big.NewInt(1)).Cmp(big.NewInt(0)) == 0 {
		s.Rsh(s, 1)
		t++
	}

	// if t == 1 { //这个if语句可以不写
	// 	///p==3 mod 4
	// 	// a^((p+1)/4)
	// 	s.Add(s, big.NewInt(1)).Rsh(s, 1)
	// 	x.Exp(a, s, p)
	// 	otherX := big.NewInt(0)
	// 	otherX.Neg(x).Add(p, otherX)
	// 	ans = append(ans, x, otherX)
	// 	return
	// }

	//找二次非剩余c
	c := big.NewInt(2)
	for ModSqrtCount(c, p) != 0 {
		c.Add(c, big.NewInt(1))
	}
	ce := big.NewInt(0).Exp(c, s, p)
	s.Add(s, big.NewInt(1)).Rsh(s, 1)
	x.Exp(a, s, p)
	for ti := t; ti >= 2; ti-- {
		aa := big.NewInt(0)
		aa.Mul(big.NewInt(0).ModInverse(a, p), x).Mod(aa, p).Mul(aa, x).Mod(aa, p)
		bb := big.NewInt(0).Exp(big.NewInt(2), big.NewInt(int64(ti-2)), nil)
		if big.NewInt(0).Exp(aa, bb, p).Cmp(big.NewInt(1)) != 0 {
			tie := big.NewInt(0).Exp(big.NewInt(2), big.NewInt(int64(t-ti)), nil)
			tie.Exp(ce, tie, p)
			x.Mul(x, tie).Mod(x, p)
		}
	}
	otherX := big.NewInt(0)
	otherX.Neg(x).Add(p, otherX)
	ans = append(ans, x, otherX)

	return
}
