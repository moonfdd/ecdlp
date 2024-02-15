package main

import (
	"fmt"
	"math/big"
	"os"
)

// https://eprint.iacr.org/2008/124 Frobenius 伪素数检验的简单推导

func JacobiSymbol(a, n *big.Int) *big.Int {
	if n.Cmp(big.NewInt(0)) <= 0 || big.NewInt(0).Mod(n, big.NewInt(2)).Cmp(big.NewInt(1)) != 0 {
		panic("不符合条件")
	}
	a = big.NewInt(0).Set(a)
	n = big.NewInt(0).Set(n)
	a.Mod(a, n)
	t := big.NewInt(1)
	var r *big.Int
	for a.Cmp(big.NewInt(0)) != 0 {
		for a.Bit(0) == 0 {
			a.Rsh(a, 1)
			r = new(big.Int).Mod(n, big.NewInt(8))
			if r.Cmp(big.NewInt(3)) == 0 || r.Cmp(big.NewInt(5)) == 0 {
				t.Neg(t)
			}
		}
		r = n
		n = a
		a = r
		if a.Bit(1) == 1 && n.Bit(1) == 1 {
			t.Neg(t)
		}
		a.Mod(a, n)
	}
	if n.Cmp(big.NewInt(1)) == 0 {
		return t
	} else {
		return big.NewInt(0)
	}
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

// https://en.wikipedia.org/wiki/Lucas_pseudoprime#Fibonacci_pseudoprimes

// var ulDmax = new(big.Int) /* tracks global max of Lucas-Selfridge |D| */

func LucasSelfridge(mpzN *big.Int) bool {
	/* Test mpzN for primality using Lucas's test with Selfridge's parameters.
	   Returns 1 if mpzN is prime or a Lucas-Selfridge pseudoprime. Returns
	   0 if mpzN is definitely composite. Note that a Lucas-Selfridge test
	   typically requires three to seven times as many bit operations as a
	   single Miller's test. The frequency of Lucas-Selfridge pseudoprimes
	   appears to be roughly four times that of base-2 strong pseudoprimes;
	   the Baillie-PSW test is based on the hope (verified by the author,
	   May, 2005, for all N < 10^13; and by Martin Fuller, January, 2007,
	   for all N < 10^15) that the two tests have no common pseudoprimes. */

	// int iComp2, iP, iJ, iSign;
	// long lDabs, lD, lQ;
	// unsigned long ulMaxBits, ulNbits, ul, ulGCD;
	// mpz_t mpzU, mpzV, mpzNplus1, mpzU2m, mpzV2m, mpzQm, mpz2Qm,
	//       mpzT1, mpzT2, mpzT3, mpzT4, mpzD;

	iComp2 := 0
	iP := big.NewInt(0)
	iJ := big.NewInt(0)
	iSign := big.NewInt(0)

	lDabs := big.NewInt(0)
	lD := big.NewInt(0)
	lQ := big.NewInt(0)

	// ulMaxBits := big.NewInt(0)
	ulNbits := 0
	// ul := big.NewInt(0)
	ulGCD := big.NewInt(0)

	mpzU := big.NewInt(0)
	mpzV := big.NewInt(0)
	mpzNplus1 := big.NewInt(0)
	mpzU2m := big.NewInt(0)
	mpzV2m := big.NewInt(0)
	mpzQm := big.NewInt(0)
	mpz2Qm := big.NewInt(0)
	mpzT1 := big.NewInt(0)
	mpzT2 := big.NewInt(0)
	mpzT3 := big.NewInt(0)
	mpzT4 := big.NewInt(0)
	mpzD := big.NewInt(0)

	// /* This implementation of the algorithm assumes N is an odd integer > 2,
	//    so we first eliminate all N < 3 and all even N. As a practical matter,
	//    we also need to filter out all perfect square values of N, such as
	//    1093^2 (a base-2 strong pseudoprime); this is because we will later
	//    require an integer D for which Jacobi(D,N) = -1, and no such integer
	//    exists if N is a perfect square. The algorithm as written would
	//    still eventually return zero in this case, but would require
	//    nearly sqrt(N)/2 iterations. */

	// iComp2=mpz_cmp_si(mpzN, 2);
	iComp2 = mpzN.Cmp(big.NewInt(2))
	// if(iComp2 < 0)return(0);
	if iComp2 < 0 {
		return false
	}
	// if(iComp2==0)return(1);
	if iComp2 == 0 {
		return true
	}
	// if(mpz_even_p(mpzN))return(0);
	if mpzN.Bit(0) == 0 {
		return false
	}
	// if(mpz_perfect_square_p(mpzN))return(0);
	sq := big.NewInt(0).Sqrt(mpzN)
	sq.Mul(sq, sq)
	if sq.Cmp(mpzN) == 0 {
		return false
	}

	// /* Allocate storage for the mpz_t variables. Most require twice
	//    the storage of mpzN, since multiplications of order O(mpzN)*O(mpzN)
	//    will be performed. */

	// ulMaxBits=2*mpz_sizeinbase(mpzN, 2) + mp_bits_per_limb;
	// mpz_init2(mpzU, ulMaxBits);
	// mpz_init2(mpzV, ulMaxBits);
	// mpz_init2(mpzNplus1, ulMaxBits);
	// mpz_init2(mpzU2m, ulMaxBits);
	// mpz_init2(mpzV2m, ulMaxBits);
	// mpz_init2(mpzQm, ulMaxBits);
	// mpz_init2(mpz2Qm, ulMaxBits);
	// mpz_init2(mpzT1, ulMaxBits);
	// mpz_init2(mpzT2, ulMaxBits);
	// mpz_init2(mpzT3, ulMaxBits);
	// mpz_init2(mpzT4, ulMaxBits);
	// mpz_init(mpzD);

	// /* Find the first element D in the sequence {5, -7, 9, -11, 13, ...}
	//    such that Jacobi(D,N) = -1 (Selfridge's algorithm). Although
	//    D will nearly always be "small" (perfect square N's having
	//    been eliminated), an overflow trap for D is present. */

	// lDabs=5;
	lDabs = big.NewInt(5)
	// iSign=1;
	iSign = big.NewInt(1)
	for {
		//   lD=iSign*lDabs;
		lD.Mul(iSign, lDabs)
		//   iSign = -iSign;
		iSign.Neg(iSign)
		//   ulGCD=mpz_gcd_ui(NULL, mpzN, lDabs);
		ulGCD.GCD(nil, nil, mpzN, lDabs)
		//   /* if 1 < GCD < N then N is composite with factor lDabs, and
		//      Jacobi(D,N) is technically undefined (but often returned
		//      as zero). */
		//   if((ulGCD > 1) && mpz_cmp_ui(mpzN, ulGCD) > 0)RETURN(0);
		if ulGCD.Cmp(big.NewInt(1)) > 0 && mpzN.Cmp(ulGCD) > 0 {
			return false
		}
		//   mpz_set_si(mpzD, lD);
		mpzD.Add(lD, big.NewInt(0))
		//   iJ=mpz_jacobi(mpzD, mpzN);
		iJ = JacobiSymbol(mpzD, mpzN)
		//   if(iJ==-1)break;
		// fmt.Println(mpzD, mpzN, iJ)
		// return false
		if iJ.Cmp(big.NewInt(-1)) == 0 {
			break
		}
		//   lDabs += 2;
		lDabs.Add(lDabs, big.NewInt(2))
		//   if(lDabs > ulDmax)ulDmax=lDabs;  /* tracks global max of |D| */
		// if lDabs.Cmp(ulDmax) > 0 {
		// 	ulDmax.Add(lDabs, big.NewInt(0))
		// }
		// if lDabs.Cmp(big.NewInt(math.MaxInt32-2)) > 0 {
		// 	fmt.Println("错误")
		// 	return false
		// }
	}

	// iP=1;         /* Selfridge's choice */
	iP.Add(big.NewInt(1), big.NewInt(0))
	// lQ=(1-lD)/4;  /* Required so D = P*P - 4*Q */
	lQ.Sub(big.NewInt(1), lD)
	lQ.Rsh(lQ, 2)
	// fmt.Println("iP = ", iP)
	// fmt.Println("lQ = ", lQ)
	// fmt.Println("lD = ", lD)

	// /* NOTE: The conditions (a) N does not divide Q, and
	//    (b) D is square-free or not a perfect square, are included by
	//    some authors; e.g., "Prime numbers and computer methods for
	//    factorization," Hans Riesel (2nd ed., 1994, Birkhauser, Boston),
	//    p. 130. For this particular application of Lucas sequences,
	//    these conditions were found to be immaterial. */

	// mpz_add_ui(mpzNplus1, mpzN, 1); /* must compute U_(N - Jacobi(D,N)) */
	mpzNplus1.Add(mpzN, big.NewInt(1))

	// /* mpzNplus1 is always even, so the accumulated values U and V
	//    are initialized to U_0 and V_0 (if the target index were odd,
	//    U and V would be initialized to U_1=1 and V_1=P). In either case,
	//    the values of U_2m and V_2m are initialized to U_1 and V_1;
	//    the FOR loop calculates in succession U_2 and V_2, U_4 and
	//    V_4, U_8 and V_8, etc. If the corresponding bits of N+1 are
	//    on, these values are then combined with the previous totals
	//    for U and V, using the composition formulas for addition
	//    of indices. */

	// mpz_set_ui(mpzU, 0);           /* U=U_0 */
	mpzU.Add(big.NewInt(0), big.NewInt(0))
	// mpz_set_ui(mpzV, 2);           /* V=V_0 */
	mpzV.Add(big.NewInt(2), big.NewInt(0))
	// mpz_set_ui(mpzU2m, 1);         /* U_1 */
	mpzU2m.Add(big.NewInt(1), big.NewInt(0))
	// mpz_set_si(mpzV2m, iP);        /* V_1 */
	mpzV2m.Add(iP, big.NewInt(0))
	// mpz_set_si(mpzQm, lQ);
	mpzQm.Add(lQ, big.NewInt(0))
	// mpz_set_si(mpz2Qm, 2*lQ);
	mpz2Qm.Add(lQ, big.NewInt(0))
	mpz2Qm.Lsh(mpz2Qm, 1)

	// ulNbits=mpz_sizeinbase(mpzNplus1, 2);
	ulNbits = mpzNplus1.BitLen()
	// for(ul=1; ul < ulNbits; ul++)  /* zero bit off, already accounted for */
	for ul := 1; ul < ulNbits; ul++ { /* zero bit off, already accounted for */
		// fmt.Println("mpzU = ", mpzU)
		// fmt.Println("mpzV = ", mpzV)
		// /* Formulas for doubling of indices (carried out mod N). Note that
		//  * the indices denoted as "2m" are actually powers of 2, specifically
		//  * 2^(ul-1) beginning each loop and 2^ul ending each loop.
		//  *
		//  * U_2m = U_m*V_m
		//  * V_2m = V_m*V_m - 2*Q^m
		//  */
		//   mpz_mul(mpzU2m, mpzU2m, mpzV2m);
		mpzU2m.Mul(mpzU2m, mpzV2m)
		//   mpz_mod(mpzU2m, mpzU2m, mpzN);
		mpzU2m.Mod(mpzU2m, mpzN)
		//   mpz_mul(mpzV2m, mpzV2m, mpzV2m);
		mpzV2m.Mul(mpzV2m, mpzV2m)
		//   mpz_sub(mpzV2m, mpzV2m, mpz2Qm);
		mpzV2m.Sub(mpzV2m, mpz2Qm)
		//   mpz_mod(mpzV2m, mpzV2m, mpzN);
		mpzV2m.Mod(mpzV2m, mpzN)
		// if mpz_tstbit(mpzNplus1, ul) {
		if mpzNplus1.Bit(ul) == 1 {
			// /* Formulas for addition of indices (carried out mod N);
			//  *
			//  * U_(m+n) = (U_m*V_n + U_n*V_m)/2
			//  * V_(m+n) = (V_m*V_n + D*U_m*U_n)/2
			//  *
			//  * Be careful with division by 2 (mod N)!
			//  */
			//     mpz_mul(mpzT1, mpzU2m, mpzV);
			mpzT1.Mul(mpzU2m, mpzV)
			//     mpz_mul(mpzT2, mpzU, mpzV2m);
			mpzT2.Mul(mpzU, mpzV2m)
			//     mpz_mul(mpzT3, mpzV2m, mpzV);
			mpzT3.Mul(mpzV2m, mpzV)
			//     mpz_mul(mpzT4, mpzU2m, mpzU);
			mpzT4.Mul(mpzU2m, mpzU)
			//     mpz_mul_si(mpzT4, mpzT4, lD);
			mpzT4.Mul(mpzT4, lD)
			//     mpz_add(mpzU, mpzT1, mpzT2);
			mpzU.Add(mpzT1, mpzT2)
			//     if(mpz_odd_p(mpzU))mpz_add(mpzU, mpzU, mpzN);
			if mpzU.Bit(0) == 1 {
				mpzU.Add(mpzU, mpzN)
			}
			//     mpz_fdiv_q_2exp(mpzU, mpzU, 1);
			mpzU.Rsh(mpzU, 1)
			//     mpz_add(mpzV, mpzT3, mpzT4);
			mpzV.Add(mpzT3, mpzT4)
			//     if(mpz_odd_p(mpzV))mpz_add(mpzV, mpzV, mpzN);
			if mpzV.Bit(0) == 1 {
				mpzV.Add(mpzV, mpzN)
			}
			//     mpz_fdiv_q_2exp(mpzV, mpzV, 1);
			mpzV.Rsh(mpzV, 1)
			//     mpz_mod(mpzU, mpzU, mpzN);
			mpzU.Mod(mpzU, mpzN)
			// fmt.Println("mpzU = ", mpzU)
			//     mpz_mod(mpzV, mpzV, mpzN);
			mpzV.Mod(mpzV, mpzN)
		}
		// /* Calculate Q^m for next bit position, doubling the exponent.
		//    The irrelevant final iteration is omitted. */
		if ul < ulNbits-1 { /* Q^m not needed for MSB. */

			//     mpz_mul(mpzQm, mpzQm, mpzQm);
			mpzQm.Mul(mpzQm, mpzQm)
			//     mpz_mod(mpzQm, mpzQm, mpzN);  /* prevents overflow */
			mpzQm.Mod(mpzQm, mpzN)
			//     mpz_add(mpz2Qm, mpzQm, mpzQm);
			mpz2Qm.Add(mpzQm, mpzQm)
		}
	}
	ll := LucasParam{iP, lQ}
	// fmt.Println(iP, lQ, mpzN)
	a, b := ll.GetUnAndVnMod(big.NewInt(0).Rsh(mpzNplus1, 0), mpzN)
	// fmt.Println("a = ", a, mpzN)
	// fmt.Println("mpzU = ", mpzU, mpzV, mpzN)
	if a.Cmp(mpzU) != 0 {
		fmt.Println("0错误", a, mpzU)
		os.Exit(0)
	}
	if b.Cmp(mpzV) != 0 {
		fmt.Println("1错误", b, mpzV)
		os.Exit(0)
	}

	// /* If U_(N - Jacobi(D,N)) is congruent to 0 mod N, then N is
	//    a prime or a Lucas pseudoprime; otherwise it is definitely
	//    composite. */

	if mpzU.Cmp(big.NewInt(0)) == 0 {
		return true
	}
	return false
	// if(mpz_sgn(mpzU)==0)RETURN(1);
	// RETURN(0);
}

func MillerRabbinA(a, num *big.Int) bool {
	if num.Cmp(big.NewInt(1)) == 0 {
		return false
	}
	if num.Cmp(big.NewInt(2)) == 0 {
		return true
	}
	if big.NewInt(0).And(num, big.NewInt(1)).Cmp(big.NewInt(0)) == 0 {
		return false
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

func Fermat(num *big.Int) bool {
	if num.Cmp(big.NewInt(1)) == 0 {
		return false
	}
	if num.Cmp(big.NewInt(2)) == 0 {
		return true
	}
	if big.NewInt(0).Exp(big.NewInt(2), big.NewInt(0).Add(num, big.NewInt(-1)), num).Cmp(big.NewInt(1)) == 0 {
		return true
	} else {
		return false
	}
}

func mpz_selfridge_prp(n *big.Int) bool {
	//   long int d = 5, p = 1, q = 0;
	d := big.NewInt(5)
	p := big.NewInt(1)
	q := big.NewInt(0)
	//   int max_d = 1000000;
	max_d := big.NewInt(1000000)
	//   int jacobi = 0;
	jacobi := big.NewInt(0)
	//   mpz_t zD;
	zD := big.NewInt(0)

	//   if (mpz_cmp_ui(n, 2) < 0)
	//     return PRP_COMPOSITE;
	if n.Cmp(big.NewInt(2)) < 0 {
		return false
	}

	//   if (mpz_divisible_ui_p(n, 2)){
	//     if (mpz_cmp_ui(n, 2) == 0)
	//       return PRP_PRIME;
	//     else
	//       return PRP_COMPOSITE;
	//   }
	if n.Cmp(big.NewInt(2)) == 0 {
		return true
	}

	if n.Bit(0) == 0 {
		return false
	}

	//   mpz_init_set_ui(zD, d);
	zD.Set(d)

	for {
		//     jacobi = mpz_jacobi(zD, n);
		jacobi = JacobiSymbol(zD, n)

		/* if jacobi == 0, d is a factor of n, therefore n is composite... */
		/* if d == n, then either n is either prime or 9... */
		if jacobi.Cmp(big.NewInt(0)) == 0 {
			//   if ((mpz_cmpabs(zD, n) == 0) && (mpz_cmp_ui(zD, 9) != 0))
			if big.NewInt(0).Abs(zD).Cmp(n) == 0 && zD.Cmp(big.NewInt(9)) != 0 {
				// mpz_clear(zD);
				// return PRP_PRIME;
				return true
			} else {
				// mpz_clear(zD);
				// return PRP_COMPOSITE;
				return false
			}
		}
		//     if (jacobi == -1)
		//       break;
		if jacobi.Cmp(big.NewInt(-1)) == 0 {
			break
		}

		/* if we get to the 5th d, make sure we aren't dealing with a square... */
		// if d == 13{
		if d.Cmp(big.NewInt(13)) == 0 {

			//   if (mpz_perfect_square_p(n))
			//   {
			//     mpz_clear(zD);
			//     return PRP_COMPOSITE;
			//   }
			sq := big.NewInt(0).Sqrt(n)
			sq.Mul(sq, sq)
			if sq.Cmp(n) == 0 {
				return false
			}
		}

		// if d < 0 {
		if d.Cmp(big.NewInt(0)) < 0 {
			// d *= -1
			d.Mul(d, big.NewInt(-1))
			// d += 2
			d.Add(d, big.NewInt(2))
		} else {
			// d += 2
			d.Add(d, big.NewInt(2))
			// d *= -1
			d.Mul(d, big.NewInt(-1))
		}

		/* make sure we don't search forever */
		//     if (d >= max_d)
		if d.Cmp(max_d) >= 0 {
			//   mpz_clear(zD);
			//   return PRP_ERROR;
			panic("make sure we don't search forever")

		}

		//     mpz_set_si(zD, d);
		zD.Set(d)
	}

	//   mpz_clear(zD);

	//   q = (1-d)/4;
	q.Sub(big.NewInt(1), d).Div(q, big.NewInt(4))

	//   return mpz_lucas_prp(n, p, q);
	// fmt.Println(p, q, n)
	return mpz_lucas_prp(n, p, q)

} /* method mpz_selfridge_prp */

/* *******************************************************************************
 * mpz_lucas_prp:
 * A "Lucas pseudoprime" with parameters (P,Q) is a composite n with D=P^2-4Q,
 * (n,2QD)=1 such that U_(n-(D/n)) == 0 mod n [(D/n) is the Jacobi symbol]
 * *******************************************************************************/
func mpz_lucas_prp(n, p, q *big.Int) bool {

	//    mpz_t zD;
	zD := big.NewInt(0)
	//    mpz_t res;
	res := big.NewInt(0)
	//    mpz_t index;
	index := big.NewInt(0)
	//    mpz_t uh, vl, vh, ql, qh, tmp; /* used for calculating the Lucas U sequence */
	uh := big.NewInt(0)
	vl := big.NewInt(0)
	vh := big.NewInt(0)
	ql := big.NewInt(0)
	qh := big.NewInt(0)
	tmp := big.NewInt(0)

	//    int s = 0, j = 0;
	var s = 0
	var j = 0
	//    int ret = 0;
	ret := big.NewInt(0)
	//    long int d = p*p - 4*q;
	d := big.NewInt(0)
	d.Mul(p, p)
	d.Sub(d, big.NewInt(0).Mul(q, big.NewInt(4)))

	//    if (d == 0) /* Does not produce a proper Lucas sequence */
	// 	 return PRP_ERROR;
	if d.Cmp(big.NewInt(0)) == 0 {
		panic("Does not produce a proper Lucas sequence")
	}

	//    if (mpz_cmp_ui(n, 2) < 0)
	// 	 return PRP_COMPOSITE;
	if n.Cmp(big.NewInt(2)) < 0 {
		return false
	}

	//    if (mpz_divisible_ui_p(n, 2))
	//    {
	// 	 if (mpz_cmp_ui(n, 2) == 0)
	// 	   return PRP_PRIME;
	// 	 else
	// 	   return PRP_COMPOSITE;
	//    }
	if n.Cmp(big.NewInt(2)) == 0 {
		return true
	}

	if n.Bit(0) == 0 {
		return false
	}

	//    mpz_init(index);
	//    mpz_init_set_si(zD, d);
	zD.Set(d)
	//    mpz_init(res);

	//    mpz_mul_si(res, zD, q);
	res.Mul(zD, q)
	//    mpz_mul_ui(res, res, 2);
	res.Mul(res, big.NewInt(2))
	//    mpz_gcd(res, res, n);
	res.GCD(nil, nil, res, n)
	//    if ((mpz_cmp(res, n) != 0) && (mpz_cmp_ui(res, 1) > 0))
	if res.Cmp(n) != 0 && res.Cmp(big.NewInt(1)) > 0 {

		// 	 mpz_clear(zD);
		// 	 mpz_clear(res);
		// 	 mpz_clear(index);
		// 	 return PRP_COMPOSITE;
		return false
	}

	//    /* index = n-(D/n), where (D/n) is the Jacobi symbol */
	//    mpz_set(index, n);
	index.Set(n)
	//    ret = mpz_jacobi(zD, n);
	ret.Set(JacobiSymbol(zD, n))
	//    if (ret == -1)
	// 	 mpz_add_ui(index, index, 1);
	//    else if (ret == 1)
	// 	 mpz_sub_ui(index, index, 1);
	if ret.Cmp(big.NewInt(-1)) == 0 {
		index.Add(index, big.NewInt(1))
	} else if ret.Cmp(big.NewInt(1)) == 0 {
		index.Sub(index, big.NewInt(1))
	}

	//    /* mpz_lucasumod(res, p, q, index, n); */
	//    mpz_init_set_si(uh, 1);
	uh.Set(big.NewInt(1))
	//    mpz_init_set_si(vl, 2);
	vl.Set(big.NewInt(2))
	//    mpz_init_set_si(vh, p);
	vh.Set(p)
	//    mpz_init_set_si(ql, 1);
	ql.Set(big.NewInt(1))
	//    mpz_init_set_si(qh, 1);
	qh.Set(big.NewInt(1))
	//    mpz_init_set_si(tmp,0);
	tmp.Set(big.NewInt(0))

	//    s = mpz_scan1(index, 0);
	s = MpzScan1(index)

	// for (j = mpz_sizeinbase(index,2)-1; j >= s+1; j--){

	for j = index.BitLen() - 1; j >= s+1; j-- {

		/* ql = ql*qh (mod n) */
		// 	 mpz_mul(ql, ql, qh);
		ql.Mul(ql, qh)
		// 	 mpz_mod(ql, ql, n);
		ql.Mod(ql, n)
		//  if (mpz_tstbit(index,j) == 1)
		if index.Bit(j) == 1 {
			//    /* qh = ql*q */
			//    mpz_mul_si(qh, ql, q);
			qh.Mul(ql, q)

			//    /* uh = uh*vh (mod n) */
			//    mpz_mul(uh, uh, vh);
			uh.Mul(uh, vh)
			//    mpz_mod(uh, uh, n);
			uh.Mod(uh, n)

			//    /* vl = vh*vl - p*ql (mod n) */
			//    mpz_mul(vl, vh, vl);
			vl.Mul(vh, vl)
			//    mpz_mul_si(tmp, ql, p);
			tmp.Mul(ql, p)
			//    mpz_sub(vl, vl, tmp);
			vl.Sub(vl, tmp)
			//    mpz_mod(vl, vl, n);
			vl.Mod(vl, n)

			//    /* vh = vh*vh - 2*qh (mod n) */
			//    mpz_mul(vh, vh, vh);
			vh.Mul(vh, vh)
			//    mpz_mul_si(tmp, qh, 2);
			tmp.Mul(qh, big.NewInt(2))
			//    mpz_sub(vh, vh, tmp);
			vh.Sub(vh, tmp)
			//    mpz_mod(vh, vh, n);
			vh.Mod(vh, n)
		} else {
			//    /* qh = ql */
			//    mpz_set(qh, ql);
			qh.Set(ql)

			/* uh = uh*vl - ql (mod n) */
			//    mpz_mul(uh, uh, vl);
			uh.Mul(uh, vl)
			//    mpz_sub(uh, uh, ql);
			uh.Sub(uh, ql)
			//    mpz_mod(uh, uh, n);
			uh.Mod(uh, n)

			//    /* vh = vh*vl - p*ql (mod n) */
			//    mpz_mul(vh, vh, vl);
			vh.Mul(vh, vl)
			//    mpz_mul_si(tmp, ql, p);
			tmp.Mul(ql, p)
			//    mpz_sub(vh, vh, tmp);
			vh.Sub(vh, tmp)
			//    mpz_mod(vh, vh, n);
			vh.Mod(vh, n)

			//    /* vl = vl*vl - 2*ql (mod n) */
			//    mpz_mul(vl, vl, vl);
			vl.Mul(vl, vl)
			//    mpz_mul_si(tmp, ql, 2);
			tmp.Mul(ql, big.NewInt(2))
			//    mpz_sub(vl, vl, tmp);
			vl.Sub(vl, tmp)
			//    mpz_mod(vl, vl, n);
			vl.Mod(vl, n)
		}
	}
	//    /* ql = ql*qh */
	//    mpz_mul(ql, ql, qh);
	ql.Mul(ql, qh)

	//    /* qh = ql*q */
	//    mpz_mul_si(qh, ql, q);
	qh.Mul(ql, q)

	//    /* uh = uh*vl - ql */
	//    mpz_mul(uh, uh, vl);
	uh.Mul(uh, vl)
	//    mpz_sub(uh, uh, ql);
	uh.Sub(uh, ql)

	//    /* vl = vh*vl - p*ql */
	//    mpz_mul(vl, vh, vl);
	vl.Mul(vh, vl)
	//    mpz_mul_si(tmp, ql, p);
	tmp.Mul(ql, p)
	//    mpz_sub(vl, vl, tmp);
	vl.Sub(vl, tmp)

	//    /* ql = ql*qh */
	//    mpz_mul(ql, ql, qh);
	ql.Mul(ql, qh)

	for j = 1; j <= s; j++ {
		//  /* uh = uh*vl (mod n) */
		//  mpz_mul(uh, uh, vl);
		uh.Mul(uh, vl)
		//  mpz_mod(uh, uh, n);
		uh.Mod(uh, n)

		//  /* vl = vl*vl - 2*ql (mod n) */
		//  mpz_mul(vl, vl, vl);
		vl.Mul(vl, vl)
		//  mpz_mul_si(tmp, ql, 2);
		tmp.Mul(ql, big.NewInt(2))
		//  mpz_sub(vl, vl, tmp);
		vl.Sub(vl, tmp)
		//  mpz_mod(vl, vl, n);
		vl.Mod(vl, n)

		//  /* ql = ql*ql (mod n) */
		//  mpz_mul(ql, ql, ql);
		ql.Mul(ql, ql)
		//  mpz_mod(ql, ql, n);
		ql.Mod(ql, n)
	}

	//    mpz_mod(res, uh, n); /* uh contains our return value */
	res.Mod(uh, n)

	ll := LucasParam{p, q}
	// fmt.Println(iP, lQ, mpzN)
	a, b := ll.GetUnAndVnMod(big.NewInt(0).Rsh(index, 0), n)
	// fmt.Println("a = ", a, mpzN)
	// fmt.Println("mpzU = ", a, b, n)
	if a.Cmp(uh) != 0 {
		fmt.Println("0错误", a, uh, n)
		os.Exit(0)
	}
	if b.Cmp(vl) != 0 {
		fmt.Println("1错误", b, vl, n)
		os.Exit(0)
	}

	//    mpz_clear(zD);
	//    mpz_clear(index);
	//    mpz_clear(uh);
	//    mpz_clear(vl);
	//    mpz_clear(vh);
	//    mpz_clear(ql);
	//    mpz_clear(qh);
	//    mpz_clear(tmp);

	//    if (mpz_cmp_ui(res, 0) == 0)
	//    {
	// 	 mpz_clear(res);
	// 	 return PRP_PRP;
	//    }
	//    else
	//    {
	// 	 mpz_clear(res);
	// 	 return PRP_COMPOSITE;
	//    }
	if res.Cmp(big.NewInt(0)) == 0 {
		return true
	} else {
		return false
	}

} /* method mpz_lucas_prp */

func MpzScan1(n *big.Int) int {
	i := 0
	for {
		if n.Bit(i) == 1 {
			return i
		}
		i++
	}
}

func SoloveyStrassen(num *big.Int) bool {
	if num.Cmp(big.NewInt(1)) == 0 {
		return false
	}
	if num.Cmp(big.NewInt(2)) == 0 {
		return true
	}
	if num.Bit(0) == 0 {
		return false
	}
	// if num.Cmp(big.NewInt(3)) == 0 {
	// 	return true
	// }
	// if big.NewInt(0).Mod(num, big.NewInt(3)).Cmp(big.NewInt(0)) == 0 {
	// 	return false
	// }
	t := big.NewInt(0).Add(num, big.NewInt(-1))
	t.Rsh(t, 1)
	j := Jacobi(big.NewInt(2), num)
	j.Mod(j, num)

	if big.NewInt(0).Exp(big.NewInt(2), t, num).Cmp(j) == 0 {
		return true
	} else {
		return false
	}
}

func Jacobi(a *big.Int, n *big.Int) *big.Int {
	a = new(big.Int).Set(a)
	n = new(big.Int).Set(n)

	//a==1或者n==1
	if a.Cmp(big.NewInt(1)) == 0 || n.Cmp(big.NewInt(1)) == 0 {
		return big.NewInt(1)
	}

	// a=0
	if a.Cmp(big.NewInt(0)) == 0 {
		if n.Cmp(big.NewInt(-1)) == 0 {
			return big.NewInt(1)
		} else {
			return big.NewInt(0)
		}
	}

	// n=0
	if n.Cmp(big.NewInt(0)) == 0 {
		if a.Cmp(big.NewInt(-1)) == 0 {
			return big.NewInt(1)
		} else {
			return big.NewInt(0)
		}
	}

	//n=-1
	if n.Cmp(big.NewInt(-1)) == 0 {
		if a.Cmp(big.NewInt(0)) < 0 {
			return big.NewInt(-1)
		} else {
			return big.NewInt(1)
		}
	}
	//a=-1
	if a.Cmp(big.NewInt(-1)) == 0 {
		isNeg := false
		if n.Cmp(big.NewInt(0)) < 0 {
			n = big.NewInt(0).Neg(n)
			isNeg = true
		}
		n = big.NewInt(0).Set(n)
		for n.Bit(0) == 0 {
			n.Rsh(n, 1)
		}
		n_1 := big.NewInt(0).Sub(n, big.NewInt(1))
		n_1.Rsh(n_1, 1)
		ans := big.NewInt(0).Exp(big.NewInt(-1), n_1, nil)
		if isNeg {
			ans.Neg(ans)
		}
		return ans
	}

	//a和n的绝对值必都大于等于2

	//a<0
	if a.Cmp(big.NewInt(0)) < 0 {
		return big.NewInt(0).Mul(Jacobi(big.NewInt(-1), n), Jacobi(big.NewInt(0).Neg(a), n))
	}

	//n<0，直接取反
	if n.Cmp(big.NewInt(0)) < 0 {
		n.Neg(n)
	}

	//最大公约数
	if new(big.Int).GCD(nil, nil, a, n).Cmp(big.NewInt(1)) > 0 {
		return big.NewInt(0)
	}

	if n.Bit(0) == 0 { //n是偶数
		if n.Cmp(big.NewInt(2)) == 0 {
			am := big.NewInt(0).Mod(a, big.NewInt(8))
			if am.Cmp(big.NewInt(3)) == 0 || am.Cmp(big.NewInt(5)) == 0 {
				return big.NewInt(-1)
			} else {
				return big.NewInt(1)
			}
		}
		n.Rsh(n, 1)
		return big.NewInt(0).Mul(Jacobi(a, big.NewInt(2)), Jacobi(a, n))
	}

	//通例n>=3的奇数
	i := big.NewInt(0)
	s := big.NewInt(0)
	a = new(big.Int).Mod(a, n)
	temp := new(big.Int).Set(a)

	for temp.Bit(0) == 0 {
		temp.Rsh(temp, 1)
		i.Add(i, big.NewInt(1))
	}
	if i.Bit(0) == 0 {
		s.SetInt64(1)
	} else {
		if new(big.Int).Mod(n, big.NewInt(8)).Cmp(big.NewInt(1)) == 0 || new(big.Int).Mod(n, big.NewInt(8)).Cmp(big.NewInt(7)) == 0 {
			s.SetInt64(1)
		} else {
			s.SetInt64(-1)
		}
	}
	if new(big.Int).Mod(n, big.NewInt(4)).Cmp(big.NewInt(3)) == 0 && new(big.Int).Mod(temp, big.NewInt(4)).Cmp(big.NewInt(3)) == 0 {
		s.Neg(s)
	}
	if temp.Cmp(big.NewInt(1)) != 0 {
		n1 := new(big.Int).Mod(n, temp)
		temp1 := new(big.Int).Mul(s, Jacobi(n1, temp))
		return temp1
	} else {
		if i.Bit(0) == 0 {
			return big.NewInt(1)
		} else {
			if new(big.Int).Mod(n, big.NewInt(8)).Cmp(big.NewInt(1)) == 0 || new(big.Int).Mod(n, big.NewInt(8)).Cmp(big.NewInt(7)) == 0 {
				return big.NewInt(1)
			} else {
				return big.NewInt(-1)
			}
		}
	}
}

func Frob(n, p, q *big.Int) bool {
	if n.Cmp(big.NewInt(2)) < 0 {
		return false
	}
	if n.Cmp(big.NewInt(2)) == 0 {
		return true
	}
	if n.Bit(0) == 0 {
		return false
	}
	d := big.NewInt(0)
	d.Mul(p, p)
	d.Sub(d, big.NewInt(0).Mul(q, big.NewInt(4)))

	if d.Cmp(big.NewInt(0)) == 0 {
		panic("Does not produce a proper Lucas sequence")
	}
	t := big.NewInt(0)
	t.Mul(q, d).Lsh(t, 1)
	t.GCD(nil, nil, n, t)
	//1
	if t.Cmp(big.NewInt(1)) != 0 {
		if t.Cmp(n) >= 0 {
			fmt.Println(n, "可能", n.ProbablyPrime(0))
			return n.ProbablyPrime(0)
		}
		return false
	}
	// if n.Cmp(big.NewInt(5)) == 0 {
	// 	fmt.Println(n, t, d)
	// }

	index := big.NewInt(0).Set(n)
	ret := JacobiSymbol(d, n)
	if ret.Cmp(big.NewInt(-1)) == 0 {
		index.Add(index, big.NewInt(1))
	} else if ret.Cmp(big.NewInt(1)) == 0 {
		index.Sub(index, big.NewInt(1))
	}

	ll := LucasParam{p, q}
	u, v := ll.GetUnAndVnMod(big.NewInt(0).Rsh(index, 0), n)

	//2.
	if u.Cmp(big.NewInt(0)) != 0 {
		return false
	}

	//3
	if true {
		vv := big.NewInt(0).Sub(big.NewInt(1), ret)
		vv.Rsh(vv, 1)
		vv.Exp(q, vv, n)
		vv.Lsh(vv, 1)
		vv.Neg(vv)
		vv.Add(vv, v)
		vv.Mod(vv, n)
		if vv.Cmp(big.NewInt(0)) != 0 {
			return false
		}
	}

	//3
	if true {
		u, v = ll.GetUnAndVnMod(n, n)
		vv := big.NewInt(0)
		vv.Sub(v, p).Mod(vv, n)
		if vv.Cmp(big.NewInt(0)) != 0 {
			return false
		}
	}

	return true
}
func Frob2(n *big.Int) bool {

	if n.Cmp(big.NewInt(2)) < 0 {
		return false
	}
	if n.Cmp(big.NewInt(2)) == 0 {
		return true
	}
	if n.Bit(0) == 0 {
		return false
	}
	// fmt.Println(n)
	q := big.NewInt(2)
	p := big.NewInt(3)
	d := big.NewInt(0)
	for {
		d.Mul(p, p)
		d.Sub(d, big.NewInt(0).Mul(q, big.NewInt(4)))
		// fmt.Println(Jacobi(d, n), d, n)
		if d.Cmp(big.NewInt(0)) != 0 && Jacobi(d, n).Cmp(big.NewInt(0)) != 0 {
			// if d.Cmp(big.NewInt(0)) != 0 && Jacobi(d, n).Cmp(big.NewInt(-1)) == 0 {
			break
		}
		p.Add(p, big.NewInt(2))
	}
	// d := big.NewInt(0)
	// d.Mul(p, p)
	// d.Sub(d, big.NewInt(0).Mul(q, big.NewInt(4)))

	// if d.Cmp(big.NewInt(0)) == 0 {
	// 	panic("Does not produce a proper Lucas sequence")
	// }
	t := big.NewInt(0)
	t.Mul(q, d).Lsh(t, 1)
	t.GCD(nil, nil, n, t)
	//1
	if t.Cmp(big.NewInt(1)) != 0 {
		if t.Cmp(n) >= 0 {
			fmt.Println(n, "可能", n.ProbablyPrime(0))
			os.Exit(0)
			return n.ProbablyPrime(0)
		}
		return false
	}
	// if n.Cmp(big.NewInt(5)) == 0 {
	// 	fmt.Println(n, t, d)
	// }

	index := big.NewInt(0).Set(n)
	ret := JacobiSymbol(d, n)
	if ret.Cmp(big.NewInt(-1)) == 0 {
		index.Add(index, big.NewInt(1))
	} else if ret.Cmp(big.NewInt(1)) == 0 {
		index.Sub(index, big.NewInt(1))
	}

	ll := LucasParam{p, q}
	u, v := ll.GetUnAndVnMod(big.NewInt(0).Rsh(index, 0), n)

	//2.
	if u.Cmp(big.NewInt(0)) != 0 {
		return false
	}

	//3
	if true {
		vv := big.NewInt(0).Sub(big.NewInt(1), ret)
		vv.Rsh(vv, 1)
		vv.Exp(q, vv, n)
		vv.Lsh(vv, 1)
		vv.Neg(vv)
		vv.Add(vv, v)
		vv.Mod(vv, n)
		if vv.Cmp(big.NewInt(0)) != 0 {
			return false
		}
	}

	//3
	if true {
		u, v = ll.GetUnAndVnMod(n, n)
		vv := big.NewInt(0)
		vv.Sub(v, p).Mod(vv, n)
		if vv.Cmp(big.NewInt(0)) != 0 {
			return false
		}
	}

	return true
}

func main() {
	if true {
		if true {
			num := big.NewInt(0)
			num.SetString("2", 10)
			count := 0
			rightLimit := big.NewInt(0)
			rightLimit.SetString("10000", 10)
			for ; num.Cmp(rightLimit) <= 0; num.Add(num, big.NewInt(1)) {

				// r := LucasSelfridge(num)
				// r := mpz_selfridge_prp(num)
				// r := Frob(num, big.NewInt(1), big.NewInt(-1))
				r := Frob(num, big.NewInt(3), big.NewInt(-1))
				// r := Frob(num, big.NewInt(3), big.NewInt(-5))
				// r := Frob2(num)
				r2 := num.ProbablyPrime(0)
				// r3 := MillerRabbinA(big.NewInt(2), num)
				// r3 := Fermat(num)
				r3 := SoloveyStrassen(num)
				if r != r3 {
					//r = false
				}

				if r == r2 {
					if r {
						// fmt.Println(num, "是素数")
					}
				} else {
					fmt.Println("测试失败", r, r2, num)
					count++
					// return
				}

			}
			fmt.Println("失败次数", count)
			return
		}
		return
	}
	if false {
		r := big.NewInt(9)
		fmt.Println(r.Bit(4))
		return
	}
	if true {
		r := LucasSelfridge(big.NewInt(105))
		fmt.Println(r)
		return
	}
	if false {
		a := big.NewInt(3)
		n := big.NewInt(97)

		j := JacobiSymbol(a, n)
		fmt.Println(j)
	}
}
