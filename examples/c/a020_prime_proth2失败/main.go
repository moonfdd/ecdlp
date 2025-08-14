package main

import (
	"fmt"
	"math/big"

	"github.com/moonfdd/ecdlp"
)

// <<Fermat数>>P10
// k*2^m+1；k<2^m;k是奇数
// file:///E:/%E5%9B%BE%E7%89%87/20250621/%E7%9B%B8%E5%86%8C/IMG_20250621_120447.jpg
func Proth(k, m *big.Int) bool {

	// num := new(big.Int).Add(new(big.Int).Mul(k, new(big.Int).Exp(big.NewInt(2), m, nil)), big.NewInt(1)) //k*2^n+1
	// num_1 := big.NewInt(0).Sub(num, big.NewInt(1))                                                       //k*2^n
	// num_1_div_2 := big.NewInt(0).Div(num_1, big.NewInt(2))                                               //k*2^(n-1)
	// t := big.NewInt(0).Exp(big.NewInt(3), num_1_div_2, num)
	// t.Add(t, big.NewInt(1)).Mod(t, num)
	// if t.Cmp(big.NewInt(0)) == 0 {
	// 	return true
	// }
	// return false

	twoExpM := big.NewInt(0).Exp(big.NewInt(2), m, nil) //twoExpM=2^m
	if k.Cmp(twoExpM) >= 0 {                            //k<2^m
		panic("not k<2^m")
	}
	N := big.NewInt(0).Mul(k, twoExpM)
	N.Add(N, big.NewInt(1)) //N=k*2^m+1
	D := big.NewInt(0)
	for D = big.NewInt(3); D.Cmp(N) < 0; D.Add(D, big.NewInt(2)) {
		if ecdlp.Jacobi(D, N).Cmp(big.NewInt(-1)) == 0 {
			break
		}
	}
	// fmt.Println(D)
	// D = big.NewInt(3)
	//D^((N-1)/2)+1==0 mod N
	N_1_Div2 := big.NewInt(0)
	N_1_Div2.Sub(N, big.NewInt(1))
	N_1_Div2.Rsh(N_1_Div2, 1)
	r := big.NewInt(0)
	r.Exp(D, N_1_Div2, N).Add(r, big.NewInt(1)).Mod(r, N)
	return r.Cmp(big.NewInt(0)) == 0
}

func main() {
	if false {
		// a*a+b 失败
		for a := big.NewInt(2); a.Cmp(big.NewInt(1000)) < 0; a.Add(a, big.NewInt(1)) {
			exp := big.NewInt(0)
			exp.Exp(a, a, nil)
			addone := big.NewInt(0)
			addone.Set(exp)
			istrue := false
			b := big.NewInt(0)
			chengji := big.NewInt(0)
			// chengji.Set(a)
			chengji.Mul(a, a)
			for b = big.NewInt(1); b.Cmp(chengji) <= 0; b.Add(b, big.NewInt(1)) {
				addone.Add(addone, big.NewInt(1))
				if addone.ProbablyPrime(0) {
					istrue = true
					break
				}
				b.Add(b, big.NewInt(1))
			}
			if !istrue {
				fmt.Printf("%v^%v + %v = %v , %v\r\n", a, a, b, addone, istrue)
				return
			} else {
				fmt.Printf("%v^%v + %v = %v , %v\r\n", a, a, b, "--", istrue)
			}
		}
		fmt.Println("success")
		return
	}
	if false {
		//a**b**c失败
		// exp := big.NewInt(0)
		// exp.Exp(big.NewInt(2), big.NewInt(2), nil)
		// exp.Exp(exp, big.NewInt(5), nil)
		// fmt.Println(exp)
		// return
		for a := big.NewInt(2); a.Cmp(big.NewInt(100)) < 0; a.Add(a, big.NewInt(1)) {
			for b := big.NewInt(2); b.Cmp(a) < 0; b.Add(b, big.NewInt(1)) {
				for c := big.NewInt(2); c.Cmp(b) < 0; c.Add(c, big.NewInt(1)) {
					exp := big.NewInt(0)
					exp.Exp(a, b, nil)
					exp.Exp(exp, c, nil)
					fmt.Println(exp, a, b, c)
					addone := big.NewInt(0)
					addone.Set(exp)
					chengji := big.NewInt(0)
					chengji.Mul(a, b)
					chengji.Mul(chengji, c)
					d := big.NewInt(0)
					istrue := false
					for d = big.NewInt(1); d.Cmp(chengji) <= 0; d.Add(d, big.NewInt(1)) {
						addone.Add(addone, big.NewInt(1))
						if addone.ProbablyPrime(0) {
							istrue = true
							break
						}
						d.Add(d, big.NewInt(1))
					}
					if !istrue {
						fmt.Printf("%v^%v^%v + %v = %v , %v  %v\r\n", a, b, c, d, addone, istrue, exp)
						return
					} else {
						fmt.Printf("%v^%v^%v + %v = %v , %v  %v\r\n", a, b, c, d, addone, istrue, exp)
					}
				}
			}
		}
		fmt.Println("success")
		return
	}
	if false {
		//2^a+c成功,c<=a*a
		B := big.NewInt(2)
		for a := big.NewInt(0).Set(B); a.Cmp(big.NewInt(10000)) < 0; a.Add(a, big.NewInt(1)) {
			addone := big.NewInt(0)
			addone.Exp(B, a, nil)
			istrue := false
			c := big.NewInt(0)
			chengji := big.NewInt(0).Set(a)
			chengji.Mul(a, a)
			// chengji.Mul(chengji, big.NewInt(2))
			// chengji.Add(chengji, big.NewInt(2))
			for c = big.NewInt(1); c.Cmp(chengji) <= 0; {
				addone.Add(addone, big.NewInt(1))
				if addone.ProbablyPrime(0) {
					istrue = true
					break
				}
				c.Add(c, big.NewInt(1))
			}
			if !istrue {
				fmt.Printf("%v^%v + %v = %v , %+v\r\n", B, a, c, "--", istrue)
				return
			} else {
				fmt.Printf("%v^%v + %v = %v , %+v\r\n", B, a, c, "--", istrue)
			}
		}
		fmt.Println("通过")
		return
	}
	if true {
		//a^2+c成功,c<=a
		// c<=a*B失败
		B := big.NewInt(2)
		for a := big.NewInt(0).Set(B); a.Cmp(big.NewInt(10000)) < 0; a.Add(a, big.NewInt(1)) {
			addone := big.NewInt(0)
			addone.Exp(a, B, nil)
			istrue := false
			c := big.NewInt(0)
			chengji := big.NewInt(0).Set(a)
			chengji.Mul(a, B)
			// chengji.Mul(chengji, big.NewInt(2))
			for c = big.NewInt(1); c.Cmp(chengji) <= 0; {
				addone.Add(addone, big.NewInt(1))
				if addone.ProbablyPrime(0) {
					istrue = true
					break
				}
				c.Add(c, big.NewInt(1))
			}
			if !istrue {
				fmt.Printf("%v^%v + %v = %v , %+v\r\n", a, B, c, addone, istrue)
				return
			} else {
				fmt.Printf("%v^%v + %v = %v , %+v\r\n", a, B, c, addone, istrue)
			}
		}
		fmt.Println("通过")
		return
	}
	if false {
		//2**2**2
		TWO := big.NewInt(2)
		two := big.NewInt(0).Set(TWO)
		tt := big.NewInt(0).Set(TWO)
		for i := big.NewInt(2); i.Cmp(big.NewInt(20)) < 0; i.Add(i, big.NewInt(1)) {
			twoExp := two.Exp(two, TWO, nil)
			tt.Mul(tt, TWO)
			// fmt.Println(twoExp)

			istrue := false
			c := big.NewInt(0)
			addone := big.NewInt(0).Set(twoExp)
			ttMultt := big.NewInt(0).Set(tt)
			for c = big.NewInt(1); c.Cmp(ttMultt) <= 0; c.Add(c, big.NewInt(1)) {
				addone.Add(addone, big.NewInt(1))
				if addone.ProbablyPrime(0) {
					istrue = true
					break
				}
			}
			if !istrue {
				fmt.Printf("%v + %v = %v , %+v,%v\r\n", i, c, addone, istrue, ttMultt)
				return
			} else {
				fmt.Printf("%v + %v = %v , %+v,%v\r\n", i, c, "---", istrue, ttMultt)
			}
		}
		fmt.Println("通过")
		return
	}
	if false {
		for m := big.NewInt(2); m.Cmp(big.NewInt(15)) < 0; m.Add(m, big.NewInt(1)) {
			for k := big.NewInt(3); k.Cmp(big.NewInt(0).Exp(big.NewInt(2), m, nil)) < 0; k.Add(k, big.NewInt(2)) {
				// if big.NewInt(0).Mod(k, big.NewInt(3)).Cmp(big.NewInt(0)) == 0 {
				// 	continue
				// }
				r1 := Proth(k, m)
				aa := new(big.Int).Add(new(big.Int).Mul(k, new(big.Int).Exp(big.NewInt(2), m, nil)), big.NewInt(1))
				r2 := aa.ProbablyPrime(0)
				if r1 != r2 {
					fmt.Printf("错误%v*2^%v+1==%v，%v，%v\r\n", k, m, aa, r1, r2)
				} else {
					if r1 {
						fmt.Printf("素数：%v*2^%v+1==%v，%v\r\n", k, m, aa, r1)
					}
				}
			}
		}
	}
}
