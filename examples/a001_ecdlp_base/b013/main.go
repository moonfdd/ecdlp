package main

import (
	"encoding/json"
	"fmt"
	"math/big"
	"os"

	"github.com/moonfdd/ecdlp"
)

func fromHex(s string) *big.Int {
	if s == "" {
		return big.NewInt(0)
	}
	r, ok := new(big.Int).SetString(s, 16)
	if !ok {
		panic("invalid hex in source file: " + s)
	}
	return r
}

// 求模立方根的个数0，1，3
func ModCbrtCount(c, p *big.Int) int {
	t := big.NewInt(0)
	t.Add(p, big.NewInt(-2))
	t.Mod(t, big.NewInt(3))
	if t.Cmp(big.NewInt(0)) == 0 {
		return 1
	}
	t = big.NewInt(0).Add(p, big.NewInt(-1))
	t.Div(t, big.NewInt(3))
	if big.NewInt(0).Exp(c, t, p).Cmp(big.NewInt(1)) == 0 {
		return 3
	} else {
		return 0
	}
}

// https://loj.ac/s/1752475
// https://eprint.iacr.org/2013/024.pdf
// Peralta Method
func ModCbrt(a, p *big.Int) (ans []*big.Int) {
	ans = make([]*big.Int, 0)
	if a.Cmp(big.NewInt(0)) == 0 {
		ans = append(ans, big.NewInt(0))
		return
	}
	count := ModCbrtCount(a, p)
	if count == 1 { //有1个解
		t := big.NewInt(0).Lsh(p, 1) //t=(2p-1)/3
		t.Mod(t, p)
		t = t.Add(t, big.NewInt(-1))
		t.Mod(t, p)
		t.Mul(t, big.NewInt(0).ModInverse(big.NewInt(3), p))
		t.Mod(t, p)
		ans = append(ans, big.NewInt(0).Exp(a, t, p))
	} else if count == 3 { //有3个解，Peralta Method算法

		//w=i^3=b^3-a
		var b *big.Int
		w := big.NewInt(0)
		for b = big.NewInt(1); b.Cmp(p) < 0; b.Add(b, big.NewInt(1)) {
			w.Exp(b, big.NewInt(3), p)
			w.Add(w, big.NewInt(0).Neg(a))
			w.Mod(w, p)
			if w.Cmp(big.NewInt(0)) != 0 && ModCbrtCount(w, p) == 0 {
				break
			}
		}

		iRoot := Ring{b, big.NewInt(-1), big.NewInt(0), w}
		pp := big.NewInt(0).Mul(p, p) // pp = (p*p+p+1)/3
		pp.Add(pp, p)
		pp.Add(pp, big.NewInt(1))
		pp.Div(pp, big.NewInt(3))
		ansr := powerModI(iRoot, pp, p)
		x := ansr.a //根x是实部

		//求周期cycle
		cycle := big.NewInt(0)
		p3 := big.NewInt(0).Add(p, big.NewInt(-1)) //(p-1)/3
		p3.Mul(p3, big.NewInt(0).ModInverse(big.NewInt(3), p))
		p3.Mod(p3, p)
		for k := big.NewInt(1); k.Cmp(p) < 0; k.Add(k, big.NewInt(1)) {
			cycle.Exp(k, p3, p)
			if cycle.Cmp(big.NewInt(1)) != 0 {
				break
			}
		}

		otherX := big.NewInt(0)
		otherX.Mul(x, cycle) //另一个根是x*cycle
		otherX.Mod(otherX, p)

		otherX2 := big.NewInt(0)
		otherX2.Mul(otherX, cycle) //另一个根是x*cycle*cycle
		otherX2.Mod(otherX2, p)

		ans = append(ans, x, otherX, otherX2)
	}
	return
}

type Ring struct {
	a *big.Int //实部
	b *big.Int //i的虚部
	c *big.Int //i^2的虚部
	w *big.Int //i^3的值
}

// 复数乘法
func mulI(x Ring, y Ring, p *big.Int) Ring {
	var res Ring
	res.a = big.NewInt(0)
	res.b = big.NewInt(0)
	res.c = big.NewInt(0)
	res.w = x.w
	w := x.w

	a1 := big.NewInt(0)
	a2 := big.NewInt(0)
	a3 := big.NewInt(0)
	a1.Mul(x.a, y.a) //x.a*y.a
	a1.Mod(a1, p)
	a2.Mul(x.b, y.c) //x.b*y.c*w
	a2.Mod(a2, p)
	a2.Mul(a2, w)
	a2.Mod(a2, p)
	a3.Mul(x.c, y.b) //x.c*y.b*w
	a3.Mod(a3, p)
	a3.Mul(a3, w)
	a3.Mod(a3, p)
	res.a.Add(a1, a2)
	res.a.Mod(res.a, p)
	res.a.Add(res.a, a3)
	res.a.Mod(res.a, p)

	b1 := big.NewInt(0)
	b2 := big.NewInt(0)
	b3 := big.NewInt(0)
	b1.Mul(x.a, y.b) //x.a*y.b
	b1.Mod(b1, p)
	b2.Mul(x.b, y.a) //x.b*y.a
	b2.Mod(b2, p)
	b3.Mul(x.c, y.c) //x.c*y.c*w
	b3.Mod(b3, p)
	b3.Mul(b3, w)
	b3.Mod(b3, p)
	res.b.Add(b1, b2)
	res.b.Mod(res.b, p)
	res.b.Add(res.b, b3)
	res.b.Mod(res.b, p)

	c1 := big.NewInt(0)
	c2 := big.NewInt(0)
	c3 := big.NewInt(0)
	c1.Mul(x.a, y.c) //x.a*y.c
	c1.Mod(c1, p)
	c2.Mul(x.b, y.b) //x.b*y.b
	c2.Mod(c2, p)
	c3.Mul(x.c, y.a) //x.c*y.a
	c3.Mod(c3, p)
	res.c.Add(c1, c2)
	res.c.Mod(res.c, p)
	res.c.Add(res.c, c3)
	res.c.Mod(res.c, p)

	return res
}

// 复数快速幂，注意b不能取模
func powerModI(a Ring, b, p *big.Int) Ring {
	res := Ring{big.NewInt(1), big.NewInt(0), big.NewInt(0), a.w}
	b = big.NewInt(0).Set(b)
	for b.Cmp(big.NewInt(0)) != 0 {
		if big.NewInt(0).Mod(b, big.NewInt(2)).Cmp(big.NewInt(1)) == 0 {
			res = mulI(res, a, p)
		}
		a = mulI(a, a, p)
		b.Rsh(b, 1)
	}
	return res
}

// https://www.secg.org/sec2-v2.pdf
func main() {
	if false {
		// pp := big.NewInt(0).Exp(big.NewInt(853), big.NewInt(2), nil)
		pp := big.NewInt(853)
		m := big.NewInt(22970879235)
		n := big.NewInt(38832563)
		n = big.NewInt(0).ModInverse(n, pp)
		m.Mul(m, n).Mod(m, pp)
		fmt.Println(m)
		return
	}

	if true {
		m := big.NewInt(0).Mul(big.NewInt(45635), big.NewInt(637))
		n := big.NewInt(0).Mul(big.NewInt(504319), big.NewInt(77))
		n = big.NewInt(0).ModInverse(n, big.NewInt(853))
		m.Mul(m, n).Mod(m, big.NewInt(853))
		fmt.Println(m)
		return
	}
	if false {
		pp := big.NewInt(0).Exp(big.NewInt(853), big.NewInt(2), nil)
		y2 := big.NewInt(0).Exp(big.NewInt(755), big.NewInt(2), nil)
		y1 := big.NewInt(0).Exp(big.NewInt(2), big.NewInt(2), nil)
		y2.Sub(y2, y1)
		x2 := big.NewInt(0).Exp(big.NewInt(563), big.NewInt(3), nil)
		x1 := big.NewInt(0).Exp(big.NewInt(0), big.NewInt(3), nil)
		x2.Sub(x2, x1)
		y2.Sub(y2, x2)
		y2.Mul(y2, big.NewInt(0).ModInverse(big.NewInt(563), pp))
		y2.Mod(y2, pp)
		fmt.Println(y2)
		return
	}
	if true {
		// // t := big.NewInt(0).ModInverse(big.NewInt(3), big.NewInt(13))
		// // t.Mul(t, big.NewInt(-1)).Mod(t, big.NewInt(13))
		// // t = t.ModSqrt(t, big.NewInt(13))
		// // fmt.Println(t)
		// // return
		// t := big.NewInt(0).ModInverse(big.NewInt(3), big.NewInt(23981))
		// t.Mul(t, big.NewInt(-17230)).Mod(t, big.NewInt(23981))
		// t = t.ModSqrt(t, big.NewInt(23981))
		// fmt.Println(t)
		// return
		// fmt.Println(big.NewInt(0).ModInverse(big.NewInt(3), big.NewInt(61)))
		// // fmt.Println(big.NewInt(0).ModInverse(big.NewInt(11), big.NewInt(61)))
		// return
	}
	if false {
		//成功
		cc := &ecdlp.CurveParams{}
		// cc.P = big.NewInt(61)
		// cc.N = big.NewInt(61)
		// cc.A = big.NewInt(0)
		// cc.B = big.NewInt(0)
		// cc.H = big.NewInt(1)
		// cc.Gx = big.NewInt(1)
		// cc.Gy = big.NewInt(1)

		//成功 4a^3+27b^2==0
		// x^3+x+3
		//对x求导，x=11，因此(-11,0)或者(2,0)
		// cc.P = big.NewInt(13)
		// cc.N = big.NewInt(7)
		// cc.A = big.NewInt(1)
		// cc.B = big.NewInt(3)
		// cc.H = big.NewInt(1)
		// cc.Gx = big.NewInt(12)
		// cc.Gy = big.NewInt(12)

		// cc.P = big.NewInt(61)
		// cc.N = big.NewInt(61)
		// cc.A = big.NewInt(0)
		// cc.B = big.NewInt(7)
		// cc.H = big.NewInt(1)
		// cc.Gx = big.NewInt(2)
		// cc.Gy = big.NewInt(25)

		// 常规
		// cc.A = big.NewInt(0)
		// cc.B = big.NewInt(7)
		// cc.P = big.NewInt(79)
		// cc.N = big.NewInt(67)
		// cc.H = big.NewInt(1)
		// cc.Gx = big.NewInt(1)
		// cc.Gy = big.NewInt(18)

		// cc.P = big.NewInt(257)
		// cc.N = big.NewInt(257)
		// cc.A = big.NewInt(223)
		// cc.B = big.NewInt(9)
		// cc.H = big.NewInt(1)
		// cc.Gx = big.NewInt(0)
		// cc.Gy = big.NewInt(3)
		// cc.Gx = big.NewInt(232)
		// cc.Gy = big.NewInt(198)

		// file:///D:/111/5/MussonTh.pdf n=p
		// cc.P = big.NewInt(727609)
		cc.P = big.NewInt(853)
		cc.N = big.NewInt(853)
		// cc.A = big.NewInt(714069)
		cc.A = big.NewInt(108)
		cc.B = big.NewInt(4)
		cc.H = big.NewInt(1)
		cc.Gx = big.NewInt(0)
		cc.Gy = big.NewInt(2)
		// cc.Gx = big.NewInt(563)
		// cc.Gy = big.NewInt(755)
		// cc.Gx = big.NewInt(232)
		// cc.Gy = big.NewInt(198)
		// https://crypto.stackexchange.com/questions/61302/how-to-solve-this-ecdlp
		// 离散对数
		// cc.P = big.NewInt(23981)
		// cc.N = big.NewInt(109)
		// cc.A = big.NewInt(17230)
		// cc.B = big.NewInt(22699)
		// cc.H = big.NewInt(1)
		// cc.Gx = big.NewInt(1451)
		// cc.Gy = big.NewInt(1362)
		for s := big.NewInt(1); s.Cmp(cc.N) < 0; s.Add(s, big.NewInt(1)) {

			qx, qy := cc.GetQ(s)
			// qyt := big.NewInt(0).ModInverse(qy, cc.P)
			// qyt.Mul(qyt, qx).Mod(qyt, cc.P) //求qx/qy

			// t2 := big.NewInt(0).Exp(qx, big.NewInt(2), cc.P)
			// t2.Mul(t2, big.NewInt(3)).Mod(t2, cc.P)
			// t2temp := big.NewInt(0).Mul(qy, big.NewInt(2))
			// t2temp = t2temp.ModInverse(t2temp, cc.P)
			// t2.Mul(t2, t2temp).Mod(t2, cc.P)

			// t3 := big.NewInt(0).ModInverse(t2, cc.P)
			// big.NewInt(0).Mul(qyt, t2)
			// t3.Mod(t3, cc.P)

			// fmt.Println(s, "-b-", qx, qy, "-a-", big.Jacobi(qx, cc.P), "-", big.Jacobi(qy, cc.P))
			// fmt.Println(s, "--", qx, qy, "-a-", qyt, "-b-", t2, "-", t3, "--")
			// t3 := big.NewInt(0).ModInverse(qy, cc.P)
			// t3.Mul(t3, qx).Mod(t3, cc.P)
			// t3.Mul(t3, big.NewInt(43)).Mod(t3, cc.P)
			t := big.NewInt(0).ModInverse(qy, cc.P)
			t.Mul(t, big.NewInt(0).Sub(qx, big.NewInt(2))).Mod(t, cc.P)
			// t.Mul(t, big.NewInt(0).ModInverse(big.NewInt(12), cc.P)).Mod(t, cc.P)
			fmt.Println(s, "-b-", qx, qy, "--", t, "--")
		}
		return
	}
	if true {
		cc := &ecdlp.CurveParams{}
		cc.P = big.NewInt(0)
		cc.N = big.NewInt(0)
		cc.A = big.NewInt(0)
		cc.B = big.NewInt(0)
		cc.H = big.NewInt(1)
		cc.Gx = big.NewInt(27)
		cc.Gy = big.NewInt(16)
		for P := big.NewInt(19); ; P.Add(P, big.NewInt(6)) {
			// 23497 1 21552 1 1945
			// 50311 1 34973 1 15338
			// 10317511 966965 1 966965 10317510

			// 23497 18050 1 18050 23496
			// 50311 50013 1 50013 50310
			// 10317511 966965 1 966965 10317510
			if !P.ProbablyPrime(0) {
				continue
			}
			cc.P.Set(P)
			cc.N.Set(P)
			isTrue := false
			for x := big.NewInt(0); x.Cmp(P) < 0; x.Add(x, big.NewInt(1)) {
				t := big.NewInt(0).Exp(x, big.NewInt(3), P)
				t.Add(t, cc.B).Mod(t, P)
				t = t.ModSqrt(t, P)
				if t != nil && t.Cmp(big.NewInt(0)) != 0 {
					cc.Gx.Set(x)
					cc.Gy.Set(t)
					isTrue = true

					break
				}
			}

			// for y := big.NewInt(1); y.Cmp(P) < 0; y.Add(y, big.NewInt(1)) {
			// 	t := big.NewInt(0).Exp(y, big.NewInt(2), P)
			// 	t.Sub(t, cc.B).Mod(t, P)
			// 	ans := ModCbrt(t, P)
			// 	if len(ans) == 3 {
			// 		cc.Gx.Set(ans[0])
			// 		cc.Gy.Set(y)
			// 		isTrue = true

			// 		break
			// 	}
			// }

			if !isTrue {
				continue
			}
			// fmt.Println(cc.P)
			// fmt.Println(cc.Gx, cc.Gy)
			// x2, y2 := cc.GetQ(big.NewInt(2))
			// x2, y2 := cc.GetQ(big.NewInt(2))

			// d, _ := json.MarshalIndent(cc, "", "  ")
			// fmt.Println(string(d))

			// fmt.Println(x2, y2)
			// if x2.Cmp(cc.Gx) == 0 {
			// 	continue
			// }
			func() {
				defer func() {
					recover()
				}()
				x2, y2 := cc.GetQ(big.NewInt(0).Sub(cc.P, big.NewInt(1)))
				if x2.Cmp(cc.Gx) == 0 && big.NewInt(0).Add(y2, cc.Gy).Cmp(cc.P) == 0 {
					x2, y2 = cc.GetQ(big.NewInt(2))
					for s := big.NewInt(3); s.Cmp(P) < 0; s.Add(s, big.NewInt(1)) {
						// x2, y2 = cc.GetQ(s) //成功

						// fmt.Println(x2, y2)
						// _ = y2
						x2, y2 = cc.PAddQ(big.NewInt(0).Set(x2), big.NewInt(0).Set(y2), big.NewInt(0).Set(cc.Gx), big.NewInt(0).Set(cc.Gy))
						if x2.Cmp(cc.Gx) == 0 {
							s.Add(s, big.NewInt(1))
							if big.NewInt(0).Sub(P, big.NewInt(1)).Cmp(s) == 0 {
								d, _ := json.MarshalIndent(cc, "", "  ")
								_ = d
								fmt.Println("//1--", P, cc.Gx, cc.Gy, x2, y2)
								os.Exit(0)
							}
							if P.Cmp(s) == 0 {
								d, _ := json.MarshalIndent(cc, "", "  ")
								_ = d
								fmt.Println("//2--", P, cc.Gx, cc.Gy, x2, y2)
								break
							}
							fmt.Println("ddd")
							os.Exit(0)
							for k := big.NewInt(1); k.Cmp(big.NewInt(6)) <= 0; k.Add(k, big.NewInt(1)) {
								fmt.Println(k)
								tt := big.NewInt(0).Exp(cc.P, k, s)
								tt.Sub(tt, big.NewInt(1)).Mod(tt, s)
								if tt.Cmp(big.NewInt(0)) != 0 {
									continue
								}
								d, _ := json.MarshalIndent(cc, "", "  ")
								_ = d
								fmt.Println("//", P, cc.Gx, cc.Gy, x2, y2)
								os.Exit(0)
								return
								break
							}
						}
					}
					// fmt.Println("找到曲线", cc.Gx, cc.Gy, x2, y2)
					// os.Exit(0)
				}
			}()
			// for s := big.NewInt(3); s.Cmp(P) < 0; s.Add(s, big.NewInt(1)) {
			// 	// x2, y2 = cc.GetQ(s) //成功

			// 	// fmt.Println(x2, y2)
			// 	// _ = y2
			// 	x2, y2 = cc.PAddQ(big.NewInt(0).Set(x2), big.NewInt(0).Set(y2), big.NewInt(0).Set(cc.Gx), big.NewInt(0).Set(cc.Gy))
			// 	if x2.Cmp(cc.Gx) == 0 {
			// 		if big.NewInt(0).Sub(cc.P, big.NewInt(1)).Cmp(s) == 0 {
			// 			fmt.Println("找到曲线", cc.Gx, cc.Gy, x2, y2)
			// 			return
			// 		} else {
			// 			break
			// 		}
			// 	}
			// }
		}
		fmt.Println("找到曲线22")
		return

		for s := big.NewInt(1); s.Cmp(big.NewInt(66)) <= 0; s.Add(s, big.NewInt(1)) {
			Rx, Ry := cc.GetQ(s)
			ss := big.NewInt(0).Mul(Ry, Ry)
			ss.Mod(ss, cc.P)
			fmt.Println(s, Rx.Text(10), Ry.Text(10), ss)
		}
		return
	}

}
