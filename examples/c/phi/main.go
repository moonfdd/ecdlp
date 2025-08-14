package main

import (
	"fmt"
	"math/big"
)

// EulerPhi 计算 n 的欧拉函数值 phi(n)
func EulerPhi(n *big.Int) *big.Int {
	S := new(big.Int).Set(n)
	R := new(big.Int).Set(n)
	m := big.NewInt(2)
	one := big.NewInt(1)
	zero := big.NewInt(0)
	tmp := new(big.Int)

	for {
		// 判断 m*m <= R
		tmp.Mul(m, m)
		if tmp.Cmp(R) > 0 {
			break
		}

		// if R % m == 0:  S = S*(m-1)/m
		tmp.Mod(R, m)
		if tmp.Cmp(zero) == 0 {
			// S = S*(m-1)/m
			S.Mul(S, new(big.Int).Sub(m, one))
			S.Div(S, m)
		}

		// while R % m == 0: R //= m
		for {
			tmp.Mod(R, m)
			if tmp.Cmp(zero) != 0 {
				break
			}
			R.Div(R, m)
		}

		// m += 1 + m%2
		tmp.Mod(m, big.NewInt(2))
		m.Add(m, new(big.Int).Add(one, tmp))
	}

	// if R > 1: S = S*(R-1)/R
	if R.Cmp(one) > 0 {
		S.Mul(S, new(big.Int).Sub(R, one))
		S.Div(S, R)
	}

	return S
}

func main() {
	for n := big.NewInt(2); n.Cmp(big.NewInt(100)) <= 0; n.Add(n, big.NewInt(1)) {
		phi := EulerPhi(n)
		fmt.Printf("Euler phi(%s) = %s\n", n.String(), phi.String())
	}
}
