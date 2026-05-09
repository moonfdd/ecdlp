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

func main() {
	fmt.Println(FactorInteger(big.NewInt(123456789)))
	fmt.Println(FactorInteger(big.NewInt(9876543210)))
}
