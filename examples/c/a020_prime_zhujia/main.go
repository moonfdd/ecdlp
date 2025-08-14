package main

import (
	"fmt"
	"math/big"
)

// file:///E:/%E5%9B%BE%E7%89%87/20250621/%E7%9B%B8%E5%86%8C/IMG_20250621_121330.jpg
// 朱家猜测
func ZhuJia(n *big.Int) bool {
	n_1 := new(big.Int).Sub(n, big.NewInt(1))
	sum := big.NewInt(1)
	for k := big.NewInt(1); k.Cmp(n) < 0; k.Add(k, big.NewInt(1)) {
		sum.Add(sum, big.NewInt(0).Exp(k, n_1, n)).Mod(sum, n)
	}
	return sum.Cmp(big.NewInt(0)) == 0
}

func main() {
	if true {
		for m := big.NewInt(0); m.Cmp(big.NewInt(5000)) < 0; m.Add(m, big.NewInt(1)) {
			r1 := ZhuJia(m)
			r2 := m.ProbablyPrime(0)
			if r1 != r2 {
				fmt.Printf("错误%v，%v，%v\r\n", m, r1, r2)
				return
			}
		}
		fmt.Println("完全正确")
	}

}
