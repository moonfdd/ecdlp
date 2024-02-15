package main

import (
	"fmt"
	"math"
)

func main() {
	if false {
		fmt.Println("根据素数定义")
		for n := 1; n < 100; n++ {
			isP := TrialDivision1(n)
			if isP {
				fmt.Print(n, " ")
			}
		}
		return
	}
	if false {
		fmt.Println("缩小试除范围")
		for n := 1; n < 100; n++ {
			isP := TrialDivision2(n)
			if isP {
				fmt.Print(n, " ")
			}
		}
		return
	}
	if false {
		fmt.Println("除数是奇数")
		for n := 1; n < 100; n++ {
			isP := TrialDivision3(n)
			if isP {
				fmt.Print(n, " ")
			}
		}
		return
	}
	if true {
		fmt.Println("除数是6n-1和6n+1的数")
		for n := 1; n < 100; n++ {
			isP := TrialDivision4(n)
			if isP {
				fmt.Print(n, " ")
			}
		}
		return
	}
	fmt.Println("")
}

// 根据素数定义
func TrialDivision1(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i <= n-1; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// 缩小试除范围
func TrialDivision2(n int) bool {
	if n <= 1 {
		return false
	}
	sq := int(math.Sqrt(float64(n)))
	for i := 2; i <= sq; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// 除数是奇数
func TrialDivision3(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	sq := int(math.Sqrt(float64(n)))
	for i := 3; i <= sq; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// 除数是6n-1和6n+1的数
func TrialDivision4(num int) bool {
	if num <= 1 {
		return false
	}
	if num == 2 {
		return true
	}
	if num == 3 {
		return true
	}
	if num%2 == 0 {
		return false
	}
	if num%3 == 0 {
		return false
	}

	// num一定是6n+1或者6n-1的数了，不需要再判断了
	// nMod6 := n % 6
	// if nMod6 != 1 && nMod6 != 5 {
	// 	return false
	// }

	sq := int(math.Sqrt(float64(num)))
	i := 1
	for {
		t := 6*i - 1
		if t <= sq {
			if num%t == 0 {
				return false
			}
		} else {
			break
		}

		t = 6*i + 1
		if t <= sq {
			if num%t == 0 {
				return false
			}
		} else {
			break
		}
		i++
	}
	return true
}
