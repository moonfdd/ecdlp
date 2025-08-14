package main

import (
	"fmt"
	"math/big"
)

// https://en.wikipedia.org/wiki/Catalan_pseudoprime
// func Catanlan(n *big.Int) bool {
// }

// 加泰罗尼亚伪素数
func main() {
	p := big.NewInt(29)
	for i := big.NewInt(1); i.Cmp(p) < 0; i.Add(i, big.NewInt(1)) {
		fmt.Println(i, big.NewInt(0).Exp(i, big.NewInt(8), p))
	}
	// // 测试输入一个奇数正整数 n
	// // 这里使用big.Int支持大整数输入
	// var nStr string
	// fmt.Print("请输入一个奇数正整数 n: ")
	// fmt.Scan(&nStr)
	// n := new(big.Int)
	// n.SetString(nStr, 10)

	// // 确保是奇数
	// if n.Bit(0) == 0 {
	// 	fmt.Println("错误: 输入必须为奇数！")
	// 	return
	// }

	// perrinCalculation(n)
}
