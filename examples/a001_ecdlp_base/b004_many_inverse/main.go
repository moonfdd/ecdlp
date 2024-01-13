package main

import (
	"fmt"
	"math/big"
)

func main() {
	//单个
	if true {
		for i := 1; i < 11; i++ {
			a := big.NewInt(0).SetInt64(int64(i))
			p := big.NewInt(11)
			fmt.Print(OneInverse(a, p), " ")
		}
		fmt.Println("")
	}

	// 批量
	if true {
		ans := ManyInverse(10, 11)
		fmt.Println(ans)
	}
	fmt.Println()
}

// 求1/a%p
// a是1到p-1之间的整数
// p是大于等于3的素数
func OneInverse(a, p *big.Int) (ans *big.Int) {
	ans = big.NewInt(1)
	atemp := big.NewInt(0).Mod(a, p)
	divValue := big.NewInt(0)
	modValue := big.NewInt(0)
	if a.Cmp(big.NewInt(1)) == 0 {
		return
	}
	for {
		divValue.Div(p, atemp)
		modValue.Mod(p, atemp)
		ans.Mul(ans, divValue).Neg(ans).Mod(ans, p)
		if modValue.Cmp(big.NewInt(1)) == 0 {
			break
		}
		if modValue.Cmp(big.NewInt(0)) == 0 {
			ans.Set(big.NewInt(0))
			break
		}
		atemp.Set(modValue)
	}
	return
}

// https://blog.csdn.net/qq_43481884/article/details/108629010 线性打表求逆元
// 批量求逆元，求1~n的逆元
// n为1到p-1的整数
// p	为素数
func ManyInverse(n, p int) (ans []int) {
	ans = make([]int, p)
	ans[1] = 1
	for i := 2; i <= p-1; i++ {
		ans[i] = ((p - p/i) * ans[p%i]) % p
	}
	ans = ans[1:]
	return
}
