package main

import (
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"time"
)

// https://blog.51cto.com/u_13424/6307746
// https://blog.csdn.net/weixin_39695712/article/details/107054736

// // 分解质因数
// func FactorInteger(num *big.Int) (factorMap map[string]*big.Int) {
// 	num = big.NewInt(0).Add(num, big.NewInt(0))
// 	factorMap = make(map[string]*big.Int)
// 	if num.Cmp(big.NewInt(1)) == 0 {
// 		return
// 	}
// 	for i := 0; i < 3; i++ {
// 		if big.NewInt(0).And(num, big.NewInt(1)).Cmp(big.NewInt(0)) != 0 {
// 			break
// 		}
// 		if factorMap["2"] == nil {
// 			factorMap["2"] = big.NewInt(1)
// 		} else {
// 			factorMap["2"].Add(factorMap["2"], big.NewInt(1))
// 		}
// 		num.Div(num, big.NewInt(2))

// 	}
// 	for i := big.NewInt(3); ; i.Add(i, big.NewInt(2)) {
// 		for big.NewInt(0).Mod(num, i).Cmp(big.NewInt(0)) == 0 {
// 			if factorMap[i.Text(10)] == nil {
// 				factorMap[i.Text(10)] = big.NewInt(1)
// 			} else {
// 				factorMap[i.Text(10)].Add(factorMap[i.Text(10)], big.NewInt(1))
// 			}
// 			num.Div(num, i)
// 		}
// 		if num.ProbablyPrime(0) {
// 			if factorMap[num.Text(10)] == nil {
// 				factorMap[num.Text(10)] = big.NewInt(1)
// 			} else {
// 				factorMap[num.Text(10)].Add(factorMap[num.Text(10)], big.NewInt(1))
// 			}
// 			break
// 		}
// 		if num.Cmp(big.NewInt(1)) == 0 {
// 			return
// 		}
// 	}
// 	return
// }

// 根据课本中的伪代码，使用Euclid算法求最大公约数。
func Euclid(a, b int) int {
	if a >= b {
		if b == 0 {
			return a
		}
		return Euclid(b, a%b)
	} else {
		if a == 0 {
			return b
		}
		return Euclid(a, b%a)
	}
}

func MultiplicativeOrder(n, r int) int {
	if Euclid(n, r) != 1 { // 如果n,r不互质，不存在Multiplicative_Order
		return -1
	} else {
		k := 1
		temp := 1
		for {
			// 根据A(k) = a^k%r = (a * a^(k-1))%r = (a%r * a^(k-1)%r)%r = (A(k-1) * a%r)%r
			temp = (temp * (n % r)) % r
			if temp == 1 {
				break
			} else {
				k++
			}
		}
		return k
	}
}

// 返回true表示素数，返回false表示合数。
func AKS_IsPrimality(num int) bool {
	// if num == 1 {
	// 	return false
	// }
	// step1: 如果num是幂，则返回合数
	// fmt.Println("第1步")
	if IsPower(num) {
		return false
	}

	// step2: 找到最小的r，使得ord_r(num) > log2(num)
	// fmt.Println("第2步")
	r := getSmallestR(num)
	// fmt.Println(r)

	// step3: 判断是否存在a ≤ r，使得1 < gcd(a, num) < num
	// fmt.Println("第3步")
	if aR(num, r) {
		return false
	}

	// step4: 如果num <= r，则返回素数
	// fmt.Println("第4步")
	if num <= r {
		return true
	}

	// step5: 判断是否满足同余条件
	// fmt.Println("第5步")
	for a := 1; a <= int(math.Floor(math.Sqrt(float64(EulerTotient(r)))*math.Log2(float64(num)))); a++ {
		// f0 := big.NewInt(0).Exp(big.NewInt(int64(a)), big.NewInt(int64(num)), big.NewInt(int64(r)))
		// f := f0.Int64()
		f := Congruence(a, num, r)
		if f == 0 {
			return false
		}
	}

	// step6: 返回素数
	// fmt.Println("第6步")
	return true
}

// 判断是否为幂
func IsPower(num int) bool {
	if num == 1 {
		return true
	} else {
		for a := 2; a <= int(math.Floor(math.Sqrt(float64(num)))); a++ {
			temp := num
			for temp%a == 0 {
				temp /= a
				if temp == 1 {
					return true
				}
			}
		}
	}
	return false
}

// 找到最小的r，使得ord_r(num) > log2(num)
func getSmallestR(num int) int {
	r := 2
	for {
		if Euclid(num, r) == 1 {
			multiOrder := MultiplicativeOrder(num, r)
			// fmt.Println("multiOrder = ", multiOrder, int(math.Pow(math.Log2(float64(num)), 2)), r)
			if multiOrder > int(math.Pow(math.Log2(float64(num)), 2)) {
				// fmt.Println("r = ", r)
				return r
			}
		}
		// fmt.Println("r0 = ", r, int(math.Pow(math.Log2(float64(num)), 2)), num)
		r++

	}
}

// 判断是否存在a ≤ r，使得1 < gcd(a, num) < num
func aR(num, r int) bool {
	for a := 1; a <= r; a++ {
		gcd := Euclid(a, num)
		if gcd > 1 && gcd < num {
			return true
		}
	}
	return false
}

// 计算欧拉函数值
func EulerTotient(r int) int {
	count := 0
	for i := 1; i <= r; i++ {
		if Euclid(i, r) == 1 {
			count++
		}
	}
	return count
}

func Congruence(a, n, r int) int {
	// modN := n
	b := pow(2, r) - 1
	c := 1 - a
	f := powerMod(c, n, b)
	e := 1
	g := powerMod(e, n, b)
	g = g - a
	if f == g {
		//fmt.Println("prime")
		return 1
	} else {
		//fmt.Println("not prime")
		return 0
	}
}

func pow(base, exponent int) int {
	result := 1
	for i := 0; i < exponent; i++ {
		result *= base
	}
	return result
}

func powerMod(base, exponent, modulus int) int {
	result := 1
	for exponent > 0 {
		if exponent%2 == 1 {
			result = (result * base) % modulus
		}
		base = (base * base) % modulus
		exponent /= 2
	}
	return result
}

// 同余判断
// func congruence(a int64, n, r *big.Int) int {
// 	b := new(big.Int).Exp(big.NewInt(10), r, nil)
// 	b.Sub(b, big.NewInt(1))
// 	c := big.NewInt(1).Sub(big.NewInt(10), big.NewInt(a))
// 	f := big.NewInt(0).Exp(c, n, b)
// 	e := big.NewInt(1)
// 	g := big.NewInt(0).Exp(e, n, b)
// 	g.Sub(g, big.NewInt(a))
// 	if f.Cmp(g) == 0 {
// 		return 1
// 	} else {
// 		return 0
// 	}
// }

// 产生随机数
func MyRand(n int) int64 {
	rand.Seed(time.Now().UnixNano())
	if n == 1 {
		return rand.Int63n(10)
	} else if n == 2 {
		return int64((rand.Intn(9)+1)*10 + rand.Intn(10))
	}
	k := int64(math.Pow10(n/2 + 1)) // k = 10^(n/2+1)
	x := (int64(rand.Intn(9)+1)*k + (int64(rand.Intn(int(k)))*int64(rand.Intn(int(k))))%k)
	if x%2 == 0 {
		x = x + 1
	}
	return x
}

func main() {

	for num := 1; num < 10000; num += 2 {
		isPrime := AKS_IsPrimality(num)
		isPrime2 := big.NewInt(int64(num)).ProbablyPrime(0)
		if isPrime {
			if isPrime == isPrime2 {

			} else {
				fmt.Println("测试失败", num)
				return
			}
		} else {
			if isPrime == isPrime2 {

			} else {
				fmt.Println("2测试失败", num)
				return
			}
		}
	}
	fmt.Println("测试成功")
	return

	// num := 123456789
	num := rand.Int63n(100000000000)
	num = 2345563853
	fmt.Println(big.NewInt(num).ProbablyPrime(0))
	isPrime := AKS_IsPrimality(int(num))
	fmt.Printf("%d is prime: %v", num, isPrime)
}
