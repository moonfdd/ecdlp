package main

import (
	"fmt"
	"math/big"
)

func main() {
	// 测试ModSqrt
	if true {

		p := big.NewInt(0)
		p.SetString("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEFFFFFC2F", 16)
		var r []*big.Int
		r = Cipolla(big.NewInt(2), p)
		fmt.Println(r)
		r = Cipolla(big.NewInt(55), big.NewInt(103))
		fmt.Println(r)
		r = Cipolla(big.NewInt(186), big.NewInt(401))
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
	t := big.NewInt(0).Add(p, big.NewInt(-1))
	t.Rsh(t, 1)
	if big.NewInt(0).Exp(a, t, p).Cmp(big.NewInt(1)) == 0 {
		return 2
	} else {
		return 0
	}
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
	b = big.NewInt(0).Set(b)
	for b.Cmp(big.NewInt(0)) != 0 {
		if big.NewInt(0).Mod(b, big.NewInt(2)).Cmp(big.NewInt(1)) == 0 {
			res = mulI(res, a, p)
		}
		a = mulI(a, a, p)
		b.Rsh(b, 1)
	}
	return res.x.Mod(res.x, p) // 只用返回实数部分，因为虚数部分没了
}

// Cipolla算法
// 求模平方根
func Cipolla(a, p *big.Int) (ans []*big.Int) {
	ans = make([]*big.Int, 0)
	x := big.NewInt(0)
	if a.Cmp(big.NewInt(0)) == 0 {
		ans = append(ans, x)
		return
	}
	if ModSqrtCount(a, p) == 0 {
		return
	}
	//i^2=w=b^2-a^2
	b := big.NewInt(0)
	w := big.NewInt(0)
	for { // 找出一个符合条件的b
		b.Add(b, big.NewInt(1)) //网上很多版本，b是取随机数
		w.Exp(b, big.NewInt(2), p).Sub(w, a).Mod(w, p)
		if ModSqrtCount(w, p) == 0 {
			break
		}
	}

	//iRoot=b+i
	iRoot := num{b, big.NewInt(1), w}
	//p2=(p+1)/2
	p2 := big.NewInt(0)
	p2.Add(p, big.NewInt(1)).Rsh(p2, 1)
	x = powerModI(iRoot, p2, p)
	otherX := big.NewInt(0)
	otherX.Neg(x).Add(p, otherX)
	ans = append(ans, x, otherX)
	return
}
