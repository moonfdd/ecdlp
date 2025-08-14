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

		// a := big.NewInt(89)
		// b := big.NewInt(618)
		// p := big.NewInt(809)
		for b := big.NewInt(0); b.Cmp(p) < 0; b.Add(b, big.NewInt(1)) {
			r := Bsgs(a, b, p)
			fmt.Println(b, r)
			r2 := PollardRho(a, b, p)
			fmt.Println(b, r2)
			// if r.Cmp(r2) != 0 {
			// 	fmt.Println("error")
			// 	return
			// }
			if r2 != nil {
				fmt.Println(big.NewInt(0).Exp(a, r2, p))
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
	// fmt.Println("---：", x, alpha, beta)
}

func PollardRho(a, b, p *big.Int) (ans *big.Int) {
	// fmt.Println(a, b, p)
	turtleX := big.NewInt(1)
	turtleAlpha := big.NewInt(0)
	turtleBeta := big.NewInt(0)
	rabbitX := big.NewInt(1)
	rabbitAlpha := big.NewInt(0)
	rabbitBeta := big.NewInt(0)
	// count := 0
	for {
		new_xab(a, b, p, turtleX, turtleAlpha, turtleBeta)
		// fmt.Println("乌龟：", turtleX, turtleAlpha, turtleBeta)
		new_xab(a, b, p, rabbitX, rabbitAlpha, rabbitBeta)
		new_xab(a, b, p, rabbitX, rabbitAlpha, rabbitBeta)
		// fmt.Println("兔子：", rabbitX, rabbitAlpha, rabbitBeta)
		// count++
		// fmt.Println("count = ", count)
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
		// d := big.NewInt(0).GCD(nil, nil, alpha, beta)
		// if d.Cmp(big.NewInt(0)) != 0 {
		// 	alpha.Div(alpha, d)
		// 	beta.Div(beta, d)
		// }

		if alpha.Cmp(big.NewInt(0)) == 0 {
			fmt.Println("alpha0")
		}
		if beta.Cmp(big.NewInt(0)) == 0 {
			fmt.Println("beta0")
		}
		fmt.Println("alpha0/beta0:", alpha, "/", beta, " mod", p_1)
		x := big.NewInt(0)
		y := big.NewInt(0)
		d := big.NewInt(0).GCD(x, y, beta, p_1)
		ans = big.NewInt(0)
		ans.Mul(alpha, x)
		// ans.Div(ans, d)
		ans.Mod(ans, p_1)
		fmt.Println("d ans1 x = ", d, ans, x)

	}

	if false {
		// ans = rabbitX
		// return
		ans = big.NewInt(0)
		ans.Sub(rabbitAlpha, turtleAlpha)
		temp := big.NewInt(0)
		temp.Sub(turtleBeta, rabbitBeta)
		p_1 := big.NewInt(0).Sub(p, big.NewInt(1))

		// p_1 := big.NewInt(382)

		// temp.ModInverse(temp, p_1)
		// ans.Mul(ans, temp).Mod(ans, p_1)
		x := big.NewInt(0)
		y := big.NewInt(0)
		temp.Mod(temp, p_1)
		if temp.Cmp(big.NewInt(0)) == 0 {
			fmt.Println("temp0")
		}
		d := big.NewInt(0).GCD(x, y, temp, p_1)
		fmt.Println("d x y= ", d, x, y)
		ans.Mod(ans, p_1)
		// fmt.Println("ans d = ", ans, d)
		// if big.NewInt(0).Mod(ans, d).Cmp(big.NewInt(0)) == 0 {
		// 	ans.Div(ans, d)
		// 	d = big.NewInt(1)
		// }
		ans.Mul(ans, x).Mod(ans, p_1)
		fmt.Println("ans d = ", ans, d)
	}

	return
}
