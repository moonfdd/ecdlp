package main

import (
	"fmt"
	"math/big"

	"github.com/moonfdd/ecdlp"
)

func main() {
	if false {
		coefficients := make([]*big.Int, 100)
		for i := 0; i < len(coefficients); i++ {
			if i%2 == 0 {
				coefficients[i] = big.NewInt(2)
			} else {
				coefficients[i] = big.NewInt(-2)
			}
		}
		for k := big.NewInt(1); k.Cmp(big.NewInt(10)) < 0; k.Add(k, big.NewInt(1)) {
			ls := ecdlp.HigherOrderLucasSequence{Coefficients: coefficients[0:10]}
			u2, v2 := ls.GetUnAndVn(k)
			fmt.Println(k, "--", u2, v2)
		}
		return
	}
	if false {
		coefficients := make([]*big.Int, 100)
		for i := 0; i < len(coefficients); i++ {
			if i%2 == 0 {
				coefficients[i] = big.NewInt(1)
			} else {
				coefficients[i] = big.NewInt(-1)
			}
		}
		const N = 4
		c := coefficients[0:2]
		fmt.Println(c, len(c))
		ls := ecdlp.HigherOrderLucasSequence{Coefficients: c}
		for k := big.NewInt(0); k.Cmp(big.NewInt(20)) < 0; k.Add(k, big.NewInt(1)) {

			u2, v2 := ls.GetUnAndVn(k)
			fmt.Println(k, "--", u2, v2)
		}
		return

	}
	if true {
		coefficients := make([]*big.Int, 100)
		for i := 0; i < len(coefficients); i++ {
			if i%2 == 0 {
				coefficients[i] = big.NewInt(1)
			} else {
				coefficients[i] = big.NewInt(-1)
			}
		}
		const N = 8
		c := coefficients[0:N]
		fmt.Println(c, len(c))
		ls := ecdlp.HigherOrderLucasSequence{Coefficients: c}
		mod := big.NewInt(10)
		count := 0
		for k := big.NewInt(0); k.Cmp(big.NewInt(1000)) < 0; k.Add(k, big.NewInt(1)) {
			u1, v1 := ls.GetUnAndVnMod(k, mod)

			u2, v2 := ls.GetUnAndVn(k)
			if big.NewInt(0).Mod(u2, mod).Cmp(u1) != 0 || big.NewInt(0).Mod(v2, mod).Cmp(v1) != 0 {
				fmt.Println("错误：", k, "--", u1, u2, "--", v1, v2)
				count++
				// return
			} else {
				// fmt.Println("正确：", k, "--", u1, u2, "--", v1, v2)
			}
		}
		fmt.Println("错误次数：", count)

		// 2
		// 0 2
		// 1 1
		// 1 3
		// 2 4
		// 3 7
		// 5 11
		// 8 18

		// 3
		// 0 3
		// 1 1
		// 1 3--
		// 2 7
		// 4 11
		// 7 21
		// 13 39

		// 4
		// 0 4
		// 1 1
		// 1 3
		// 2 7--
		// 4 15
		// 8 26
		// 15 51

		// 5
		// 0 5
		// 1 1
		// 1 3
		// 2 7
		// 4 15--
		// 8 31
		// 16 57
		// 31 113

		// 6
		// 0 6
		// 1 1
		// 1 3
		// 2 7
		// 4 15
		// 8 31--
		// 16 63
		// 32 120

		// 7
		// 0 7
		// 1 1
		// 1 3
		// 2 7
		// 4 15
		// 8 31
		// 16 63--
		// 32 127

		// 2
		// 0 2
		// 1 1
		// 1 3
		// 2 4
		// 3 7
		// 5 11
		// 8 18
		// 13 29

		// 2
		// 0--0 2
		// 1--1 2
		// 2--2 8
		// 3--6 20
		// 4--16 56
		// 5--44 152
		// 6--120 416
		// 7--328 1136
		// 8--896 3104

		// 3
		// 0--0 3
		// 1--1 2
		// 2--2 8--
		// 3--6 26
		// 4--18 72

		// 4
		// 0--0 4
		// 1--1 2
		// 2--2 8
		// 3--6 26--
		// 4--18 76

	}
}
