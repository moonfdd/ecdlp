package ecdlp

import "math/big"

type LucasParam struct {
	P *big.Int
	Q *big.Int
}

func (that *LucasParam) GetUnAndVnMod(k *big.Int, N *big.Int) (*big.Int, *big.Int) {
	k = big.NewInt(0).Set(k)
	N = big.NewInt(0).Set(N)
	if k.Cmp(big.NewInt(0)) == 0 {
		return big.NewInt(0), big.NewInt(2)
	}
	if k.Cmp(big.NewInt(1)) == 0 {
		return big.NewInt(1), big.NewInt(0).Mod(that.P, N)
	}
	ansU := big.NewInt(0)
	ansV := big.NewInt(0)
	var tempAnsU *big.Int
	var tempAndV *big.Int
	doubleU := big.NewInt(1)
	doubleV := big.NewInt(0).Set(that.P)
	temp2U := big.NewInt(0)
	temp2V := big.NewInt(0)
	d := big.NewInt(0).Exp(that.P, big.NewInt(2), nil)
	d.Sub(d, big.NewInt(0).Lsh(that.Q, 2))
	d.Mod(d, N)
	kBitLen := k.BitLen()
	for i := 0; i < kBitLen; i++ {
		if k.Bit(i) != 0 {
			if tempAnsU == nil {
				tempAnsU = big.NewInt(0).Set(doubleU)
				tempAndV = big.NewInt(0).Set(doubleV)
			} else {
				tempAnsU.Mul(ansU, doubleV).Add(tempAnsU, big.NewInt(0).Mul(ansV, doubleU))
				tempAnsU.Mul(tempAnsU, big.NewInt(0).ModInverse(big.NewInt(2), N))

				tempAndV.Mul(d, ansU).Mul(tempAndV, doubleU).Add(tempAndV, big.NewInt(0).Mul(ansV, doubleV))
				tempAndV.Mul(tempAndV, big.NewInt(0).ModInverse(big.NewInt(2), N))

				tempAnsU.Mod(tempAnsU, N)
				tempAndV.Mod(tempAndV, N)
			}
			ansU.Set(tempAnsU)
			ansV.Set(tempAndV)
		}
		temp2U.Mul(doubleU, doubleV)
		temp2V.Mul(d, doubleU).Mul(temp2V, doubleU).Add(temp2V, big.NewInt(0).Exp(doubleV, big.NewInt(2), nil))
		temp2V.Mul(temp2V, big.NewInt(0).ModInverse(big.NewInt(2), N))

		temp2U.Mod(temp2U, N)
		temp2V.Mod(temp2V, N)

		doubleU.Set(temp2U)
		doubleV.Set(temp2V)

	}
	return ansU, ansV
}

func (that *LucasParam) TwoUAndTwoVMod(u *big.Int, v *big.Int, N *big.Int) (*big.Int, *big.Int) {
	d := big.NewInt(0).Exp(that.P, big.NewInt(2), nil)
	d.Sub(d, big.NewInt(0).Lsh(that.Q, 2))
	d.Mod(d, N)

	twoU := big.NewInt(0)
	twoV := big.NewInt(0)

	twoU.Mul(u, v)
	twoV.Mul(d, u).Mul(twoV, u).Add(twoV, big.NewInt(0).Exp(v, big.NewInt(2), nil))
	twoV.Mul(twoV, big.NewInt(0).ModInverse(big.NewInt(2), N))

	twoU.Mod(twoU, N)
	twoV.Mod(twoV, N)

	return twoU, twoV
}

func (that *LucasParam) GetUnAndVn(k *big.Int) (*big.Int, *big.Int) {
	k = big.NewInt(0).Set(k)
	if k.Cmp(big.NewInt(0)) == 0 {
		return big.NewInt(0), big.NewInt(2)
	}
	if k.Cmp(big.NewInt(1)) == 0 {
		return big.NewInt(1), big.NewInt(0).Set(that.P)
	}
	ansU := big.NewInt(0)
	ansV := big.NewInt(0)
	var tempAnsU *big.Int
	var tempAndV *big.Int
	doubleU := big.NewInt(1)
	doubleV := big.NewInt(0).Set(that.P)
	temp2U := big.NewInt(0)
	temp2V := big.NewInt(0)
	d := big.NewInt(0).Exp(that.P, big.NewInt(2), nil)
	d.Sub(d, big.NewInt(0).Lsh(that.Q, 2))
	kBitLen := k.BitLen()
	for i := 0; i < kBitLen; i++ {
		if k.Bit(i) != 0 {
			if tempAnsU == nil {
				tempAnsU = big.NewInt(0).Set(doubleU)
				tempAndV = big.NewInt(0).Set(doubleV)
			} else {
				tempAnsU.Mul(ansU, doubleV).Add(tempAnsU, big.NewInt(0).Mul(ansV, doubleU)).Rsh(tempAnsU, 1)
				tempAndV.Mul(d, ansU).Mul(tempAndV, doubleU).Add(tempAndV, big.NewInt(0).Mul(ansV, doubleV)).Rsh(tempAndV, 1)
			}
			ansU.Set(tempAnsU)
			ansV.Set(tempAndV)
		}
		temp2U.Mul(doubleU, doubleV)
		temp2V.Mul(d, doubleU).Mul(temp2V, doubleU).Add(temp2V, big.NewInt(0).Exp(doubleV, big.NewInt(2), nil)).Rsh(temp2V, 1)
		doubleU.Set(temp2U)
		doubleV.Set(temp2V)

	}
	return ansU, ansV
}
