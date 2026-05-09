package main

import (
	"fmt"
	"math/big"
)

// 《Fermat数》P7
// 康继鼎 判别法
// file:///E:/%E5%9B%BE%E7%89%87/20250621/%E7%9B%B8%E5%86%8C/IMG_20250621_120240.jpg
func KangJiDing(m *big.Int) bool {
	F := big.NewInt(0)
	F.Exp(big.NewInt(2), m, nil)
	F.Exp(big.NewInt(2), F, nil)
	F.Add(F, big.NewInt(1))
	F_1 := big.NewInt(0).Sub(F, big.NewInt(1))
	sum := big.NewInt(1)
	for k := big.NewInt(1); k.Cmp(F) < 0; k.Add(k, big.NewInt(1)) {
		sum.Add(sum, big.NewInt(0).Exp(k, F_1, F)).Mod(sum, F)
	}
	return sum.Cmp(big.NewInt(0)) == 0
}

func main() {
	if true {
		for m := big.NewInt(0); m.Cmp(big.NewInt(6)) < 0; m.Add(m, big.NewInt(1)) {
			r1 := KangJiDing(m)
			F := big.NewInt(0)
			F.Exp(big.NewInt(2), m, nil)
			F.Exp(big.NewInt(2), F, nil)
			F.Add(F, big.NewInt(1))
			r2 := F.ProbablyPrime(0)
			if r1 != r2 {
				fmt.Printf("错误%v，%v，%v\r\n", m, r1, r2)
				return
			}
		}
		fmt.Println("完全正确")
	}

}
