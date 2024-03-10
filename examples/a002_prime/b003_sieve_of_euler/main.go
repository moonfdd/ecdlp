package main

import (
	"fmt"
)

// https://www.geeksforgeeks.org/sieve-eratosthenes-0n-time-complexity/?ref=lbp
func SieveOfEuler1(N int) []bool {
	isprime := make([]bool, N+1)
	prime := []int{}        //素数表
	SPF := make([]int, N+1) //最小的素因数表

	for i := 0; i < N+1; i++ {
		isprime[i] = true
	}

	isprime[0] = false
	isprime[1] = false

	for i := 2; i < N; i++ {
		if isprime[i] {
			prime = append(prime, i)
			SPF[i] = i
		}

		for j := 0; j < len(prime) && i*prime[j] <= N && prime[j] <= SPF[i]; j++ {
			isprime[i*prime[j]] = false
			SPF[i*prime[j]] = prime[j]
		}
	}
	return isprime
}

const maxn = 100000000

var pri [maxn]int  //素数表
var now = 0        //素数表中的素数个数
var vis [maxn]bool //标记表

// 含break
func SieveOfEuler2(n int) {
	now = 0
	for i := 2; i <= n; i++ {
		if !vis[i] {
			pri[now] = i // 是质数
			now++
		}
		nDivI := n / i
		for j := 0; j < now && pri[j] <= nDivI; j++ {
			vis[pri[j]*i] = true

			if i%pri[j] == 0 {
				break // 提前break，避免重复筛
			}
		}
	}
}

// https://www.geeksforgeeks.org/sieve-eratosthenes-0n-time-complexity/?ref=lbp
func main() {
	if false {
		fmt.Println("最小素因数表")
		N := 100

		isprime := SieveOfEuler1(N)

		for i := 0; i < len(isprime); i++ {
			if isprime[i] {
				fmt.Printf("%d ", i)
			}
		}
		return
	}
	if true {
		fmt.Println("模除")
		N := 100

		SieveOfEuler2(N)

		fmt.Println(pri[0:now])
		return
	}
}
