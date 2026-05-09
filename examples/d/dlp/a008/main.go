package main

import (
	"fmt"
	"math/big"
)

var (
	n     = big.NewInt(28)
	N     = big.NewInt(29) // prime
	alpha = big.NewInt(2)  // generator
	beta  = big.NewInt(5)  // 2^{10} = 1024 ≡ 5 mod N
)

func new_xab(x, a, b *big.Int) {
	mod3 := new(big.Int).Mod(x, big.NewInt(3)).Int64()
	switch mod3 {
	case 0:
		// x = x*x % N
		x.Mul(x, x).Mod(x, N)
		// a = a*2 % n
		a.Mul(a, big.NewInt(2)).Mod(a, n)
		// b = b*2 % n
		b.Mul(b, big.NewInt(2)).Mod(b, n)
	case 1:
		// x = x*alpha % N
		x.Mul(x, alpha).Mod(x, N)
		// a = (a+1) % n
		a.Add(a, big.NewInt(1)).Mod(a, n)
	case 2:
		// x = x*beta % N
		x.Mul(x, beta).Mod(x, N)
		// b = (b+1) % n
		b.Add(b, big.NewInt(1)).Mod(b, n)
	}
}

func main() {
	x := big.NewInt(1)
	a := big.NewInt(0)
	b := big.NewInt(0)

	X := new(big.Int).Set(x)
	A := new(big.Int).Set(a)
	B := new(big.Int).Set(b)

	// one := big.NewInt(1)
	for i := int64(1); i < 1018; i++ {
		new_xab(x, a, b)
		new_xab(X, A, B)
		new_xab(X, A, B)
		fmt.Printf("%3d  %4s %3s %3s  %4s %3s %3s\n",
			i, x.String(), a.String(), b.String(), X.String(), A.String(), B.String())
		if x.Cmp(X) == 0 {
			fmt.Println("结果：", x)
			break
		}
	}
}
