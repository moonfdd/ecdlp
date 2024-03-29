package main

import (
	"fmt"
	"math/big"
)

func main() {

	if false {
		p := big.NewInt(0)
		p.SetString("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEFFFFFC2F", 16)
		for a := big.NewInt(30000); a.Cmp(big.NewInt(30300)) <= 0; a.Add(a, big.NewInt(1)) {
			fmt.Println("a = ", a, "-------------")
			r := ModCbrt(a, p)
			fmt.Println("答案：", r)
			for i := 0; i < len(r); i++ {
				if big.NewInt(0).Exp(r[i], big.NewInt(3), p).Cmp(a) == 0 {

				} else {
					fmt.Println("答案错误", r[i], "a = ", big.NewInt(0).Exp(r[i], big.NewInt(3), p))
					return
				}
			}
		}
		return
	}
	if false {
		p := big.NewInt(0)
		p.SetString("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEBAAEDCE6AF48A03BBFD25E8CD0364141", 16)
		for a := big.NewInt(30000); a.Cmp(big.NewInt(30300)) <= 0; a.Add(a, big.NewInt(1)) {
			fmt.Println("a = ", a, "-------------")
			r := ModCbrt(a, p)
			fmt.Println("答案：", r)
			for i := 0; i < len(r); i++ {
				if big.NewInt(0).Exp(r[i], big.NewInt(3), p).Cmp(a) == 0 {

				} else {
					fmt.Println("答案错误", r[i], "a = ", big.NewInt(0).Exp(r[i], big.NewInt(3), p))
					return
				}
			}
		}
		return
	}
	if false {
		p := big.NewInt(11)
		for a := big.NewInt(1); a.Cmp(big.NewInt(0).Add(p, big.NewInt(-1))) <= 0; a.Add(a, big.NewInt(1)) {
			fmt.Println("a = ", a, "-------------")
			r := ModCbrt(a, p)
			fmt.Println("答案：", r)
			for i := 0; i < len(r); i++ {
				if big.NewInt(0).Exp(r[i], big.NewInt(3), p).Cmp(a) == 0 {

				} else {
					fmt.Println("答案错误", r[i], "a = ", big.NewInt(0).Exp(r[i], big.NewInt(3), p))
					return
				}
			}
		}
		return
	}

	if true {
		p := big.NewInt(997)
		for a := big.NewInt(1); a.Cmp(big.NewInt(0).Add(p, big.NewInt(-1))) <= 0; a.Add(a, big.NewInt(1)) {
			fmt.Println("a = ", a, "-------------")
			r := ModCbrt(a, p)
			fmt.Println("答案：", r)
			for i := 0; i < len(r); i++ {
				if big.NewInt(0).Exp(r[i], big.NewInt(3), p).Cmp(a) == 0 {

				} else {
					fmt.Println("答案错误", r[i], "a = ", big.NewInt(0).Exp(r[i], big.NewInt(3), p))
					return
				}
			}
		}
	}
	fmt.Println("")
}

// 求模立方根的个数0，1，3
func ModCbrtCount(c, p *big.Int) int {
	t := big.NewInt(0)
	t.Add(p, big.NewInt(-2))
	t.Mod(t, big.NewInt(3))
	if t.Cmp(big.NewInt(0)) == 0 {
		return 1
	}
	t = big.NewInt(0).Add(p, big.NewInt(-1))
	t.Div(t, big.NewInt(3))
	if big.NewInt(0).Exp(c, t, p).Cmp(big.NewInt(1)) == 0 {
		return 3
	} else {
		return 0
	}
}

// https://loj.ac/s/1752475
// https://eprint.iacr.org/2013/024.pdf
// Peralta Method
func ModCbrt(a, p *big.Int) (ans []*big.Int) {
	ans = make([]*big.Int, 0)
	if a.Cmp(big.NewInt(0)) == 0 {
		ans = append(ans, big.NewInt(0))
		return
	}
	count := ModCbrtCount(a, p)
	if count == 1 { //有1个解
		t := big.NewInt(0).Lsh(p, 1) //t=(2p-1)/3
		t.Mod(t, p)
		t = t.Add(t, big.NewInt(-1))
		t.Mod(t, p)
		t.Mul(t, big.NewInt(0).ModInverse(big.NewInt(3), p))
		t.Mod(t, p)
		ans = append(ans, big.NewInt(0).Exp(a, t, p))
	} else if count == 3 { //有3个解，Peralta Method算法

		//w=i^3=b^3-a
		var b *big.Int
		w := big.NewInt(0)
		for b = big.NewInt(1); b.Cmp(p) < 0; b.Add(b, big.NewInt(1)) {
			w.Exp(b, big.NewInt(3), p)
			w.Add(w, big.NewInt(0).Neg(a))
			w.Mod(w, p)
			if w.Cmp(big.NewInt(0)) != 0 && ModCbrtCount(w, p) == 0 {
				break
			}
		}

		iRoot := Ring{b, big.NewInt(-1), big.NewInt(0), w}
		pp := big.NewInt(0).Mul(p, p) // pp = (p*p+p+1)/3
		pp.Add(pp, p)
		pp.Add(pp, big.NewInt(1))
		pp.Div(pp, big.NewInt(3))
		ansr := powerModI(iRoot, pp, p)
		x := ansr.a //根x是实部

		//求周期cycle
		cycle := big.NewInt(0)
		p3 := big.NewInt(0).Add(p, big.NewInt(-1)) //(p-1)/3
		p3.Mul(p3, big.NewInt(0).ModInverse(big.NewInt(3), p))
		p3.Mod(p3, p)
		for k := big.NewInt(1); k.Cmp(p) < 0; k.Add(k, big.NewInt(1)) {
			cycle.Exp(k, p3, p)
			if cycle.Cmp(big.NewInt(1)) != 0 {
				break
			}
		}

		otherX := big.NewInt(0)
		otherX.Mul(x, cycle) //另一个根是x*cycle
		otherX.Mod(otherX, p)

		otherX2 := big.NewInt(0)
		otherX2.Mul(otherX, cycle) //另一个根是x*cycle*cycle
		otherX2.Mod(otherX2, p)

		ans = append(ans, x, otherX, otherX2)
	}
	return
}

type Ring struct {
	a *big.Int //实部
	b *big.Int //i的虚部
	c *big.Int //i^2的虚部
	w *big.Int //i^3的值
}

// 复数乘法
func mulI(x Ring, y Ring, p *big.Int) Ring {
	var res Ring
	res.a = big.NewInt(0)
	res.b = big.NewInt(0)
	res.c = big.NewInt(0)
	res.w = x.w
	w := x.w

	a1 := big.NewInt(0)
	a2 := big.NewInt(0)
	a3 := big.NewInt(0)
	a1.Mul(x.a, y.a) //x.a*y.a
	a1.Mod(a1, p)
	a2.Mul(x.b, y.c) //x.b*y.c*w
	a2.Mod(a2, p)
	a2.Mul(a2, w)
	a2.Mod(a2, p)
	a3.Mul(x.c, y.b) //x.c*y.b*w
	a3.Mod(a3, p)
	a3.Mul(a3, w)
	a3.Mod(a3, p)
	res.a.Add(a1, a2)
	res.a.Mod(res.a, p)
	res.a.Add(res.a, a3)
	res.a.Mod(res.a, p)

	b1 := big.NewInt(0)
	b2 := big.NewInt(0)
	b3 := big.NewInt(0)
	b1.Mul(x.a, y.b) //x.a*y.b
	b1.Mod(b1, p)
	b2.Mul(x.b, y.a) //x.b*y.a
	b2.Mod(b2, p)
	b3.Mul(x.c, y.c) //x.c*y.c*w
	b3.Mod(b3, p)
	b3.Mul(b3, w)
	b3.Mod(b3, p)
	res.b.Add(b1, b2)
	res.b.Mod(res.b, p)
	res.b.Add(res.b, b3)
	res.b.Mod(res.b, p)

	c1 := big.NewInt(0)
	c2 := big.NewInt(0)
	c3 := big.NewInt(0)
	c1.Mul(x.a, y.c) //x.a*y.c
	c1.Mod(c1, p)
	c2.Mul(x.b, y.b) //x.b*y.b
	c2.Mod(c2, p)
	c3.Mul(x.c, y.a) //x.c*y.a
	c3.Mod(c3, p)
	res.c.Add(c1, c2)
	res.c.Mod(res.c, p)
	res.c.Add(res.c, c3)
	res.c.Mod(res.c, p)

	return res
}

// 复数快速幂，注意b不能取模
func powerModI(a Ring, b, p *big.Int) Ring {
	res := Ring{big.NewInt(1), big.NewInt(0), big.NewInt(0), a.w}
	b = big.NewInt(0).Set(b)
	for b.Cmp(big.NewInt(0)) != 0 {
		if big.NewInt(0).Mod(b, big.NewInt(2)).Cmp(big.NewInt(1)) == 0 {
			res = mulI(res, a, p)
		}
		a = mulI(a, a, p)
		b.Rsh(b, 1)
	}
	return res
}
