package main

import (
	"fmt"
	"math/big"

	"github.com/moonfdd/ecdlp"
)

func main() {
	if true {
		// 斐波那契数列   卢卡斯数列
		ll := ecdlp.LucasParam{big.NewInt(1), big.NewInt(-1)}
		for k := big.NewInt(0); k.Cmp(big.NewInt(11)) <= 0; k.Add(k, big.NewInt(1)) {
			fmt.Println(ll.GetUnAndVn(k))
		}
		return
	}
	if true {
		// 梅森数   2^n+1包含费马数
		ll := ecdlp.LucasParam{big.NewInt(3), big.NewInt(2)}
		for k := big.NewInt(0); k.Cmp(big.NewInt(11)) <= 0; k.Add(k, big.NewInt(1)) {
			fmt.Println(ll.GetUnAndVn(k))
		}
		return
	}
	fmt.Println("")
}
