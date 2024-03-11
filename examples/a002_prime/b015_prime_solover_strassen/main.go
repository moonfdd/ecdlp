package main

import (
	"fmt"
	"math/big"

	"github.com/moonfdd/ecdlp"
)

func SoloveyStrassen(a, num *big.Int) bool {
	if num.Cmp(big.NewInt(1)) <= 0 {
		return false
	}
	if num.Cmp(big.NewInt(2)) == 0 {
		return true
	}
	if num.Bit(0) == 0 {
		return false
	}
	j := ecdlp.Jacobi(a, num)
	if j.Cmp(big.NewInt(0)) == 0 {
		// fmt.Println("a", a, "num", num, "j", j)
		return false
	}
	j.Mod(j, num)
	t := big.NewInt(0).Add(num, big.NewInt(-1))
	t.Rsh(t, 1)
	if big.NewInt(0).Exp(a, t, num).Cmp(j) == 0 {
		return true
	} else {
		return false
	}
}

func main() {

	if true {
		fmt.Println("以2为底")
		left := big.NewInt(4)
		count := 0
		right := big.NewInt(100000)
		for num := big.NewInt(1).Set(left); num.Cmp(right) <= 0; num.Add(num, big.NewInt(1)) {
			r1 := SoloveyStrassen(big.NewInt(2), num)
			r2 := num.ProbablyPrime(0)
			if r1 == r2 {

			} else {
				count++
				if r1 {
					fmt.Print(num, " ")
				} else {
					fmt.Print(num, "是合数")
					break
				}
			}
		}
		fmt.Println()
		fmt.Println("错判次数", count)
		return
	}
	if true {
		fmt.Println("以3为底")
		left := big.NewInt(4)
		count := 0
		right := big.NewInt(100000)
		for num := big.NewInt(1).Set(left); num.Cmp(right) <= 0; num.Add(num, big.NewInt(1)) {
			r1 := SoloveyStrassen(big.NewInt(3), num)
			r2 := num.ProbablyPrime(0)
			if r1 == r2 {

			} else {
				count++
				if r1 {
					fmt.Print(num, " ")
				} else {
					fmt.Print(num, "是合数")
					break
				}
			}
		}
		fmt.Println()
		fmt.Println("错判次数", count)
		return
	}
	if true {
		fmt.Println("以2到(log(log n))^2 为底")
		left := big.NewInt(4)
		count := 0
		right := big.NewInt(100000)
		left.SetString("3317044064679887385961981", 10)
		right.SetString("3317044064679887386061981", 10)
		for num := big.NewInt(0).Set(left); num.Cmp(right) <= 0; num.Add(num, big.NewInt(1)) {
			r1 := true
			sqnum := big.NewInt(0).SetInt64(int64(num.BitLen()))
			sqnum = big.NewInt(0).SetInt64(int64(sqnum.BitLen()))
			sqnum.Mul(sqnum, sqnum)
			if sqnum.Cmp(num) > 0 {
				sqnum.Set(num)
			}
			for a := big.NewInt(2); a.Cmp(sqnum) <= 0; a.Add(a, big.NewInt(1)) {
				r1 = SoloveyStrassen(a, num)
				if !r1 {
					break
				}
			}
			r2 := num.ProbablyPrime(0)
			if r1 == r2 {

			} else {
				count++
				if r1 {
					fmt.Print(num, " ")
				} else {
					fmt.Print(num, "是合数")
					break
				}
			}
		}
		fmt.Println()
		fmt.Println("错判次数", count)
		return
	}

	fmt.Println("")
}
