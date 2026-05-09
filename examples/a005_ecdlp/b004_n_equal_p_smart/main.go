package main

import (
	"fmt"
	"math/big"

	"github.com/moonfdd/ecdlp"
)

func main() {
	cc := &ecdlp.CurveParams{}
	cc.A = big.NewInt(108)
	cc.B = big.NewInt(4)
	cc.P = big.NewInt(853)
	cc.N = big.NewInt(853)
	cc.H = big.NewInt(1)
	cc.Gx = big.NewInt(0)
	cc.Gy = big.NewInt(2)

	// cc.A = big.NewInt(223)
	// cc.B = big.NewInt(9)
	// cc.P = big.NewInt(257)
	// cc.N = big.NewInt(257)
	// cc.H = big.NewInt(1)
	// cc.Gx = big.NewInt(0)
	// cc.Gy = big.NewInt(3)

	i := big.NewInt(0)
	for i = big.NewInt(2); i.Cmp(cc.P) < 0; i.Add(i, big.NewInt(1)) {
		qx, qy := cc.GetQ(i)
		fmt.Println("入参：", i, qx.Text(10), qy.Text(10))
		res := NEqualPSmart(cc, qx, qy)
		fmt.Println("结果：", res)
		if res.Cmp(i) != 0 {
			fmt.Println("失败")
			return
		}
		fmt.Println("--------------------")
	}

}

func NEqualPSmart(cc *ecdlp.CurveParams, Qx, Qy *big.Int) (ans *big.Int) {
	// 只适用于 anomalous curve：#E(Fp) = p
	if cc.N.Cmp(cc.P) != 0 {
		return nil
	}

	p := new(big.Int).Set(cc.P)
	pp := new(big.Int).Mul(p, p)
	p_1 := new(big.Int).Sub(p, big.NewInt(1))

	// 求A和B
	x1 := new(big.Int).Set(cc.Gx)
	y1 := new(big.Int).Set(cc.Gy)
	x2 := new(big.Int).Set(Qx)
	y2 := new(big.Int).Set(Qy)

	// 计算 y2^2 - y1^2
	y2Sq := new(big.Int).Mul(y2, y2)
	y1Sq := new(big.Int).Mul(y1, y1)
	diffYSq := new(big.Int).Sub(y2Sq, y1Sq)

	// 计算 x2^3 - x1^3
	x2Cube := new(big.Int).Mul(x2, x2)
	x2Cube.Mul(x2Cube, x2)
	x1Cube := new(big.Int).Mul(x1, x1)
	x1Cube.Mul(x1Cube, x1)
	diffXCube := new(big.Int).Sub(x2Cube, x1Cube)

	// 计算分子: y2^2 - y1^2 - (x2^3 - x1^3)
	numerator := new(big.Int).Sub(diffYSq, diffXCube)

	// 计算分母: x2 - x1
	denominator := new(big.Int).Sub(x2, x1)

	// 计算分母模pp的逆元
	denominatorInv := new(big.Int).ModInverse(denominator, pp)

	// 计算 A = numerator * denominatorInv mod pp
	A := new(big.Int).Mul(numerator, denominatorInv)
	A.Mod(A, pp)

	// 计算 B = y1^2 - x1^3 - A*x1 mod pp
	y1SqMod := new(big.Int).Mod(y1Sq, pp)
	x1CubeMod := new(big.Int).Mod(x1Cube, pp)
	Ax1 := new(big.Int).Mul(A, x1)
	Ax1.Mod(Ax1, pp)

	B := new(big.Int).Sub(y1SqMod, x1CubeMod)
	B.Sub(B, Ax1)
	B.Mod(B, pp)

	// fmt.Println("Lifted A:", A.Text(10))
	// fmt.Println("Lifted B:", B.Text(10))

	// 根据A、B、Gx、Gy，和Qx、Qy计算Gx2=Gx、Gy2，和Qx2=Qx、Qy2，其中Gy=Gy2 mod cc.P,Qy=Qy2 mod cc.P //失败，暂时不做

	// 创建升腾后的曲线（模 p²）
	liftedCurve := &ecdlp.CurveParams{
		P:  pp,
		N:  p, // 不需要
		A:  A,
		B:  B,
		Gx: new(big.Int).Set(cc.Gx),
		Gy: new(big.Int).Set(cc.Gy),
		H:  big.NewInt(1),
	}

	// 求(n-1)G和(n-1)Q ，mod pp的。
	Gx2, Gy2 := liftedCurve.GetQ(p_1)
	Qx2, Qy2 := liftedCurve.GetQBase(p_1, Qx, Qy)
	// fmt.Println("(p-1)G mod p²:", Gx2.Text(10), Gy2.Text(10))
	// fmt.Println("(p-1)Q mod p²:", Qx2.Text(10), Qy2.Text(10))

	// m/n就是私钥
	mz := big.NewInt(0).Sub(Gy2, cc.Gy)
	mz.Mod(mz, pp)
	mm := big.NewInt(0).Sub(Gx2, cc.Gx)
	mm.Mod(mm, pp)

	nz := big.NewInt(0).Sub(Qy2, Qy)
	nz.Mod(nz, pp)
	nm := big.NewInt(0).Sub(Qx2, Qx)
	nm.Mod(nm, pp)

	mz.Mul(mz, nm).Mod(mz, pp)
	nz.Mul(nz, mm).Mod(nz, pp)

	gcd := big.NewInt(0).GCD(nil, nil, mz, nz)
	mz.Div(mz, gcd)
	nz.Div(nz, gcd)

	nz.ModInverse(nz, pp)
	mz.Mul(mz, nz).Mod(mz, pp).Mod(mz, p)

	ans = mz

	return
}
