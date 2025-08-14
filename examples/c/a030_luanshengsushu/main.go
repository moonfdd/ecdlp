package main

import (
	"fmt"
	"math"
	"math/big"
)

// 判断是否为完全平方数
func isPerfectSquare(x int) bool {
	y := int(math.Sqrt(float64(x)))
	return y*y == x
}

// 判断是否能表示为 4^k*(8m+7)
func checkAnswer4(x int) bool {
	for x%4 == 0 {
		x /= 4
	}
	return x%8 == 7
}

func numSquares(n int) int {
	if isPerfectSquare(n) {
		return 1
	}
	if checkAnswer4(n) {
		return 4
	}
	for i := 1; i*i <= n; i++ {
		j := n - i*i
		if isPerfectSquare(j) {
			return 2
		}
	}
	return 3
}

func numSquares2(n *big.Int) (ans int) {
	n = big.NewInt(0).Set(n)
	ans = 1
	for n.Bit(0) == 0 && n.Bit(1) == 0 {
		n.Rsh(n, 2)
	}
	if n.Cmp(big.NewInt(1)) == 0 {
		return
	}
	if big.NewInt(0).Mod(n, big.NewInt(8)).Cmp(big.NewInt(7)) == 0 {
		ans = 4
		return
	}
	if big.NewInt(0).Mod(n, big.NewInt(4)).Cmp(big.NewInt(3)) == 0 {
		ans = 3
		return
	}
	m := FactorInteger(n)
	pb := big.NewInt(0)
	for p, c := range m {
		if c.Bit(0) == 0 {
			continue
		}
		switch ans {
		case 1:
			pb.SetString(p, 10)
			ans = primeSquares(pb)
		case 2:
			pb.SetString(p, 10)
			ans2 := primeSquares(pb)
			switch ans2 {
			case 2:
				ans = 2
			case 3:
				ans = 3
			case 4:
				ans = 3
			default:
			}
		case 3:
			pb.SetString(p, 10)
			ans2 := primeSquares(pb)
			switch ans2 {
			case 2:
				ans = 3
			case 3:
				ans = 3
			case 4:
				ans = 3
			default:
			}
		default:
			pb.SetString(p, 10)
			ans2 := primeSquares(pb)
			switch ans2 {
			case 2:
				ans = 3
			case 3:
				ans = 3
			case 4:
				ans = 3
			default:
			}
		}
	}
	return
}

// 分解质因数
func FactorInteger(num *big.Int) (factorMap map[string]*big.Int) {
	num = big.NewInt(0).Add(num, big.NewInt(0))
	factorMap = make(map[string]*big.Int)
	if num.Cmp(big.NewInt(1)) == 0 {
		return
	}
	for {
		if big.NewInt(0).And(num, big.NewInt(1)).Cmp(big.NewInt(0)) != 0 {
			break
		}
		if factorMap["2"] == nil {
			factorMap["2"] = big.NewInt(1)
		} else {
			factorMap["2"].Add(factorMap["2"], big.NewInt(1))
		}
		num.Div(num, big.NewInt(2))

	}
	for i := big.NewInt(3); ; i.Add(i, big.NewInt(2)) {
		for big.NewInt(0).Mod(num, i).Cmp(big.NewInt(0)) == 0 {
			if factorMap[i.Text(10)] == nil {
				factorMap[i.Text(10)] = big.NewInt(1)
			} else {
				factorMap[i.Text(10)].Add(factorMap[i.Text(10)], big.NewInt(1))
			}
			num.Div(num, i)
		}
		if num.ProbablyPrime(0) {
			if factorMap[num.Text(10)] == nil {
				factorMap[num.Text(10)] = big.NewInt(1)
			} else {
				factorMap[num.Text(10)].Add(factorMap[num.Text(10)], big.NewInt(1))
			}
			break
		}
		if num.Cmp(big.NewInt(1)) == 0 {
			return
		}
	}
	return
}

func primeSquares(p *big.Int) (ans int) {
	if p.Cmp(big.NewInt(2)) == 0 {
		ans = 2
		return
	}
	if big.NewInt(0).Mod(p, big.NewInt(8)).Cmp(big.NewInt(7)) == 0 {
		ans = 4
		return
	}
	if big.NewInt(0).Mod(p, big.NewInt(4)).Cmp(big.NewInt(3)) == 0 {
		ans = 3
		return
	}
	ans = 2
	return
}

func main() {
	// 4n+1
	// 12
	// 12 23
	// 12 32 22
	// 12 32 23 23
	// 12 23 22 32 22
	// 12 33 23 22 22 33
	// 12 32 23 22 32 33 32
	// 12 23 22 33 22 22 32 23
	if false {
		count := 0
		for i := 1; i <= 1000; i += 1 {
			ret := numSquares2(big.NewInt(int64(i)))
			if ret == 1 {
				fmt.Println()
			}
			count++
			fmt.Print(ret)
			// if count == 2 {
			// 	fmt.Print(" ")
			// 	count = 0
			// }
		}
		return
	}
	// 4n+1
	// 12
	// 12 23
	// 12 32 22
	// 12 32 23 23
	// 12 23 22 32 22
	// 12 33 23 22 22 33
	// 12 32 23 22 32 33 32
	// 12 23 22 33 22 22 32 23
	if true {
		count := 0
		for i := 1; i <= 1000; i += 4 {
			ret := numSquares2(big.NewInt(int64(i)))
			if ret == 1 {
				fmt.Println()
			}
			count++
			fmt.Print(ret)
			if count == 2 {
				fmt.Print(" ")
				count = 0
			}
		}
		return
	}
	// 2n+1
	// 1324
	// 1324 2334
	// 1324 3324 2324
	// 1324 3324 2334 2334
	// 1324 2334 2324 3324 2324
	// 1324 3334 2334 2324 2324 3334
	if true {
		count := 0
		for i := 1; i <= 1000; i += 2 {
			ret := numSquares2(big.NewInt(int64(i)))
			if ret == 1 {
				fmt.Println()
			}
			count++
			fmt.Print(ret)
			if count == 4 {
				fmt.Print(" ")
				count = 0
			}
		}
		return
	}
	// 对数器
	if true {
		count := 0
		for i := 1; i <= 1000000; i += 1 {
			ans1 := numSquares(i)
			ans2 := numSquares2(big.NewInt(int64(i)))
			if ans1 != ans2 {
				count++
				fmt.Println(i, ans1, ans2)
			}
		}
		fmt.Println("错误：", count)
		return
	}
	// 测试并观察规律
	if true {
		for i := 1; i <= 100000; i += 2 {
			r := numSquares(i)
			if r == 3 {
				// if big.NewInt(int64(i)).ProbablyPrime(0) {
				// 	continue
				// }
				fmt.Println(i, numSquares(i), i%4)
			} else if r == 2 {
				// if big.NewInt(int64(i)).ProbablyPrime(0) {
				// 	continue
				// }
				fmt.Println(i, numSquares(i), i%4)

			} else {
				fmt.Println(i, numSquares(i))
			}
		}
		return
	}
	for i := big.NewInt(1); i.Cmp(big.NewInt(300)) <= 0; i.Add(i, big.NewInt(1)) {
		left := big.NewInt(0).Exp(big.NewInt(2), i, nil)
		right := big.NewInt(0).Lsh(left, 1)
		// left.Exp(big.NewInt(2), left, nil)
		// right.Exp(big.NewInt(2), right, nil)
		n := big.NewInt(0)
		count := 0
		for n.Set(left); n.Cmp(right) < 0; n.Add(n, big.NewInt(1)) {
			isPrime1 := n.ProbablyPrime(0)
			isPrime2 := big.NewInt(0).Add(n, big.NewInt(2)).ProbablyPrime(0)
			if isPrime1 && isPrime2 {
				count++
				// fmt.Println(i,"是素数", left,n)
				// break
			}
		}
		fmt.Println(i, ":", count)

		// if count==0 {
		// 	fmt.Println("没有素数", i)
		// }
	}
	fmt.Println("完成")
}
