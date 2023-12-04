package main

import (
	"fmt"
	"math/big"
)

func main() {

	if false {

		p := big.NewInt(19)

		for i := big.NewInt(1); i.Cmp(big.NewInt(0).Add(p, big.NewInt(-1))) <= 0; i.Add(i, big.NewInt(1)) {
			fmt.Println(i, big.NewInt(1).Exp(big.NewInt(2), i, p))
		}

		return
	}

	if true {
		for p := big.NewInt(3); p.Cmp(big.NewInt(300)) <= 0; p.Add(p, big.NewInt(2)) {
			if p.ProbablyPrime(99) {
				for i := big.NewInt(1); i.Cmp(big.NewInt(0).Add(p, big.NewInt(-1))) <= 0; i.Add(i, big.NewInt(1)) {
					if big.NewInt(1).Exp(big.NewInt(2), i, p).Cmp(big.NewInt(1)) == 0 {
						fmt.Println(p, big.NewInt(0).Div(big.NewInt(0).Add(p, big.NewInt(-1)), i))
						break
					}
				}
				fmt.Println("----")
			}
		}
	}
	fmt.Println("")
}
