package main

import (
	"fmt"
	"math/big"
)

func main() {
	if true {
		//自定义求模除a/b%p
		r := MulAndModInverse(big.NewInt(4), big.NewInt(5), big.NewInt(7))
		fmt.Println(r)
	}
	if true {
		r3 := big.NewInt(0)
		r3.ModInverse(big.NewInt(5), big.NewInt(7))      //求逆 1/5
		r3.Mul(r3, big.NewInt(4)).Mod(r3, big.NewInt(7)) // 4*1/5 %p
		fmt.Println(r3)
	}
	fmt.Println("")
}

// 求模除a/b%p
// a是正整数
// b是非零整数
// p是大于等于3的质数
func MulAndModInverse(a, b, p *big.Int) (ans *big.Int) {
	ans = big.NewInt(0).Exp(b, big.NewInt(0).Add(p, big.NewInt(-2)), p) //ans = 1/b等价于ans=b^(p-2)
	ans.Mul(ans, a).Mod(ans, p)                                         //ans*=a ans%=p
	//返回ans
	return
}
