package main

import (
	"fmt"
	"math/big"
)

func main() {

	// 测试ModSqrt
	if true {

		p := big.NewInt(0)
		p.SetString("9929", 10)
		// p.SetString("11", 10)
		// p.SetString("13", 10)
		// p.SetString("17", 10)
		// p.SetString("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEFFFFFC2F", 16)
		a := big.NewInt(13)
		for a = big.NewInt(1); a.Cmp(p) < 0; a.Add(a, big.NewInt(1)) {
			r := PocklingtonSqrt(a, p)
			fmt.Println("结果：", a, r)
			if len(r) > 0 {
				tt := big.NewInt(0)
				if tt.Exp(r[0], big.NewInt(2), p).Mod(tt, p).Cmp(a) == 0 {
				} else {
					fmt.Println("出错了")
					return
				}
			}
			fmt.Println("---------------------")
		}

	}

	fmt.Println("")
}

// 求模平方根的个数
func ModSqrtCount(a, p *big.Int) int {
	t := big.NewInt(0).Add(p, big.NewInt(-1)) //t=(p-1)/2
	t.Rsh(t, 1)
	if big.NewInt(0).Exp(a, t, p).Cmp(big.NewInt(1)) == 0 {
		return 2
	} else {
		return 0
	}
}

// https://en.wikipedia.org/wiki/Pocklington%27s_algorithm#Solution_method
// 求模平方根(lnp)^2
func PocklingtonSqrt(a, p *big.Int) (ans []*big.Int) {
	ans = make([]*big.Int, 0)
	x := big.NewInt(0)

	//a==0
	if a.Cmp(big.NewInt(0)) == 0 {
		ans = append(ans, x)
		return
	}

	//欧拉判别法
	if ModSqrtCount(a, p) == 0 {
		return
	}

	//存在模平方根，拆解成s和t
	//p-1=s*(2^t)  s是奇数
	t := 0
	s := big.NewInt(0).Add(p, big.NewInt(-1))
	for big.NewInt(0).And(s, big.NewInt(1)).Cmp(big.NewInt(0)) == 0 {
		s.Rsh(s, 1)
		t++
	}

	// 情况1
	if true {
		if t == 1 {
			///p==3 mod 4
			// a^((p+1)/4)
			s.Add(s, big.NewInt(1)).Rsh(s, 1)
			x.Exp(a, s, p)
			otherX := big.NewInt(0)
			otherX.Neg(x).Add(p, otherX)
			ans = append(ans, x, otherX)
			return
		}
	}

	if true { // 可省略
		m, r := big.NewInt(0).DivMod(p, big.NewInt(8), big.NewInt(0))
		// fmt.Println(m, r)
		// 情况2
		if r.Cmp(big.NewInt(5)) == 0 {
			m2Add1 := big.NewInt(2)
			m2Add1.Mul(m2Add1, m)
			m2Add1.Add(m2Add1, big.NewInt(1))
			aExpM2Add1 := big.NewInt(0).Exp(a, m2Add1, p)

			if aExpM2Add1.Cmp(big.NewInt(1)) == 0 {
				x.Exp(a, m, p).Mul(x, a).Mod(x, p)
				otherX := big.NewInt(0)
				otherX.Neg(x).Add(p, otherX)
				ans = append(ans, x, otherX)
				return
			}
			if aExpM2Add1.Cmp(big.NewInt(0).Add(p, big.NewInt(-1))) == 0 {
				mAdd1 := big.NewInt(0).Add(m, big.NewInt(1))
				y := big.NewInt(0).Mul(a, big.NewInt(4))
				y.Exp(y, mAdd1, p)
				if y.Bit(0) == 1 {
					y.Neg(y).Mod(y, p)
				}
				x.Div(y, big.NewInt(2))
				otherX := big.NewInt(0)
				otherX.Neg(x).Add(p, otherX)
				ans = append(ans, x, otherX)
				return
			}
			panic(fmt.Sprint("not impl ", aExpM2Add1)) // 正常来说，不应该会执行
		}
	}

	// 情况3
	var t1, u1 *big.Int
	for t1 = big.NewInt(1); t1.Cmp(p) < 0; t1.Add(t1, big.NewInt(1)) {
		isBreak := false
		for u1 = big.NewInt(1); u1.Cmp(p) < 0; u1.Add(u1, big.NewInt(1)) {
			N := big.NewInt(0)
			N.Mul(t1, t1).Mod(N, p)

			temp2 := big.NewInt(0)
			temp2.Mul(u1, u1).Mod(temp2, p).Mul(temp2, a).Mod(temp2, p)

			N.Add(N, temp2).Mod(N, p)

			if ModSqrtCount(N, p) == 0 {
				tp_1, up_1 := PowTU(t1, u1, a, p, big.NewInt(0).Sub(p, big.NewInt(1)))
				if tp_1.Cmp(big.NewInt(1)) == 0 && up_1.Cmp(big.NewInt(0)) == 0 {
					isBreak = true
					break
				}
			}
		}
		if isBreak {
			break
		}
	}

	ts, us := PowTU(t1, u1, a, p, s)
	ts2, us2 := TwoTU(ts, us, a, p)
	for {
		if ts2.Cmp(big.NewInt(0)) == 0 {
			break
		}
		ts, us = ts2, us2
		ts2, us2 = TwoTU(ts, us, a, p)
	}

	// fmt.Printf("%d/%d mod %d\r\n", ts, us, p)
	x.ModInverse(us, p)
	x.Mul(x, ts).Mod(x, p) //x=ts/us
	otherX := big.NewInt(0)
	otherX.Neg(x).Add(p, otherX)
	ans = append(ans, x, otherX)

	return
}

func AddTU(t1, u1, t2, u2, a, p *big.Int) (t3, u3 *big.Int) {
	temp := big.NewInt(0)
	temp.Mul(u1, u2).Mod(temp, p).Mul(temp, a).Neg(temp).Mod(temp, p) // -a*u1*u2

	t3 = big.NewInt(0)
	t3.Mul(t1, t2).Mod(t3, p).Add(t3, temp).Mod(t3, p) // t1*t2 -a*u1*u2

	temp = big.NewInt(0)
	temp.Mul(u1, t2).Mod(temp, p) // u1*t2
	u3 = big.NewInt(0)
	u3.Mul(u2, t1).Mod(u3, p).Add(u3, temp).Mod(u3, p) // u2*t1 + u1*t2
	return
}

func TwoTU(t1, u1, a, p *big.Int) (t2, u2 *big.Int) {
	t2, u2 = AddTU(t1, u1, t1, u1, a, p)
	return
}

func PowTU(t1, u1, a, p *big.Int, k *big.Int) (tk, uk *big.Int) {
	if k.Cmp(big.NewInt(1)) == 0 {
		tk = big.NewInt(0).Set(t1)
		uk = big.NewInt(0).Set(u1)
		return
	}
	if k.Cmp(big.NewInt(2)) == 0 {
		tk, uk = TwoTU(t1, u1, a, p)
		return
	}
	var tkTemp, ukTemp *big.Int
	t2 := big.NewInt(0).Set(t1)
	u2 := big.NewInt(0).Set(u1)
	for i := 0; i < k.BitLen(); i++ {
		if k.Bit(i) != 0 {
			tkTemp, ukTemp = tk, uk
			if tkTemp == nil {
				tk, uk = t2, u2
			} else {
				tk, uk = AddTU(tkTemp, ukTemp, t2, u2, a, p)
			}
		}
		t2, u2 = TwoTU(t2, u2, a, p)
	}
	return
}
