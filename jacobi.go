package ecdlp

import "math/big"

func Jacobi(a *big.Int, n *big.Int) *big.Int {
	a = new(big.Int).Set(a)
	n = new(big.Int).Set(n)

	//a==1或者n==1
	if a.Cmp(big.NewInt(1)) == 0 || n.Cmp(big.NewInt(1)) == 0 {
		return big.NewInt(1)
	}

	// a=0
	if a.Cmp(big.NewInt(0)) == 0 {
		if n.Cmp(big.NewInt(-1)) == 0 {
			return big.NewInt(1)
		} else {
			return big.NewInt(0)
		}
	}

	// n=0
	if n.Cmp(big.NewInt(0)) == 0 {
		if a.Cmp(big.NewInt(-1)) == 0 {
			return big.NewInt(1)
		} else {
			return big.NewInt(0)
		}
	}

	//n=-1
	if n.Cmp(big.NewInt(-1)) == 0 {
		if a.Cmp(big.NewInt(0)) < 0 {
			return big.NewInt(-1)
		} else {
			return big.NewInt(1)
		}
	}
	//a=-1
	if a.Cmp(big.NewInt(-1)) == 0 {
		isNeg := false
		if n.Cmp(big.NewInt(0)) < 0 {
			n = big.NewInt(0).Neg(n)
			isNeg = true
		}
		n = big.NewInt(0).Set(n)
		for n.Bit(0) == 0 {
			n.Rsh(n, 1)
		}
		n_1 := big.NewInt(0).Sub(n, big.NewInt(1))
		n_1.Rsh(n_1, 1)
		ans := big.NewInt(0).Exp(big.NewInt(-1), n_1, nil)
		if isNeg {
			ans.Neg(ans)
		}
		return ans
	}

	//a和n的绝对值必都大于等于2

	//a<0
	if a.Cmp(big.NewInt(0)) < 0 {
		return big.NewInt(0).Mul(Jacobi(big.NewInt(-1), n), Jacobi(big.NewInt(0).Neg(a), n))
	}

	//n<0，直接取反
	if n.Cmp(big.NewInt(0)) < 0 {
		n.Neg(n)
	}

	//最大公约数
	if new(big.Int).GCD(nil, nil, a, n).Cmp(big.NewInt(1)) > 0 {
		return big.NewInt(0)
	}

	if n.Bit(0) == 0 { //n是偶数
		if n.Cmp(big.NewInt(2)) == 0 {
			am := big.NewInt(0).Mod(a, big.NewInt(8))
			if am.Cmp(big.NewInt(3)) == 0 || am.Cmp(big.NewInt(5)) == 0 {
				return big.NewInt(-1)
			} else {
				return big.NewInt(1)
			}
		}
		n.Rsh(n, 1)
		return big.NewInt(0).Mul(Jacobi(a, big.NewInt(2)), Jacobi(a, n))
	}

	//通例n>=3的奇数
	i := big.NewInt(0)
	s := big.NewInt(0)
	a = new(big.Int).Mod(a, n)
	temp := new(big.Int).Set(a)

	for temp.Bit(0) == 0 {
		temp.Rsh(temp, 1)
		i.Add(i, big.NewInt(1))
	}
	if i.Bit(0) == 0 {
		s.SetInt64(1)
	} else {
		if new(big.Int).Mod(n, big.NewInt(8)).Cmp(big.NewInt(1)) == 0 || new(big.Int).Mod(n, big.NewInt(8)).Cmp(big.NewInt(7)) == 0 {
			s.SetInt64(1)
		} else {
			s.SetInt64(-1)
		}
	}
	if new(big.Int).Mod(n, big.NewInt(4)).Cmp(big.NewInt(3)) == 0 && new(big.Int).Mod(temp, big.NewInt(4)).Cmp(big.NewInt(3)) == 0 {
		s.Neg(s)
	}
	if temp.Cmp(big.NewInt(1)) != 0 {
		n1 := new(big.Int).Mod(n, temp)
		temp1 := new(big.Int).Mul(s, Jacobi(n1, temp))
		return temp1
	} else {
		if i.Bit(0) == 0 {
			return big.NewInt(1)
		} else {
			if new(big.Int).Mod(n, big.NewInt(8)).Cmp(big.NewInt(1)) == 0 || new(big.Int).Mod(n, big.NewInt(8)).Cmp(big.NewInt(7)) == 0 {
				return big.NewInt(1)
			} else {
				return big.NewInt(-1)
			}
		}
	}
}
