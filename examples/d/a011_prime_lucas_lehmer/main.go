package main

import (
	"fmt"
	"math/big"
)

// https://en.wikipedia.org/wiki/Lucas%E2%80%93Lehmer_primality_test

// Lucas-Lehmer素性测试的步骤如下：

// 选择一个梅森数，表示为M = 2^p - 1，其中p是质数。

// 初始化变量s为4。

// 执行以下操作n-2次：

// 计算s = (s^2 - 2) mod M。
// 判断结果：如果s等于0，则梅森数对应的p可能是一个素数。否则，该梅森数对应的p不是一个素数。

// 下面是用Go语言编写Lucas-Lehmer素性测试的示例代码：
// 卢卡斯-莱墨素性检验
func lucasLehmerTest(p int) bool {
	m := big.NewInt(0).Exp(big.NewInt(2), big.NewInt(int64(p)), nil)
	m.Sub(m, big.NewInt(1))
	s := big.NewInt(4)

	for i := 0; i < p-2; i++ {
		s.Mul(s, s)
		s.Sub(s, big.NewInt(2))
		s.Mod(s, m)
		if s.Cmp(big.NewInt(0)) == 0 {
			return true
		}
	}

	return s.Cmp(big.NewInt(0)) == 0
}

func main() {
	if true {
		for n := 3; n < 10000; n++ {

			isPrime := lucasLehmerTest(int(n))
			aa := big.NewInt(0).Exp(big.NewInt(2), big.NewInt(int64(n)), nil)
			aa.Sub(aa, big.NewInt(1))
			rr := aa.ProbablyPrime(0)
			if rr == isPrime {
				//fmt.Println("正确", n, isPrime)
				if isPrime {
					fmt.Printf("The number N = 2^%d - 1 is prime: %v\n", n, isPrime)
				}
			} else {
				fmt.Println("错误", n, isPrime)
				return
			}

			// if isPrime {
			// 	fmt.Printf("The number N = %d * 2^%d - 1 is prime: %v\n", k, n, isPrime)
			// }

		}
	}
	if false {
		p := 127 // 示例中选择的质数p
		isPrime := lucasLehmerTest(p)
		if isPrime {
			fmt.Printf("Mersenne number with p=%d is prime\n", p)
		} else {
			fmt.Printf("Mersenne number with p=%d is not prime\n", p)
		}
	}
}
