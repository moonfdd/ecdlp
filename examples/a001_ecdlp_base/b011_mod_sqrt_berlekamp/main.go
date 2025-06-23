package main

import (
	"fmt"
	"math/big"

	"github.com/moonfdd/ecdlp"
)

func main() {
	if true {
		p := big.NewInt(0)
		p.SetString("9929", 10)
		// p.SetString("9973", 10)
		// p.SetString("11", 10)
		// p.SetString("13", 10)
		// p.SetString("997", 10)
		// p.SetString("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEFFFFFC2F", 16)
		a := big.NewInt(13)
		//481899
		for a = big.NewInt(1); a.Cmp(p) < 0; a.Add(a, big.NewInt(1)) {
			r := BerlekampSqrt(a, p)
			fmt.Println("结果：", a, r)
			if len(r) > 0 {
				tt := big.NewInt(0)
				if tt.Exp(r[0], big.NewInt(2), p).Mod(tt, p).Cmp(a) == 0 {
				} else {
					fmt.Println("出错了")
					return
				}
			}
			fmt.Println("---------------------")
		}

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

// https://en.wikipedia.org/wiki/Berlekamp%E2%80%93Rabin_algorithm
// 求模平方根(lnp)^2
func BerlekampSqrt(a, p *big.Int) (ans []*big.Int) {
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

	m := big.NewInt(0)
	m.Add(p, big.NewInt(-1)).Rsh(m, 1)

	for z := big.NewInt(0); z.Cmp(p) < 0; z.Add(z, big.NewInt(1)) {
		// fmt.Println("z = ", z)
		polynomial1 := []*big.Int{big.NewInt(1), big.NewInt(0).Neg(z)}   //x-z
		polynomial1 = ecdlp.PolynomialMul(polynomial1, polynomial1, p)   //(x-z)^2
		polynomial1 = ecdlp.PolynomialSub(polynomial1, []*big.Int{a}, p) //(x-z)^2-a

		polynomial2 := ecdlp.PolynomialExpMod([]*big.Int{big.NewInt(1), big.NewInt(0)}, m, polynomial1, p) //x^m mod (x-z)^2-a
		polynomial2 = ecdlp.PolynomialSub(polynomial2, []*big.Int{big.NewInt(1)}, p)                       //x^m-1 mod (x-z)^2-a
		gcd := ecdlp.PolynomialGcd(polynomial1, polynomial2, p)
		if len(gcd) == 2 {
			// fmt.Println("z = ", z)
			x.Add(gcd[1], z).Mod(x, p)
			otherX := big.NewInt(0)
			otherX.Neg(x).Add(p, otherX)
			ans = append(ans, x, otherX)
			return
		}
	}

	return
}
