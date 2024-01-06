package main

import (
	"fmt"
	"math/big"
)

func main() {
	if true {
		//自定义快速幂PowerMod
		r := PowerMod(big.NewInt(5), big.NewInt(3), big.NewInt(7))
		fmt.Println(r)
	}
	if true {
		// Exp就是求快速幂
		r := big.NewInt(0).Exp(big.NewInt(5), big.NewInt(3), big.NewInt(7))
		fmt.Println(r)
	}
	fmt.Println("")
}

// 求快速幂a^b%p
// a是正整数
// b是非负整数
// p是大于等于3的质数
func PowerMod(a, b, p *big.Int) (ans *big.Int) {
	ans = big.NewInt(1)         //ans=1
	a = big.NewInt(0).Mod(a, p) //a=a%p
	b = big.NewInt(0).Set(b)    //循环里会改变b的值
	// b.Mod(b, big.NewInt(0).Sub(p, big.NewInt(1))) //b=b%(p-1)
	for b.Cmp(big.NewInt(0)) != 0 { //b!=0
		if b.Bit(0) == 1 { //b&1!=0
			ans.Mul(ans, a).Mod(ans, p) //ans*=a ans%=p
		}
		b.Rsh(b, 1)           //b>>=1
		a.Mul(a, a).Mod(a, p) //a*=a a%=p
	}
	//返回ans
	return
}
