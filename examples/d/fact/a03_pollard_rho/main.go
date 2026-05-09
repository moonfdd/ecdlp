// 算法导论第三版 P571
// 时间复杂度：O(N^(1/4))
// 额外空间复杂度：O(1)
package main

import (
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

// https://en.oi-wiki.org/math/pollard-rho/
func PollardRho(N *big.Int) *big.Int {
	i := big.NewInt(1)
	rr := rand.New(rand.NewSource(time.Now().Unix()))
	x := new(big.Int).Rand(rr, N)
	// x = big.NewInt(2) //随机数
	y := big.NewInt(0).Set(x)
	k := big.NewInt(2)
	for {
		fmt.Println("i=", i, k)
		i.Add(i, big.NewInt(1))
		x.Exp(x, big.NewInt(2), N).Sub(x, big.NewInt(1)).Mod(x, N)
		d := big.NewInt(0).GCD(nil, nil, big.NewInt(0).Sub(y, x), N)
		if d.Cmp(big.NewInt(1)) != 0 && d.Cmp(N) != 0 {
			return d
		}
		if i.Cmp(k) == 0 {
			y.Set(x)
			k.Lsh(k, 1)
		}
	}
}

func main() {
	N := big.NewInt(18848997157)
	result := PollardRho(N)
	fmt.Println(result.String())
}
