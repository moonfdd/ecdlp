package main

import (
	"fmt"
	"math/big"

	"github.com/moonfdd/ecdlp"
)

func main() {
	if false {
		fmt.Println("判断num是否是a的b次方")
		fmt.Println(IsPower(big.NewInt(0).Exp(big.NewInt(7), big.NewInt(7), nil)))
		return
	}
	if false {
		fmt.Println(GetSmallestR(big.NewInt(31)))
		return
	}
	if true {
		num := big.NewInt(0)
		num.SetString("2", 10)
		// num.SetString("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEBAAEDCE6AF48A03BBFD25E8CD0364141", 16)
		// right := big.NewInt(200000)
		// right.SetString("1000000", 10)

		for ; ; /*num.Cmp(right) <= 0*/ num.Add(num, big.NewInt(1)) {
			r := Aks(num)
			r2 := num.ProbablyPrime(0)

			if r == r2 {
				if r {
					fmt.Println(num, "是素数")
				}
			} else {
				fmt.Println("测试失败", r, r2, num)
				return
			}

		}
		fmt.Println("测试成功")
		return
	}
}

func Aks(num *big.Int) bool {
	if num.Cmp(big.NewInt(1)) <= 0 {
		return false
	}

	//step1
	if IsPower(num) {
		return false
	}
	//step2
	r := GetSmallestR(num)

	//step3,存在a ≤ r，使得1 < gcd(a,n) < n 成立。返回合数,
	if Ar(num, r) {
		return false
	}

	//step4
	if num.Cmp(r) <= 0 {
		return true
	}

	//step5
	if PolynomialComputation(r, num) {
		return false
	}

	//step6
	return true
}

// step1：
// 判断num是否是a的b次方
func IsPower(num *big.Int) (ans bool) {
	if num.Cmp(big.NewInt(3)) <= 0 {
		return
	}
	bRight, r1 := LogRangeByBase(num, big.NewInt(2))
	if bRight.Cmp(r1) == 0 {
		ans = true
		return
	}
	for b := big.NewInt(2); b.Cmp(bRight) <= 0; b.Add(b, big.NewInt(1)) {
		n0, n1 := NthRootRange(num, b, bRight)
		if n0.Cmp(n1) == 0 {
			ans = true
			return
		} else if n0.Cmp(big.NewInt(1)) == 0 {
			return
		}
	}

	return
}

// 求num的对数，以base为底。num>0，base>=2
func LogRangeByBase(num *big.Int, base *big.Int) (ans1, ans2 *big.Int) {
	ans1 = big.NewInt(0)
	ans2 = big.NewInt(0)
	num = big.NewInt(0).Add(num, big.NewInt(0))
	isMod := true
	for num.Cmp(big.NewInt(0)) > 0 {
		if num.Cmp(big.NewInt(1)) == 0 {
			if !isMod {
				ans2.Add(ans2, big.NewInt(1))
			}
			break
		} else if num.Cmp(base) < 0 {
			ans2.Add(ans2, big.NewInt(1))
			break
		} else {
			_, mod := num.DivMod(num, base, big.NewInt(0))
			if mod.Cmp(big.NewInt(0)) != 0 {
				isMod = false
			}
			ans1.Add(ans1, big.NewInt(1))
			ans2.Add(ans2, big.NewInt(1))
		}
	}
	return
}

// n次方根，rightLimit是二分法的右边界。num>=0，n>=2，rightLimit>0
func NthRootRange(num *big.Int, n *big.Int, rightLimit *big.Int) (ans1, ans2 *big.Int) {
	ans1 = big.NewInt(0)
	ans2 = big.NewInt(0)
	if num.Cmp(big.NewInt(0)) == 0 {
		return
	} else if num.Cmp(big.NewInt(1)) == 0 {
		ans1 = big.NewInt(1)
		ans2 = big.NewInt(1)
		return
	}
	left := big.NewInt(1)
	right := big.NewInt(0)
	if rightLimit == nil {
		right.Add(num, big.NewInt(0))
	} else {
		right.Add(rightLimit, big.NewInt(0))
	}

	for {
		mid := big.NewInt(0).Add(left, right)
		mid.Rsh(mid, 1)
		midexpcmp := big.NewInt(0).Exp(mid, n, nil).Cmp(num)
		if midexpcmp == -1 { //小于
			left = mid
		} else if midexpcmp == 0 { //等于
			ans1.Add(ans1, mid)
			ans2.Add(ans2, mid)
			break
		} else { //大于
			right = mid
		}
		if big.NewInt(0).Add(left, big.NewInt(1)).Cmp(right) == 0 {
			ans1.Add(ans1, left)
			ans2.Add(ans2, right)
			break
		}

	}

	return
}

// step2：
// https://blog.csdn.net/weixin_39695712/article/details/107054736
// https://en.wikipedia.org/wiki/AKS_primality_test#Example_1:_n_=_31_is_prime
// 找到最小的r，使得ord_r(num) > (log(num))^2
func GetSmallestR(num *big.Int) (r *big.Int) {
	maxk, maxr := LogRangeByBase(num, big.NewInt(2))
	maxk.Exp(maxk, big.NewInt(2), nil)
	maxr.Exp(maxr, big.NewInt(5), nil) // maxr 实际上并不是必须的
	if maxr.Cmp(big.NewInt(3)) < 0 {
		maxr = big.NewInt(3)
	}
	nextR := true
	for r = big.NewInt(2); nextR && r.Cmp(maxr) < 0; r.Add(r, big.NewInt(1)) {
		nextR = false
		for k := big.NewInt(1); (!nextR) && k.Cmp(maxk) <= 0; k.Add(k, big.NewInt(1)) {
			temp := big.NewInt(0).Exp(num, k, r)
			nextR = temp.Cmp(big.NewInt(0)) == 0 || temp.Cmp(big.NewInt(1)) == 0
		}
	}
	// r = r - 1 # 循环多增加了一层
	r.Sub(r, big.NewInt(1))
	return

}

// step3：
// https://blog.51cto.com/u_13424/6307746
// 返回true表示合数，返回false表示素数。
func Ar(num, r *big.Int) bool {
	flag := false
	for a := big.NewInt(0).Set(r); a.Cmp(big.NewInt(1)) > 0; a.Sub(a, big.NewInt(1)) {
		gcdValue := big.NewInt(0).GCD(nil, nil, a, num)
		if gcdValue.Cmp(big.NewInt(1)) > 0 && gcdValue.Cmp(num) < 0 {
			flag = true
			break
		}
	}
	return flag
}

// step5：
func PolynomialComputation(r, n *big.Int) (ans bool) {
	xr := make([]*big.Int, r.Int64()+1)
	for i := 0; i < len(xr); i++ {
		xr[i] = big.NewInt(0)
	}
	xr[0] = big.NewInt(1)
	xr[r.Int64()] = big.NewInt(-1)
	_, b := NthRootRange(EulerTotient(r), big.NewInt(2), nil) //sqrt(phi(r))
	_, t := LogRangeByBase(n, big.NewInt(2))                  //log2_n
	b.Mul(b, t)

	for a := big.NewInt(1); a.Cmp(b) <= 0; a.Add(a, big.NewInt(1)) {
		polynomial1 := ecdlp.PolynomialExpMod([]*big.Int{big.NewInt(1), a}, n, xr, n)             //(x+a)^r mod x^r-1
		polynomial2 := ecdlp.PolynomialExpMod([]*big.Int{big.NewInt(1), big.NewInt(0)}, n, xr, n) //x^r mod x^r-1
		polynomial2 = ecdlp.PolynomialAdd(polynomial2, []*big.Int{a}, n)                          //x^r+a mod x^r-1
		polynomialCha := ecdlp.PolynomialSub(polynomial1, polynomial2, n)
		if len(polynomialCha) != 1 || polynomialCha[0].Cmp(big.NewInt(0)) != 0 {
			return true
		}
	}
	return false
}

//https://blog.51cto.com/u_13424/6307746
/*欧拉函数
摘自维基百科：在数论中，对正整数n，欧拉函数是小于或等于n的正整数中与n互质的数的数目。
例如Euler_Totient(8)=4，因为1,3,5,7均和8互质。
*/
func EulerTotient(r *big.Int) *big.Int {
	count := big.NewInt(0)
	for i := big.NewInt(1); i.Cmp(r) <= 0; i.Add(i, big.NewInt(1)) {
		if big.NewInt(0).GCD(nil, nil, i, r).Cmp(big.NewInt(1)) == 0 {
			count.Add(count, big.NewInt(1))
		}
	}
	return count
}
