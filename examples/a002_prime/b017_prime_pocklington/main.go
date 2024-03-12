package main

import (
	"fmt"
	"math/big"
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
	if big.NewInt(0).Exp(a, big.NewInt(0).Add(num, big.NewInt(-1)), num).Cmp(big.NewInt(1)) == 0 {
		return true
	} else {
		return false
	}
}

func Pocklington(num *big.Int) bool {
	if num.Cmp(big.NewInt(1)) <= 0 {
		return false
	}
	if num.Cmp(big.NewInt(2)) == 0 {
		return true
	}
	if num.Bit(0) == 0 {
		return false
	}
	//6log(log n)
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
			return false
		}
		isAllPass := true
		for f, _ := range factorMap {
			factor := big.NewInt(0)
			factor.SetString(f, 10)

			if true {
				//gcd[a^((num-1)/p)-1,num]==1才有可能是素数
				cc := big.NewInt(0).Exp(a, big.NewInt(0).Div(num_1, factor), num)
				cc.Sub(cc, big.NewInt(1)).Mod(cc, num)
				if big.NewInt(0).GCD(nil, nil, cc, num).Cmp(big.NewInt(1)) != 0 {
					isAllPass = false
					break
				}
			}
		}
		if isAllPass {
			// fmt.Println(a, num, "素数", num_2)
			return true
		}
	}
	return false
}

func main() {

	if true {
		left := big.NewInt(4)
		count := 0
		right := big.NewInt(100000)
		for num := big.NewInt(1).Set(left); num.Cmp(right) <= 0; num.Add(num, big.NewInt(1)) {
			r1 := Pocklington(num)
			r2 := num.ProbablyPrime(0)
			if r1 == r2 {

			} else {
				count++
				if r1 {
					fmt.Print(num, " ")
				} else {
					fmt.Print(num, "是合数")
					break
				}
			}
		}
		fmt.Println()
		fmt.Println("错判次数", count)
		return
	}

	fmt.Println("")
}
