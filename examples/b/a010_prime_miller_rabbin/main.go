package main

import (
	"fmt"
	"math/big"
)

func MillerRabbin(num *big.Int) bool {
	if num.Cmp(big.NewInt(1)) == 0 {
		return false
	}
	if num.Cmp(big.NewInt(2)) == 0 {
		return true
	}
	if big.NewInt(0).And(num, big.NewInt(1)).Cmp(big.NewInt(0)) == 0 {
		return false
	}

	aList := []*big.Int{big.NewInt(2)}
	// aList := []*big.Int{big.NewInt(2), big.NewInt(325), big.NewInt(9375), big.NewInt(28178), big.NewInt(450775), big.NewInt(9780504), big.NewInt(1795265022)}

	// https://qa.1r1g.com/sf/ask/1686743271/
	//aList := []*big.Int{big.NewInt(2), big.NewInt(3), big.NewInt(5), big.NewInt(7), big.NewInt(11), big.NewInt(13), big.NewInt(17), big.NewInt(19), big.NewInt(23)}

	// aList := []*big.Int{big.NewInt(2), big.NewInt(3)}
	// if num.Cmp(big.NewInt(1373653)) < 0 {
	// 	aList = []*big.Int{big.NewInt(2), big.NewInt(3)}
	// } else if num.Cmp(big.NewInt(9080191)) < 0 {
	// 	aList = []*big.Int{big.NewInt(31), big.NewInt(73)}
	// } else if num.Cmp(big.NewInt(4759123141)) < 0 {
	// 	aList = []*big.Int{big.NewInt(2), big.NewInt(7), big.NewInt(61)}
	// } else if num.Cmp(big.NewInt(2152302898747)) < 0 {
	// 	aList = []*big.Int{big.NewInt(2), big.NewInt(3), big.NewInt(5), big.NewInt(7), big.NewInt(11)}
	// }
	// aList := []*big.Int{big.NewInt(2)}

	b := false
	for _, a := range aList {
		// if a.Cmp(num) < 0 {
		if big.NewInt(0).GCD(nil, nil, a, num).Cmp(big.NewInt(1)) == 0 {
			b = MillerRabbinA(a, num)
			if !b {
				return false
			}
		}
	}
	return b
}

// 算法导论第三版567
func MillerRabbinA(a, num *big.Int) bool {
	if num.Cmp(big.NewInt(1)) == 0 {
		return false
	}
	if num.Cmp(big.NewInt(2)) == 0 {
		return true
	}
	if big.NewInt(0).And(num, big.NewInt(1)).Cmp(big.NewInt(0)) == 0 {
		return false
	}
	t := big.NewInt(0)
	u := big.NewInt(0).Sub(num, big.NewInt(1))
	for big.NewInt(0).And(u, big.NewInt(1)).Cmp(big.NewInt(0)) == 0 {
		t.Add(t, big.NewInt(1))
		u.Rsh(u, 1)
	}
	x := big.NewInt(0).Exp(a, u, num)
	var xtemp *big.Int
	for i := big.NewInt(0); i.Cmp(t) < 0; i.Add(i, big.NewInt(1)) {
		xtemp = big.NewInt(0).Exp(x, big.NewInt(2), num)
		if xtemp.Cmp(big.NewInt(1)) == 0 && x.Cmp(big.NewInt(1)) != 0 && x.Cmp(big.NewInt(0).Sub(num, big.NewInt(1))) != 0 {
			return false
		}
		x = xtemp
	}
	if x.Cmp(big.NewInt(1)) != 0 {
		return false
	}
	return true
}

func main() {
	// 算法导论第三版566
	if true {
		num := big.NewInt(0)
		num.SetString("1", 10)
		count := 0
		rightLimit := big.NewInt(0)
		rightLimit.SetString("10000", 10)
		for ; ; /*num.Cmp(rightLimit) <= 0*/ num.Add(num, big.NewInt(1)) {

			r := MillerRabbin(num)
			r2 := num.ProbablyPrime(0)

			if r == r2 {
				if r {
					// fmt.Println(num, "是素数")
				}
			} else {
				fmt.Println("测试失败", r, r2, num)
				count++
				return
			}

		}
		fmt.Println("失败次数", count)
		return
	}
	// https://www.cnblogs.com/pengchujie/p/17801601.html
	if true {
		//488
		num := big.NewInt(0)
		num.SetString("1", 10)
		// num.SetString("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEFFFFFC2F", 16)
		// num.SetString("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEBAAEDCE6AF48A03BBFD25E8CD0364141", 16)
		for ; ; num.Add(num, big.NewInt(1)) {

			r := MillerRabbin(num)
			r2 := num.ProbablyPrime(0)

			if r == r2 {
				if r {
					// fmt.Println(num, "是素数")
				}
			} else {
				fmt.Println("测试失败", r, r2, num)
				// return
			}

		}
	}
	fmt.Println("")
}
