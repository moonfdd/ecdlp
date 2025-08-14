package main

import (
	"fmt"
	"math/big"

	"github.com/moonfdd/ecdlp"
)

func main() {
	if true {
		//w^2=3 mod 101
		f := ecdlp.FieldParam{M: 2, N: big.NewInt(101), D: big.NewInt(3)}
		//(1+w)^2
		ans := f.ExpMod([]*big.Int{big.NewInt(1), big.NewInt(1)}, big.NewInt(2))
		fmt.Println(ans)
	}

}
