package main

import (
	"fmt"
	"math/big"
)

// https://zhuanlan.zhihu.com/p/132603308
// https://codeleading.com/article/14084397800/
func main() {
	// if true {
	// 	ans := big.NewInt(0)
	// 	ans.Mul(big.NewInt(618), big.NewInt(618)).Mod(ans, big.NewInt(809))
	// 	fmt.Println(ans)
	// 	ans.Mul(ans, big.NewInt(618)).Mod(ans, big.NewInt(809))
	// 	fmt.Println(ans)
	// 	ans.Mul(ans, big.NewInt(618)).Mod(ans, big.NewInt(809))
	// 	fmt.Println(ans)

	// 	ans.Mul(big.NewInt(555), big.NewInt(555)).Mod(ans, big.NewInt(809))
	// 	fmt.Println(ans)
	// }
	// return
	if true {
		a := big.NewInt(2)
		// b := big.NewInt(228)
		p := big.NewInt(1019)
		p = big.NewInt(101)
		p = big.NewInt(383)
		p = big.NewInt(29)
		p = big.NewInt(31) //不一定成功
		p = big.NewInt(1907)

		for b := big.NewInt(0); b.Cmp(p) < 0; b.Add(b, big.NewInt(1)) {
			r1 := Bsgs(a, b, p)
			fmt.Println(b, r1)
			r2 := PollardRho(a, b, p)
			fmt.Println(b, r2)
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
		ans = PollardRho(a, b, m)
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

func new_xab(a, b, p, x, alpha, beta *big.Int) {
	p_1 := big.NewInt(0).Sub(p, big.NewInt(1))
	switch fmt.Sprint(big.NewInt(0).Mod(x, big.NewInt(3))) {
	case "1":
		x.Mul(x, b).Mod(x, p)
		beta.Add(beta, big.NewInt(1)).Mod(beta, p_1)
	case "0":
		x.Exp(x, big.NewInt(2), p)
		alpha.Add(alpha, alpha).Mod(alpha, p_1)
		beta.Add(beta, beta).Mod(beta, p_1)
	case "2":
		x.Mul(x, a).Mod(x, p)
		alpha.Add(alpha, big.NewInt(1)).Mod(alpha, p_1)
	}
}

func PollardRho(a, b, p *big.Int) (ans *big.Int) {
	turtleX := big.NewInt(1)
	turtleAlpha := big.NewInt(0)
	turtleBeta := big.NewInt(0)
	rabbitX := big.NewInt(1)
	rabbitAlpha := big.NewInt(0)
	rabbitBeta := big.NewInt(0)
	for {
		new_xab(a, b, p, turtleX, turtleAlpha, turtleBeta)
		new_xab(a, b, p, rabbitX, rabbitAlpha, rabbitBeta)
		new_xab(a, b, p, rabbitX, rabbitAlpha, rabbitBeta)
		if turtleX.Cmp(rabbitX) == 0 {
			break
		}
	}

	if true {
		p_1 := big.NewInt(0).Sub(p, big.NewInt(1))
		alpha := big.NewInt(0).Sub(rabbitAlpha, turtleAlpha)
		alpha.Mod(alpha, p_1)
		beta := big.NewInt(0).Sub(turtleBeta, rabbitBeta)
		beta.Mod(beta, p_1)

		if beta.Cmp(big.NewInt(0)) == 0 {
			fmt.Println("beta = ", 0)
			return
		}
		x := big.NewInt(0)
		y := big.NewInt(0)
		d := big.NewInt(0).GCD(x, y, beta, p_1)
		m := big.NewInt(0).Sqrt(p_1)
		if d.Cmp(m) == 1 {
			fmt.Println("d = ", d)
			return
		}

		ans = big.NewInt(0)
		ans.Mul(alpha, x).Mod(ans, p_1)

		// fmt.Println("d ans1 x = ", d, ans, x)
		ans.Div(ans, d)
		p_d_div_d := big.NewInt(0).Div(p_1, d)
		for di := big.NewInt(0); di.Cmp(d) < 0; di.Add(di, big.NewInt(1)) {
			if big.NewInt(0).Exp(a, ans, p).Cmp(b) == 0 {
				return
			}
			ans.Add(ans, p_d_div_d)
		}
		ans = nil
	}
	return
}
