package main

import (
	"fmt"
	"math/big"

	"github.com/moonfdd/ecdlp"
)

// https://www.secg.org/sec2-v2.pdf
func main() {
	//自定义椭圆曲线y^2 ≡ x^3+7 mod 79
	if true {
		cc := &ecdlp.CurveParams{}
		cc.A = big.NewInt(0)
		cc.B = big.NewInt(7)
		cc.P = big.NewInt(79)
		cc.N = big.NewInt(67)
		cc.H = big.NewInt(1)
		cc.Gx = big.NewInt(27)
		cc.Gy = big.NewInt(16)
		cc.Gx = big.NewInt(1)
		cc.Gy = big.NewInt(18)
		// mapa := make(map[string]int)
		// maps := make(map[string]int)
		// sl := make([]string, 0)
		// ssl := make([]string, 0)
		for s := big.NewInt(1); s.Cmp(big.NewInt(66)) <= 0; s.Add(s, big.NewInt(1)) {
			Rx, Ry := cc.GetQ(s)
			// ss := big.NewInt(0).Mul(Rx, big.NewInt(0).ModInverse(Ry, cc.P))
			// ss.Mod(ss, cc.P)
			// ss.Mul(ss, big.NewInt(0).ModInverse(big.NewInt(22), cc.P))
			// ss.Mod(ss, cc.P)

			sx := big.NewInt(0)
			sx.Mul(Rx, big.NewInt(0).ModInverse(cc.Gx, cc.P))
			sx.Mod(sx, cc.P)

			sy := big.NewInt(0)
			sy.Mul(Ry, big.NewInt(0).ModInverse(cc.Gy, cc.P))
			sy.Mod(sy, cc.P)

			fmt.Println(s, Rx.Text(10), Ry, "----", sx, sy)
			// if mapa[Rx.Text(10)] == 0 {
			// 	sl = append(sl, Rx.Text(10))
			// }
			// mapa[Rx.Text(10)]++

			// if maps[ss.Text(10)] == 0 {
			// 	ssl = append(ssl, ss.Text(10))
			// }
			// maps[ss.Text(10)]++
		}
		// fmt.Println(mapa)
		// fmt.Println(sl)

		// fmt.Println(maps)
		// sort.Slice(ssl, func(i, j int) bool {
		// 	a1, _ := strconv.Atoi(ssl[i])
		// 	a2, _ := strconv.Atoi(ssl[j])
		// 	return a1 <= a2
		// })
		// fmt.Println(ssl)
		// return
	}

}

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
