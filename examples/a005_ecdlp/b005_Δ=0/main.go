package main

import (
	"fmt"
	"math/big"

	"github.com/moonfdd/ecdlp"
)

func main() {
	cc := &ecdlp.CurveParams{}
	cc.A = big.NewInt(17230)
	cc.B = big.NewInt(22699)
	cc.P = big.NewInt(23981)
	cc.N = big.NewInt(109)
	cc.H = big.NewInt(1)
	cc.Gx = big.NewInt(1451)
	cc.Gy = big.NewInt(1362)

	for i := big.NewInt(1); i.Cmp(cc.N) < 0; i.Add(i, big.NewInt(1)) {
		qx, qy := cc.GetQ(i)
		fmt.Println("入参：", i, qx.Text(10), qy.Text(10))
		res := DeltaEqualZero(cc, qx, qy)
		fmt.Println("结果：", res)
		if res.Cmp(i) != 0 {
			fmt.Println("失败")
			return
		}
		fmt.Println("--------------------")
	}

}

// https://crypto.stackexchange.com/questions/61302/how-to-solve-this-ecdlp
func DeltaEqualZero(ll *ecdlp.CurveParams, Qx, Qy *big.Int) (ans *big.Int) {
	// 3x^2+a=0，求x
	// 解方程 3x^2 + a ≡ 0 (mod p)
	// 即 x^2 ≡ -a * inv(3) (mod p)

	// 1. 计算 -a mod p
	negA := new(big.Int).Sub(ll.P, ll.A) // -a ≡ p-a (mod p)

	// 2. 计算 3 的模逆元
	three := big.NewInt(3)
	inv3 := new(big.Int).ModInverse(three, ll.P)

	// 3. 计算右边: (-a) * inv(3) mod p
	rhs := new(big.Int).Mul(negA, inv3)
	rhs.Mod(rhs, ll.P)

	// 4. 模平方根：x ≡ ±√(rhs) mod p
	// 需要判断 rhs 是否为模 p 的二次剩余
	x := new(big.Int).ModSqrt(rhs, ll.P)
	if x == nil {
		// 没有平方根，说明该曲线没有水平方向的奇异点
		// 实际上 Δ=0 时一定有解，但这里做防御性处理
		return nil
	}

	// 将x带入椭圆曲线方程中，看看y是否等于0，如果不等于0，x=p-x
	// 计算 y^2 ≡ x^3 + a*x + b (mod p)
	x3 := new(big.Int).Exp(x, three, ll.P) // x^3 mod p
	ax := new(big.Int).Mul(ll.A, x)        // a*x
	ax.Mod(ax, ll.P)
	rhsY := new(big.Int).Add(x3, ax) // x^3 + a*x
	rhsY.Add(rhsY, ll.B)             // + b
	rhsY.Mod(rhsY, ll.P)

	// fmt.Println(x, rhsY, ll.A, ll.B)
	// 检查 rhsY 是否为 0（即 y=0 是否满足方程）
	if rhsY.Cmp(big.NewInt(0)) != 0 {
		// 当前 x 不对应 y=0，换成另一个根 p - x
		x.Sub(ll.P, x)
		// 重新计算 y^2
		x3.Exp(x, three, ll.P)
		ax.Mul(ll.A, x)
		ax.Mod(ax, ll.P)
		rhsY.Add(x3, ax)
		rhsY.Add(rhsY, ll.B)
		rhsY.Mod(rhsY, ll.P)
		// fmt.Println(x, rhsY)
		// 理论上此时 rhsY 应为 0，若不是，说明出现意外
		if rhsY.Cmp(big.NewInt(0)) != 0 {
			return nil // 防御：找不到奇异点
		}
	}
	// fmt.Println(x)

	f := []*big.Int{big.NewInt(1), big.NewInt(0), ll.A, ll.B}
	g := []*big.Int{big.NewInt(0), big.NewInt(0), big.NewInt(1), x}
	m := []*big.Int{big.NewInt(1), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0)}
	po := ecdlp.PolynomialCompose(f, g, m, ll.P)
	c := po[1]
	// x 现在是 sqrt(c)，其中平移后的曲线为 y^2 = X^2 (X + c)
	t := new(big.Int).ModSqrt(c, ll.P)
	// t.Neg(t).Mod(t, ll.P)//另一个值的测试，测试通过
	// fmt.Println("c = ", c)

	// 平移基点 G -> G'
	gx := new(big.Int).Sub(ll.Gx, x)
	gx.Mod(gx, ll.P)
	gy := new(big.Int).Set(ll.Gy)

	// 平移目标点 Q -> Q'
	qx := new(big.Int).Sub(Qx, x)
	qx.Mod(qx, ll.P)
	qy := new(big.Int).Set(Qy)

	// 节点曲线映射：
	// u = (y + t*x) / (y - t*x) mod p
	mapPoint := func(px, py *big.Int) *big.Int {
		tx := new(big.Int).Mul(t, px)
		tx.Mod(tx, ll.P)

		num := new(big.Int).Add(py, tx)
		num.Mod(num, ll.P)

		den := new(big.Int).Sub(py, tx)
		den.Mod(den, ll.P)

		denInv := new(big.Int).ModInverse(den, ll.P)
		if denInv == nil {
			return nil
		}

		res := new(big.Int).Mul(num, denInv)
		res.Mod(res, ll.P)
		return res
	}

	u := mapPoint(gx, gy)
	v := mapPoint(qx, qy)
	if u == nil || v == nil {
		return nil
	}

	// 在乘法群里暴力求离散对数：v = u^x
	ans = big.NewInt(0)
	cur := big.NewInt(1)

	for x := big.NewInt(0); x.Cmp(ll.N) < 0; x.Add(x, big.NewInt(1)) {
		if cur.Cmp(v) == 0 {
			ans.Set(x)
			return
		}
		cur.Mul(cur, u)
		cur.Mod(cur, ll.P)
	}

	return nil
}
