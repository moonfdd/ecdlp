package ecdlp

import (
	"fmt"
	"math/big"
)

// polynomial^n mod modPolynomial,modN
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

// polynomial1*polynomial2 mod modN
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

// polynomial1+polynomial2 mod modN
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

// polynomial1-polynomial2 mod modN
func PolynomialSub(polynomial1 []*big.Int, polynomial2 []*big.Int, modN *big.Int) (ans []*big.Int) {
	polynomial2 = PolynomialNeg(polynomial2, modN)
	ans = PolynomialAdd(polynomial1, polynomial2, modN)
	return
}

// polynomial%modPolynomial mod modN,其中modPolynomial的最高次项必须是1
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

// -polynomial mod modN
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

// gcd(polynomial1,polynomial2) mod modN
func PolynomialGcd(polynomial1 []*big.Int, polynomial2 []*big.Int, modN *big.Int) (ans []*big.Int) {
	if len(polynomial1) < len(polynomial2) {
		polynomial1, polynomial2 = polynomial2, polynomial1
	}
	polynomial2 = polynomialCopy(polynomial2)
	if len(polynomial2) == 1 && polynomial2[0].Cmp(big.NewInt(0)) == 0 {
		ans = polynomialCopy(polynomial1)
		return
	}
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

// PolynomialCompose 计算 f(g(x)) mod (modPolynomial, modN)
// f: 多项式 f 的系数，从常数项到最高次项
// g: 多项式 g 的系数，从常数项到最高次项
// modPolynomial: 模多项式（最高次项系数为1）
// modN: 系数模数（一般为素数）
func PolynomialCompose2(f []*big.Int, g []*big.Int, modPolynomial []*big.Int, modN *big.Int) []*big.Int {
	if len(f) == 0 {
		return []*big.Int{}
	}

	// 处理最高次项
	// 如果最高次项系数为0，则结果初始化为零多项式
	highestCoeff := new(big.Int).Set(f[len(f)-1])
	if highestCoeff.Cmp(big.NewInt(0)) == 0 {
		// 如果最高次项系数为0，从次高次项开始
		if len(f) == 1 {
			return []*big.Int{big.NewInt(0)}
		}
		// 递归调用处理去掉最高次项的多项式
		return PolynomialCompose(f[:len(f)-1], g, modPolynomial, modN)
	}

	// 初始化结果为最高次项系数 * g(x)
	result := make([]*big.Int, len(g))
	for i := 0; i < len(g); i++ {
		result[i] = new(big.Int).Mul(highestCoeff, g[i])
		result[i].Mod(result[i], modN)
	}
	result = PolynomialMod(result, modPolynomial, modN)

	// 从次高次项开始到常数项
	for i := len(f) - 2; i >= 0; i-- {
		// 先乘 g
		result = PolynomialMul(result, g, modN)
		result = PolynomialMod(result, modPolynomial, modN)

		// 再加当前的系数 a_i
		coeff := new(big.Int).Set(f[i])
		if coeff.Cmp(big.NewInt(0)) != 0 {
			coeffPoly := []*big.Int{coeff}
			result = PolynomialAdd(result, coeffPoly, modN)
			result = PolynomialMod(result, modPolynomial, modN)
		}
	}

	return result
}

// PolynomialCompose 计算 f(g(x)) mod (modPolynomial, modN)
// 注意：本项目中的多项式系数顺序是“最高次项在前，常数项在后”。
// 例如：
//
//	[]{1,0}   表示 x
//	[]{1,1}   表示 x+1
func PolynomialCompose(f []*big.Int, g []*big.Int, modPolynomial []*big.Int, modN *big.Int) []*big.Int {
	if len(f) == 0 {
		return []*big.Int{big.NewInt(0)}
	}

	// Horner法：
	// 若 f(x)=a0 x^n + a1 x^(n-1) + ... + an
	// 则 f(g)=(((a0)g + a1)g + a2 ... )g + an
	result := []*big.Int{new(big.Int).Set(f[0])}

	for i := 1; i < len(f); i++ {
		// result = result * g
		result = PolynomialMul(result, g, modN)
		result = PolynomialMod(result, modPolynomial, modN)

		// result = result + f[i]
		result = PolynomialAdd(result, []*big.Int{new(big.Int).Set(f[i])}, modN)
		result = PolynomialMod(result, modPolynomial, modN)
	}

	return result
}
