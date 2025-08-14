package main

import (
	"fmt"
	"math/big"
)

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

func ExtraStrongLucas(mpzN, lB *big.Int) bool {
	/* Test N for primality using the extra strong Lucas test with base B,
	   as formulated by Zhaiyu Mo and James P. Jones ("A new primality test
	   using Lucas sequences," preprint, circa 1997), and described by Jon
	   Grantham in "Frobenius pseudoprimes," (preprint, 16 July 1998),
	   available at <http://www.pseudoprime.com/pseudo1.ps>.

	   Returns 1 if N is prime or an extra strong Lucas pseudoprime (base B).
	   Returns 0 if N is definitely composite.

	   Even N and N < 3 are eliminated before applying the Lucas test.

	   In this implementation of the algorithm, Q=1, and B is an integer
	   in 2 < B < INT32_MAX (2147483647 on 32-bit machines); the default value
	   is B=3. B is incremented as necessary if the values of B and N are
	   inconsistent with the hypotheses of Jones and Mo: P=B, Q=1,
	   D=P*P - 4*Q, GCD(N,2D)=1, Jacobi(D,N) <> 0.

	   Since the base B is used solely to calculate the discriminant
	   D=B*B - 4, negative values of B are redundant. The bases B=0 and
	   B=1 are excluded because they produce huge numbers of pseudoprimes,
	   and B=2 is excluded because the resulting D=0 fails the Jones-Mo
	   hypotheses.

	   Note that the choice Q=1 eliminates the computation of powers of Q
	   which appears in the weak and strong Lucas tests.

	   The running time of the extra strong Lucas-Selfridge test is, on
	   average, roughly 80 % that of the standard Lucas-Selfridge test
	   or 2 to 6 times that of a single Miller's test. This is superior
	   in speed to both the standard and strong Lucas-Selfridge tests. The
	   frequency of extra strong Lucas pseudoprimes also appears to be
	   about 80 % that of the strong Lucas-Selfridge test and 30 % that of
	   the standard Lucas-Selfridge test, comparable to the frequency of
	   spsp(2).

	   Unfortunately, the apparent superior peformance of the extra strong
	   Lucas test is offset by the fact that it is not "backwards compatible"
	   with the Lucas-Selfridge tests, due to the differing choice of
	   parameters: P=B and Q=1 in the extra strong test, while P=1 and
	   Q=(1 - D)/4 in the standard and strong Lucas-Selfridge tests (with D
	   chosen from the sequence 5, -7, 9, ...). Thus, although every extra
	   strong Lucas pseudoprime to base B is also both a strong and standard
	   Lucas pseudoprime with parameters P=B and Q=1, the extra strong
	   pseudoprimes do *NOT* constitute a proper subset of the Lucas-Selfridge
	   standard and strong pseudoprimes. As a specific example, 4181 is an
	   extra strong Lucas pseudoprime to base 3, but is neither a standard
	   nor strong Lucas-Selfridge pseudoprime.

	   As a result, the corresponding Baillie-PSW test is fatally flawed.
	   Regardless of the base chosen for the extra strong Lucas test, it
	   appears that there exist numerous N for which the corresponding
	   extra strong Lucas pseudoprimes (xslpsp) will also be strong
	   pseudoprimes to base 2 (or any other particular Miller's base).
	   For example, 6368689 is both spsp(2) and xslpsp(3); 8725753
	   is both spsp(2) and xslpsp(11); 80579735209 is spsp(2) and
	   simultaneously xslpsp for the bases 3, 5, and 7; 105919633 is
	   spsp(3) and xslpsp(11); 1121176981 is spsp(19) and xslpsp(31);
	   and so on. Perhaps some combination of the extra strong test
	   and multiple Miller's tests could match the performance of the
	   Lucas-Selfridge BPSW tests, but the prospects do not look bright.
	*/

	// int iComp2, iJ;
	iComp2 := 0
	iJ := big.NewInt(0)
	// long lD, lP, lQ;
	lD := big.NewInt(0)
	lP := big.NewInt(0)
	lQ := big.NewInt(0)

	// unsigned long ulMaxBits, uldbits, ul, ulGCD, r, s;
	uldbits := 0
	ulGCD := big.NewInt(0)
	r := 0
	s := 0

	// mpz_t mpzU, mpzV, mpzM, mpzU2m, mpzV2m, mpzT1, mpzT2, mpzT3, mpzT4,
	//       mpzD, mpzd, mpzTwo, mpzMinusTwo;
	mpzU := big.NewInt(0)
	mpzV := big.NewInt(0)
	mpzM := big.NewInt(0)
	mpzU2m := big.NewInt(0)
	mpzV2m := big.NewInt(0)
	mpzT1 := big.NewInt(0)
	mpzT2 := big.NewInt(0)
	mpzT3 := big.NewInt(0)
	mpzT4 := big.NewInt(0)
	mpzD := big.NewInt(0)
	mpzd := big.NewInt(0)
	mpzTwo := big.NewInt(0)
	mpzMinusTwo := big.NewInt(0)
	// #undef RETURN
	// #define RETURN(n)           \
	//   {                         \
	//   mpz_clear(mpzU);          \
	//   mpz_clear(mpzV);          \
	//   mpz_clear(mpzM);          \
	//   mpz_clear(mpzU2m);        \
	//   mpz_clear(mpzV2m);        \
	//   mpz_clear(mpzT1);         \
	//   mpz_clear(mpzT2);         \
	//   mpz_clear(mpzT3);         \
	//   mpz_clear(mpzT4);         \
	//   mpz_clear(mpzD);          \
	//   mpz_clear(mpzd);          \
	//   mpz_clear(mpzTwo);        \
	//   mpz_clear(mpzMinusTwo);   \
	//   return(n);                \
	//   }

	/* This implementation of the algorithm assumes N is an odd integer > 2,
	   so we first eliminate all N < 3 and all even N. */

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
	if mpzN.Bit(0) == 0 {
		return false
	}
	// if(mpz_even_p(mpzN))return(0);
	sq := big.NewInt(0).Sqrt(mpzN)
	sq.Mul(sq, sq)
	if sq.Cmp(mpzN) == 0 {
		return false
	}

	/* Allocate storage for the mpz_t variables. Most require twice
	   the storage of mpzN, since multiplications of order O(mpzN)*O(mpzN)
	   will be performed. */

	// ulMaxBits=2*mpz_sizeinbase(mpzN, 2) + mp_bits_per_limb;
	// mpz_init2(mpzU, ulMaxBits);
	// mpz_init2(mpzV, ulMaxBits);
	// mpz_init2(mpzM, ulMaxBits);
	// mpz_init2(mpzU2m, ulMaxBits);
	// mpz_init2(mpzV2m, ulMaxBits);
	// mpz_init2(mpzT1, ulMaxBits);
	// mpz_init2(mpzT2, ulMaxBits);
	// mpz_init2(mpzT3, ulMaxBits);
	// mpz_init2(mpzT4, ulMaxBits);
	// mpz_init(mpzD);
	// mpz_init2(mpzd, ulMaxBits);
	// mpz_init_set_si(mpzTwo, 2);
	mpzTwo.SetInt64(2)
	// mpz_init_set_si(mpzMinusTwo, -2);
	mpzMinusTwo.SetInt64(-2)

	/* The parameters specified by Zhaiyu Mo and James P. Jones,
	   as set forth in Grantham's paper, are P=B, Q=1, D=P*P - 4*Q,
	   with (N,2D)=1 so that Jacobi(D,N) <> 0. As explained above,
	   bases B < 3 are excluded. */

	if lB.Cmp(big.NewInt(3)) < 0 {
		//   lP=3;
		lP.Set(big.NewInt(3))
	} else {
		//   lP=lB;
		lP.Set(lB)
	}
	// lQ=1;
	lQ.SetInt64(1)

	/* We check to make sure that N and D are relatively prime. If not,
	   then either 1 < (D,N) < N, in which case N is composite with
	   divisor (D,N); or N = (D,N), in which case N divides D and may be
	   either prime or composite, so we increment the base B=P and
	   try again. */

	for {
		//   lD=lP*lP - 4*lQ;
		lD.Mul(lP, lP).Sub(lD, big.NewInt(0).Mul(big.NewInt(4), lQ))
		//   ulGCD=mpz_gcd_ui(NULL, mpzN, labs(lD));
		ulGCD.GCD(nil, nil, mpzN, big.NewInt(0).Abs(lD))
		//   if(ulGCD==1)break;
		if ulGCD.Cmp(big.NewInt(1)) == 0 {
			break
		}
		//   if(mpz_cmp_ui(mpzN, ulGCD) > 0)RETURN(0);
		if mpzN.Cmp(ulGCD) > 0 {
			return false
		}
		//   lP++;
		lP.Add(lP, big.NewInt(1))
	}

	/* Now calculate M = N - Jacobi(D,N) (M even), and calculate the
	   odd positive integer d and positive integer s for which
	   M = 2^s*d (similar to the step for N - 1 in Miller's
	   test). The extra strong Lucas-Selfridge test then returns N as
	   an extra strong Lucas probable prime (eslprp) if any of the
	   following conditions is met: U_d=0 and V_d瘃2; or V_d=0; or
	   V_2d=0, V_4d=0, V_8d=0, V_16d=0, ..., etc., ending with
	   V_{2^(s-2)*d}=V_{M/4}? (all equalities mod N). Thus d is the
	   highest index of U that must be computed (since V_2m is
	   independent of U), compared to U_M for the standard Lucas
	   test; and no index of V beyond M/4 is required, compared to
	   M/2 for the standard and strong Lucas tests. Furthermore,
	   since Q=1, the powers of Q required in the standard and
	   strong Lucas tests can be dispensed with. The result is that
	   the extra strong Lucas test has a running time shorter than
	   that of either the standard or strong Lucas-Selfridge tests
	   (roughly two to six times that of a single Miller's test).
	   The extra strong test also produces fewer pseudoprimes.
	   Unfortunately, the pseudoprimes produced are *NOT* a subset
	   of the standard or strong Lucas-Selfridge pseudoprimes (due
	   to the incompatible parameters P and Q), and consequently the
	   extra strong test does not combine with a single Miller's test
	   to produce a Baillie-PSW test of the reliability level of the
	   BPSW tests based on the standard or strong Lucas-Selfridge tests. */

	// mpz_set_si(mpzD, lD);
	mpzD.Set(lD)
	// iJ=mpz_jacobi(mpzD, mpzN);
	iJ = JacobiSymbol(mpzD, mpzN)
	fmt.Println(lP, lQ, mpzD, mpzN, iJ)
	// assert(iJ != 0);
	if iJ.Cmp(big.NewInt(1)) == 0 {
		//   mpz_sub_ui(mpzM, mpzN, 1);
		mpzM.Sub(mpzN, big.NewInt(1))
	} else {
		//   mpz_add_ui(mpzM, mpzN, 1);
		mpzM.Add(mpzN, big.NewInt(1))
	}

	// s=mpz_scan1(mpzM, 0);
	// mpz_tdiv_q_2exp(mpzd, mpzM, s);
	for s = 0; mpzM.Bit(s) == 0; s++ {
	}
	mpzd.Rsh(mpzM, uint(s))

	/* We must now compute U_d and V_d. Since d is odd, the accumulated
	   values U and V are initialized to U_1 and V_1 (if the target
	   index were even, U and V would be initialized instead to U_0=0
	   and V_0=2). The values of U_2m and V_2m are also initialized to
	   U_1 and V_1; the FOR loop calculates in succession U_2 and V_2,
	   U_4 and V_4, U_8 and V_8, etc. If the corresponding bits
	   (1, 2, 3, ...) of t are on (the zero bit having been accounted
	   for in the initialization of U and V), these values are then
	   combined with the previous totals for U and V, using the
	   composition formulas for addition of indices. */

	// mpz_set_ui(mpzU, 1);                       /* U=U_1 */
	mpzU.SetInt64(1)
	// mpz_set_si(mpzV, lP);                      /* V=V_1 */
	mpzV.Set(lP)
	// mpz_set_ui(mpzU2m, 1);                     /* U_1 */
	mpzU2m.SetInt64(1)
	// mpz_set_si(mpzV2m, lP);                    /* V_1 */
	mpzV2m.Set(lP)

	// uldbits=mpz_sizeinbase(mpzd, 2);
	uldbits = mpzd.BitLen()
	for ul := 1; ul < uldbits; ul++ { /* zero bit on, already accounted for */

		/* Formulas for doubling of indices (carried out mod N). Note that
		 * the indices denoted as "2m" are actually powers of 2, specifically
		 * 2^(ul-1) beginning each loop and 2^ul ending each loop.
		 *
		 * U_2m = U_m*V_m
		 * V_2m = V_m*V_m - 2*Q^m
		 */
		//   mpz_mul(mpzU2m, mpzU2m, mpzV2m);
		mpzU2m.Mul(mpzU2m, mpzV2m)
		//   mpz_mod(mpzU2m, mpzU2m, mpzN);
		mpzU2m.Mod(mpzU2m, mpzN)
		//   mpz_mul(mpzV2m, mpzV2m, mpzV2m);
		mpzV2m.Mul(mpzV2m, mpzV2m)
		//   mpz_sub_ui(mpzV2m, mpzV2m, 2);
		mpzV2m.Sub(mpzV2m, big.NewInt(2))
		//   mpz_mod(mpzV2m, mpzV2m, mpzN);
		mpzV2m.Mod(mpzV2m, mpzN)
		if mpzd.Bit(ul) == 1 {
			//   if(mpz_tstbit(mpzd, ul))
			//     {
			/* Formulas for addition of indices (carried out mod N);
			 *
			 * U_(m+n) = (U_m*V_n + U_n*V_m)/2
			 * V_(m+n) = (V_m*V_n + D*U_m*U_n)/2
			 *
			 * Be careful with division by 2 (mod N)!
			 */
			// mpz_mul(mpzT1, mpzU2m, mpzV);
			mpzT1.Mul(mpzU2m, mpzV)
			// mpz_mul(mpzT2, mpzU, mpzV2m);
			mpzT2.Mul(mpzU, mpzV2m)
			// mpz_mul(mpzT3, mpzV2m, mpzV);
			mpzT3.Mul(mpzV2m, mpzV)
			// mpz_mul(mpzT4, mpzU2m, mpzU);
			mpzT4.Mul(mpzU2m, mpzU)
			// mpz_mul_si(mpzT4, mpzT4, lD);
			mpzT4.Mul(mpzT4, lD)
			// mpz_add(mpzU, mpzT1, mpzT2);
			mpzU.Add(mpzT1, mpzT2)
			// if(mpz_odd_p(mpzU))mpz_add(mpzU, mpzU, mpzN);
			if mpzU.Bit(0) == 1 {
				mpzU.Add(mpzU, mpzN)
			}
			// mpz_fdiv_q_2exp(mpzU, mpzU, 1);
			mpzU.Rsh(mpzU, 1)
			// mpz_add(mpzV, mpzT3, mpzT4);
			mpzV.Add(mpzT3, mpzT4)
			// if(mpz_odd_p(mpzV))mpz_add(mpzV, mpzV, mpzN);
			if mpzV.Bit(0) == 1 {
				mpzV.Add(mpzV, mpzN)
			}
			// mpz_fdiv_q_2exp(mpzV, mpzV, 1);
			mpzV.Rsh(mpzV, 1)
			// mpz_mod(mpzU, mpzU, mpzN);
			mpzU.Mod(mpzU, mpzN)
			// mpz_mod(mpzV, mpzV, mpzN);
			mpzV.Mod(mpzV, mpzN)
		}
	}
	if true {
		mpzU.Mod(mpzU, mpzN)
		mpzV.Mod(mpzV, mpzN)
	}

	ll := LucasParam{lP, lQ}
	a, b := ll.GetUnAndVnMod(big.NewInt(0).Rsh(mpzd, 0), mpzN)
	// fmt.Println("a = ", a, mpzN)
	// fmt.Println("mpzU = ", mpzU, mpzN)
	if a.Cmp(mpzU) != 0 {
		fmt.Println("0错误", a, mpzU)
		//os.Exit(0)
	}
	if b.Cmp(mpzV) != 0 {
		fmt.Println("1错误", b, mpzV, mpzN)
		//os.Exit(0)
	}
	// fmt.Println(mpzU, mpzV, mpzN)
	/* N first passes the extra strong Lucas test if V_d?, or if V_d瘃2
	   and U_d?.  U and V are tested for divisibility by N, rather than
	   zero, in case the previous FOR is a zero-iteration loop.*/

	// if(mpz_divisible_p(mpzV, mpzN))RETURN(1);

	if big.NewInt(0).Mod(mpzV, mpzN).Cmp(big.NewInt(0)) == 0 {
		return true
	}

	// if(mpz_divisible_p(mpzU, mpzN))
	if big.NewInt(0).Mod(mpzU, mpzN).Cmp(big.NewInt(0)) == 0 {
		//   if(mpz_congruent_p(mpzV, mpzTwo, mpzN))RETURN(1);
		if big.NewInt(0).Mod(mpzV, mpzN).Cmp(big.NewInt(0).Mod(mpzTwo, mpzN)) == 0 {
			return true
		}
		//   if(mpz_congruent_p(mpzV, mpzMinusTwo, mpzN))RETURN(1);
		if big.NewInt(0).Mod(mpzV, mpzN).Cmp(big.NewInt(0).Mod(mpzMinusTwo, mpzN)) == 0 {
			return true
		}
	}

	/* Otherwise, we must compute V_2d, V_4d, V_8d, ..., V_{2^(s-2)*d}
	   by repeated use of the formula V_2m = V_m*V_m - 2*Q^m. If any of
	   these are congruent to 0 mod N, then N is a prime or an extra
	   strong Lucas pseudoprime. */

	for r = 1; r < s-1; r++ {
		//   mpz_mul(mpzV, mpzV, mpzV);
		mpzV.Mul(mpzV, mpzV)
		//   mpz_sub_ui(mpzV, mpzV, 2);
		mpzV.Sub(mpzV, big.NewInt(2))
		//   mpz_mod(mpzV, mpzV, mpzN);
		mpzV.Mod(mpzV, mpzN)
		//   if(mpz_sgn(mpzV)==0)RETURN(1);
		if mpzV.Cmp(big.NewInt(0)) == 0 {
			return true
		}
	}

	/* Otherwise N is definitely composite. */

	return false
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

/* *******************************************************************************************
 * mpz_extrastronglucas_prp:
 * Let U_n = LucasU(p,1), V_n = LucasV(p,1), and D=p^2-4.
 * An "extra strong Lucas pseudoprime" to the base p is a composite n = (2^r)*s+(D/n), where
 * s is odd and (n,2D)=1, such that either U_s == 0 mod n and V_s == +/-2 mod n, or
 * V_((2^t)*s) == 0 mod n for some t with 0 <= t < r-1 [(D/n) is the Jacobi symbol]
 * *******************************************************************************************/
func mpz_extrastronglucas_prp(n, p *big.Int) bool {
	//    mpz_t zD;
	zD := big.NewInt(0)
	//    mpz_t s;
	s := big.NewInt(0)
	//    mpz_t nmj; /* n minus jacobi(D/n) */
	nmj := big.NewInt(0)
	//    mpz_t res;
	res := big.NewInt(0)
	//    mpz_t uh, vl, vh, ql, qh, tmp; /* these are needed for the LucasU and LucasV part of this function */
	uh := big.NewInt(0)
	vl := big.NewInt(0)
	vh := big.NewInt(0)
	ql := big.NewInt(0)
	qh := big.NewInt(0)
	tmp := big.NewInt(0)
	//    long int d = p*p - 4;
	d := big.NewInt(0)
	d.Mul(p, p)
	d.Sub(d, big.NewInt(4))
	//    long int q = 1;
	q := big.NewInt(1)
	//    unsigned long int r = 0;
	r := 0
	//    int ret = 0;
	ret := big.NewInt(0)
	//    unsigned long int j = 0;
	j := 0

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

	//    mpz_init_set_si(zD, d);
	zD.Set(d)
	//    mpz_init(res);

	//    mpz_mul_ui(res, zD, 2);
	res.Mul(zD, big.NewInt(2))
	//    mpz_gcd(res, res, n);
	res.GCD(nil, nil, res, n)
	//    if ((mpz_cmp(res, n) != 0) && (mpz_cmp_ui(res, 1) > 0))
	//    {
	// 	 mpz_clear(zD);
	// 	 mpz_clear(res);
	// 	 return PRP_COMPOSITE;
	//    }
	if res.Cmp(n) != 0 && res.Cmp(big.NewInt(1)) > 0 {
		return false
	}

	//    mpz_init(s);
	//    mpz_init(nmj);

	//    /* nmj = n - (D/n), where (D/n) is the Jacobi symbol */
	//    mpz_set(nmj, n);
	nmj.Set(n)
	//    ret = mpz_jacobi(zD, n);
	ret = Jacobi(zD, n)
	//    if (ret == -1)
	// 	 mpz_add_ui(nmj, nmj, 1);
	//    else if (ret == 1)
	// 	 mpz_sub_ui(nmj, nmj, 1);
	if ret.Cmp(big.NewInt(-1)) == 0 {
		nmj.Add(nmj, big.NewInt(1))
	} else if ret.Cmp(big.NewInt(1)) == 0 {
		nmj.Sub(nmj, big.NewInt(1))
	}

	//    r = mpz_scan1(nmj, 0);
	r = MpzScan1(nmj)
	//    mpz_fdiv_q_2exp(s, nmj, r);
	s.Rsh(nmj, uint(r))

	//    /* make sure that either (U_s == 0 mod n and V_s == +/-2 mod n), or */
	//    /* V_((2^t)*s) == 0 mod n for some t with 0 <= t < r-1           */
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

	//    for (j = mpz_sizeinbase(s,2)-1; j >= 1; j--)
	for j = s.BitLen() - 1; j >= 1; j-- {
		// 	 /* ql = ql*qh (mod n) */
		// 	 mpz_mul(ql, ql, qh);
		ql.Mul(ql, qh)
		// 	 mpz_mod(ql, ql, n);
		ql.Mod(ql, n)
		// 	 if (mpz_tstbit(s,j) == 1)
		if s.Bit(j) == 1 {
			// 	   /* qh = ql*q */
			// 	   mpz_mul_si(qh, ql, q);
			qh.Mul(ql, q)

			// 	   /* uh = uh*vh (mod n) */
			// 	   mpz_mul(uh, uh, vh);
			uh.Mul(uh, vh)
			// 	   mpz_mod(uh, uh, n);
			uh.Mod(uh, n)

			// 	   /* vl = vh*vl - p*ql (mod n) */
			// 	   mpz_mul(vl, vh, vl);
			vl.Mul(vh, vl)
			// 	   mpz_mul_si(tmp, ql, p);
			tmp.Mul(ql, p)
			// 	   mpz_sub(vl, vl, tmp);
			vl.Sub(vl, tmp)
			// 	   mpz_mod(vl, vl, n);
			vl.Mod(vl, n)

			// 	   /* vh = vh*vh - 2*qh (mod n) */
			// 	   mpz_mul(vh, vh, vh);
			vh.Mul(vh, vh)
			// 	   mpz_mul_si(tmp, qh, 2);
			tmp.Mul(qh, big.NewInt(2))
			// 	   mpz_sub(vh, vh, tmp);
			vh.Sub(vh, tmp)
			// 	   mpz_mod(vh, vh, n);
			vh.Mod(vh, n)
		} else {
			// 	   /* qh = ql */
			// 	   mpz_set(qh, ql);
			qh.Set(ql)

			// 	   /* uh = uh*vl - ql (mod n) */
			// 	   mpz_mul(uh, uh, vl);
			uh.Mul(uh, vl)
			// 	   mpz_sub(uh, uh, ql);
			uh.Sub(uh, ql)
			// 	   mpz_mod(uh, uh, n);
			uh.Mod(uh, n)

			// 	   /* vh = vh*vl - p*ql (mod n) */
			// 	   mpz_mul(vh, vh, vl);
			vh.Mul(vh, vl)
			// 	   mpz_mul_si(tmp, ql, p);
			tmp.Mul(ql, p)
			// 	   mpz_sub(vh, vh, tmp);
			vh.Sub(vh, tmp)
			// 	   mpz_mod(vh, vh, n);
			vh.Mod(vh, n)

			// 	   /* vl = vl*vl - 2*ql (mod n) */
			// 	   mpz_mul(vl, vl, vl);
			vl.Mul(vl, vl)
			// 	   mpz_mul_si(tmp, ql, 2);
			tmp.Mul(ql, big.NewInt(2))
			// 	   mpz_sub(vl, vl, tmp);
			vl.Sub(vl, tmp)
			// 	   mpz_mod(vl, vl, n);
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

	//    mpz_mod(uh, uh, n);
	uh.Mod(uh, n)
	//    mpz_mod(vl, vl, n);
	vl.Mod(vl, n)

	ll := LucasParam{p, q}
	a, b := ll.GetUnAndVnMod(s, n)
	if a.Cmp(uh) != 0 {
		fmt.Println("0错误", a, uh)
		//os.Exit(0)
	}
	if b.Cmp(vl) != 0 {
		fmt.Println("1错误", b, vl)
		//os.Exit(0)
	}
	// fmt.Println(a, b, n)

	//    /* tmp = n-2, for the following comparison */
	//    mpz_sub_ui(tmp, n, 2);
	tmp.Sub(n, big.NewInt(2))

	// //mpz_aprcl.c的代码里没有这段代码
	// if vl.Cmp(big.NewInt(0)) == 0 {
	// 	return true
	// }

	//    /* uh contains LucasU_s and vl contains LucasV_s */
	//    if ((mpz_cmp_ui(uh, 0) == 0) && ((mpz_cmp(vl, tmp) == 0) || (mpz_cmp_si(vl, 2) == 0)))
	if uh.Cmp(big.NewInt(0)) == 0 && (vl.Cmp(tmp) == 0 || vl.Cmp(big.NewInt(2)) == 0) {
		// 	 mpz_clear(zD);
		// 	 mpz_clear(s);
		// 	 mpz_clear(nmj);
		// 	 mpz_clear(res);
		// 	 mpz_clear(uh);
		// 	 mpz_clear(vl);
		// 	 mpz_clear(vh);
		// 	 mpz_clear(ql);
		// 	 mpz_clear(qh);
		// 	 mpz_clear(tmp);
		// 	 return PRP_PRP;
		return true
	}

	// if vl.Cmp(big.NewInt(0)) == 0 {
	// 	return true
	// }
	if vl.Cmp(big.NewInt(0)) == 0 {
		return true
	}
	// if uh.Cmp(big.NewInt(0)) == 0 {
	// 	return true
	// }
	if uh.Cmp(big.NewInt(0)) == 0 {
		return true
	}
	//    for (j = 1; j < r-1; j++)
	for j = 1; j < r-1; j++ {
		// 	 /* vl = vl*vl - 2*ql (mod n) */
		// 	 mpz_mul(vl, vl, vl);
		vl.Mul(vl, vl)
		// 	 mpz_mul_si(tmp, ql, 2);
		tmp.Mul(ql, big.NewInt(2))
		// 	 mpz_sub(vl, vl, tmp);
		vl.Sub(vl, tmp)
		// 	 mpz_mod(vl, vl, n);
		vl.Mod(vl, n)

		// 	 /* ql = ql*ql (mod n) */
		// 	 mpz_mul(ql, ql, ql);
		ql.Mul(ql, ql)
		// 	 mpz_mod(ql, ql, n);
		ql.Mod(ql, n)

		// 	 if (mpz_cmp_ui(vl, 0) == 0)
		if vl.Cmp(big.NewInt(0)) == 0 {
			// 	   mpz_clear(zD);
			// 	   mpz_clear(s);
			// 	   mpz_clear(nmj);
			// 	   mpz_clear(res);
			// 	   mpz_clear(uh);
			// 	   mpz_clear(vl);
			// 	   mpz_clear(vh);
			// 	   mpz_clear(ql);
			// 	   mpz_clear(qh);
			// 	   mpz_clear(tmp);
			// 	   return PRP_PRP;
			return true
		}
	}

	//    mpz_clear(zD);
	//    mpz_clear(s);
	//    mpz_clear(nmj);
	//    mpz_clear(res);
	//    mpz_clear(uh);
	//    mpz_clear(vl);
	//    mpz_clear(vh);
	//    mpz_clear(ql);
	//    mpz_clear(qh);
	//    mpz_clear(tmp);
	//    return PRP_COMPOSITE;
	return false

} /* method mpz_extrastronglucas_prp */

func MpzScan1(n *big.Int) int {
	i := 0
	for {
		if n.Bit(i) == 1 {
			return i
		}
		i++
	}
}

func main() {
	if true {
		if true {

			num := big.NewInt(0)
			num.SetString("2", 10)
			count := 0
			rightLimit := big.NewInt(0)
			rightLimit.SetString("100", 10)
			for ; num.Cmp(rightLimit) <= 0; num.Add(num, big.NewInt(1)) {

				r := ExtraStrongLucas(num, big.NewInt(3))
				// r := mpz_extrastronglucas_prp(num, big.NewInt(3))
				r2 := num.ProbablyPrime(0)
				// r3 := MillerRabbinA(big.NewInt(2), num)
				// r3 := Fermat(num)

				// r3 := SoloveyStrassen(num)
				// if r != r3 {
				// 	//r = false
				// }

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
			fmt.Println("结束")
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
		r := ExtraStrongLucas(big.NewInt(105), big.NewInt(3))
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
