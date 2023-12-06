package main

import (
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/btcsuite/btcd/btcec/v2"
)

func main() {
	//自定义椭圆曲线
	if false {
		cc := &MyCurveParams{}
		cc.A = big.NewInt(0)
		cc.B = big.NewInt(7)
		cc.P = big.NewInt(79)
		cc.N = big.NewInt(67)
		cc.H = big.NewInt(1)
		cc.Gx = big.NewInt(1)
		cc.Gy = big.NewInt(18)

		for k := big.NewInt(1); k.Cmp(big.NewInt(66)) <= 0; k.Add(k, big.NewInt(1)) {
			Rx, Ry := cc.GetQ(k)
			fmt.Println(k, Rx, Ry)
		}
		return
	}
	//secp256k1
	if false {
		cc := &MyCurveParams{}
		cc.A = big.NewInt(0)
		cc.B = big.NewInt(7)
		cc.P = fromHex("fffffffffffffffffffffffffffffffffffffffffffffffffffffffefffffc2f")
		cc.N = fromHex("fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364141")
		cc.H = big.NewInt(1)
		cc.Gx = fromHex("79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798")
		cc.Gy = fromHex("483ada7726a3c4655da4fbfc0e1108a8fd17b448a68554199c47d08ffb10d4b8")

		Rx, Ry := cc.GetQ(fromHex("ff"))
		fmt.Println(Rx.Text(16), Ry.Text(16))

	}
	//btcec的secp256k1
	if false {
		hexString := "ff" // 16进制字符串
		privateKeyBytes, err := hex.DecodeString(hexString)
		if err != nil {
			fmt.Println("解码出错：", err)
			return
		}
		privKey, _ := btcec.PrivKeyFromBytes(privateKeyBytes)
		publicKey := privKey.PubKey()
		fmt.Println(publicKey.X().Text(16), publicKey.Y().Text(16))
		return
	}

	//对数器
	if true {
		cc := &MyCurveParams{}
		cc.A = big.NewInt(0)
		cc.B = big.NewInt(7)
		cc.P = fromHex("fffffffffffffffffffffffffffffffffffffffffffffffffffffffefffffc2f")
		cc.N = fromHex("fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364141")
		cc.H = big.NewInt(1)
		cc.Gx = fromHex("79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798")
		cc.Gy = fromHex("483ada7726a3c4655da4fbfc0e1108a8fd17b448a68554199c47d08ffb10d4b8")
		for k := big.NewInt(10000); k.Cmp(big.NewInt(20000)) <= 0; k.Add(k, big.NewInt(1)) {
			Rx, Ry := cc.GetQ(k)
			kk := k.Text(16)
			if len(kk)%2 == 1 {
				kk = "0" + kk
			}
			privateKeyBytes, _ := hex.DecodeString(kk)
			privKey, _ := btcec.PrivKeyFromBytes(privateKeyBytes)
			publicKey := privKey.PubKey()
			if publicKey.X().Text(16) == Rx.Text(16) && publicKey.Y().Text(16) == Ry.Text(16) {

			} else {
				fmt.Println(k)
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

type MyCurveParams struct {
	P      *big.Int
	N      *big.Int
	A      *big.Int
	B      *big.Int
	Gx, Gy *big.Int
	H      *big.Int
}

func (this *MyCurveParams) GetQ(k *big.Int) (Rx, Ry *big.Int) {
	k = big.NewInt(0).Mod(k, this.N)
	doubleX := big.NewInt(0).Add(this.Gx, big.NewInt(0))
	doubleY := big.NewInt(0).Add(this.Gy, big.NewInt(0))

	for k.Cmp(big.NewInt(0)) != 0 {
		if big.NewInt(1).And(k, big.NewInt(1)).Cmp(big.NewInt(1)) == 0 {
			if Rx == nil {
				Rx = big.NewInt(0).Add(doubleX, big.NewInt(0))
				Ry = big.NewInt(0).Add(doubleY, big.NewInt(0))
			} else {
				Rx, Ry = this.PAddQ(Rx, Ry, doubleX, doubleY)
			}
		}
		doubleX, doubleY = this.DoubleP(doubleX, doubleY)
		k.Rsh(k, 1)
	}
	return
}

func (this *MyCurveParams) PAddQ(Px, Py, Qx, Qy *big.Int) (Rx, Ry *big.Int) {
	// 计算斜率
	//k=(Qy-Py)/(Qx-Px)
	k := big.NewInt(0)
	k.Add(Qy, big.NewInt(0).Neg(Py))
	k.Mod(k, this.P)
	k.Mul(k, big.NewInt(0).ModInverse(big.NewInt(0).Add(Qx, big.NewInt(0).Neg(Px)), this.P))
	k.Mod(k, this.P)

	// 计算Rx
	// Rx=k^2-Px-Qx
	Rx = big.NewInt(0)
	Rx.Exp(k, big.NewInt(2), this.P)
	Rx.Add(Rx, big.NewInt(0).Neg(Px))
	Rx.Add(Rx, big.NewInt(0).Neg(Qx))
	Rx.Mod(Rx, this.P)
	// 计算Ry
	// Ry=k(Px-Rx)-Py
	Ry = big.NewInt(0)
	Ry.Add(Px, big.NewInt(0).Neg(Rx))
	Ry.Mul(Ry, k)
	Ry.Mod(Ry, this.P)
	Ry.Add(Ry, big.NewInt(0).Neg(Py))
	Ry.Mod(Ry, this.P)
	return
}

func (this *MyCurveParams) DoubleP(Px, Py *big.Int) (Rx, Ry *big.Int) {
	// 计算斜率
	// k=(3Px^2+a)/2Py
	k := big.NewInt(0)
	k.Mul(big.NewInt(3), Px)
	k.Mod(k, this.P)
	k.Mul(k, Px)
	k.Mod(k, this.P)
	k.Add(k, this.A)
	k.Mod(k, this.P)
	k.Mul(k, big.NewInt(0).ModInverse(big.NewInt(0).Mul(big.NewInt(2), Py), this.P))
	k.Mod(k, this.P)
	// 计算Rx
	// Rx=k^2-2Px
	Rx = big.NewInt(0)
	Rx.Exp(k, big.NewInt(2), this.P)
	Rx.Add(Rx, big.NewInt(0).Neg(Px))
	Rx.Add(Rx, big.NewInt(0).Neg(Px))
	Rx.Mod(Rx, this.P)
	// 计算Ry
	// Ry=k(Px-Rx)-Py
	Ry = big.NewInt(0)
	Ry.Add(Px, big.NewInt(0).Neg(Rx))
	Ry.Mul(Ry, k)
	Ry.Mod(Ry, this.P)
	Ry.Add(Ry, big.NewInt(0).Neg(Py))
	Ry.Mod(Ry, this.P)
	return
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
