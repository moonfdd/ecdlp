package main

import (
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/btcsuite/btcd/btcec/v2"
)

func main() {

	if false {
		// cc := &MyCurveParams{}
		// cc.A = big.NewInt(0)
		// cc.B = big.NewInt(7)
		// cc.P = big.NewInt(79)
		// cc.N = big.NewInt(67)
		// cc.H = big.NewInt(1)
		// cc.Gx = big.NewInt(1)
		// cc.Gy = big.NewInt(18)
		// Qx := big.NewInt(49)
		// Qy := big.NewInt(5)
		cc := &MyCurveParams{}
		cc.A = big.NewInt(0)
		cc.B = big.NewInt(7)
		cc.P = fromHex("fffffffffffffffffffffffffffffffffffffffffffffffffffffffefffffc2f")
		cc.N = fromHex("fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364141")
		cc.H = big.NewInt(1)
		cc.Gx = fromHex("79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798")
		cc.Gy = fromHex("483ada7726a3c4655da4fbfc0e1108a8fd17b448a68554199c47d08ffb10d4b8")
		Qx := fromHex("cb4ca5829ee4ba374980131a5afc773d7ee5088ab6875d21c649ba63937237e5")
		Qy := fromHex("46022aaf894bcbba192677772fd09e72cbb2d1541d2b1dcd2c7e4e93777aa5c5")
		i := 0
		m := make(map[string]struct{})
		for {
			if Qy.Cmp(cc.Gy) == 0 {
				fmt.Println("1破解成功", i, Qx, Qy)
				return
			}
			if _, ok := m[Qy.Text(16)]; ok {
				fmt.Println("2破解成功", i, Qx, Qy)
				return
			}
			m[Qy.Text(16)] = struct{}{}

			i++

			// if big.NewInt(0).Mod(Qy, big.NewInt(2)).Cmp(big.NewInt(1)) == 0 {
			// 	Qx, Qy = cc.PAddQ(Qx, Qy, cc.Gx, big.NewInt(0).Neg(cc.Gy))
			// 	fmt.Println("减1", i)
			// } else {
			// 	Qx, Qy = cc.GetQ2(big.NewInt(0).ModInverse(big.NewInt(2), cc.P), Qx, Qy)
			// 	fmt.Println("除以2", i)
			// }

			if big.NewInt(0).Mod(Qy, big.NewInt(2)).Cmp(big.NewInt(1)) == 0 {
				Qy = Qy.Neg(Qy)
				Qy.Mod(Qy, cc.P)
				fmt.Println("取反", i)
			} else {
				// Qx, Qy = cc.GetQ2(big.NewInt(2), Qx, Qy)
				// fmt.Println("乘以2", i)
				Qx, Qy = cc.GetQ2(big.NewInt(0).ModInverse(big.NewInt(2), cc.P), Qx, Qy)
				fmt.Println("除以2", i)
			}
			fmt.Println(Qx.Text(10), Qy.Text(10))
			fmt.Println("--------------------")

			// if big.NewInt(0).Mod(Qy, big.NewInt(2)).Cmp(big.NewInt(1)) == 0 {
			// 	Qx, Qy = cc.PAddQ(Qx, Qy, cc.Gx, big.NewInt(0).Neg(cc.Gy))
			// 	fmt.Println("加1", i)
			// } else {
			// 	Qx, Qy = cc.GetQ2(big.NewInt(2), Qx, Qy)
			// 	fmt.Println("乘以2", i)
			// }
			// fmt.Println(Qx.Text(16), Qy.Text(16))
			// fmt.Println("--------------------")

		}
		return
	}
	//自定义椭圆曲线
	if false {
		cc := &MyCurveParams{}
		cc.A = big.NewInt(0)
		cc.B = big.NewInt(7)
		cc.P = big.NewInt(79)
		cc.N = big.NewInt(67)
		cc.H = big.NewInt(1)
		cc.Gx = big.NewInt(27)
		cc.Gy = big.NewInt(16)

		// pqx, pgy := cc.PAddQ(big.NewInt(1), big.NewInt(18), big.NewInt(49), big.NewInt(5))
		// fmt.Println(pqx, pgy)
		// pqx, pgy = cc.PAddQ2(big.NewInt(1), big.NewInt(18), big.NewInt(49), big.NewInt(5))
		// fmt.Println(pqx, pgy)
		// return

		// return

		// m := make(map[string]struct{})
		// n := make(map[string]struct{})
		// k := big.NewInt(1)
		// for i := 1; i < 100; i++ {
		// 	Rx, Ry := cc.GetQ(k)
		// 	m[Rx.Text(10)] = struct{}{}
		// 	n[Ry.Text(10)] = struct{}{}
		// 	fmt.Println(i, k, Rx, Ry, big.NewInt(0).Mod(Ry, big.NewInt(2)))
		// 	k.Mul(k, big.NewInt(2))
		// 	k.Mod(k, cc.N)
		// }
		// fmt.Println("x的个数是", len(m))
		// fmt.Println("y的个数是", len(n))
		// fmt.Println(n)
		// return

		// m := make(map[string]struct{})
		// n := make(map[string]struct{})
		// for k := big.NewInt(1); k.Cmp(big.NewInt(66)) <= 0; k.Add(k, big.NewInt(1)) {
		// 	Rx, Ry := cc.GetQ(k)
		// 	m[Rx.Text(10)] = struct{}{}
		// 	n[Ry.Text(10)] = struct{}{}
		// 	fmt.Println(k, Rx, Ry, big.NewInt(0).Mod(Ry, big.NewInt(2)))
		// }
		// fmt.Println("x的个数是", len(m))
		// fmt.Println("y的个数是", len(n))
		// fmt.Println(n)
		// return
	}
	//secp256k1 x打圈
	if false {
		fmt.Println("x打圈")
		cc := &MyCurveParams{}
		cc.A = big.NewInt(0)
		cc.B = big.NewInt(7)
		cc.P = fromHex("fffffffffffffffffffffffffffffffffffffffffffffffffffffffefffffc2f")
		cc.N = fromHex("fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364141")
		cc.H = big.NewInt(1)
		cc.Gx = fromHex("79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798")
		cc.Gy = fromHex("483ada7726a3c4655da4fbfc0e1108a8fd17b448a68554199c47d08ffb10d4b8")

		k := big.NewInt(55)
		Rx, Ry := cc.GetQ(k)
		fmt.Println(k, Rx.Text(16), Ry.Text(16))
		fmt.Println("--------")
		kk := big.NewInt(0)
		kk.SetString("38597363079105398474523661669562635950945854759691634794201721047172720498112", 10) //(N-1)/3//x打圈了
		kk = big.NewInt(1).Exp(big.NewInt(2), kk, cc.N)
		fmt.Println("kk=", kk)

		for i := 0; i < 10; i++ {
			k2 := big.NewInt(0).Mul(k, kk)
			k2.Mod(k2, cc.N)
			Rx, Ry = cc.GetQ(k2)
			fmt.Println(k2)
			fmt.Println(Rx.Text(16), Ry.Text(16))
			fmt.Println("----")

			k = k2
		}
		return

	}
	//btcec的secp256k1
	if true {
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
		for k := big.NewInt(1); k.Cmp(big.NewInt(1000)) <= 0; k.Add(k, big.NewInt(1)) {
			Rx, Ry := cc.GetQ(k)
			kk := k.Text(16)
			if len(kk)%2 == 1 {
				kk = "0" + kk
			}
			privateKeyBytes, _ := hex.DecodeString(kk)
			privKey, _ := btcec.PrivKeyFromBytes(privateKeyBytes)
			publicKey := privKey.PubKey()
			if publicKey.X().Text(16) == Rx.Text(16) && publicKey.Y().Text(16) == Ry.Text(16) {
				fmt.Println(k, publicKey.X().Text(16), publicKey.Y().Text(16))
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

func (this *MyCurveParams) GetQ2(k *big.Int, Gx, Gy *big.Int) (Rx, Ry *big.Int) {
	k = big.NewInt(0).Mod(k, this.N)
	doubleX := big.NewInt(0).Add(Gx, big.NewInt(0))
	doubleY := big.NewInt(0).Add(Gy, big.NewInt(0))

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

func (this *MyCurveParams) PAddQ2(Px, Py, Qx, Qy *big.Int) (Rx, Ry *big.Int) {
	t1 := big.NewInt(2)
	t1.Mul(t1, this.B)
	t1.Mod(t1, this.P)
	t2 := big.NewInt(0).Add(Px, big.NewInt(0))
	t2.Mul(t2, Px)
	t2.Mul(t2, Qx)
	t2.Mod(t2, this.P)
	t3 := big.NewInt(0).Add(Px, big.NewInt(0))
	t3.Mul(t3, Qx)
	t3.Mul(t3, Qx)
	t3.Mod(t3, this.P)

	t1_3 := big.NewInt(0)
	t1_3.Add(t1_3, t1)
	t1_3.Add(t1_3, t2)
	t1_3.Add(t1_3, t3)
	t1_3.Mod(t1_3, this.P)

	t4 := big.NewInt(0).Add(this.B, big.NewInt(0))
	t4.Mul(t4, this.B)
	t4.Mod(t4, this.P)
	t5 := big.NewInt(0).Add(this.B, big.NewInt(0))
	t5.Mul(t5, Px)
	t5.Mul(t5, Px)
	t5.Mul(t5, Px)
	t5.Mod(t5, this.P)
	t6 := big.NewInt(0).Add(this.B, big.NewInt(0))
	t6.Mul(t6, Qx)
	t6.Mul(t6, Qx)
	t6.Mul(t6, Qx)
	t6.Mod(t6, this.P)
	t7 := big.NewInt(0).Add(Px, big.NewInt(0))
	t7.Mul(t7, Px)
	t7.Mul(t7, Px)
	t7.Mul(t7, Qx)
	t7.Mul(t7, Qx)
	t7.Mul(t7, Qx)
	t7.Mod(t7, this.P)

	//sqrt
	t47 := big.NewInt(0)
	t47.Add(t47, t4)
	t47.Add(t47, t5)
	t47.Add(t47, t6)
	t47.Add(t47, t7)
	sq1 := big.NewInt(0).ModSqrt(t47, this.P)
	sq2 := big.NewInt(0).Add(this.P, big.NewInt(0))
	sq2.Add(sq2, big.NewInt(0).Neg(sq1))

	t8 := big.NewInt(0).Add(Px, big.NewInt(0))
	t8.Mul(t8, Px)
	t8.Mod(t8, this.P)

	t9 := big.NewInt(0).Mul(big.NewInt(-2), Px)
	t9.Mul(t9, Qx)
	t9.Mod(t9, this.P)

	t10 := big.NewInt(0).Add(Qx, big.NewInt(0))
	t10.Mul(t10, Qx)
	t10.Mod(t10, this.P)

	t8_10 := big.NewInt(0)
	t8_10.Add(t8_10, t8)
	t8_10.Add(t8_10, t9)
	t8_10.Add(t8_10, t10)
	t8_10.Mod(t8_10, this.P)

	Rx = big.NewInt(0)
	tt1 := big.NewInt(0).Mul(sq1, big.NewInt(2))
	// tt1.Neg(tt1)
	Rx.Add(t1_3, tt1)
	Rx.Mul(Rx, big.NewInt(0).ModInverse(t8_10, this.P))
	Rx.Mod(Rx, this.P)

	Ry = big.NewInt(0)
	tt2 := big.NewInt(0).Mul(sq2, big.NewInt(2))
	// tt2.Neg(tt2)
	Ry.Add(t1_3, tt2)
	Ry.Mul(Ry, big.NewInt(0).ModInverse(t8_10, this.P))
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
