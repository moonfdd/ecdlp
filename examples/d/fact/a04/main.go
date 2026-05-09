package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// f(x) = (x*x + c) mod N
func f(x, c, N *big.Int) *big.Int {
	x2 := new(big.Int).Mul(x, x) // x * x
	x2.Add(x2, c)                // x*x + c
	x2.Mod(x2, N)                // (x*x + c) mod N
	return x2
}

// 计算非负的 |a - b|
func absSub(a, b *big.Int) *big.Int {
	diff := new(big.Int).Sub(a, b)
	return diff.Abs(diff)
}

// 计算 gcd
func gcd(a, b *big.Int) *big.Int {
	return new(big.Int).GCD(nil, nil, a, b)
}

// 生成大于0且小于N的随机数
func randInt(n *big.Int) (*big.Int, error) {
	one := big.NewInt(1)
	max := new(big.Int).Sub(n, one) // N-1
	r, err := rand.Int(rand.Reader, max)
	if err != nil {
		return nil, err
	}
	r.Add(r, one) // 保证是 [1, N-1]
	return r, nil
}

// Pollard-Rho主函数
func PollardRho(N *big.Int) (*big.Int, error) {
	if N.Cmp(big.NewInt(1)) <= 0 {
		return nil, fmt.Errorf("N must be greater than 1")
	}

	c, err := randInt(N)
	if err != nil {
		return nil, err
	}

	// t = f(0, c, N)
	t := f(big.NewInt(0), c, N)
	// r = f(f(0, c, N), c, N)
	r := f(f(big.NewInt(0), c, N), c, N)

	one := big.NewInt(1)

	for t.Cmp(r) != 0 {
		diff := absSub(t, r)
		d := gcd(diff, N)
		if d.Cmp(one) == 1 && d.Cmp(N) == -1 {
			return d, nil
		}
		t = f(t, c, N)
		r = f(f(r, c, N), c, N)
	}
	return new(big.Int).Set(N), nil
}

// https://en.oi-wiki.org/math/pollard-rho/
// Floyd 判环
func main() {
	N := big.NewInt(18848997157)
	factor, err := PollardRho(N)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Found factor:", factor.String())
}
