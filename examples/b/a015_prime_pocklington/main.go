package main

import (
	"fmt"
	"math/big"
	"os"
)

// 分解质因数
func FactorInteger(num *big.Int) (factorMap map[string]*big.Int) {
	num = big.NewInt(0).Add(num, big.NewInt(0))
	factorMap = make(map[string]*big.Int)
	if num.Cmp(big.NewInt(1)) == 0 {
		return
	}
	for {
		if big.NewInt(0).And(num, big.NewInt(1)).Cmp(big.NewInt(0)) != 0 {
			break
		}
		if factorMap["2"] == nil {
			factorMap["2"] = big.NewInt(1)
		} else {
			factorMap["2"].Add(factorMap["2"], big.NewInt(1))
		}
		num.Div(num, big.NewInt(2))

	}
	for i := big.NewInt(3); ; i.Add(i, big.NewInt(2)) {
		for big.NewInt(0).Mod(num, i).Cmp(big.NewInt(0)) == 0 {
			if factorMap[i.Text(10)] == nil {
				factorMap[i.Text(10)] = big.NewInt(1)
			} else {
				factorMap[i.Text(10)].Add(factorMap[i.Text(10)], big.NewInt(1))
			}
			num.Div(num, i)
		}
		if num.ProbablyPrime(0) {
			if factorMap[num.Text(10)] == nil {
				factorMap[num.Text(10)] = big.NewInt(1)
			} else {
				factorMap[num.Text(10)].Add(factorMap[num.Text(10)], big.NewInt(1))
			}
			break
		}
		if num.Cmp(big.NewInt(1)) == 0 {
			return
		}
	}
	return
}

func Fermat(a, num *big.Int) bool {
	if num.Cmp(big.NewInt(1)) <= 0 {
		return false
	}
	if num.Cmp(big.NewInt(2)) == 0 {
		return true
	}
	if num.Bit(0) == 0 {
		return false
	}
	if big.NewInt(0).Exp(a, big.NewInt(0).Add(num, big.NewInt(-1)), num).Cmp(big.NewInt(1)) == 0 {
		return true
	} else {
		return false
	}
}

func t(num *big.Int) bool {
	if num.Cmp(big.NewInt(1)) <= 0 {
		return false
	}
	if num.Cmp(big.NewInt(2)) == 0 {
		return true
	}
	if num.Bit(0) == 0 {
		return false
	}
	num_1 := big.NewInt(0).Sub(num, big.NewInt(1))
	factorMap := FactorInteger(num_1)
	num_2 := big.NewInt(0)
	num_2.SetInt64(int64(num_1.BitLen()))
	num_3 := big.NewInt(0)
	num_3.SetInt64(int64(num_2.BitLen()))
	num_2.Mul(num_3, num_3)
	num_2.Mul(num_2, big.NewInt(6))
	// num_2.Mul(num_2, num_3)

	for a := big.NewInt(2); a.Cmp(num_2) <= 0; a.Add(a, big.NewInt(1)) {
		if !Fermat(a, num) {
			// fmt.Println("a1")
			return false
		}
		bb := true
		for f, _ := range factorMap {
			factor := big.NewInt(0)
			factor.SetString(f, 10)

			if false {
				// a^((num-1)/p)!=1才有可能是素数
				// fmt.Println(a, num_1, factor, big.NewInt(0).Div(num_1, factor))
				if big.NewInt(0).Exp(a, big.NewInt(0).Div(num_1, factor), num).Cmp(big.NewInt(1)) == 0 {
					bb = false
					// fmt.Println(a, num_1, factor, big.NewInt(0).Div(num_1, factor))
					break
				}
			}
			if true {
				//gcd[a^((num-1)/p)-1,num]==1才有可能是素数
				cc := big.NewInt(0).Exp(a, big.NewInt(0).Div(num_1, factor), num)
				cc.Sub(cc, big.NewInt(1))
				if big.NewInt(0).GCD(nil, nil, cc, num).Cmp(big.NewInt(1)) != 0 {
					bb = false
					break
				}
			}

		}
		if bb {
			fmt.Println(a, num, "素数", num_2)
			return true
		}
	}
	// fmt.Println("a3")
	return false
}

func main() {
	if false {
		fmt.Println(big.NewInt(19).Text(2))
		fmt.Println(big.NewInt(607).Text(2))
		fmt.Println(big.NewInt(89).Text(2))
		fmt.Println(big.NewInt(22783).Text(2))
		return
	}
	if false {

		rightLimit := big.NewInt(0)
		rightLimit.SetString("10000", 10)
		num := big.NewInt(2)
		jinzhi := big.NewInt(100)
		jinzhi_1 := big.NewInt(0).Sub(jinzhi, big.NewInt(1))
		for ; num.Cmp(rightLimit) <= 0; num.Add(num, big.NewInt(1)) {
			x := big.NewInt(0).Exp(jinzhi, num, nil)
			x.Sub(x, big.NewInt(1))
			if big.NewInt(0).Mod(x, jinzhi_1).Cmp(big.NewInt(0)) != 0 {
				continue
			}
			x.Div(x, jinzhi_1)

			r := x.ProbablyPrime(0)
			if r {
				// fmt.Println(x, "是素数", len(x.Text(int(jinzhi.Int64()))))
				fmt.Println(x, "是素数")
			}

		}
		return
	}
	if false {
		num := big.NewInt(0)
		num.SetString("1", 10)
		count := 0
		rightLimit := big.NewInt(0)
		rightLimit.SetString("1000", 10)
		for ; num.Cmp(rightLimit) <= 0; num.Add(num, big.NewInt(1)) {
			r := num.ProbablyPrime(0)
			if r {
				fmt.Println("是素数", num.Text(3), num)
			}

		}
		fmt.Println("失败次数", count)
		return
	}
	if true {
		if false {
			// fmt.Println(big.NewInt(31).ProbablyPrime(0))
			// return
			p := big.NewInt(3)
			for b := big.NewInt(2); b.Cmp(p) < 0; b.Add(b, big.NewInt(1)) {
				// for b := big.NewInt(0).Sub(p, big.NewInt(1)); b.Cmp(big.NewInt(2)) >= 0; b.Sub(b, big.NewInt(1)) {
				count := 0
				for a := big.NewInt(1); a.Cmp(p) < 0; a.Add(a, big.NewInt(1)) {
					if big.NewInt(0).Exp(b, a, p).Cmp(big.NewInt(1)) == 0 {
						// fmt.Println(b, a, big.NewInt(0).Exp(b, a, p))
						count++
					}
				}
				if count == 1 {
					fmt.Println(b, "素数", p)
					// return
				}

			}
			fmt.Println("合数", p)
			return
		}
		num := big.NewInt(0)
		num.SetString("1", 10)
		count := 0
		rightLimit := big.NewInt(0)
		rightLimit.SetString("51097210000", 10) // 5109721 94
		for ; num.Cmp(rightLimit) <= 0; num.Add(num, big.NewInt(1)) {
			r := t(num)
			r2 := num.ProbablyPrime(0)

			if r == r2 {
				if r {
					fmt.Println(num, "是素数")
				}
			} else {
				fmt.Println("测试失败", r, r2, num)
				os.Exit(0)
				count++
				// return
			}

		}
		fmt.Println("失败次数", count)
		return
	}
	fmt.Println("")
}
