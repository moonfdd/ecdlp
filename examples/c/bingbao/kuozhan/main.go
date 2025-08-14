package main

import "fmt"

func main() {
	for i := 1; i < 100; i += 2 {
		fmt.Print(i, ":")
		fmt.Println(T_bx_Jia_c2(i, 5, 3))
		fmt.Println()
	}
}

//3x+1
func T_3x_Jia_1(a int) bool {
	set := make(map[int]struct{})
	set[a] = struct{}{}
	for {
		if a&1 == 1 {
			a *= 3
			a++
			a >>= 1
		} else {
			a >>= 1
		}
		if a == 1 {
			return true
		} else {
			if _, ok := set[a]; ok {
				return false
			}
			set[a] = struct{}{}
		}
	}
}

//3x-1
func T_3x_Jian_1(a int) bool {
	set := make(map[int]struct{})
	set[a] = struct{}{}
	for {
		if a&1 == 1 {
			a *= 3
			a--
		} else {
			a >>= 1
		}
		if a == 1 {
			return true
		} else {
			if _, ok := set[a]; ok {
				return false
			}
			set[a] = struct{}{}
		}
	}
}

//3x+3
func T_3x_Jia_3(a int) bool {
	set := make(map[int]struct{})
	set[a] = struct{}{}
	for {
		if a&1 == 1 {
			a *= 3
			a += 3
		} else {
			a >>= 1
		}
		if a == 1 {
			return true
		} else {
			if _, ok := set[a]; ok {
				return false
			}
			set[a] = struct{}{}
		}
	}
}

//3x+5
func T_3x_Jia_5(a int) bool {
	set := make(map[int]struct{})
	set[a] = struct{}{}
	for {
		if a&1 == 1 {
			a *= 3
			a += 5
		} else {
			a >>= 1
		}
		if a == 1 {
			return true
		} else {
			if _, ok := set[a]; ok {
				return false
			}
			set[a] = struct{}{}
		}
	}
}

//3x+b
func T_3x_Jia_b(a int, b int) bool {
	set := make(map[int]struct{})
	set[a] = struct{}{}
	for {
		if a&1 == 1 {
			a *= 3
			a += b
		} else {
			a >>= 1
		}
		if a == 1 {
			return true
		} else {
			if _, ok := set[a]; ok {
				return false
			}
			set[a] = struct{}{}
		}
	}
}

//bx+1
func T_bx_Jia_1(a int, b int) bool {
	set := make(map[int]struct{})
	set[a] = struct{}{}
	for {
		if a&1 == 1 {
			a *= b
			a += 1
		} else {
			a >>= 1
		}
		if a == 1 {
			return true
		} else {
			if _, ok := set[a]; ok {
				return false
			}
			set[a] = struct{}{}
		}
	}
}

//bx+c
func T_bx_Jia_c(a int, b int, c int) bool {
	set := make(map[int]struct{})
	set[a] = struct{}{}
	for {
		// fmt.Println(a)
		if a&1 == 1 {
			a *= b
			a += c
		} else {
			a >>= 1
		}
		if a == 1 {
			return true
		} else {
			if _, ok := set[a]; ok {
				return false
			}
			set[a] = struct{}{}
		}
	}
}

//bx+c
func T_bx_Jia_c2(a int, b int, c int) bool {
	set := make(map[int]struct{})
	set[a] = struct{}{}
	count := 0
	for {
		count++
		if a&1 == 1 {
			a *= b
			a += c
			a >>= 1
			fmt.Print(a, " ")
		} else {
			a >>= 1
			fmt.Print(a, " ")
		}
		if a <= 1 {
			fmt.Print(a, " ")
			return true
		} else {
			if _, ok := set[a]; ok {
				// fmt.Print(a, " ")
				// a *= b
				// a += c
				// a >>= 1
				// fmt.Print(a, " ")
				return false
			}
			set[a] = struct{}{}
		}
		if count == 10000 {
			return false
		}
	}
}
