package ecdlp

import "math/big"

// 椭圆曲线相关参数，基于y^2≡x^3+ax+b mod p
type CurveParams struct {
	P      *big.Int //坐标点(x,y)的有限域
	N      *big.Int //私钥的有限域
	A      *big.Int
	B      *big.Int
	Gx, Gy *big.Int //私钥为1的坐标点，这是规定
	H      *big.Int //余因子，这里都是为1
}

//sG
func (that *CurveParams) GetQ(s *big.Int) (Rx, Ry *big.Int) {
	Rx, Ry = that.GetQBase(s, that.Gx, that.Gy)
	return
}

//sP
func (that *CurveParams) GetQBase(s *big.Int, Px, Py *big.Int) (Rx, Ry *big.Int) {
	s = big.NewInt(0).Mod(s, that.N)
	doubleX := big.NewInt(0).Add(Px, big.NewInt(0))
	doubleY := big.NewInt(0).Add(Py, big.NewInt(0))

	for s.Cmp(big.NewInt(0)) != 0 {
		if big.NewInt(1).And(s, big.NewInt(1)).Cmp(big.NewInt(1)) == 0 {
			if Rx == nil {
				Rx = big.NewInt(0).Add(doubleX, big.NewInt(0))
				Ry = big.NewInt(0).Add(doubleY, big.NewInt(0))
			} else {
				Rx, Ry = that.PAddQ(Rx, Ry, doubleX, doubleY)
			}
		}
		doubleX, doubleY = that.TwoP(doubleX, doubleY)
		s.Rsh(s, 1)
	}
	return
}

// P+Q
func (that *CurveParams) PAddQ(Px, Py, Qx, Qy *big.Int) (Rx, Ry *big.Int) {
	// 计算斜率
	//k=(Qy-Py)/(Qx-Px)
	k := big.NewInt(0)
	k.Add(Qy, big.NewInt(0).Neg(Py))
	k.Mod(k, that.P)
	k.Mul(k, big.NewInt(0).ModInverse(big.NewInt(0).Add(Qx, big.NewInt(0).Neg(Px)), that.P))
	k.Mod(k, that.P)

	// 计算Rx
	// Rx=k^2-Px-Qx
	Rx = big.NewInt(0)
	Rx.Exp(k, big.NewInt(2), that.P)
	Rx.Add(Rx, big.NewInt(0).Neg(Px))
	Rx.Add(Rx, big.NewInt(0).Neg(Qx))
	Rx.Mod(Rx, that.P)
	// 计算Ry
	// Ry=k(Px-Rx)-Py
	Ry = big.NewInt(0)
	Ry.Add(Px, big.NewInt(0).Neg(Rx))
	Ry.Mul(Ry, k)
	Ry.Mod(Ry, that.P)
	Ry.Add(Ry, big.NewInt(0).Neg(Py))
	Ry.Mod(Ry, that.P)
	return
}

// 2P
func (that *CurveParams) TwoP(Px, Py *big.Int) (Rx, Ry *big.Int) {
	// 计算斜率
	// k=(3Px^2+a)/2Py
	k := big.NewInt(0)
	k.Mul(big.NewInt(3), Px)
	k.Mod(k, that.P)
	k.Mul(k, Px)
	k.Mod(k, that.P)
	k.Add(k, that.A)
	k.Mod(k, that.P)
	k.Mul(k, big.NewInt(0).ModInverse(big.NewInt(0).Mul(big.NewInt(2), Py), that.P))
	k.Mod(k, that.P)
	// 计算Rx
	// Rx=k^2-2Px
	Rx = big.NewInt(0)
	Rx.Exp(k, big.NewInt(2), that.P)
	Rx.Add(Rx, big.NewInt(0).Neg(Px))
	Rx.Add(Rx, big.NewInt(0).Neg(Px))
	Rx.Mod(Rx, that.P)
	// 计算Ry
	// Ry=k(Px-Rx)-Py
	Ry = big.NewInt(0)
	Ry.Add(Px, big.NewInt(0).Neg(Rx))
	Ry.Mul(Ry, k)
	Ry.Mod(Ry, that.P)
	Ry.Add(Ry, big.NewInt(0).Neg(Py))
	Ry.Mod(Ry, that.P)
	return
}
