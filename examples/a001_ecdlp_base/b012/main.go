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

	if true { //失败
		cc := &ecdlp.CurveParams{}
		cc.P = big.NewInt(23497)
		cc.N = big.NewInt(23497)
		cc.A = big.NewInt(0)
		cc.B = big.NewInt(7)
		cc.H = big.NewInt(1)
		cc.Gx = big.NewInt(1)
		cc.Gy = big.NewInt(1945)
		for s := big.NewInt(16325); s.Cmp(big.NewInt(16326)) < 0; s.Add(s, big.NewInt(1)) {
			qx, qy := cc.GetQ(s)
			fmt.Println(s, qx, qy)
		}
		//15061,12928
		xq := big.NewInt(15061)
		yq := big.NewInt(12928)
		xxx := big.NewInt(0).Mul(xq, cc.Gy)
		yyy := big.NewInt(0).Mul(yq, cc.Gx)
		// xxx := big.NewInt(0).Mul(yq, cc.Gx)
		// yyy := big.NewInt(0).Mul(xq, cc.Gy)
		yyy = yyy.ModInverse(yyy, cc.P)
		xxx = xxx.Mul(xxx, yyy)
		xxx.Mod(xxx, cc.P) //.Neg(xxx).Mod(xxx, cc.P)
		// xxx = xxx.ModInverse(xxx, cc.P)
		xxx.Mul(xxx, big.NewInt(3)).Mod(xxx, cc.P)
		xxx.Mul(xxx, big.NewInt(0).ModInverse(big.NewInt(2), cc.P)).Mod(xxx, cc.P)
		fmt.Println(xxx)
		return
	}
	if true {
		cc := &ecdlp.CurveParams{}
		cc.P = big.NewInt(0)
		cc.N = big.NewInt(0)
		cc.A = big.NewInt(0)
		cc.B = big.NewInt(7)
		cc.H = big.NewInt(1)
		cc.Gx = big.NewInt(27)
		cc.Gy = big.NewInt(16)
		for P := big.NewInt(23497); ; P.Add(P, big.NewInt(6)) {
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
							if big.NewInt(0).Sub(cc.P, big.NewInt(1)).Cmp(s) == 0 {
								d, _ := json.MarshalIndent(cc, "", "  ")
								// fmt.Println(string(d))
								_ = d
								// fmt.Println("找到曲线", P, cc.Gx, cc.Gy, x2, y2)
								fmt.Println("//", P, cc.Gx, cc.Gy, x2, y2)
								// os.Exit(0)
								return
								break
							} else {
								break
							}
						}
					}
					fmt.Println("找到曲线", cc.Gx, cc.Gy, x2, y2)
					os.Exit(0)
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

//61 2 36 2 25
//127 1 32 1 95
//397 3 35 3 362
//547 2 62 2 485
//4447 1 854 1 3593
//6211 3 3174 3 3037
//23497 1 21552 1 1945
//24571 3 16546 3 8025
//27361 1 10576 1 16785
//50311 1 34973 1 15338
//73477 3 72813 3 664
//89269 2 51844 2 37425
//96661 2 48175 2 48486
//170647 1 7285 1 163362
//180811 5 17178 5 163633
//200467 2 2936 2 197531
//329677 4 42975 4 286702
//347821 2 1865 2 345956
//351919 1 278668 1 73251
//436627 2 425511 2 11116
//650071 1 69015 1 581056
//742519 1 452815 1 289704
//990151 1 863141 1 127010
//997057 1 580202 1 416855
//1021417 1 726653 1 294764
//1085407 1 596836 1 488571
//1177507 2 797126 2 380381
//1328671 1 962476 1 366195
//1372957 3 956026 3 416931
//1409731 3 69283 3 1340448
//1653919 1 1228552 1 425367
//1662841 1 593282 1 1069559
//2139541 2 500775 2 1638766
//2149687 1 1010972 1 1138715
//2185387 2 903241 2 1282146
//2625481 1 8574 1 2616907
//2636719 1 1999321 1 637398
//2967091 7 2765725 7 201366
//3329587 2 2126543 2 1203044

//y
//2967091 2137860 2 2137860 2967089

//3489487 1487445 1 1487445 3489486
//3653137 3239158 3 3239158 3653134
//3834091 3245241 1 3245241 3834090
//4117237 956837 1 956837 4117236
//4194919 79890 3 79890 4194916
//4476187 2683484 5 2683484 4476182
//4850137 172194 3 172194 4850134
// 5310691 4600872 1 4600872 5310690
// 5382781 1958598 5 1958598 5382776
// 5455357 904342 1 904342 5455356
// 5471551 4724491 1 4724491 5471550
// 5733919 1620061 1 1620061 5733918
// 6002431 3325820 5 3325820 6002426
// 6019417 5746972 3 5746972 6019414
// 6061987 1069434 1 1069434 6061986
// 6156169 5614268 5 5614268 6156164
// 6799591 532402 3 532402 6799588
// 6817669 3197217 2 3197217 6817667
// 6899317 4647138 2 4647138 6899315
// 6963157 2295805 3 2295805 6963154
// 7512919 6902090 1 6902090 7512918
// 7684801 4616497 1 4616497 7684800
// 7771471 4348814 5 4348814 7771466
// 8083567 4987557 1 4987557 8083566
// 8331667 2915035 5 2915035 8331662
// 8401807 7472469 1 7472469 8401806
// 8421901 8284611 1 8284611 8421900
// 9266419 8796759 3 8796759 9266416
// 9287521 17549 2 17549 9287519
// 9382777 3490040 2 3490040 9382775
// 9404011 2881381 3 2881381 9404008
// 10118197 34900 1 34900 10118196
// 10173367 2024120 3 2024120 10173364
// 10195477 6296798 5 6296798 10195472
// 10295269 7361106 1 7361106 10295268
// 10317511 966965 1 966965 10317510
// 10756027 5053775 1 5053775 10756026
// 11331577 11236843 1 11236843 11331576
// 11354911 9544181 1 9544181 11354910
