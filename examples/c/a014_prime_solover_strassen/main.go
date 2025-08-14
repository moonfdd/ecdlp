package main

import (
	"fmt"
	"math/big"

	"github.com/moonfdd/ecdlp"
)

// ecpp.pdf 55页 101
// 看到P59
func SoloveyStrassen(num *big.Int) bool {
	if num.Cmp(big.NewInt(1)) == 0 {
		return false
	}
	if num.Cmp(big.NewInt(2)) == 0 {
		return true
	}
	if num.Bit(0) == 0 {
		return false
	}
	// if num.Cmp(big.NewInt(3)) == 0 {
	// 	return true
	// }
	// if big.NewInt(0).Mod(num, big.NewInt(3)).Cmp(big.NewInt(0)) == 0 {
	// 	return false
	// }
	t := big.NewInt(0).Add(num, big.NewInt(-1))
	t.Rsh(t, 1)
	j := Jacobi(big.NewInt(2), num)
	j.Mod(j, num)

	if big.NewInt(0).Exp(big.NewInt(2), t, num).Cmp(j) == 0 {
		return true
	} else {
		return false
	}
}

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

func Jacobi3(a *big.Int, n *big.Int) *big.Int {
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
	if false {
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

// http://jpenne.free.fr/llr4/llr405src.zip jacobi.c

func Jacobi2(a *big.Int, b *big.Int) *big.Int {
	/* Computes Jacobi (a, b) */

	jdvs := big.NewInt(0)
	jdvd := big.NewInt(0)
	jq := big.NewInt(0)
	jr := big.NewInt(0)
	resul := big.NewInt(0)
	s := big.NewInt(0)
	t := big.NewInt(0)
	u := big.NewInt(0)
	v := big.NewInt(0)
	jdvs.Set(a)
	jdvd.Set(b)
	resul.Set(big.NewInt(1))
	zero := big.NewInt(0)
	one := big.NewInt(1)
	for jdvs.Cmp(big.NewInt(0)) != 0 {
		if jdvs.Cmp(big.NewInt(1)) == 0 { /* Finished ! */
			return resul
		} else {
			v.Set(jdvd)
			s.Set(v).Sub(s, big.NewInt(1)).Rsh(s, 1)
			t.Set(v).Add(t, big.NewInt(1)).Rsh(t, 1)
			for new(big.Int).And(jdvs, one).Cmp(zero) == 0 { // While dvs is even
				if new(big.Int).And(t, one).Cmp(one) == 0 {
					if new(big.Int).And(new(big.Int).Rsh(s, 1), one).Cmp(one) == 0 {
						resul.Neg(resul)
					}
				} else {
					if new(big.Int).And(new(big.Int).Rsh(t, 1), one).Cmp(one) == 0 {
						resul.Neg(resul)
					}
				}
				jdvs.Rsh(jdvs, 1) // dvs /= 2
			}

			if jdvs.Cmp(one) == 0 {
				return resul // Finished!
			} else {
				//u.Rsh(new(big.Int).Sub(jdvs, one), 1) // (dvs-1)/2
				u.Set(jdvs).Sub(u, big.NewInt(1)).Rsh(u, 1)

				sAndU := new(big.Int).And(s, u)
				if sAndU.And(sAndU, one).Cmp(one) == 0 {
					resul.Neg(resul)
				}

				jq.Div(jdvd, jdvs)
				jr.Mod(jdvd, jdvs)

				jdvd.Set(jdvs)
				jdvs.Set(jr)
			}
		}
	}
	return (jdvd) /* a and b are not coprime, */
	/* so, return their gcd. */
}

func main() {
	if false {
		str1 := "0 1 0 0 0 -1 0 1 0 0 0 -1 0 1 0 0 0 -1 0 1 0 0 0 -1 0 -1 0 1 -1 1 1 1 -1 -1 -1 1 -1 0 1 -1 1 1 1 -1 -1 -1 1 -1 0 1 0 -1 0 -1 0 -1 0 0 0 1 0 -1 0 1 0 -1 0 0 0 1 0 1 0 1 0 0 1 -1 0 -1 1 0 -1 -1 0 -1 -1 0 1 1 0 1 1 0 -1 1 0 1 -1 0 0 -1 0 -1 0 1 0 1 0 -1 0 -1 0 1 0 1 0 -1 0 -1 0 1 0 1 0 1 -1 1 -1 -1 0 1 1 -1 1 -1 -1 0 1 1 -1 1 -1 -1 0 1 1 -1 1 -1 0 -1 0 0 0 -1 0 -1 0 0 0 -1 0 1 0 0 0 1 0 1 0 0 0 1 0 -1 1 0 -1 1 -1 1 0 -1 -1 1 -1 0 1 -1 1 1 0 -1 1 -1 1 0 -1 1 0 1 0 -1 0 1 0 -1 0 1 0 -1 0 1 0 -1 0 1 0 -1 0 1 0 -1 0 0 1 -1 0 1 -1 0 1 -1 0 1 -1 0 1 -1 0 1 -1 0 1 -1 0 1 -1 0 0 -1 0 -1 0 1 0 1 0 -1 0 -1 0 1 0 1 0 -1 0 -1 0 1 0 1 0 1 1 -1 -1 -1 1 1 -1 -1 1 -1 -1 1 1 1 -1 1 1 -1 -1 1 1 1 -1 -1 0 0 0 0 0 0 0 0 0 0 0 1 0 1 0 0 0 0 0 0 0 0 0 0 0 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 0 -1 0 1 0 1 0 -1 0 -1 0 1 0 1 0 -1 0 -1 0 1 0 1 0 -1 0 0 1 1 0 -1 -1 0 -1 1 0 -1 1 0 1 -1 0 1 -1 0 -1 -1 0 1 1 0 0 1 0 1 0 1 0 1 0 1 0 1 0 1 0 1 0 1 0 1 0 1 0 1 0 -1 1 0 1 -1 -1 1 0 1 -1 -1 1 0 1 -1 -1 1 0 1 -1 -1 1 0 1 -1 0 -1 0 0 0 -1 0 1 0 0 0 1 0 1 0 0 0 1 0 -1 0 0 0 -1 0 1 -1 -1 1 1 0 1 -1 1 1 1 1 0 1 1 1 1 -1 1 0 1 1 -1 -1 1 0 -1 0 1 0 1 0 -1 0 -1 0 1 0 1 0 -1 0 -1 0 1 0 1 0 -1 0 0 1 1 0 1 1 0 1 1 0 1 1 0 1 1 0 1 1 0 1 1 0 1 1 0 0 -1 0 1 0 -1 0 0 0 1 0 1 0 1 0 1 0 0 0 -1 0 1 0 -1 0 -1 0 -1 1 -1 1 1 1 1 -1 -1 1 0 1 -1 -1 1 1 1 1 -1 1 -1 0 -1 0 1 0 0 0 -1 0 -1 0 0 0 1 0 1 0 0 0 -1 0 -1 0 0 0 1 0"

		str2 := "0 1 0 0 0 -1 0 1 0 0 0 -1 0 1 0 0 0 -1 0 1 0 0 0 -1 0 -1 0 1 -1 1 1 1 -1 -1 -1 1 -1 0 1 -1 1 1 1 -1 -1 -1 1 -1 0 1 0 -1 0 -1 0 -1 0 0 0 1 0 -1 0 1 0 -1 0 0 0 1 0 1 0 1 0 0 1 -1 0 -1 1 0 -1 -1 0 -1 -1 0 1 1 0 1 1 0 -1 1 0 1 -1 0 0 -1 0 -1 0 1 0 1 0 -1 0 -1 0 1 0 1 0 -1 0 -1 0 1 0 1 0 1 -1 1 -1 -1 0 1 1 -1 1 -1 -1 0 1 1 -1 1 -1 -1 0 1 1 -1 1 -1 0 -1 0 0 0 -1 0 -1 0 0 0 -1 0 1 0 0 0 1 0 1 0 0 0 1 0 -1 1 0 -1 1 -1 1 0 -1 -1 1 -1 0 1 -1 1 1 0 -1 1 -1 1 0 -1 1 0 1 0 -1 0 1 0 -1 0 1 0 -1 0 1 0 -1 0 1 0 -1 0 1 0 -1 0 0 1 -1 0 1 -1 0 1 -1 0 1 -1 0 1 -1 0 1 -1 0 1 -1 0 1 -1 0 0 -1 0 -1 0 1 0 1 0 -1 0 -1 0 1 0 1 0 -1 0 -1 0 1 0 1 0 1 1 -1 -1 -1 1 1 -1 -1 1 -1 -1 1 1 1 -1 1 1 -1 -1 1 1 1 -1 -1 0 0 0 0 0 0 0 0 0 0 0 1 0 1 0 0 0 0 0 0 0 0 0 0 0 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 0 -1 0 1 0 1 0 -1 0 -1 0 1 0 1 0 -1 0 -1 0 1 0 1 0 -1 0 0 1 1 0 -1 -1 0 -1 1 0 -1 1 0 1 -1 0 1 -1 0 -1 -1 0 1 1 0 0 1 0 1 0 1 0 1 0 1 0 1 0 1 0 1 0 1 0 1 0 1 0 1 0 -1 1 0 1 -1 -1 1 0 1 -1 -1 1 0 1 -1 -1 1 0 1 -1 -1 1 0 1 -1 0 -1 0 0 0 -1 0 1 0 0 0 1 0 1 0 0 0 1 0 -1 0 0 0 -1 0 1 -1 -1 1 1 0 1 -1 1 1 1 1 0 1 1 1 1 -1 1 0 1 1 -1 -1 1 0 -1 0 1 0 1 0 -1 0 -1 0 1 0 1 0 -1 0 -1 0 1 0 1 0 -1 0 0 1 1 0 1 1 0 1 1 0 1 1 0 1 1 0 1 1 0 1 1 0 1 1 0 0 -1 0 1 0 -1 0 0 0 1 0 1 0 1 0 1 0 0 0 -1 0 1 0 -1 0 -1 0 -1 1 -1 1 1 1 1 -1 -1 1 0 1 -1 -1 1 1 1 1 -1 1 -1 0 -1 0 1 0 0 0 -1 0 -1 0 0 0 1 0 1 0 0 0 -1 0 -1 0 0 0 1 0"
		fmt.Println(str1 == str2)
		return
	}
	if false {
		for a := big.NewInt(-12); a.Cmp(big.NewInt(12)) <= 0; a.Add(a, big.NewInt(1)) {
			for p := big.NewInt(-12); p.Cmp(big.NewInt(12)) <= 0; p.Add(p, big.NewInt(1)) {
				r := Jacobi(a, p)
				fmt.Print(r, " ")
			}
		}
		return
	}
	if true {
		//jacobi jacobi3
		if true {
			for a := big.NewInt(-12); a.Cmp(big.NewInt(12)) <= 0; a.Add(a, big.NewInt(1)) {
				for p := big.NewInt(-12); p.Cmp(big.NewInt(12)) <= 0; p.Add(p, big.NewInt(1)) {
					r := ecdlp.Jacobi(a, p)
					r2 := Jacobi3(a, p)
					if r.Cmp(r2) != 0 {
						fmt.Println("错误", a, p, r, r2)
						return
					} else {
						// fmt.Println("正确", a, p, r, r2)
					}
				}
			}
			return
		}
	}
	if true {
		for a := big.NewInt(-12); a.Cmp(big.NewInt(12)) <= 0; a.Add(a, big.NewInt(1)) {
			for p := big.NewInt(-12); p.Cmp(big.NewInt(12)) <= 0; p.Add(p, big.NewInt(1)) {
				r := Jacobi(a, p)
				r2 := Jacobi2(a, p)
				if r.Cmp(r2) != 0 {
					fmt.Println("错误", a, p, r, r2)
				} else {
					fmt.Println("正确", a, p, r, r2)
				}
			}
		}
		return
	}

	if true {
		num := big.NewInt(0)
		num.SetString("1", 10)
		count := 0
		rightLimit := big.NewInt(0)
		rightLimit.SetString("10000", 10)
		for ; num.Cmp(rightLimit) <= 0; num.Add(num, big.NewInt(1)) {

			r := SoloveyStrassen(num)
			r2 := num.ProbablyPrime(0)

			if r == r2 {
				if r {
					// fmt.Println(num, "是素数")
				}
			} else {
				fmt.Println("测试失败", r, r2, num)
				count++
				// return
			}

		}
		fmt.Println("失败次数", count)
		return
	}

	if true {
		result := Jacobi(big.NewInt(14414442441), big.NewInt(1626616611101713))
		fmt.Println("Jacobi 符号结果为:", result)
		return
	}
	//a=2 n<=3奇数成功
	if false {
		a := big.NewInt(-2)
		for n := big.NewInt(-39); n.Cmp(big.NewInt(-1)) <= 0; n.Add(n, big.NewInt(2)) {
			result := Jacobi(a, n)
			fmt.Println(a, n, "Jacobi 符号结果为:", result)
		}

		return
	}
	//a=-1成功
	if false {
		a := big.NewInt(-1)
		for n := big.NewInt(-20); n.Cmp(big.NewInt(20)) <= 0; n.Add(n, big.NewInt(1)) {
			result := Jacobi(a, n)
			fmt.Println(a, n, "Jacobi 符号结果为:", result)
		}

		return
	}
	// n=1成功
	if false {
		for a := big.NewInt(-10); a.Cmp(big.NewInt(10)) <= 0; a.Add(a, big.NewInt(1)) {

			result := Jacobi(a, big.NewInt(1))
			fmt.Println(a, big.NewInt(1), "Jacobi 符号结果为:", result)

		}
		return
	}
	// n = 0成功
	if false {
		n := big.NewInt(0)
		for a := big.NewInt(-10); a.Cmp(big.NewInt(10)) <= 0; a.Add(a, big.NewInt(1)) {

			result := Jacobi(a, n)
			fmt.Println(a, n, "Jacobi 符号结果为:", result)

		}
		return
	}
	// n为偶数
	if true {
		for a := big.NewInt(-10); a.Cmp(big.NewInt(10)) <= 0; a.Add(a, big.NewInt(1)) {
			for n := big.NewInt(-12); n.Cmp(big.NewInt(12)) <= 0; n.Add(n, big.NewInt(1)) {
				result := Jacobi(a, n)
				fmt.Println(a, n, "Jacobi 符号结果为:", result)
			}
		}
		return
	}
	//n为奇数的整数，通过
	if true {
		for a := big.NewInt(-10); a.Cmp(big.NewInt(10)) <= 0; a.Add(a, big.NewInt(1)) {
			for n := big.NewInt(-12); n.Cmp(big.NewInt(12)) <= 0; n.Add(n, big.NewInt(1)) {
				result := Jacobi(a, n)
				fmt.Println(a, n, "Jacobi 符号结果为:", result)
			}
		}
	}
}
