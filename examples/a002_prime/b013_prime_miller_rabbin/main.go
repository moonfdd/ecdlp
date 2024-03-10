package main

import (
	"fmt"
	"math/big"
)

func MillerRabbin(a, num *big.Int) bool {
	if num.Cmp(big.NewInt(1)) <= 0 {
		return false
	}
	if num.Cmp(big.NewInt(2)) == 0 {
		return true
	}
	t := big.NewInt(0)
	u := big.NewInt(0).Sub(num, big.NewInt(1))
	for big.NewInt(0).And(u, big.NewInt(1)).Cmp(big.NewInt(0)) == 0 {
		t.Add(t, big.NewInt(1))
		u.Rsh(u, 1)
	}
	x := big.NewInt(0).Exp(a, u, num)
	var xtemp *big.Int
	for i := big.NewInt(0); i.Cmp(t) < 0; i.Add(i, big.NewInt(1)) {
		xtemp = big.NewInt(0).Exp(x, big.NewInt(2), num)
		if xtemp.Cmp(big.NewInt(1)) == 0 && x.Cmp(big.NewInt(1)) != 0 && x.Cmp(big.NewInt(0).Sub(num, big.NewInt(1))) != 0 {
			return false
		}
		x = xtemp
	}
	if x.Cmp(big.NewInt(1)) != 0 {
		return false
	}
	return true
}

func main() {
	if false {
		nn := big.NewInt(0)
		nn.SetString("3317044064679887385961981", 10)
		nn.SetInt64(int64(nn.BitLen()))
		nn.SetInt64(int64(nn.BitLen()))
		fmt.Println(nn)
		return
	}

	if false {
		fmt.Println("以2为底")
		left := big.NewInt(4)
		count := 0
		right := big.NewInt(100000)
		for num := big.NewInt(1).Set(left); num.Cmp(right) <= 0; num.Add(num, big.NewInt(1)) {
			r1 := MillerRabbin(big.NewInt(2), num)
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

	if false {
		fmt.Println("以3为底")
		left := big.NewInt(4)
		count := 0
		right := big.NewInt(100000)
		for num := big.NewInt(1).Set(left); num.Cmp(right) <= 0; num.Add(num, big.NewInt(1)) {
			r1 := MillerRabbin(big.NewInt(3), num)
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
				r1 = MillerRabbin(a, num)
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
