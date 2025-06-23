package ecdlp

import (
	"fmt"
	"math/big"
)

// a^n mod e,n
func PolynomialExpMod(polynomial []*big.Int, n *big.Int, modPolynomial []*big.Int, modN *big.Int) (ans []*big.Int) {
	b := big.NewInt(0).Add(n, big.NewInt(0))
	ans = []*big.Int{big.NewInt(1)}
	a := polynomial
	for b.Cmp(big.NewInt(0)) != 0 { //b!=0
		if big.NewInt(0).And(b, big.NewInt(1)).Cmp(big.NewInt(0)) != 0 { //b&1!=0
			ans = PolynomialMul(ans, a, modN)
			ans = PolynomialMod(ans, modPolynomial, modN)
		}
		b.Rsh(b, 1) //b>>=1
		a = PolynomialMul(a, a, modN)
		a = PolynomialMod(a, modPolynomial, modN)
	}
	return
}

// a*b mod n
func PolynomialMul(polynomial1 []*big.Int, polynomial2 []*big.Int, modN *big.Int) (ans []*big.Int) {
	ans = make([]*big.Int, (len(polynomial1)-1)+(len(polynomial2)-1)+1)
	for i := 0; i < len(ans); i++ {
		ans[i] = big.NewInt(0)
	}
	for i := 0; i < len(polynomial1); i++ {
		for j := 0; j < len(polynomial2); j++ {
			ans[i+j].Add(ans[i+j], big.NewInt(0).Mul(polynomial1[i], polynomial2[j]))
			ans[i+j].Mod(ans[i+j], modN)
		}
	}
	k := 0
	for k < len(ans)-1 {
		if ans[k].Cmp(big.NewInt(0)) != 0 {
			break
		}
		k++
	}
	ans = ans[k:]
	return
}

// a+b mod n
func PolynomialAdd(polynomial1 []*big.Int, polynomial2 []*big.Int, modN *big.Int) (ans []*big.Int) {
	//假设第1个多项式的长度大于第2个多项式的长度
	if len(polynomial1) < len(polynomial2) {
		polynomial1, polynomial2 = polynomial2, polynomial1
	}
	temp := polynomialCopy(polynomial1)
	for i := 0; i < len(polynomial2); i++ {
		temp[i+len(polynomial1)-len(polynomial2)].Add(temp[i+len(polynomial1)-len(polynomial2)], polynomial2[i])

		temp[i+len(polynomial1)-len(polynomial2)].Mod(temp[i+len(polynomial1)-len(polynomial2)], modN)
	}
	ans = temp
	k := 0
	for k < len(ans)-1 {
		if ans[k].Cmp(big.NewInt(0)) != 0 {
			break
		}
		k++
	}
	ans = ans[k:]
	return
}

// a-b mod n
func PolynomialSub(polynomial1 []*big.Int, polynomial2 []*big.Int, modN *big.Int) (ans []*big.Int) {
	polynomial2 = PolynomialNeg(polynomial2, modN)
	ans = PolynomialAdd(polynomial1, polynomial2, modN)
	return
}

// a%b mod n,b的最高次项必须是1
func PolynomialMod(polynomial []*big.Int, modPolynomial []*big.Int, modN *big.Int) (ans []*big.Int) {
	zero := big.NewInt(0)
	if len(modPolynomial) == 1 {
		ans = []*big.Int{big.NewInt(0).Add(polynomial[len(polynomial)-1], zero)}
		ans[0].Mod(ans[0], modN)
		return
	}
	temp := polynomialCopy(polynomial)
	for i := 0; i <= len(temp)-len(modPolynomial); i++ {
		if temp[i].Cmp(zero) == 0 {
			continue
		}
		//消减多项式
		// 确定消减系数
		coefficient := big.NewInt(0).Neg(temp[i])
		// 遍历消减多项式
		for j := 0; j < len(modPolynomial); j++ {
			temp[i+j].Add(temp[i+j], big.NewInt(0).Mul(coefficient, modPolynomial[j]))
			temp[i+j].Mod(temp[i+j], modN)
		}
	}
	if len(temp)-len(modPolynomial) >= 0 {
		k := len(temp) - len(modPolynomial) + 1
		for ; k < len(temp)-1; k++ {
			if temp[k].Cmp(big.NewInt(0)) != 0 {
				break
			}
		}
		ans = temp[k:]
	} else {
		for i := 0; i < len(temp); i++ {
			temp[i].Mod(temp[i], modN)
		}
		ans = temp
	}
	return
}

// -a mod n
func PolynomialNeg(polynomial []*big.Int, modN *big.Int) (ans []*big.Int) {
	ans = polynomialCopy(polynomial)
	for i := 0; i < len(ans); i++ {
		ans[i].Neg(ans[i]).Mod(ans[i], modN)
	}
	return
}

// copy
func polynomialCopy(polynomial []*big.Int) (ans []*big.Int) {
	ans = make([]*big.Int, len(polynomial))
	for i := 0; i < len(ans); i++ {
		ans[i] = big.NewInt(0).Set(polynomial[i])
	}
	return
}

// gcd(a,b) mod n
func PolynomialGcd(polynomial1 []*big.Int, polynomial2 []*big.Int, modN *big.Int) (ans []*big.Int) {
	if len(polynomial1) < len(polynomial2) {
		polynomial1, polynomial2 = polynomial2, polynomial1
	}
	polynomial2 = polynomialCopy(polynomial2)
	for {
		if len(polynomial2) == 1 {
			if polynomial2[0].Cmp(big.NewInt(0)) == 0 {
				ans = polynomialCopy(polynomial1)

				if ans[0].Cmp(big.NewInt(1)) != 0 {
					ansLeft := big.NewInt(0).Set(ans[0])
					ansLeft.ModInverse(ansLeft, modN)
					for i := 0; i < len(ans); i++ {
						ans[i].Mul(ans[i], ansLeft).Mod(ans[i], modN)
					}
				}

			} else {
				ans = []*big.Int{big.NewInt(1)}
			}
			return
		}

		polynomial2Left := big.NewInt(0).Set(polynomial2[0])
		polynomial2Left.ModInverse(polynomial2Left, modN)
		for i := 0; i < len(polynomial2); i++ {
			polynomial2[i].Mul(polynomial2[i], polynomial2Left).Mod(polynomial2[i], modN)
		}

		temp := PolynomialMod(polynomial1, polynomial2, modN)
		polynomial1, polynomial2 = polynomial2, temp
	}
}

// 多项式字符串
func PolynomialString(polynomial []*big.Int) (ans string) {
	if len(polynomial) == 1 {
		ans = polynomial[0].String()
		return
	}
	v := ""
	c := 0
	for i := 0; i < len(polynomial)-2; i++ {
		v = polynomial[i].String()
		if v == "1" {
			v = ""
		} else if v == "-1" {
			v = "-"
		}
		c = polynomial[i].Cmp(big.NewInt(0))
		if c < 0 {
			ans += v + "x^" + fmt.Sprint(len(polynomial)-i-1)
		} else if c > 0 {
			ans += "+" + v + "x^" + fmt.Sprint(len(polynomial)-i-1)
		}
	}
	c = polynomial[len(polynomial)-2].Cmp(big.NewInt(0))
	v = polynomial[len(polynomial)-2].String()
	if v == "1" {
		v = ""
	} else if v == "-1" {
		v = "-"
	}
	if c < 0 {
		ans += v + "x"
	} else if c > 0 {
		ans += "+" + v + "x"
	}
	c = polynomial[len(polynomial)-1].Cmp(big.NewInt(0))
	v = polynomial[len(polynomial)-1].String()
	if c < 0 {
		ans += v
	} else if c > 0 {
		ans += "+" + v
	}
	if ans[0] == '+' {
		ans = ans[1:]
	}
	return
}
