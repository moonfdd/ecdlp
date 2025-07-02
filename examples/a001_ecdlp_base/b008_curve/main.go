package main

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"sort"
	"strconv"

	"github.com/btcsuite/btcd/btcec/v2"
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
		mapa := make(map[string]int)
		maps := make(map[string]int)
		sl := make([]string, 0)
		ssl := make([]string, 0)
		for s := big.NewInt(1); s.Cmp(big.NewInt(66)) <= 0; s.Add(s, big.NewInt(1)) {
			Rx, Ry := cc.GetQ(s)
			ss := big.NewInt(0).Mul(Ry, Ry)
			ss.Mod(ss, cc.P)
			fmt.Println(s, Rx.Text(10), Ry.Text(10), ss)
			if mapa[Rx.Text(10)] == 0 {
				sl = append(sl, Rx.Text(10))
			}
			mapa[Rx.Text(10)]++

			if maps[ss.Text(10)] == 0 {
				ssl = append(ssl, ss.Text(10))
			}
			maps[ss.Text(10)]++
		}
		fmt.Println(mapa)
		fmt.Println(sl)

		fmt.Println(maps)
		sort.Slice(ssl, func(i, j int) bool {
			a1, _ := strconv.Atoi(ssl[i])
			a2, _ := strconv.Atoi(ssl[j])
			return a1 <= a2
		})
		fmt.Println(ssl)
		return
	}

	//对数器
	if false {
		cc := &ecdlp.CurveParams{}
		cc.A = big.NewInt(0)
		cc.B = big.NewInt(7)
		cc.P = fromHex("fffffffffffffffffffffffffffffffffffffffffffffffffffffffefffffc2f")
		cc.N = fromHex("fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364141")
		cc.H = big.NewInt(1)
		cc.Gx = fromHex("79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798")
		cc.Gy = fromHex("483ada7726a3c4655da4fbfc0e1108a8fd17b448a68554199c47d08ffb10d4b8")
		for s := big.NewInt(1); s.Cmp(big.NewInt(1000)) <= 0; s.Add(s, big.NewInt(1)) {
			Rx, Ry := cc.GetQ(s)
			ss := s.Text(16)
			if len(ss)%2 == 1 {
				ss = "0" + ss
			}
			privateKeyBytes, _ := hex.DecodeString(ss)
			privKey, _ := btcec.PrivKeyFromBytes(privateKeyBytes)
			publicKey := privKey.PubKey()
			if publicKey.X().Text(16) == Rx.Text(16) && publicKey.Y().Text(16) == Ry.Text(16) {
				fmt.Println(s, publicKey.X().Text(16), publicKey.Y().Text(16))
			} else {
				fmt.Println(s)
				fmt.Println(publicKey.X().Text(16))
				fmt.Println(publicKey.Y().Text(16))
				fmt.Println(Rx.Text(16))
				fmt.Println(Ry.Text(16))
				fmt.Println("验证失败")
				return
			}
		}
		fmt.Println("验证成功")
		return
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
