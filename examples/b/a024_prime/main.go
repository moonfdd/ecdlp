package main

import (
	"fmt"
	"math"
	"time"
)

// https://www.luogu.com.cn/problem/P3383
func main() {
	if false {
		fmt.Println(isPrime(97))
		return
	}
	if true {
		dt := time.Now()
		r := sieve(10000 - 1)
		if r == 0 {

		}
		fmt.Println(time.Now().Sub(dt).String())
		fmt.Println(prime[0:r])
		return
	}
	if true {
		dt := time.Now()
		sieve2(100000000 - 1)
		fmt.Println(time.Now().Sub(dt).String())
		// fmt.Println(pri[0 : now+1])
	}
	fmt.Println("")
}

// https://www.bilibili.com/video/BV1LC4y117H4/?spm_id_from=333.337.search-card.all.click&vd_source=25bced4af8c6d5f851758632d0ca8444
// 01:31
func isPrime(n int) bool {
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

const maxn = 100000000

var prime [maxn]int

var is_prime [maxn + 1]bool
var p = 0

//03:12
func sieve(n int) int {
	p = 0
	for i := 0; i <= n; i++ {
		is_prime[i] = true
	}
	is_prime[0] = false
	is_prime[1] = false
	sq := int(math.Sqrt(float64(n)))
	for i := 2; i <= sq; i++ {
		if is_prime[i] {
			prime[p] = i
			p++
			for j := i * i; j <= n; j += i {
				is_prime[j] = false
			}
		}
	}
	for i := sq + 1; i <= n; i++ {
		if is_prime[i] {
			prime[p] = i
			p++
		}
	}
	return p
}
func sieve1(n int) int {
	p = 0
	for i := 0; i <= n; i++ {
		is_prime[i] = true
	}
	is_prime[0] = false
	is_prime[1] = false
	for i := 2; i <= n; i++ {
		if is_prime[i] {
			prime[p] = i
			p++
			nDivI := n / i
			for j := i; j <= nDivI; j += 1 {
				is_prime[i*j] = false
			}
		}
	}
	return p
}

//04:52
var pri [maxn]int
var now = 0
var vis [maxn]bool

func sieve2(n int) {
	now = 0
	for i := 2; i <= n; i++ {
		if !vis[i] {
			now++
			pri[now] = i // 是质数
		}
		nDivI := n / i
		for j := 1; j <= now && pri[j] <= nDivI; j++ {
			vis[pri[j]*i] = true

			if i%pri[j] == 0 {
				break // 提前break，避免重复筛
			}
		}
	}
}
