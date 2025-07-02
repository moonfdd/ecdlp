package main

import (
	"fmt"
	"math/big"

	"github.com/moonfdd/ecdlp"
)

func main() {
	if true {
		fmt.Println("测试加法")
		poly1 := []*big.Int{big.NewInt(3), big.NewInt(5), big.NewInt(7)}
		poly2 := []*big.Int{big.NewInt(4), big.NewInt(2), big.NewInt(9)}
		modN := big.NewInt(5)
		res := ecdlp.PolynomialAdd(poly1, poly2, modN)
		fmt.Println(ecdlp.PolynomialString(res))
	}
	if true {
		fmt.Println("测试减法")
		poly1 := []*big.Int{big.NewInt(3), big.NewInt(5), big.NewInt(7)}
		poly2 := []*big.Int{big.NewInt(4), big.NewInt(2), big.NewInt(9)}
		modN := big.NewInt(5)
		res := ecdlp.PolynomialSub(poly1, poly2, modN)
		fmt.Println(ecdlp.PolynomialString(res))
	}
	if true {
		fmt.Println("测试乘法")
		poly1 := []*big.Int{big.NewInt(3), big.NewInt(5), big.NewInt(7)}
		poly2 := []*big.Int{big.NewInt(4), big.NewInt(2), big.NewInt(9)}
		modN := big.NewInt(5)
		res := ecdlp.PolynomialMul(poly1, poly2, modN)
		fmt.Println(ecdlp.PolynomialString(res))
	}
	if true {
		fmt.Println("测试余除")
		poly1 := []*big.Int{big.NewInt(1), big.NewInt(-2), big.NewInt(0), big.NewInt(-1)}
		polyMod := []*big.Int{big.NewInt(1), big.NewInt(0), big.NewInt(0), big.NewInt(0)}
		modN := big.NewInt(5)
		res := ecdlp.PolynomialMod(poly1, polyMod, modN)
		fmt.Println(ecdlp.PolynomialString(res))
	}
	if true {
		fmt.Println("测试快速幂")
		poly1 := []*big.Int{big.NewInt(1), big.NewInt(-1)}
		polyMod := []*big.Int{big.NewInt(1), big.NewInt(0), big.NewInt(0), big.NewInt(0)}
		exp := big.NewInt(4)
		modN := big.NewInt(5)
		res := ecdlp.PolynomialExpMod(poly1, exp, polyMod, modN)
		fmt.Println(ecdlp.PolynomialString(res))
	}
	if true {
		fmt.Println("最大公约数")
		poly1 := []*big.Int{big.NewInt(1), big.NewInt(-2), big.NewInt(1)}
		poly2 := []*big.Int{big.NewInt(1), big.NewInt(-1)}
		modN := big.NewInt(5)
		res := ecdlp.PolynomialGcd(poly1, poly2, modN)
		fmt.Println(ecdlp.PolynomialString(res))
	}
}
