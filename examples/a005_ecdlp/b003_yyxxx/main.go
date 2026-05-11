package main

import (
	"fmt"
	"math/big"

	"github.com/moonfdd/ecdlp"
)

func main() {
	if true {
		cc := &ecdlp.CurveParams{}
		cc.A = big.NewInt(0)
		cc.B = big.NewInt(0)
		cc.P = big.NewInt(79)
		cc.N = big.NewInt(79)
		cc.H = big.NewInt(1)
		cc.Gx = big.NewInt(1)
		cc.Gy = big.NewInt(1)

		i := big.NewInt(0)
		for i = big.NewInt(1); i.Cmp(cc.N) < 0; i.Add(i, big.NewInt(1)) {
			qx, qy := cc.GetQ(i)
			fmt.Println("入参：", i, qx.Text(10), qy.Text(10))
			res := Yyxxx(cc, qx, qy)
			fmt.Println("结果：", res)
			if res.Cmp(i) != 0 {
				fmt.Println("失败")
				return
			}
			fmt.Println("--------------------")
		}
	}
	if false {
		cc := &ecdlp.CurveParams{}
		cc.A = big.NewInt(0)
		cc.B = big.NewInt(0)
		cc.P = fromHex("fffffffffffffffffffffffffffffffffffffffffffffffffffffffefffffc2f")
		cc.N = fromHex("fffffffffffffffffffffffffffffffffffffffffffffffffffffffefffffc2f")
		cc.H = big.NewInt(1)
		cc.Gx = big.NewInt(1)
		cc.Gy = big.NewInt(1)

		i := big.NewInt(0)
		for i = fromHex("fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364141"); i.Cmp(cc.N) < 0; i.Add(i, big.NewInt(1)) {
			qx, qy := cc.GetQ(i)
			fmt.Println("入参：", i, qx.Text(10), qy.Text(10))
			res := Yyxxx(cc, qx, qy)
			fmt.Println("结果：", res)
			if res.Cmp(i) != 0 {
				fmt.Println("失败")
				return
			}
			fmt.Println("--------------------")
		}
	}

}

func Yyxxx(cc *ecdlp.CurveParams, Qx, Qy *big.Int) (ans *big.Int) {
	// 计算 Gx/Gy 在有限域中的值
	// 即 Gx * (Gy的模逆) mod P
	GyInv := new(big.Int).ModInverse(cc.Gy, cc.P)
	GDiv := new(big.Int).Mul(cc.Gx, GyInv)
	GDiv.Mod(GDiv, cc.P)

	// 计算 Qx/Qy 在有限域中的值
	// 即 Qx * (Qy的模逆) mod P
	QyInv := new(big.Int).ModInverse(Qy, cc.P)
	QDiv := new(big.Int).Mul(Qx, QyInv)
	QDiv.Mod(QDiv, cc.P)

	// 计算 n/m = (Qx/Qy) / (Gx/Gy) = QDiv * (GDiv的模逆) mod P
	GDivInv := new(big.Int).ModInverse(GDiv, cc.P)
	ans = new(big.Int).Mul(QDiv, GDivInv)
	ans.Mod(ans, cc.P)

	return ans
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
