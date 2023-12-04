package main

import (
	"fmt"
	"math/big"
)

func main() {
	// 测试Cipolla
	if true {
		if false {
			p := big.NewInt(0)
			p.SetString("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEFFFFFC2F", 16)
			for c := big.NewInt(20000); c.Cmp(big.NewInt(30000)) <= 0; c.Add(c, big.NewInt(1)) {
				fmt.Println("c = ", c, "-------------")
				r := Cipolla(c, p)
				fmt.Println("答案：", r)
				for i := 0; i < len(r); i++ {
					if big.NewInt(0).Exp(r[i], big.NewInt(2), p).Cmp(c) == 0 {

					} else {
						fmt.Println("答案错误", r[i], "，c = ", big.NewInt(0).Exp(r[i], big.NewInt(2), p))
						return
					}
				}
			}
			return
		}
		if false {
			p := big.NewInt(103)
			for c := big.NewInt(1); c.Cmp(big.NewInt(0).Add(p, big.NewInt(-1))) <= 0; c.Add(c, big.NewInt(1)) {
				fmt.Println("c = ", c, "-------------")
				r := Cipolla(c, p)
				fmt.Println("答案：", r)
				for i := 0; i < len(r); i++ {
					if big.NewInt(0).Exp(r[i], big.NewInt(2), p).Cmp(c) == 0 {

					} else {
						fmt.Println("答案错误", r[i], "，c = ", big.NewInt(0).Exp(r[i], big.NewInt(2), p))
						return
					}
				}
			}
			return
		}
		if true {
			p := big.NewInt(401)
			for c := big.NewInt(1); c.Cmp(big.NewInt(0).Add(p, big.NewInt(-1))) <= 0; c.Add(c, big.NewInt(1)) {
				fmt.Println("c = ", c, "-------------")
				r := Cipolla(c, p)
				fmt.Println("答案：", r)
				for i := 0; i < len(r); i++ {
					if big.NewInt(0).Exp(r[i], big.NewInt(2), p).Cmp(c) == 0 {

					} else {
						fmt.Println("答案错误", r[i], "，c = ", big.NewInt(0).Exp(r[i], big.NewInt(2), p))
						return
					}
				}
			}
			return
		}
	}
	// 测试ModSqrt
	if false {
		if true {
			p := big.NewInt(0)
			p.SetString("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEFFFFFC2F", 16)
			for c := big.NewInt(20000); c.Cmp(big.NewInt(30000)) <= 0; c.Add(c, big.NewInt(1)) {
				fmt.Println("c = ", c, "-------------")
				r := ModSqrt(c, p)
				fmt.Println("答案：", r)
				for i := 0; i < len(r); i++ {
					if big.NewInt(0).Exp(r[i], big.NewInt(2), p).Cmp(c) == 0 {

					} else {
						fmt.Println("答案错误", r[i], "，c = ", big.NewInt(0).Exp(r[i], big.NewInt(2), p))
						return
					}
				}
			}
			return
		}
		if false {
			p := big.NewInt(103)
			for c := big.NewInt(1); c.Cmp(big.NewInt(0).Add(p, big.NewInt(-1))) <= 0; c.Add(c, big.NewInt(1)) {
				fmt.Println("c = ", c, "-------------")
				r := ModSqrt(c, p)
				fmt.Println("答案：", r)
				for i := 0; i < len(r); i++ {
					if big.NewInt(0).Exp(r[i], big.NewInt(2), p).Cmp(c) == 0 {

					} else {
						fmt.Println("答案错误", r[i], "，c = ", big.NewInt(0).Exp(r[i], big.NewInt(2), p))
						return
					}
				}
			}
			return
		}
		if false {
			p := big.NewInt(401)
			for c := big.NewInt(1); c.Cmp(big.NewInt(0).Add(p, big.NewInt(-1))) <= 0; c.Add(c, big.NewInt(1)) {
				fmt.Println("c = ", c, "-------------")
				r := ModSqrt(c, p)
				fmt.Println("答案：", r)
				for i := 0; i < len(r); i++ {
					if big.NewInt(0).Exp(r[i], big.NewInt(2), p).Cmp(c) == 0 {

					} else {
						fmt.Println("答案错误", r[i], "，c = ", big.NewInt(0).Exp(r[i], big.NewInt(2), p))
						return
					}
				}
			}
			return
		}

	}
	// 测试系统自带的ModSqrt
	if true {
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

// 求模平方根的个数
func ModSqrtCount(c, p *big.Int) int {
	t := big.NewInt(0).Add(p, big.NewInt(-1))
	t.Rsh(t, 1)
	if big.NewInt(0).Exp(c, t, p).Cmp(big.NewInt(1)) == 0 {
		return 2
	} else {
		return 0
	}
}

// Tonelli–Shanks算法
// 求模平方根
func ModSqrt(c, p *big.Int) (ans []*big.Int) {
	ans = make([]*big.Int, 0)
	ans0 := big.NewInt(0)
	if ModSqrtCount(c, p) == 0 {
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
		ans0.Exp(c, s, p)
		ans1 := big.NewInt(0)
		ans1.Neg(ans0)
		ans1.Add(p, ans1)
		ans = append(ans, ans0, ans1)
	} else if t.Cmp(big.NewInt(2)) >= 0 {
		x_ := big.NewInt(0).Exp(c, big.NewInt(0).Add(p, big.NewInt(-2)), p)
		n := big.NewInt(1)
		for ModSqrtCount(n, p) != 0 {
			n.Add(n, big.NewInt(1))
		}
		b := big.NewInt(0).Exp(n, s, p)
		s.Add(s, big.NewInt(1))
		s.Rsh(s, 1)
		ans0.Exp(c, s, p)
		t_ := big.NewInt(0)
		for t.Cmp(big.NewInt(1)) > 0 {
			aa := big.NewInt(0).Mul(x_, ans0)
			aa.Mod(aa, p)
			aa.Mul(aa, ans0)
			aa.Mod(aa, p)
			bb := big.NewInt(0).Exp(big.NewInt(2), big.NewInt(0).Add(t, big.NewInt(-2)), p)
			if big.NewInt(0).Exp(aa, bb, p).Cmp(big.NewInt(1)) != 0 {
				tt := big.NewInt(0).Exp(big.NewInt(2), t_, p)
				tt.Exp(b, tt, p)
				ans0.Mul(ans0, tt)
				ans0.Mod(ans0, p)
			}
			t.Add(t, big.NewInt(-1))
			t_.Add(t_, big.NewInt(1))
		}
		ans1 := big.NewInt(0)
		ans1.Neg(ans0)
		ans1.Add(p, ans1)
		ans = append(ans, ans0, ans1)
	}
	return
}

type num struct {
	x *big.Int //实部
	y *big.Int // 虚部(即虚数单位√w的系数)
	w *big.Int
}

// 复数乘法
func mulI(a num, b num, p *big.Int) num {
	var res num
	res.x = big.NewInt(0)
	res.y = big.NewInt(0)
	res.w = a.w

	res.x.Mul(a.x, b.x) //a.x*b.x
	res.x.Mod(res.x, p)
	x2 := big.NewInt(0)
	x2.Mul(a.y, b.y) //a.y*b.y*w
	x2.Mod(x2, p)
	x2.Mul(x2, res.w)
	x2.Mod(x2, p)
	res.x.Add(res.x, x2)
	res.x.Mod(res.x, p)

	res.y.Mul(a.x, b.y) //a.x*b.y
	res.y.Mod(res.y, p)
	y2 := big.NewInt(0)
	y2.Mul(a.y, b.x) //a.y*b.x
	y2.Mod(y2, p)
	res.y.Add(res.y, y2)
	res.y.Mod(res.y, p)
	return res
}

// 复数快速幂，注意b不能取模
func powerModI(a num, b, p *big.Int) *big.Int {
	res := num{big.NewInt(1), big.NewInt(0), a.w}
	for b.Cmp(big.NewInt(0)) != 0 {
		if big.NewInt(0).Mod(b, big.NewInt(2)).Cmp(big.NewInt(1)) == 0 {
			res = mulI(res, a, p)
		}
		a = mulI(a, a, p)
		b.Rsh(b, 1)
	}
	return res.x.Mod(res.x, p) // 只用返回实数部分，因为虚数部分没了
}

// https://www.luogu.com.cn/blog/shaymin5216/quadratic-residue-solution
// Cipolla算法
// 求模平方根
func Cipolla(n, p *big.Int) (ans []*big.Int) {
	ans = make([]*big.Int, 0)
	var ans0 *big.Int
	if ModSqrtCount(n, p) == 0 {
		return
	}

	a := big.NewInt(0)
	w := big.NewInt(0)
	for { // 找出一个符合条件的a
		a.Add(a, big.NewInt(1)) //网上很多版本，a是取随机数
		w = w.Exp(a, big.NewInt(2), p)
		w.Add(w, big.NewInt(0).Neg(n))
		w.Mod(w, p)
		if ModSqrtCount(w, p) == 0 {
			break
		}
	}

	x := num{a, big.NewInt(1), w}
	p2 := big.NewInt(0).Add(p, big.NewInt(1))
	p2.Rsh(p2, 1)
	ans0 = powerModI(x, p2, p)
	ans1 := big.NewInt(0)
	ans1.Neg(ans0)
	ans1.Add(p, ans1)
	ans = append(ans, ans0, ans1)
	return
}
