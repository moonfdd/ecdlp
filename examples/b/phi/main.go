package main

import "fmt"

// 计算欧拉函数的函数
func eulerPhi(n int) int {
	result := n // 初始化结果为n

	// 对于每个小于n且与n互质的数i，将结果减去i的倍数
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			// i是n的一个因子
			result -= result / i
			for n%i == 0 {
				n /= i
			}
		}
	}

	// 如果n大于1，则n是一个质数
	if n > 1 {
		result -= result / n
	}

	return result
}

func main() {
	// 测试欧拉函数计算
	for n:=1;n<100;n++{
	fmt.Printf("欧拉函数φ(%d)：%d\n",n, eulerPhi(n))
	}
}
