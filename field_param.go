package ecdlp

import "math/big"

type FieldParam struct {
	M int      //如果是2，就是二次域
	D *big.Int //i^M=D
	N *big.Int // 有限域
}

// a^k
func (that *FieldParam) ExpMod(a []*big.Int, k *big.Int) (ans []*big.Int) {
	ans = that.create()
	ans[0].SetInt64(1)
	twoA := that.copyA(a)
	for i := 0; i < k.BitLen(); i++ {
		if k.Bit(i) == 1 {
			ans = that.MulMod(ans, twoA)
		}
		twoA = that.MulMod(twoA, twoA)
	}
	return
}

// a*b
func (that *FieldParam) MulMod(a, b []*big.Int) (ans []*big.Int) {
	ans = that.create()
	for i := 0; i < that.M; i++ {
		for j := 0; j < that.M; j++ {
			mod := (i + j) % that.M
			div := (i + j) / that.M
			mul := big.NewInt(0).Mul(a[i], b[j]) //a[i]*b[i]
			mul.Mod(mul, that.N)
			if div == 1 {
				mul.Mul(mul, that.D) //a[i]*b[i]*d
				mul.Mod(mul, that.N)
			}
			ans[mod].Add(ans[mod], mul)
			ans[mod].Mod(ans[mod], that.N)
		}
	}
	return
}

// a+b
func (that *FieldParam) AddMod(a, b []*big.Int) (ans []*big.Int) {
	ans = that.create()
	for i := 0; i < that.M; i++ {
		ans[i].Add(a[i], b[i])
		ans[i].Mod(ans[i], that.N)
	}
	return
}

// a-b
func (that *FieldParam) SubMod(a, b []*big.Int) (ans []*big.Int) {
	ans = that.create()
	for i := 0; i < that.M; i++ {
		ans[i].Sub(a[i], b[i])
		ans[i].Mod(ans[i], that.N)
	}
	return
}

// -a
func (that *FieldParam) NegMod(a []*big.Int) (ans []*big.Int) {
	ans = that.create()
	for i := 0; i < that.M; i++ {
		ans[i].Neg(a[i])
		ans[i].Mod(ans[i], that.N)
	}
	return
}

func (that *FieldParam) create() (ans []*big.Int) {
	ans = make([]*big.Int, that.M)
	for i := 0; i < that.M; i++ {
		ans[i] = big.NewInt(0)
	}
	return
}

func (that *FieldParam) copyA(a []*big.Int) (ans []*big.Int) {
	ans = make([]*big.Int, that.M)
	for i := 0; i < that.M; i++ {
		ans[i] = big.NewInt(0)
		ans[i].Set(a[i])
	}
	return
}
