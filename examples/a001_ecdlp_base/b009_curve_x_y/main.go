package main

import (
	"fmt"
	"math/big"

	"github.com/moonfdd/ecdlp"
)

func main() {
	cc := &ecdlp.CurveParams{}
	cc.A = big.NewInt(0)
	cc.B = big.NewInt(7)
	cc.P = fromHex("fffffffffffffffffffffffffffffffffffffffffffffffffffffffefffffc2f")
	cc.N = fromHex("fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364141")
	cc.H = big.NewInt(1)
	cc.Gx = fromHex("79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798")
	cc.Gy = fromHex("483ada7726a3c4655da4fbfc0e1108a8fd17b448a68554199c47d08ffb10d4b8")
	//求周期pcycle
	pcycle := big.NewInt(0)
	if true {
		p3 := big.NewInt(0).Add(cc.P, big.NewInt(-1)) //(p-1)/3
		p3.Mul(p3, big.NewInt(0).ModInverse(big.NewInt(3), cc.P))
		p3.Mod(p3, cc.P)
		for k := big.NewInt(1); k.Cmp(cc.P) < 0; k.Add(k, big.NewInt(1)) {
			pcycle.Exp(k, p3, cc.P)
			if pcycle.Cmp(big.NewInt(1)) != 0 {
				break
			}
		}
	}

	//求周期ncycle
	ncycle := big.NewInt(0)
	if true {
		p3 := big.NewInt(0).Add(cc.N, big.NewInt(-1)) //(p-1)/3
		p3.Mul(p3, big.NewInt(0).ModInverse(big.NewInt(3), cc.N))
		p3.Mod(p3, cc.N)
		for k := big.NewInt(1); k.Cmp(cc.N) < 0; k.Add(k, big.NewInt(1)) {
			ncycle.Exp(k, p3, cc.N)
			if ncycle.Cmp(big.NewInt(1)) != 0 {
				break
			}
		}
	}
	fmt.Println("pcycle = ", pcycle)
	fmt.Println("ncycle = ", ncycle)

	//原坐标
	Px := fromHex("9680241112d370b56da22eb535745d9e314380e568229e09f7241066003bc471")
	Py := fromHex("ddac2d377f03c201ffa0419d6596d10327d6c70313bb492ff495f946285d8f38")
	s := big.NewInt(999)
	fmt.Println("原坐标：", s, Px, Py)
	fmt.Println("-----------")
	//根据pcycle计算坐标点x
	x2 := big.NewInt(0).Mul(Px, pcycle)
	x2.Mod(x2, cc.P)
	x3 := big.NewInt(0).Mul(x2, pcycle)
	x3.Mod(x3, cc.P)
	fmt.Println("根据周期算出来的三个坐标x：")
	fmt.Println(Px)
	fmt.Println(x2)
	fmt.Println(x3)
	fmt.Println("-----------")
	//根据ncycle计算私钥
	s2 := big.NewInt(0).Mul(s, ncycle)
	s2.Mod(s2, cc.N)
	s3 := big.NewInt(0).Mul(s2, ncycle)
	s3.Mod(s3, cc.N)
	fmt.Println("根据周期算出来的三个私钥：", s, s2, s3)
	fmt.Println("-----------")
	//根据3个私钥计算坐标点
	fmt.Println("根据3个私钥求坐标点：")
	fmt.Println(cc.GetQ(s))
	fmt.Println(cc.GetQ(s2))
	fmt.Println(cc.GetQ(s3))

	fmt.Println("")
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
