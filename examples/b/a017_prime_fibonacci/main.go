package main

import (
	"fmt"
	"math/big"
	"os"
)

// https://sourceforge.net/projects/mpzaprcl/files/latest/download
// mpz_aprcl.c文件中的mpz_fibonacci_prp
// https://en.wikipedia.org/wiki/Lucas_pseudoprime#Fibonacci_pseudoprimes  伪素数705开头
func main() {
	if true {
		errCount := 0
		for n := big.NewInt(2); n.Cmp(big.NewInt(10000)) <= 0; n.Add(n, big.NewInt(1)) {
			//https://oeis.org/A005845 布鲁克曼-卢卡斯伪素数
			// r := FibonacciPrp(n, big.NewInt(1), big.NewInt(-1))
			// https://en.wikipedia.org/wiki/Lucas_pseudoprime#Fibonacci_pseudoprimes 佩尔伪素数第三个定义
			r := FibonacciPrp(n, big.NewInt(2), big.NewInt(-1))
			r2 := n.ProbablyPrime(0)
			if r != r2 {
				errCount++
				fmt.Println("错误", n, r, r2)
			} else {
				if r {
					//fmt.Println("素数", n)
				}
			}
		}
		fmt.Println("错误次数", errCount)
	}
	fmt.Println("")
}

type LucasParam struct {
	P *big.Int
	Q *big.Int
}

func (that *LucasParam) GetUnAndVnMod(k *big.Int, N *big.Int) (*big.Int, *big.Int) {
	if k.Cmp(big.NewInt(0)) == 0 {
		return big.NewInt(0), big.NewInt(2)
	}
	if k.Cmp(big.NewInt(1)) == 0 {
		return big.NewInt(1), big.NewInt(0).Mod(that.P, N)
	}
	ansU := big.NewInt(0)
	ansV := big.NewInt(0)
	var tempAnsU *big.Int
	var tempAndV *big.Int
	doubleU := big.NewInt(1)
	doubleV := big.NewInt(0).Set(that.P)
	temp2U := big.NewInt(0)
	temp2V := big.NewInt(0)
	d := big.NewInt(0).Exp(that.P, big.NewInt(2), nil)
	d.Sub(d, big.NewInt(0).Lsh(that.Q, 2))
	d.Mod(d, N)
	kBitLen := k.BitLen()
	for i := 0; i < kBitLen; i++ {
		if k.Bit(i) != 0 {
			if tempAnsU == nil {
				tempAnsU = big.NewInt(0).Set(doubleU)
				tempAndV = big.NewInt(0).Set(doubleV)
			} else {
				tempAnsU.Mul(ansU, doubleV).Add(tempAnsU, big.NewInt(0).Mul(ansV, doubleU))
				tempAnsU.Mul(tempAnsU, big.NewInt(0).ModInverse(big.NewInt(2), N))

				tempAndV.Mul(d, ansU).Mul(tempAndV, doubleU).Add(tempAndV, big.NewInt(0).Mul(ansV, doubleV))
				tempAndV.Mul(tempAndV, big.NewInt(0).ModInverse(big.NewInt(2), N))

				tempAnsU.Mod(tempAnsU, N)
				tempAndV.Mod(tempAndV, N)
			}
			ansU.Set(tempAnsU)
			ansV.Set(tempAndV)
		}
		temp2U.Mul(doubleU, doubleV)
		temp2V.Mul(d, doubleU).Mul(temp2V, doubleU).Add(temp2V, big.NewInt(0).Exp(doubleV, big.NewInt(2), nil))
		temp2V.Mul(temp2V, big.NewInt(0).ModInverse(big.NewInt(2), N))

		temp2U.Mod(temp2U, N)
		temp2V.Mod(temp2V, N)

		doubleU.Set(temp2U)
		doubleV.Set(temp2V)

	}
	return ansU, ansV
}

func FibonacciPrp(n, p, q *big.Int) bool {
	//   mpz_t pmodn, zP;
	pmodn := big.NewInt(0)
	zP := big.NewInt(0)
	//   mpz_t vl, vh, ql, qh, tmp; /* used for calculating the Lucas V sequence */
	vl := big.NewInt(0)
	vh := big.NewInt(0)
	ql := big.NewInt(0)
	qh := big.NewInt(0)
	tmp := big.NewInt(0)
	//   int s = 0, j = 0;
	s := 0
	j := 0
	d := big.NewInt(0).Exp(p, big.NewInt(2), nil)
	d.Sub(d, big.NewInt(0).Lsh(q, 2))
	//   if (p*p-4*q == 0)
	//     return PRP_ERROR;
	if d.Cmp(big.NewInt(0)) == 0 {
		panic("p*p-4*q==0，不符合条件")
	}

	//   if (((q != 1) && (q != -1)) || (p <= 0))
	//     return PRP_ERROR;
	if (q.Cmp(big.NewInt(1)) != 0 && q.Cmp(big.NewInt(-1)) != 0) || p.Cmp(big.NewInt(0)) <= 0 {
		panic("p和q不符合条件")
	}

	//   if (mpz_cmp_ui(n, 2) < 0)
	//     return PRP_COMPOSITE;
	if n.Cmp(big.NewInt(2)) < 0 {
		return false
	}

	//   if (mpz_divisible_ui_p(n, 2))
	//   {
	//     if (mpz_cmp_ui(n, 2) == 0)
	//       return PRP_PRIME;
	//     else
	//       return PRP_COMPOSITE;
	//   }
	if n.Bit(0) == 0 {
		if n.Cmp(big.NewInt(2)) == 0 {
			return true
		} else {
			return false
		}
	}

	//   mpz_init_set_ui(zP, p);
	//   mpz_init(pmodn);
	//   mpz_mod(pmodn, zP, n);
	zP.Set(p)
	pmodn.Mod(zP, n)

	//   /* mpz_lucasvmod(res, p, q, n, n); */
	//   mpz_init_set_si(vl, 2);
	//   mpz_init_set_si(vh, p);
	//   mpz_init_set_si(ql, 1);
	//   mpz_init_set_si(qh, 1);
	//   mpz_init_set_si(tmp,0);
	vl.Set(big.NewInt(2))
	vh.Set(p)
	ql.Set(big.NewInt(1))
	qh.Set(big.NewInt(1))
	tmp.Set(big.NewInt(0))

	//   s = mpz_scan1(n, 0);
	s = MpzScan1(n)
	for j = n.BitLen() - 1; j >= s+1; j-- {
		//     /* ql = ql*qh (mod n) */
		//     mpz_mul(ql, ql, qh);
		ql.Mul(ql, qh)
		//     mpz_mod(ql, ql, n);
		ql.Mod(ql, n)
		// if mpz_tstbit(n, j) == 1 {
		if n.Bit(j) == 1 {
			//       /* qh = ql*q */
			//       mpz_mul_si(qh, ql, q);
			qh.Mul(ql, q)

			//       /* vl = vh*vl - p*ql (mod n) */
			//       mpz_mul(vl, vh, vl);
			vl.Mul(vh, vl)
			//       mpz_mul_si(tmp, ql, p);
			tmp.Mul(ql, p)
			//       mpz_sub(vl, vl, tmp);
			vl.Sub(vl, tmp)
			//       mpz_mod(vl, vl, n);
			vl.Mod(vl, n)

			//       /* vh = vh*vh - 2*qh (mod n) */
			//       mpz_mul(vh, vh, vh);
			vh.Mul(vh, vh)
			//       mpz_mul_si(tmp, qh, 2);
			tmp.Mul(qh, big.NewInt(2))
			//       mpz_sub(vh, vh, tmp);
			vh.Sub(vh, tmp)
			//       mpz_mod(vh, vh, n);
			vh.Mod(vh, n)
		} else {
			//       /* qh = ql */
			//       mpz_set(qh, ql);
			qh.Set(ql)

			//       /* vh = vh*vl - p*ql (mod n) */
			//       mpz_mul(vh, vh, vl);
			vh.Mul(vh, vl)
			//       mpz_mul_si(tmp, ql, p);
			tmp.Mul(ql, p)
			//       mpz_sub(vh, vh, tmp);
			vh.Sub(vh, tmp)
			//       mpz_mod(vh, vh, n);
			vh.Mod(vh, n)

			//       /* vl = vl*vl - 2*ql (mod n) */
			//       mpz_mul(vl, vl, vl);
			vl.Mul(vl, vl)
			//       mpz_mul_si(tmp, ql, 2);
			tmp.Mul(ql, big.NewInt(2))
			//       mpz_sub(vl, vl, tmp);
			vl.Sub(vl, tmp)
			//       mpz_mod(vl, vl, n);
			vl.Mod(vl, n)
		}
	}
	//   /* ql = ql*qh */
	//   mpz_mul(ql, ql, qh);
	ql.Mul(ql, qh)

	//   /* qh = ql*q */
	//   mpz_mul_si(qh, ql, q);
	qh.Mul(ql, q)

	//   /* vl = vh*vl - p*ql */
	//   mpz_mul(vl, vh, vl);
	vl.Mul(vh, vl)
	//   mpz_mul_si(tmp, ql, p);
	tmp.Mul(ql, p)
	//   mpz_sub(vl, vl, tmp);
	vl.Sub(vl, tmp)

	//   /* ql = ql*qh */
	//   mpz_mul(ql, ql, qh);
	ql.Mul(ql, qh)
	for j = 1; j <= s; j++ {
		//     /* vl = vl*vl - 2*ql (mod n) */
		//     mpz_mul(vl, vl, vl);
		vl.Mul(vl, vl)
		//     mpz_mul_si(tmp, ql, 2);
		tmp.Mul(ql, big.NewInt(2))
		//     mpz_sub(vl, vl, tmp);
		vl.Mul(vl, tmp)
		//     mpz_mod(vl, vl, n);
		vl.Mod(vl, n)

		//     /* ql = ql*ql (mod n) */
		//     mpz_mul(ql, ql, ql);
		ql.Mul(ql, ql)
		//     mpz_mod(ql, ql, n);
		ql.Mod(ql, n)
	}

	//   mpz_mod(vl, vl, n); /* vl contains our return value */
	vl.Mod(vl, n)
	ll := LucasParam{p, q}
	nn := big.NewInt(0).Add(n, big.NewInt(0))

	_, a := ll.GetUnAndVnMod(nn, n)
	if a.Cmp(vl) != 0 {
		fmt.Println(n, "错误", a, vl)
		os.Exit(0)
	} else {
		//fmt.Println(n, "正确", a, vl)
	}
	//   if (mpz_cmp(vl, pmodn) == 0)
	//   {
	//     mpz_clear(zP);
	//     mpz_clear(pmodn);
	//     mpz_clear(vl);
	//     mpz_clear(vh);
	//     mpz_clear(ql);
	//     mpz_clear(qh);
	//     mpz_clear(tmp);
	//     return PRP_PRP;
	//   }
	//   mpz_clear(zP);
	//   mpz_clear(pmodn);
	//   mpz_clear(vl);
	//   mpz_clear(vh);
	//   mpz_clear(ql);
	//   mpz_clear(qh);
	//   mpz_clear(tmp);
	//   return PRP_COMPOSITE;
	if vl.Cmp(pmodn) == 0 {
		return true
	} else {
		return false
	}

} /* method mpz_fibonacci_prp */

func MpzScan1(n *big.Int) int {
	i := 0
	for {
		if n.Bit(i) == 1 {
			return i
		}
	}
}
