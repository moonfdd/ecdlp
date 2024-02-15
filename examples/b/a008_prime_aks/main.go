package main

import (
	"fmt"
	"math/big"
)

// https://eprint.iacr.org/2013/449.pdf
// https://eprint.iacr.org/2006/232 AKS算法的改进

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

func MultiplicativeOrder(n, r *big.Int) *big.Int {
	if big.NewInt(0).GCD(nil, nil, n, r).Cmp(big.NewInt(1)) != 0 { // 如果n,r不互质，不存在Multiplicative_Order
		return big.NewInt(-1)
	} else {
		k := big.NewInt(1)
		temp := big.NewInt(1)
		for {
			// 根据A(k) = a^k%r = (a * a^(k-1))%r = (a%r * a^(k-1)%r)%r = (A(k-1) * a%r)%r
			temp.Mul(temp, big.NewInt(0).Mod(n, r))
			temp.Mod(temp, r)
			if temp.Cmp(big.NewInt(1)) == 0 {
				break
			} else {
				k.Add(k, big.NewInt(1))
			}
		}
		return k
	}
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
		n0, n1 := NthRootRange(num, b, nil)
		if n0.Cmp(n1) == 0 {
			ans = true
			return
		} else if n0.Cmp(big.NewInt(1)) == 0 {
			return
		}
	}

	return
}

// https://blog.51cto.com/u_13424/6307746
// 找到最小的r，使得ord_r(num) > log2(num)
func GetSmallestR(num *big.Int) (r *big.Int) {
	r = big.NewInt(2)
	for {
		if big.NewInt(0).GCD(nil, nil, num, r).Cmp(big.NewInt(1)) == 0 {
			multiOrder := MultiplicativeOrder(num, r)
			_, t := LogRangeByBase(num, big.NewInt(2))
			if multiOrder.Cmp(big.NewInt(0).Exp(t, big.NewInt(2), nil)) >= 0 {
				return r
			}
		}
		r.Add(r, big.NewInt(1))
	}
}

// step2：
// https://blog.csdn.net/weixin_39695712/article/details/107054736
// 找到最小的r，使得ord_r(num) > log2(num)
func GetSmallestR2(num *big.Int) (r *big.Int) {
	maxk, maxr := LogRangeByBase(num, big.NewInt(2))
	maxk.Exp(maxk, big.NewInt(2), nil)
	maxr.Exp(maxr, big.NewInt(5), nil)
	if maxr.Cmp(big.NewInt(3)) < 0 {
		maxr = big.NewInt(3)
	}
	nextR := true
	for r = big.NewInt(2); r.Cmp(maxr) < 0; r.Add(r, big.NewInt(1)) {
		if !nextR {
			break
		}
		nextR = false
		for k := big.NewInt(1); k.Cmp(maxk) <= 0; k.Add(k, big.NewInt(1)) {
			if nextR {
				break
			}
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
	for a := big.NewInt(1); a.Cmp(r) <= 0; a.Add(a, big.NewInt(1)) {
		gcdValue := big.NewInt(0).GCD(nil, nil, a, num)
		if gcdValue.Cmp(big.NewInt(1)) > 0 && gcdValue.Cmp(num) < 0 {
			flag = true
			break
		}
	}
	return flag
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

func EulerTotient2(r *big.Int) *big.Int {
	res := big.NewInt(0).Add(r, big.NewInt(0))
	a := big.NewInt(0).Add(r, big.NewInt(0))
	for i := big.NewInt(2); i.Cmp(a) <= 0; i.Add(i, big.NewInt(1)) {
		if big.NewInt(0).Mod(a, i).Cmp(big.NewInt(0)) == 0 {
			res.Div(res, big.NewInt(0).Mul(i, big.NewInt(0).Sub(i, big.NewInt(1))))
			for big.NewInt(0).Mod(a, i).Cmp(big.NewInt(0)) == 0 {
				a.Div(a, i)
			}
		}
	}
	if a.Cmp(big.NewInt(1)) > 0 {
		res.Div(res, big.NewInt(0).Mul(a, big.NewInt(0).Sub(a, big.NewInt(1))))
	}
	return res
}

// https://blog.csdn.net/weixin_39695712/article/details/107054736
// step5
// 返回true表示合数，返回false表示素数。
func PolynomialComputation(r, num *big.Int) bool {
	// b, _ := NthRootRange(EulerTotient(r), big.NewInt(2), nil)
	// t, _ := LogRangeByBase(num, big.NewInt(2))
	_, b := NthRootRange(EulerTotient(r), big.NewInt(2), nil)
	_, t := LogRangeByBase(num, big.NewInt(2))
	b.Mul(b, t)
	oper := NewPolynomialOperation(r, num)
	for a := big.NewInt(1); a.Cmp(b) < 0; a.Add(a, big.NewInt(1)) {

		if true {
			xAddA := []*big.Int{big.NewInt(1), a}
			xAddA = oper.PowerMod(xAddA, num)
			x := []*big.Int{big.NewInt(1), big.NewInt(0)}
			x = oper.PowerMod(x, num)
			result := oper.AddMod(xAddA, oper.NegMod(x))
			result = oper.AddMod(result, oper.NegMod([]*big.Int{a}))
			if len(result) != 1 || result[0].Cmp(big.NewInt(0)) != 0 {
				return true
			}
		}
		//x=0，a变成-a，一个特例
		if false {
			zuo := big.NewInt(0).Exp(big.NewInt(0).Sub(big.NewInt(0), a), num, num)
			you := big.NewInt(0).Sub(big.NewInt(0), a)
			you.Mod(you, num)
			if zuo.Cmp(you) != 0 {
				return true
			}
		}
	}
	return false
}

func Aks(num *big.Int) bool {
	if num.Cmp(big.NewInt(1)) == 0 {
		return false
	}
	//step1
	// fmt.Println("第1步")
	if IsPower(num) {
		return false
	}
	//step2
	// fmt.Println("第2步")
	r := GetSmallestR(num)
	//step3,存在a ≤ r，使得1 < gcd(a,n) < n 成立。返回合数,
	// fmt.Println("第3步")
	if Ar(num, r) {
		return false
	}
	//step4
	// fmt.Println("第4步")
	if num.Cmp(r) <= 0 {
		return true
	}
	// fmt.Println("第5步")
	if PolynomialComputation(r, num) {
		return false
	}
	// fmt.Println("第6步")
	return true
}

// 多项式相关操作
type PolynomialOperation struct {
	// 系数取模
	Num *big.Int
	//多项式取模x^r-1
	Xr []*big.Int
}

// r>=1
func NewPolynomialOperation(r, num *big.Int) *PolynomialOperation {
	ans := &PolynomialOperation{}
	ans.Num = num
	ans.Xr = make([]*big.Int, r.Int64()+1)
	ans.Xr[0] = big.NewInt(1)
	ans.Xr[r.Int64()] = big.NewInt(-1)
	for i := 1; i < len(ans.Xr)-1; i++ {
		ans.Xr[i] = big.NewInt(0)
	}
	return ans
}

// 多项式幂模 num>=1
func (this *PolynomialOperation) PowerMod(polynomial []*big.Int, num *big.Int) (ans []*big.Int) {
	b := big.NewInt(0).Add(num, big.NewInt(0))
	ans = []*big.Int{big.NewInt(1)}
	a := polynomial
	for b.Cmp(big.NewInt(0)) != 0 { //b!=0
		if big.NewInt(0).And(b, big.NewInt(1)).Cmp(big.NewInt(0)) != 0 { //b&1!=0
			ans = this.MulMod(ans, a)
		}
		b.Rsh(b, 1) //b>>=1
		a = this.MulMod(a, a)
	}
	return
}

// 两个多项式相乘，并且求模
func (this *PolynomialOperation) MulMod(polynomial1 []*big.Int, polynomial2 []*big.Int) (ans []*big.Int) {
	if len(polynomial1) == 0 {
		panic("第1个多项式不能为空")
	}
	if len(polynomial2) == 0 {
		panic("第2个多项式不能为空")
	}
	polynomial1 = this.Mod(polynomial1)
	polynomial2 = this.Mod(polynomial2)
	temp := make([]*big.Int, (len(polynomial1)-1)+(len(polynomial2)-1)+1)
	for i := 0; i < len(temp); i++ {
		temp[i] = big.NewInt(0)
	}
	for i := 0; i < len(polynomial1); i++ {
		for j := 0; j < len(polynomial2); j++ {
			temp[i+j].Add(temp[i+j], big.NewInt(0).Mul(polynomial1[i], polynomial2[j]))
			temp[i+j].Mod(temp[i+j], this.Num)
		}
	}
	ans = this.Mod(temp)
	k := 0
	for k < len(ans)-1 {
		if ans[k].Cmp(big.NewInt(0)) != 0 {
			break
		}
		k++
	}
	ans = ans[k:]
	return
}

// 两个多项式相加，并且求模
func (this *PolynomialOperation) AddMod(polynomial1 []*big.Int, polynomial2 []*big.Int) (ans []*big.Int) {
	if len(polynomial1) == 0 {
		panic("第1个多项式不能为空")
	}
	if len(polynomial2) == 0 {
		panic("第2个多项式不能为空")
	}
	//假设第1个多项式的长度大于第2个多项式的长度
	if len(polynomial1) < len(polynomial2) {
		polynomial1, polynomial2 = polynomial2, polynomial1
	}

	temp := this.copy(polynomial1)

	for i := 0; i < len(polynomial2); i++ {
		temp[i+len(polynomial1)-len(polynomial2)].Add(temp[i+len(polynomial1)-len(polynomial2)], polynomial2[i])

		temp[i+len(polynomial1)-len(polynomial2)].Mod(temp[i+len(polynomial1)-len(polynomial2)], this.Num)
	}
	ans = this.Mod(temp)
	k := 0
	for k < len(ans)-1 {
		if ans[k].Cmp(big.NewInt(0)) != 0 {
			break
		}
		k++
	}
	ans = ans[k:]
	return
}

// 取反
func (this *PolynomialOperation) NegMod(polynomial []*big.Int) (ans []*big.Int) {
	if len(polynomial) == 0 {
		panic("多项式不能为空")
	}

	ans = this.copy(polynomial)
	for i := 0; i < len(ans); i++ {
		ans[i].Neg(ans[i])
		ans[i].Mod(ans[i], this.Num)
	}
	k := 0
	for k < len(ans)-1 {
		if ans[k].Cmp(big.NewInt(0)) != 0 {
			break
		}
		k++
	}
	ans = ans[k:]
	return
}

// 求模
func (this *PolynomialOperation) Mod(polynomial []*big.Int) (ans []*big.Int) {
	if len(polynomial) == 0 {
		panic("多项式不能为空")
	}
	zero := big.NewInt(0)
	if len(this.Xr) == 1 {
		ans = []*big.Int{big.NewInt(0).Add(polynomial[len(polynomial)-1], zero)}
		ans[0].Mod(ans[0], this.Num)
		return
	}

	temp := this.copy(polynomial)
	for i := 0; i <= len(temp)-len(this.Xr); i++ {
		if temp[i].Cmp(zero) == 0 {
			continue
		}
		//消减多项式
		// 确定消减系数
		coefficient := big.NewInt(0).Neg(temp[i])
		// 遍历消减多项式
		for j := 0; j < len(this.Xr); j++ {
			temp[i+j].Add(temp[i+j], big.NewInt(0).Mul(coefficient, this.Xr[j]))
			temp[i+j].Mod(temp[i+j], this.Num)
		}
	}
	if len(temp)-len(this.Xr) >= 0 {
		k := len(temp) - len(this.Xr) + 1
		for ; k < len(temp)-1; k++ {
			if temp[k].Cmp(big.NewInt(0)) != 0 {
				break
			}
		}
		ans = temp[k:]
	} else {
		for i := 0; i < len(temp); i++ {
			temp[i].Mod(temp[i], this.Num)
		}
		ans = temp
	}
	return
}

// 拷贝
func (this *PolynomialOperation) copy(polynomial []*big.Int) []*big.Int {
	ans := make([]*big.Int, len(polynomial))
	for i := 0; i < len(ans); i++ {
		ans[i] = big.NewInt(0)
		ans[i].Add(polynomial[i], big.NewInt(0))
	}
	return ans
}

func main() {
	if true {
		//96091
		//95477
		num := big.NewInt(0)
		// num.SetString("9495007", 10)
		num.SetString("1", 10)
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
