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
	cc.P = big.NewInt(79)
	cc.N = big.NewInt(67)
	cc.H = big.NewInt(1)
	cc.Gx = big.NewInt(27)
	cc.Gy = big.NewInt(16)

	for i := big.NewInt(1); i.Cmp(big.NewInt(66)) <= 0; i.Add(i, big.NewInt(1)) {
		qx, qy := cc.GetQ(i)
		fmt.Println("入参：", i, qx.Text(10), qy.Text(10))
		res := Bsgs(cc, qx, qy)
		fmt.Println("结果：", res)
	}

}
func Bsgs(ll *ecdlp.CurveParams, Qx, Qy *big.Int) (ans *big.Int) { //(im+j
	m := big.NewInt(0)
	m.Sqrt(ll.N).Add(m, big.NewInt(1))

	mapRightJ := make(map[string]*big.Int)
	for j := big.NewInt(1); j.Cmp(m) <= 0; j.Add(j, big.NewInt(1)) { //1<=j<=m
		sx, sy := ll.GetQ(j)
		mapRightJ[sx.Text(10)+"_"+sy.Text(10)] = big.NewInt(0).Set(j)
	}

	mgx, mgy := ll.GetQ(m)
	if jj, ok := mapRightJ[Qx.Text(10)+"_"+Qy.Text(10)]; ok {
		ans = jj
		return
	}

	for i := big.NewInt(1); i.Cmp(m) <= 0; i.Add(i, big.NewInt(1)) { //1<=i<=m
		sx, sy2 := ll.GetQBase(i, mgx, mgy)
		if Qx.Cmp(sx) == 0 {
			ans = i.Mul(i, m)
			if Qy.Cmp(sy2) != 0 {
				ans.Neg(ans).Mod(ans, ll.N)
			}
			return
		}
		sy := big.NewInt(0)
		sy.Neg(sy2).Mod(sy, ll.P)
		sx, sy = ll.PAddQ(Qx, Qy, sx, sy)
		if j, ok := mapRightJ[sx.Text(10)+"_"+sy.Text(10)]; ok {
			ans = big.NewInt(0)
			ans.Mul(i, m).Add(ans, j)
			return
		}
	}
	return
}
