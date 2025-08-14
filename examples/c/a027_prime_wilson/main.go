package main

import (
	"fmt"
	"math/big"
)

// file:///E:/%E5%9B%BE%E7%89%87/20250621/%E7%9B%B8%E5%86%8C/IMG_20250621_122525.jpg
func Wilson(n *big.Int) bool {
	sum := big.NewInt(1)
	for i := big.NewInt(2); i.Cmp(n) < 0; i.Add(i, big.NewInt(1)) {
		sum.Mul(sum, i).Mod(sum, n)
	}
	sum.Add(sum, big.NewInt(1)).Mod(sum, n)
	return sum.Cmp(big.NewInt(0)) == 0
}
func main() {
	if true {
		errCount := 0
		for n := big.NewInt(2); n.Cmp(big.NewInt(5000)) <= 0; n.Add(n, big.NewInt(1)) {
			r := Wilson(n)
			r2 := n.ProbablyPrime(0)
			if r != r2 {
				errCount++
				fmt.Println("错误", n, r, r2)
			} else {
				if r {
					//fmt.Println("素数", n)
				}
			}
		}
		fmt.Println("错误次数", errCount)
	}
	fmt.Println("")
}
