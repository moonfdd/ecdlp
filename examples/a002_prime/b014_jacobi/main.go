package main

import (
	"fmt"
	"math/big"

	"github.com/moonfdd/ecdlp"
)

func main() {
	if true {
		for a := big.NewInt(-12); a.Cmp(big.NewInt(12)) <= 0; a.Add(a, big.NewInt(1)) {
			for p := big.NewInt(-12); p.Cmp(big.NewInt(12)) <= 0; p.Add(p, big.NewInt(1)) {
				r := ecdlp.Jacobi(a, p)
				fmt.Print(r, " ")
			}
		}
		return
	}
	fmt.Println("")
}
