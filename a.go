package main

import (
	"fmt"
)

type tup struct {
	a, b, c int
}

func mul1(a, b, mod int) int {
	return int(((int64(a) * int64(b)) % int64(mod)))
}

func mul2(x, y tup, key, mod int) tup {
	var ret tup
	ret.a = mul1(x.a, y.a, mod)
	ret.a = (ret.a + mul1(mul1(x.b, y.c, mod), key, mod)) % mod
	ret.a = (ret.a + mul1(mul1(x.c, y.b, mod), key, mod)) % mod

	ret.b = mul1(x.a, y.b, mod)
	ret.b = (ret.b + mul1(x.b, y.a, mod)) % mod
	ret.b = (ret.b + mul1(mul1(x.c, y.c, mod), key, mod)) % mod

	ret.c = mul1(x.a, y.c, mod)
	ret.c = (ret.c + mul1(x.b, y.b, mod)) % mod
	ret.c = (ret.c + mul1(x.c, y.a, mod)) % mod
	return ret
}

func quick_pow1(x int, p int64, mod int) int {
	ans := 1

	for p > 0 {
		if p&1 == 1 {
			ans = mul1(ans, x, mod)
		}
		x = mul1(x, x, mod)
		p >>= 1
	}

	return ans
}

func quick_pow2(a tup, p int64, key, mod int) tup {
	ret := tup{1, 0, 0}

	for p > 0 {
		if p&1 == 1 {
			ret = mul2(ret, a, key, mod)
		}
		a = mul2(a, a, key, mod)
		p >>= 1
	}

	return ret
}

func exgcd(a, b int, x, y *int) int {
	if b == 0 {
		*x = 1
		*y = 0
		return a
	}

	q := exgcd(b, a%b, y, x)
	*y -= a / b * *x
	return q
}

func trim(x, mod int) int {
	if x >= 0 {
		return x % mod
	}
	return mod - (((-x)%mod)+mod)%mod
}

func add(a, b, mod int) int {
	if (a + b) >= mod {
		return a + b - mod
	}
	return a + b
}

func sub(a, b, mod int) int {
	if a < b {
		return a - b + mod
	}
	return a - b
}

func cbrtMod(A, P int) []int {
	var ret []int

	if A == 0 {
		ret = append(ret, 0, P)
	} else if P == 3 {
		ret = append(ret, A)
	} else if P%3 != 1 {
		x := 0
		y := 0
		exgcd(3, P-1, &x, &y)
		x = trim(x, P-1)
		ret = append(ret, quick_pow1(A, int64(x), P))
	} else if quick_pow1(A, int64((P-1)/3), P) == 1 {
		w := 0

		for i := 2; i < P; i++ {
			w = quick_pow1(i, int64((P-1)/3), P)
			if w != 1 {
				break
			}
		}

		w2 := mul1(w, w, P)
		x := 0
		key := 0

		for x = 2; x < P; x++ {
			key = sub(mul1(x, mul1(x, x, P), P), A, P)
			if quick_pow1(key, int64((P-1)/3), P) != 1 {
				break
			}
		}

		ans := quick_pow2(tup{x, P - 1, 0}, (int64(P)*int64(P)+int64(P)+1)/3, key, P).a
		ret = append(ret, ans, mul1(ans, w, P), mul1(ans, w2, P))
	}

	return ret
}

func main() {
	A := 5
	P := 13
	ans := cbrtMod(A, P)
	for i := 0; i < len(ans); i++ {
		if mul1(ans[i], mul1(ans[i], ans[i], P), P) != A {
			fmt.Println("Wrong Answer")
		}
	}
	fmt.Printf("%d %d %d", ans[0], ans[1], ans[2])
}
