package ecdlp

import (
	"math/big"
)

// 矩阵操作
type MatrixOperation struct {
}

func (that *MatrixOperation) create(m, n int) (ans [][]*big.Int) {
	ans = make([][]*big.Int, m)
	for i := 0; i < m; i++ {
		ans[i] = make([]*big.Int, n)
		for j := 0; j < n; j++ {
			ans[i][j] = big.NewInt(0)
		}
	}
	return
}

func (that *MatrixOperation) copy(a [][]*big.Int) (ans [][]*big.Int) {
	m := len(a)
	n := len(a[0])
	ans = that.create(m, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			ans[i][j].Set(a[i][j])
		}
	}

	return
}

// a*b
func (that *MatrixOperation) Mul(a, b [][]*big.Int) (ans [][]*big.Int) {
	m := len(a)
	t := len(a[0])
	n := len(b[0])
	ans = that.create(m, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < t; k++ {
				ans[i][j].Add(ans[i][j], big.NewInt(0).Mul(a[i][k], b[k][j]))
			}
		}
	}
	return
}

// a*b mod modN
func (that *MatrixOperation) MulMod(a, b [][]*big.Int, modN *big.Int) (ans [][]*big.Int) {
	m := len(a)
	t := len(a[0])
	n := len(b[0])
	ans = that.create(m, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < t; k++ {
				ans[i][j].Add(ans[i][j], big.NewInt(0).Mul(a[i][k], b[k][j]))
				ans[i][j].Mod(ans[i][j], modN)
			}
		}
	}
	return
}

// a一定要方阵
func (that *MatrixOperation) Exp(a [][]*big.Int, k *big.Int) (ans [][]*big.Int) {
	m := len(a)
	ans = that.create(m, m) // 初始化单位矩阵
	for i := 0; i < m; i++ {
		ans[i][i].SetInt64(1)
	}
	two := that.copy(a)
	for i := 0; i < k.BitLen(); i++ {
		if k.Bit(i) == 1 {
			ans = that.Mul(ans, two)
		}
		two = that.Mul(two, two)
	}
	return
}

// a一定要方阵
func (that *MatrixOperation) ExpMod(a [][]*big.Int, k *big.Int, modN *big.Int) (ans [][]*big.Int) {
	m := len(a)
	ans = that.create(m, m) // 初始化单位矩阵
	for i := 0; i < m; i++ {
		ans[i][i].SetInt64(1)
	}
	two := that.copy(a)
	for i := 0; i < k.BitLen(); i++ {
		if k.Bit(i) == 1 {
			// ans = that.Mul(ans, two)
			ans = that.MulMod(ans, two, modN)
		}
		// two = that.Mul(two, two)
		two = that.MulMod(two, two, modN)
	}
	return
}

// Q矩阵，快速幂矩阵
func (that *MatrixOperation) CreateQMatrix(a []*big.Int) (ans [][]*big.Int) {
	m := len(a)
	ans = that.create(m, m)
	ONE := big.NewInt(1)
	for i := 0; i < m; i++ {
		ans[0][i].Set(a[i])
	}
	for i := 1; i < m; i++ {
		ans[i][i-1].Set(ONE)
	}
	return
}
