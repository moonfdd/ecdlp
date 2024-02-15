package main

import (
	"fmt"
	"math/big"
)

// https://en.wikipedia.org/wiki/Lucas%E2%80%93Lehmer%E2%80%93Riesel_test
// https://www.rieselprime.de/ziki/LLR
// https://link.springer.com/article/10.1007/BF01935653
// https://github.com/arcetri/goprime/blob/master/rieseltest/rieseltest.go
// https://www.emis.de/journals/INTEGERS/papers/n15/n15.pdf
// https://www.semanticscholar.org/topic/Lucas%E2%80%93Lehmer%E2%80%93Riesel-test/2238438

// Performs the Lucas-Lehmer-Riesel test for N = k * 2^n - 1.
// https://eprint.iacr.org/2023/195.pdf
func TestLLR(k, n int64) bool {
	N := new(big.Int).Sub(new(big.Int).Mul(big.NewInt(k), new(big.Int).Exp(big.NewInt(2), big.NewInt(n), nil)), big.NewInt(1))
	// N = k * 2^n - 1
	if true {
		kk := k % 6
		if (kk == 1 && n&1 == 0) || (kk == 5 && n&1 == 1) {
			if N.Cmp(big.NewInt(3)) == 0 {
				return true
			}
			return false
		}
	}

	s := big.NewInt(4)
	if k != 1 {
		P := big.NewInt(0)
		for {
			// fmt.Println(P, N)
			// fmt.Println(Jacobi(big.NewInt(0).Sub(P, big.NewInt(2)), N))
			// fmt.Println(Jacobi(big.NewInt(0).Add(P, big.NewInt(2)), N))
			// fmt.Println("----------------")
			if Jacobi(big.NewInt(0).Sub(P, big.NewInt(2)), N).Cmp(big.NewInt(1)) == 0 && Jacobi(big.NewInt(0).Add(P, big.NewInt(2)), N).Cmp(big.NewInt(-1)) == 0 {
				break
			}
			P.Add(P, big.NewInt(1))
		}

		// fmt.Println("P = ", P)
		// Start with s = 4.
		// s := big.NewInt(4)
		// s := big.NewInt(5778)
		ll := &LucasParam{P, big.NewInt(1)}
		// s := big.NewInt(5778)
		_, s = ll.GetUnAndVnMod(big.NewInt(k), N)
		// fmt.Println("s = ", s)
		// s.Mod(s, N)
		// fmt.Println("s2 = ", s)
	}
	// fmt.Println("s2 = ", s)
	// Repeat for n-2 iterations.
	for i := int64(0); i < n-2; i++ {
		// s = (s^2 - 2) % N
		s.Exp(s, big.NewInt(2), N).Sub(s, big.NewInt(2)).Mod(s, N)
	}

	// If we reach here, N passed the test - it might be prime.
	return s.Cmp(big.NewInt(0)) == 0
}

// 来自solover_strassen
func Jacobi(a *big.Int, n *big.Int) *big.Int {
	a = new(big.Int).Set(a)
	n = new(big.Int).Set(n)

	//a==1或者n==1
	if a.Cmp(big.NewInt(1)) == 0 || n.Cmp(big.NewInt(1)) == 0 {
		return big.NewInt(1)
	}

	// a=0
	if a.Cmp(big.NewInt(0)) == 0 {
		if n.Cmp(big.NewInt(-1)) == 0 {
			return big.NewInt(1)
		} else {
			return big.NewInt(0)
		}
	}

	// n=0
	if n.Cmp(big.NewInt(0)) == 0 {
		if a.Cmp(big.NewInt(-1)) == 0 {
			return big.NewInt(1)
		} else {
			return big.NewInt(0)
		}
	}

	//n=-1
	if n.Cmp(big.NewInt(-1)) == 0 {
		if a.Cmp(big.NewInt(0)) < 0 {
			return big.NewInt(-1)
		} else {
			return big.NewInt(1)
		}
	}
	//a=-1
	if a.Cmp(big.NewInt(-1)) == 0 {
		isNeg := false
		if n.Cmp(big.NewInt(0)) < 0 {
			n = big.NewInt(0).Neg(n)
			isNeg = true
		}
		n = big.NewInt(0).Set(n)
		for n.Bit(0) == 0 {
			n.Rsh(n, 1)
		}
		n_1 := big.NewInt(0).Sub(n, big.NewInt(1))
		n_1.Rsh(n_1, 1)
		ans := big.NewInt(0).Exp(big.NewInt(-1), n_1, nil)
		if isNeg {
			ans.Neg(ans)
		}
		return ans
	}

	//a和n的绝对值必都大于等于2

	//a<0
	if a.Cmp(big.NewInt(0)) < 0 {
		return big.NewInt(0).Mul(Jacobi(big.NewInt(-1), n), Jacobi(big.NewInt(0).Neg(a), n))
	}

	//n<0，直接取反
	if n.Cmp(big.NewInt(0)) < 0 {
		n.Neg(n)
	}

	//最大公约数
	if new(big.Int).GCD(nil, nil, a, n).Cmp(big.NewInt(1)) > 0 {
		return big.NewInt(0)
	}

	if n.Bit(0) == 0 { //n是偶数
		if n.Cmp(big.NewInt(2)) == 0 {
			am := big.NewInt(0).Mod(a, big.NewInt(8))
			if am.Cmp(big.NewInt(3)) == 0 || am.Cmp(big.NewInt(5)) == 0 {
				return big.NewInt(-1)
			} else {
				return big.NewInt(1)
			}
		}
		n.Rsh(n, 1)
		return big.NewInt(0).Mul(Jacobi(a, big.NewInt(2)), Jacobi(a, n))
	}

	//通例n>=3的奇数
	i := big.NewInt(0)
	s := big.NewInt(0)
	a = new(big.Int).Mod(a, n)
	temp := new(big.Int).Set(a)

	for temp.Bit(0) == 0 {
		temp.Rsh(temp, 1)
		i.Add(i, big.NewInt(1))
	}
	if i.Bit(0) == 0 {
		s.SetInt64(1)
	} else {
		if new(big.Int).Mod(n, big.NewInt(8)).Cmp(big.NewInt(1)) == 0 || new(big.Int).Mod(n, big.NewInt(8)).Cmp(big.NewInt(7)) == 0 {
			s.SetInt64(1)
		} else {
			s.SetInt64(-1)
		}
	}
	if new(big.Int).Mod(n, big.NewInt(4)).Cmp(big.NewInt(3)) == 0 && new(big.Int).Mod(temp, big.NewInt(4)).Cmp(big.NewInt(3)) == 0 {
		s.Neg(s)
	}
	if temp.Cmp(big.NewInt(1)) != 0 {
		n1 := new(big.Int).Mod(n, temp)
		temp1 := new(big.Int).Mul(s, Jacobi(n1, temp))
		return temp1
	} else {
		if i.Bit(0) == 0 {
			return big.NewInt(1)
		} else {
			if new(big.Int).Mod(n, big.NewInt(8)).Cmp(big.NewInt(1)) == 0 || new(big.Int).Mod(n, big.NewInt(8)).Cmp(big.NewInt(7)) == 0 {
				return big.NewInt(1)
			} else {
				return big.NewInt(-1)
			}
		}
	}
}

type LucasParam struct {
	P *big.Int
	Q *big.Int
}

func (that *LucasParam) GetUnAndVnMod(k *big.Int, N *big.Int) (*big.Int, *big.Int) {
	if k.Cmp(big.NewInt(0)) == 0 {
		return big.NewInt(0), big.NewInt(2)
	}
	if k.Cmp(big.NewInt(1)) == 0 {
		return big.NewInt(1), big.NewInt(0).Mod(that.P, N)
	}
	ansU := big.NewInt(0)
	ansV := big.NewInt(0)
	var tempAnsU *big.Int
	var tempAndV *big.Int
	doubleU := big.NewInt(1)
	doubleV := big.NewInt(0).Set(that.P)
	temp2U := big.NewInt(0)
	temp2V := big.NewInt(0)
	d := big.NewInt(0).Exp(that.P, big.NewInt(2), nil)
	d.Sub(d, big.NewInt(0).Lsh(that.Q, 2))
	d.Mod(d, N)
	kBitLen := k.BitLen()
	for i := 0; i < kBitLen; i++ {
		if k.Bit(i) != 0 {
			if tempAnsU == nil {
				tempAnsU = big.NewInt(0).Set(doubleU)
				tempAndV = big.NewInt(0).Set(doubleV)
			} else {
				tempAnsU.Mul(ansU, doubleV).Add(tempAnsU, big.NewInt(0).Mul(ansV, doubleU))
				tempAnsU.Mul(tempAnsU, big.NewInt(0).ModInverse(big.NewInt(2), N))

				tempAndV.Mul(d, ansU).Mul(tempAndV, doubleU).Add(tempAndV, big.NewInt(0).Mul(ansV, doubleV))
				tempAndV.Mul(tempAndV, big.NewInt(0).ModInverse(big.NewInt(2), N))

				tempAnsU.Mod(tempAnsU, N)
				tempAndV.Mod(tempAndV, N)
			}
			ansU.Set(tempAnsU)
			ansV.Set(tempAndV)
		}
		temp2U.Mul(doubleU, doubleV)
		temp2V.Mul(d, doubleU).Mul(temp2V, doubleU).Add(temp2V, big.NewInt(0).Exp(doubleV, big.NewInt(2), nil))
		temp2V.Mul(temp2V, big.NewInt(0).ModInverse(big.NewInt(2), N))

		temp2U.Mod(temp2U, N)
		temp2V.Mod(temp2V, N)

		doubleU.Set(temp2U)
		doubleV.Set(temp2V)

	}
	return ansU, ansV
}

func (that *LucasParam) GetUnAndVn(k *big.Int) (*big.Int, *big.Int) {
	if k.Cmp(big.NewInt(0)) == 0 {
		return big.NewInt(0), big.NewInt(2)
	}
	if k.Cmp(big.NewInt(1)) == 0 {
		return big.NewInt(1), big.NewInt(0).Set(that.P)
	}
	ansU := big.NewInt(0)
	ansV := big.NewInt(0)
	var tempAnsU *big.Int
	var tempAndV *big.Int
	doubleU := big.NewInt(1)
	doubleV := big.NewInt(0).Set(that.P)
	temp2U := big.NewInt(0)
	temp2V := big.NewInt(0)
	d := big.NewInt(0).Exp(that.P, big.NewInt(2), nil)
	d.Sub(d, big.NewInt(0).Lsh(that.Q, 2))
	kBitLen := k.BitLen()
	for i := 0; i < kBitLen; i++ {
		if k.Bit(i) != 0 {
			if tempAnsU == nil {
				tempAnsU = big.NewInt(0).Set(doubleU)
				tempAndV = big.NewInt(0).Set(doubleV)
			} else {
				tempAnsU.Mul(ansU, doubleV).Add(tempAnsU, big.NewInt(0).Mul(ansV, doubleU)).Rsh(tempAnsU, 1)
				tempAndV.Mul(d, ansU).Mul(tempAndV, doubleU).Add(tempAndV, big.NewInt(0).Mul(ansV, doubleV)).Rsh(tempAndV, 1)
			}
			ansU.Set(tempAnsU)
			ansV.Set(tempAndV)
		}
		temp2U.Mul(doubleU, doubleV)
		temp2V.Mul(d, doubleU).Mul(temp2V, doubleU).Add(temp2V, big.NewInt(0).Exp(doubleV, big.NewInt(2), nil)).Rsh(temp2V, 1)
		doubleU.Set(temp2U)
		doubleV.Set(temp2V)

	}
	return ansU, ansV
}

func main() {
	if false {
		ll := LucasParam{big.NewInt(1), big.NewInt(-1)}
		for k := big.NewInt(0); k.Cmp(big.NewInt(11)) <= 0; k.Add(k, big.NewInt(1)) {
			fmt.Println(ll.GetUnAndVn(k))
		}
		return
	}
	if true {
		// Example: Test LLR for k = 3, n = 5.
		// k := int64(3)
		// n := int64(5)
		k := int64(3)
		n := int64(5)
		primeCount := 0
		for n = 1; n < 6; n++ {
			for k = 1; k < 1<<n; k += 2 {
				// fmt.Println("开始", k, n)
				isPrime := TestLLR(k, n)
				aa := new(big.Int).Sub(new(big.Int).Mul(big.NewInt(k), new(big.Int).Exp(big.NewInt(2), big.NewInt(n), nil)), big.NewInt(1))
				rr := aa.ProbablyPrime(0)
				if rr == isPrime {
					//fmt.Printf("正确The number N = %d * 2^%d - 1 is prime: %v %v %v\n", k, n, isPrime, rr, aa)
				} else {
					fmt.Printf("错误The number N = %d * 2^%d - 1 is prime: %v %v %v\n", k, n, isPrime, rr, aa)
					return
				}

				if isPrime {
					primeCount++
					fmt.Printf("The number N = %d * 2^%d - 1 is prime: %v %v\n", k, n, isPrime, aa)
				}
				// break
			}
		}
		fmt.Println("完全正确", primeCount)
	}
}
