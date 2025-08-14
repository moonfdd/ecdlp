package main

import (
	"fmt"
	"math/big"
	"math/rand"
	"time"

	"github.com/moonfdd/ecdlp"
)

var r = rand.New(rand.NewSource(time.Now().Unix()))

func main() {
	// x := new(big.Int).Rand(r, big.NewInt(100))
	// fmt.Println(x)
	// return
	cc := &ecdlp.CurveParams{}
	cc.A = big.NewInt(0)
	cc.B = big.NewInt(7)
	cc.P = big.NewInt(79)
	cc.N = big.NewInt(67)
	cc.H = big.NewInt(1)
	cc.Gx = big.NewInt(27)
	cc.Gy = big.NewInt(16)
	i := big.NewInt(4)
	for i = big.NewInt(1); i.Cmp(big.NewInt(66)) <= 0; i.Add(i, big.NewInt(1)) {
		qx, qy := cc.GetQ(i)
		fmt.Println("入参：", i, qx.Text(10), qy.Text(10))
		res := PollardRho(cc, qx, qy)
		fmt.Println("结果：", res)
		fmt.Println("--------------------")
	}

}

func PollardRho(cc *ecdlp.CurveParams, Qx, Qy *big.Int) (ans *big.Int) {
	var err interface{}
	isNeg := false
	for {
		ans, err = PollardRhoErr(cc, Qx, Qy)
		if err == nil {
			break
		}
		isNeg = true
		Qy.Neg(Qy)
	}
	if isNeg {
		Qy.Neg(Qy)
		ans.Neg(ans).Mod(ans, cc.N)

	}
	return
}
func PollardRhoErr(cc *ecdlp.CurveParams, Qx, Qy *big.Int) (ans *big.Int, err2 interface{}) {
	defer func() {
		if err := recover(); err != nil {
			err2 = err
		}
	}()
	ans = PollardRhoInternal(cc, Qx, Qy)
	return
}
func PollardRhoInternal(cc *ecdlp.CurveParams, Qx, Qy *big.Int) (ans *big.Int) {
	// 1.确定随机数alpha, beta
	N_1 := big.NewInt(0).Sub(cc.N, big.NewInt(1))
	turtleAlpha := big.NewInt(0).Rand(r, N_1)
	turtleAlpha.Add(turtleAlpha, big.NewInt(1))
	turtleBeta := big.NewInt(0).Rand(r, N_1)
	turtleBeta.Add(turtleBeta, big.NewInt(1))

	rabbitAlpha := big.NewInt(0).Set(turtleAlpha)
	rabbitBeta := big.NewInt(0).Set(turtleBeta)

	// 2.确定初始点
	alphaX, alphaY := cc.GetQ(turtleAlpha)
	betaX, betaY := cc.GetQBase(turtleBeta, Qx, Qy)
	turtleX, turtleY := cc.PAddQ(alphaX, alphaY, betaX, betaY)
	rabbitX := big.NewInt(0).Set(turtleX)
	rabbitY := big.NewInt(0).Set(turtleY)
	// 3.龟兔赛跑
	count := 0
	for {
		count++
		new_xab(cc, Qx, Qy, turtleX, turtleY, turtleAlpha, turtleBeta)
		new_xab(cc, Qx, Qy, rabbitX, rabbitY, rabbitAlpha, rabbitBeta)
		new_xab(cc, Qx, Qy, rabbitX, rabbitY, rabbitAlpha, rabbitBeta)
		if rabbitX.Cmp(turtleX) == 0 && rabbitY.Cmp(turtleY) == 0 {
			// fmt.Println(count, "龟兔相遇", rabbitX, rabbitY)
			break
		}
	}
	// turtleAlpha*P+turtleBeta*Q=rabbitAlpha*P+rabbitBeta*Q
	//(turtleAlpha-rabbitAlpha)P=(rabbitBeta-turtleBeta)Q
	//Q=(turtleAlpha-rabbitAlpha)P/(rabbitBeta-turtleBeta)

	// 4.计算结果
	// fmt.Printf("(%v-%v)/(%v-%v) mod %v\r\n", rabbitAlpha, turtleAlpha, turtleBeta, rabbitBeta, cc.N)

	alphaCha := big.NewInt(0).Sub(rabbitAlpha, turtleAlpha)
	// fmt.Println("alphaCha:", alphaCha)
	alphaCha.Mod(alphaCha, cc.N)
	betaCha := big.NewInt(0).Sub(turtleBeta, rabbitBeta)
	betaCha.Mod(betaCha, cc.N)
	// fmt.Println("betaCha:", betaCha)
	betaCha.ModInverse(betaCha, cc.N)
	ans = big.NewInt(0)
	ans.Mul(alphaCha, betaCha).Mod(ans, cc.N)
	return
}

func new_xab(cc *ecdlp.CurveParams, Qx, Qy, Rx, Ry, alpha, beta *big.Int) {
	switch fmt.Sprint(big.NewInt(0).Mod(Rx, big.NewInt(3))) {
	case "0":
		RxTemp, RyTemp := cc.PAddQ(Rx, Ry, Qx, Qy)
		Rx.Set(RxTemp)
		Ry.Set(RyTemp)
		beta.Add(beta, big.NewInt(1)).Mod(beta, cc.N)
	case "1":
		RxTemp, RyTemp := cc.TwoP(Rx, Ry)
		Rx.Set(RxTemp)
		Ry.Set(RyTemp)
		alpha.Lsh(alpha, 1).Mod(alpha, cc.N)
		beta.Lsh(beta, 1).Mod(beta, cc.N)
	case "2":
		RxTemp, RyTemp := cc.PAddQ(Rx, Ry, cc.Gx, cc.Gy)
		Rx.Set(RxTemp)
		Ry.Set(RyTemp)
		alpha.Add(alpha, big.NewInt(1)).Mod(alpha, cc.N)
	}
}
