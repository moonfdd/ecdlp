package ecdlp

import (
	"math/big"
)

type HigherOrderLucasSequence struct {
	// 一元高次方程的系数x^2-p[0]x+p[1]
	Coefficients []*big.Int
}

func (that *HigherOrderLucasSequence) GetUnAndVnMod(k *big.Int, N *big.Int) (ansU *big.Int, ansV *big.Int) {
	ansU = big.NewInt(0)
	ansV = big.NewInt(0)
	m := len(that.Coefficients)
	if k.Cmp(big.NewInt(0)) == 0 {
		ansV.SetInt64(int64(m))
		ansV.Mod(ansV, N)
		return
	}
	if k.Cmp(big.NewInt(1)) == 0 {
		ansU.SetInt64(1)
		ansV.Set(that.Coefficients[0])
		ansV.Mod(ansV, N)
		return
	}
	us := make([]*big.Int, m)
	vs := make([]*big.Int, m)
	for i := 0; i < m; i++ {
		us[i] = big.NewInt(0)
		vs[i] = big.NewInt(0)
	}
	// us[0] = big.NewInt(0)
	vs[0].SetInt64(int64(m)).Mod(vs[0], N)
	us[1].SetInt64(1)
	vs[1].Set(that.Coefficients[0])
	current := big.NewInt(0)
	for i := 2; i < m; i++ {
		current.SetInt64(int64(i))

		vs[0].Set(current).Mod(vs[0], N)
		sign := big.NewInt(1)

		temp := big.NewInt(0)
		for j := 0; j < i; j++ {
			temp.Mul(us[i-j-1], that.Coefficients[j])
			temp.Mul(temp, sign)
			us[i].Add(us[i], temp).Mod(us[i], N)
			sign.Neg(sign)
		}

		sign.SetInt64(1)
		for j := 0; j < i; j++ {
			temp.Mul(vs[i-j-1], that.Coefficients[j])
			temp.Mul(temp, sign)
			vs[i].Add(vs[i], temp).Mod(vs[i], N)
			sign.Neg(sign)
		}

		if current.Cmp(k) == 0 {
			ansU.Set(us[i])
			ansV.Set(vs[i])
			return
		}
	}

	vs[0].SetInt64(int64(m)).Mod(vs[0], N)
	mBig := big.NewInt(int64(m))
	kMinusM := big.NewInt(0).Sub(k, mBig)
	kMinusM.Add(kMinusM, big.NewInt(1))
	matrixOperation := MatrixOperation{}
	qMatrix := matrixOperation.CreateQMatrix(that.Coefficients)
	for i := 1; i < m; i += 2 {
		qMatrix[0][i].Neg(qMatrix[0][i])
	}
	qMatrix = matrixOperation.ExpMod(qMatrix, kMinusM, N)
	// qMatrix = matrixOperation.Exp(qMatrix, kMinusM)

	// u
	for i := 0; i < m; i++ {
		ansU.Add(ansU, big.NewInt(0).Mul(us[m-i-1], qMatrix[0][i])).Mod(ansU, N)
	}
	// v
	for i := 0; i < m; i++ {
		ansV.Add(ansV, big.NewInt(0).Mul(vs[m-i-1], qMatrix[0][i])).Mod(ansV, N)
	}

	return
}

func (that *HigherOrderLucasSequence) GetUnAndVn(k *big.Int) (ansU *big.Int, ansV *big.Int) {
	ansU = big.NewInt(0)
	ansV = big.NewInt(0)
	m := len(that.Coefficients)
	if k.Cmp(big.NewInt(0)) == 0 {
		ansV.SetInt64(int64(m))
		return
	}
	if k.Cmp(big.NewInt(1)) == 0 {
		ansU.SetInt64(1)
		ansV.Set(that.Coefficients[0])
		return
	}
	us := make([]*big.Int, m)
	vs := make([]*big.Int, m)
	for i := 0; i < m; i++ {
		us[i] = big.NewInt(0)
		vs[i] = big.NewInt(0)
	}
	// us[0] = big.NewInt(0)
	vs[0].SetInt64(int64(m))
	us[1].SetInt64(1)
	vs[1].Set(that.Coefficients[0])
	current := big.NewInt(0)
	for i := 2; i < m; i++ {
		current.SetInt64(int64(i))

		vs[0].Set(current)
		sign := big.NewInt(1)

		temp := big.NewInt(0)
		for j := 0; j < i; j++ {
			temp.Mul(us[i-j-1], that.Coefficients[j])
			temp.Mul(temp, sign)
			us[i].Add(us[i], temp)
			sign.Neg(sign)
		}

		sign.SetInt64(1)
		for j := 0; j < i; j++ {
			temp.Mul(vs[i-j-1], that.Coefficients[j])
			temp.Mul(temp, sign)
			vs[i].Add(vs[i], temp)
			sign.Neg(sign)
		}

		if current.Cmp(k) == 0 {
			ansU.Set(us[i])
			ansV.Set(vs[i])
			return
		}
	}

	vs[0].SetInt64(int64(m))
	mBig := big.NewInt(int64(m))
	kMinusM := big.NewInt(0).Sub(k, mBig)
	kMinusM.Add(kMinusM, big.NewInt(1))
	matrixOperation := MatrixOperation{}
	qMatrix := matrixOperation.CreateQMatrix(that.Coefficients)
	for i := 1; i < m; i += 2 {
		qMatrix[0][i].Neg(qMatrix[0][i])
	}
	qMatrix = matrixOperation.Exp(qMatrix, kMinusM)

	// u
	for i := 0; i < m; i++ {
		ansU.Add(ansU, big.NewInt(0).Mul(us[m-i-1], qMatrix[0][i]))
	}
	// v
	for i := 0; i < m; i++ {
		ansV.Add(ansV, big.NewInt(0).Mul(vs[m-i-1], qMatrix[0][i]))
	}

	return
}
