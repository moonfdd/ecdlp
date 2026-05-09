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

// https://en.wikipedia.org/wiki/Lucas_pseudoprime#Fibonacci_pseudoprimes

// var ulDmax = new(big.Int) /* tracks global max of Lucas-Selfridge |D| */

func StrongLucasSelfridge(mpzN *big.Int) bool {
	/* Test N for primality using the strong Lucas test with Selfridge's
	   parameters. Returns 1 if N is prime or a strong Lucas-Selfridge
	   pseudoprime (in which case N is also a pseudoprime to the standard
	   Lucas-Selfridge test). Returns 0 if N is definitely composite.

	   The running time of the strong Lucas-Selfridge test is, on average,
	   roughly 10 % greater than the running time for the standard
	   Lucas-Selfridge test (3 to 7 times that of a single Miller's test).
	   However, the frequency of strong Lucas pseudoprimes appears to be
	   only (roughly) 30 % that of (standard) Lucas pseudoprimes, and only
	   slightly greater than the frequency of base-2 strong pseudoprimes,
	   indicating that the strong Lucas-Selfridge test is more computationally
	   effective than the standard version. */

	iComp2 := 0
	iP := big.NewInt(0)
	iJ := big.NewInt(0)
	iSign := big.NewInt(0)

	lDabs := big.NewInt(0)
	lD := big.NewInt(0)
	lQ := big.NewInt(0)
	// unsigned long ulMaxBits, uldbits, ul, ulGCD, r, s;
	uldbits := 0
	ulGCD := big.NewInt(0)
	r := 0
	s := 0

	// mpz_t mpzU, mpzV, mpzNplus1, mpzU2m, mpzV2m, mpzQm, mpz2Qm,
	//       mpzT1, mpzT2, mpzT3, mpzT4, mpzD, mpzd, mpzQkd, mpz2Qkd;
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
	mpzd := big.NewInt(0)
	mpzQkd := big.NewInt(0)
	mpz2Qkd := big.NewInt(0)

	// #undef RETURN
	// #define RETURN(n)           \
	//   {                         \
	//   mpz_clear(mpzU);          \
	//   mpz_clear(mpzV);          \
	//   mpz_clear(mpzNplus1);     \
	//   mpz_clear(mpzU2m);        \
	//   mpz_clear(mpzV2m);        \
	//   mpz_clear(mpzQm);         \
	//   mpz_clear(mpz2Qm);        \
	//   mpz_clear(mpzT1);         \
	//   mpz_clear(mpzT2);         \
	//   mpz_clear(mpzT3);         \
	//   mpz_clear(mpzT4);         \
	//   mpz_clear(mpzD);          \
	//   mpz_clear(mpzd);          \
	//   mpz_clear(mpzQkd);        \
	//   mpz_clear(mpz2Qkd);       \
	//   return(n);                \
	//   }

	/* This implementation of the algorithm assumes N is an odd integer > 2,
	   so we first eliminate all N < 3 and all even N. As a practical matter,
	   we also need to filter out all perfect square values of N, such as
	   1093^2 (a base-2 strong pseudoprime); this is because we will later
	   require an integer D for which Jacobi(D,N) = -1, and no such integer
	   exists if N is a perfect square. The algorithm as written would
	   still eventually return zero in this case, but would require
	   nearly sqrt(N)/2 iterations. */

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

	/* Allocate storage for the mpz_t variables. Most require twice
	   the storage of mpzN, since multiplications of order O(mpzN)*O(mpzN)
	   will be performed. */

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
	// mpz_init2(mpzd, ulMaxBits);
	// mpz_init2(mpzQkd, ulMaxBits);
	// mpz_init2(mpz2Qkd, ulMaxBits);

	/* Find the first element D in the sequence {5, -7, 9, -11, 13, ...}
	   such that Jacobi(D,N) = -1 (Selfridge's algorithm). Theory
	   indicates that, if N is not a perfect square, D will "nearly
	   always" be "small." Just in case, an overflow trap for D is
	   included. */

	// lDabs=5;
	lDabs.Set(big.NewInt(5))
	// iSign=1;
	iSign.Set(big.NewInt(1))
	for {
		//   lD=iSign*lDabs;
		lD.Mul(iSign, lDabs)
		//   iSign = -iSign;
		iSign.Neg(iSign)
		//   ulGCD=mpz_gcd_ui(NULL, mpzN, lDabs);
		ulGCD.GCD(nil, nil, mpzN, lDabs)
		/* if 1 < GCD < N then N is composite with factor lDabs, and
		   Jacobi(D,N) is technically undefined (but often returned
		   as zero). */
		//   if((ulGCD > 1) && mpz_cmp_ui(mpzN, ulGCD) > 0)RETURN(0);
		if ulGCD.Cmp(big.NewInt(1)) > 0 && mpzN.Cmp(ulGCD) > 0 {
			return false
		}
		//   mpz_set_si(mpzD, lD);
		mpzD.Add(lD, big.NewInt(0))
		//   iJ=mpz_jacobi(mpzD, mpzN);
		iJ = JacobiSymbol(mpzD, mpzN)
		//   if(iJ==-1)break;
		if iJ.Cmp(big.NewInt(-1)) == 0 {
			break
		}
		//   lDabs += 2;
		lDabs.Add(lDabs, big.NewInt(2))
		//   if(lDabs > ulDmax)ulDmax=lDabs;  /* tracks global max of |D| */
		//   if(lDabs > INT32_MAX-2)
		//     {
		//     fprintf(stderr,
		//       "\n ERROR: D overflows signed long in Lucas-Selfridge test.");
		//     fprintf(stderr, "\n N=");
		//     mpz_out_str(stderr, 10, mpzN);
		//     fprintf(stderr, "\n |D|=%ld\n\n", lDabs);
		//     exit(EXIT_FAILURE);
		//     }
	}

	// iP=1;         /* Selfridge's choice */
	iP.Set(big.NewInt(1))
	// lQ=(1-lD)/4;  /* Required so D = P*P - 4*Q */
	lQ.Sub(big.NewInt(1), lD)
	lQ.Rsh(lQ, 2)

	/* NOTE: The conditions (a) N does not divide Q, and
	   (b) D is square-free or not a perfect square, are included by
	   some authors; e.g., "Prime numbers and computer methods for
	   factorization," Hans Riesel (2nd ed., 1994, Birkhauser, Boston),
	   p. 130. For this particular application of Lucas sequences,
	   these conditions were found to be immaterial. */

	/* Now calculate N - Jacobi(D,N) = N + 1 (even), and calculate the
	   odd positive integer d and positive integer s for which
	   N + 1 = 2^s*d (similar to the step for N - 1 in Miller's test).
	   The strong Lucas-Selfridge test then returns N as a strong
	   Lucas probable prime (slprp) if any of the following
	   conditions is met: U_d=0, V_d=0, V_2d=0, V_4d=0, V_8d=0,
	   V_16d=0, ..., etc., ending with V_{2^(s-1)*d}=V_{(N+1)/2}=0
	   (all equalities mod N). Thus d is the highest index of U that
	   must be computed (since V_2m is independent of U), compared
	   to U_{N+1} for the standard Lucas-Selfridge test; and no
	   index of V beyond (N+1)/2 is required, just as in the
	   standard Lucas-Selfridge test. However, the quantity Q^d must
	   be computed for use (if necessary) in the latter stages of
	   the test. The result is that the strong Lucas-Selfridge test
	   has a running time only slightly greater (order of 10 %) than
	   that of the standard Lucas-Selfridge test, while producing
	   only (roughly) 30 % as many pseudoprimes (and every strong
	   Lucas pseudoprime is also a standard Lucas pseudoprime). Thus
	   the evidence indicates that the strong Lucas-Selfridge test is
	   more effective than the standard Lucas-Selfridge test, and a
	   Baillie-PSW test based on the strong Lucas-Selfridge test
	   should be more reliable. */

	// mpz_add_ui(mpzNplus1, mpzN, 1);
	mpzNplus1.Add(mpzN, big.NewInt(1))

	// s = mpz_scan1(mpzNplus1, 0)
	// mpz_tdiv_q_2exp(mpzd, mpzNplus1, s)
	for s = 0; mpzNplus1.Bit(s) == 0; s++ {
	}
	mpzd.Rsh(mpzNplus1, uint(s))
	// fmt.Println(mpzNplus1.Text(2), mpzd.Text(2))
	// os.Exit(0)

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

	// mpz_set_ui(mpzU, 1);                      /* U=U_1 */
	mpzU.SetInt64(1)
	// mpz_set_ui(mpzV, iP);                     /* V=V_1 */
	mpzV.Set(iP)
	// mpz_set_ui(mpzU2m, 1);                    /* U_1 */
	mpzU2m.SetInt64(1)
	// mpz_set_si(mpzV2m, iP);                   /* V_1 */
	mpzV2m.Set(iP)
	// mpz_set_si(mpzQm, lQ);
	mpzQm.Set(lQ)
	// mpz_set_si(mpz2Qm, 2*lQ);
	mpz2Qm.Add(lQ, big.NewInt(0))
	mpz2Qm.Lsh(mpz2Qm, 1)
	// mpz_set_si(mpzQkd, lQ);  /* Initializes calculation of Q^d */
	mpzQkd.Set(lQ)

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
		//   mpz_sub(mpzV2m, mpzV2m, mpz2Qm);
		mpzV2m.Sub(mpzV2m, mpz2Qm)
		//   mpz_mod(mpzV2m, mpzV2m, mpzN);
		mpzV2m.Mod(mpzV2m, mpzN)
		/* Must calculate powers of Q for use in V_2m, also for Q^d later */
		//   mpz_mul(mpzQm, mpzQm, mpzQm);
		mpzQm.Mul(mpzQm, mpzQm)
		//   mpz_mod(mpzQm, mpzQm, mpzN);  /* prevents overflow */
		mpzQm.Mod(mpzQm, mpzN)
		//   mpz_mul_2exp(mpz2Qm, mpzQm, 1);
		mpz2Qm.Lsh(mpzQm, 1)
		if mpzd.Bit(ul) == 1 {
			//   if(mpz_tstbit(mpzd, ul))
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
			// mpz_mul(mpzQkd, mpzQkd, mpzQm);  /* Calculating Q^d for later use */
			mpzQkd.Mul(mpzQkd, mpzQm)
			// mpz_mod(mpzQkd, mpzQkd, mpzN);
			mpzQkd.Mod(mpzQkd, mpzN)
		}
	}

	ll := LucasParam{iP, lQ}
	a, b := ll.GetUnAndVnMod(mpzd, mpzN)
	if a.Cmp(mpzU) != 0 {
		fmt.Println("0错误", a, mpzU)
		//os.Exit(0)
	}
	if b.Cmp(mpzV) != 0 {
		fmt.Println("1错误", b, mpzV)
		//os.Exit(0)
	}

	/* If U_d or V_d is congruent to 0 mod N, then N is a prime or a
	   strong Lucas pseudoprime. */

	// if(mpz_sgn(mpzU)==0)RETURN(1);
	if mpzU.Cmp(big.NewInt(0)) == 0 {
		return true
	}
	// if(mpz_sgn(mpzV)==0)RETURN(1);
	if mpzV.Cmp(big.NewInt(0)) == 0 {
		return true
	}

	/* NOTE: Ribenboim ("The new book of prime number records," 3rd ed.,
	   1995/6) omits the condition V? on p.142, but includes it on
	   p. 130. The condition is NECESSARY; otherwise the test will
	   return false negatives---e.g., the primes 29 and 2000029 will be
	   returned as composite. */

	/* Otherwise, we must compute V_2d, V_4d, V_8d, ..., V_{2^(s-1)*d}
	   by repeated use of the formula V_2m = V_m*V_m - 2*Q^m. If any of
	   these are congruent to 0 mod N, then N is a prime or a strong
	   Lucas pseudoprime. */

	// mpz_mul_2exp(mpz2Qkd, mpzQkd, 1);  /* Initialize 2*Q^(d*2^r) for V_2m */
	mpz2Qkd.Lsh(mpzQkd, 1)

	for r = 1; r < s; r++ {
		//   mpz_mul(mpzV, mpzV, mpzV);
		mpzV.Mul(mpzV, mpzV)
		//   mpz_sub(mpzV, mpzV, mpz2Qkd);
		mpzV.Sub(mpzV, mpz2Qkd)
		//   mpz_mod(mpzV, mpzV, mpzN);
		mpzV.Mod(mpzV, mpzN)
		//   if(mpz_sgn(mpzV)==0)RETURN(1);
		if mpzV.Cmp(big.NewInt(0)) == 0 {
			return true
		}
		/* Calculate Q^{d*2^r} for next r (final iteration irrelevant). */
		if r < s-1 {
			// mpz_mul(mpzQkd, mpzQkd, mpzQkd);
			mpzQkd.Mul(mpzQkd, mpzQkd)
			// mpz_mod(mpzQkd, mpzQkd, mpzN);
			mpzQkd.Mod(mpzQkd, mpzN)
			// mpz_mul_2exp(mpz2Qkd, mpzQkd, 1);
			mpz2Qkd.Lsh(mpzQkd, 1)
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

/* *********************************************************************************************************
 * mpz_strongselfridge_prp:
 * A "strong Lucas-Selfridge pseudoprime" n is a "strong Lucas pseudoprime" using Selfridge parameters of:
 * Find the first element D in the sequence {5, -7, 9, -11, 13, ...} such that Jacobi(D,n) = -1
 * Then use P=1 and Q=(1-D)/4 in the strong Lucase pseudoprime test.
 * Make sure n is not a perfect square, otherwise the search for D will only stop when D=n.
 * **********************************************************************************************************/
func mpz_strongselfridge_prp(n *big.Int) bool {
	//    long int d = 5, p = 1, q = 0;
	d := big.NewInt(5)
	p := big.NewInt(1)
	q := big.NewInt(0)
	//    int max_d = 1000000;
	max_d := big.NewInt(1000000)
	//    int jacobi = 0;
	jacobi := big.NewInt(0)
	//    mpz_t zD;
	zD := new(big.Int)

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

	//    mpz_init_set_ui(zD, d);
	zD.Set(d)

	for {
		// 	 jacobi = mpz_jacobi(zD, n);
		jacobi.Set(Jacobi(zD, n))

		// 	 /* if jacobi == 0, d is a factor of n, therefore n is composite... */
		// 	 /* if d == n, then either n is either prime or 9... */
		// 	 if (jacobi == 0)
		if jacobi.Cmp(big.NewInt(0)) == 0 {
			// 	   if ((mpz_cmpabs(zD, n) == 0) && (mpz_cmp_ui(zD, 9) != 0))
			// 	   {
			// 		 mpz_clear(zD);
			// 		 return PRP_PRIME;
			// 	   }
			// 	   else
			// 	   {
			// 		 mpz_clear(zD);
			// 		 return PRP_COMPOSITE;
			// 	   }
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
		// 	 if (jacobi == -1)
		// 	   break;
		if jacobi.Cmp(big.NewInt(-1)) == 0 {
			break
		}
		// 	 /* if we get to the 5th d, make sure we aren't dealing with a square... */
		// 	 if (d == 13)
		// 	 {
		// 	   if (mpz_perfect_square_p(n))
		// 	   {
		// 		 mpz_clear(zD);
		// 		 return PRP_COMPOSITE;
		// 	   }
		// 	 }
		if d.Cmp(big.NewInt(13)) == 0 {
			sq := big.NewInt(0).Sqrt(n)
			sq.Mul(sq, sq)
			if sq.Cmp(n) == 0 {
				return false
			}
		}

		// 	 if (d < 0)
		// 	 {
		// 	   d *= -1;
		// 	   d += 2;
		// 	 }
		// 	 else
		// 	 {
		// 	   d += 2;
		// 	   d *= -1;
		// 	 }
		if d.Cmp(big.NewInt(0)) < 0 {
			d.Mul(d, big.NewInt(-1))
			d.Add(d, big.NewInt(2))
		} else {
			d.Add(d, big.NewInt(2))
			d.Mul(d, big.NewInt(-1))
		}

		// 	 /* make sure we don't search forever */
		// 	 if (d >= max_d)
		// 	 {
		// 	   mpz_clear(zD);
		// 	   return PRP_ERROR;
		// 	 }
		if d.Cmp(max_d) >= 0 {
			panic("make sure we don't search forever")

		}

		// 	 mpz_set_si(zD, d);
		zD.Set(d)
	}
	//    mpz_clear(zD);

	//    q = (1-d)/4;
	q.Sub(big.NewInt(1), d).Div(q, big.NewInt(4))

	return mpz_stronglucas_prp(n, p, q)

} /* method mpz_strongselfridge_prp */

/* *********************************************************************************************
 * mpz_stronglucas_prp:
 * A "strong Lucas pseudoprime" with parameters (P,Q) is a composite n = (2^r)*s+(D/n), where
 * s is odd, D=P^2-4Q, and (n,2QD)=1 such that either U_s == 0 mod n or V_((2^t)*s) == 0 mod n
 * for some t, 0 <= t < r. [(D/n) is the Jacobi symbol]
 * *********************************************************************************************/
func mpz_stronglucas_prp(n, p, q *big.Int) bool {
	//   mpz_t zD;
	zD := big.NewInt(0)
	//   mpz_t s;
	s := big.NewInt(0)
	//   mpz_t nmj; /* n minus jacobi(D/n) */
	nmj := big.NewInt(0)
	//   mpz_t res;
	res := big.NewInt(0)
	//   mpz_t uh, vl, vh, ql, qh, tmp; /* these are needed for the LucasU and LucasV part of this function */
	uh := big.NewInt(0)
	vl := big.NewInt(0)
	vh := big.NewInt(0)
	ql := big.NewInt(0)
	qh := big.NewInt(0)
	tmp := big.NewInt(0)
	//   long int d = p*p - 4*q;
	d := big.NewInt(0)
	d.Mul(p, p)
	d.Sub(d, big.NewInt(0).Mul(q, big.NewInt(4)))
	//   unsigned long int r = 0;
	r := 0
	//   int ret = 0;
	ret := big.NewInt(0)
	//   unsigned long int j = 0;
	j := 0

	//   if (d == 0) /* Does not produce a proper Lucas sequence */
	//     return PRP_ERROR;
	if d.Cmp(big.NewInt(0)) == 0 {
		panic("Does not produce a proper Lucas sequence")
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
	if n.Cmp(big.NewInt(2)) == 0 {
		return true
	}

	if n.Bit(0) == 0 {
		return false
	}

	//   mpz_init_set_si(zD, d);
	zD.Set(d)
	//   mpz_init(res);

	//   mpz_mul_si(res, zD, q);
	res.Mul(zD, q)
	//   mpz_mul_ui(res, res, 2);
	res.Mul(res, big.NewInt(2))
	//   mpz_gcd(res, res, n);
	res.GCD(nil, nil, res, n)
	//   if ((mpz_cmp(res, n) != 0) && (mpz_cmp_ui(res, 1) > 0))
	//   {
	//     mpz_clear(zD);
	//     mpz_clear(res);
	//     return PRP_COMPOSITE;
	//   }
	if res.Cmp(n) != 0 && res.Cmp(big.NewInt(1)) > 0 {
		return false
	}

	//   mpz_init(s);
	//   mpz_init(nmj);

	//   /* nmj = n - (D/n), where (D/n) is the Jacobi symbol */
	//   mpz_set(nmj, n);
	nmj.Set(n)
	//   ret = mpz_jacobi(zD, n);
	ret = Jacobi(zD, n)
	//   if (ret == -1)
	//     mpz_add_ui(nmj, nmj, 1);
	//   else if (ret == 1)
	//     mpz_sub_ui(nmj, nmj, 1);
	if ret.Cmp(big.NewInt(-1)) == 0 {
		nmj.Add(nmj, big.NewInt(1))
	} else if ret.Cmp(big.NewInt(1)) == 0 {
		nmj.Sub(nmj, big.NewInt(1))
	}

	//   r = mpz_scan1(nmj, 0);
	r = MpzScan1(nmj)
	//   mpz_fdiv_q_2exp(s, nmj, r);
	s.Rsh(nmj, uint(r))

	//   /* make sure U_s == 0 mod n or V_((2^t)*s) == 0 mod n, for some t, 0 <= t < r */
	//   mpz_init_set_si(uh, 1);
	uh.Set(big.NewInt(1))
	//   mpz_init_set_si(vl, 2);
	vl.Set(big.NewInt(2))
	//   mpz_init_set_si(vh, p);
	vh.Set(p)
	//   mpz_init_set_si(ql, 1);
	ql.Set(big.NewInt(1))
	//   mpz_init_set_si(qh, 1);
	qh.Set(big.NewInt(1))
	//   mpz_init_set_si(tmp,0);
	tmp.Set(big.NewInt(0))

	//   for (j = mpz_sizeinbase(s,2)-1; j >= 1; j--)
	for j = s.BitLen() - 1; j >= 1; j-- {
		//     /* ql = ql*qh (mod n) */
		//     mpz_mul(ql, ql, qh);
		ql.Mul(ql, qh)
		//     mpz_mod(ql, ql, n);
		ql.Mod(ql, n)
		//     if (mpz_tstbit(s,j) == 1)
		if s.Bit(j) == 1 {
			//       /* qh = ql*q */
			//       mpz_mul_si(qh, ql, q);
			qh.Mul(ql, q)

			//       /* uh = uh*vh (mod n) */
			//       mpz_mul(uh, uh, vh);
			uh.Mul(uh, vh)
			//       mpz_mod(uh, uh, n);
			uh.Mod(uh, n)

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

			//       /* uh = uh*vl - ql (mod n) */
			//       mpz_mul(uh, uh, vl);
			uh.Mul(uh, vl)
			//       mpz_sub(uh, uh, ql);
			uh.Sub(uh, ql)
			//       mpz_mod(uh, uh, n);
			uh.Mod(uh, n)

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

	//   /* uh = uh*vl - ql */
	//   mpz_mul(uh, uh, vl);
	uh.Mul(uh, vl)
	//   mpz_sub(uh, uh, ql);
	uh.Sub(uh, ql)

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

	//   mpz_mod(uh, uh, n);
	uh.Mod(uh, n)
	//   mpz_mod(vl, vl, n);
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

	//   /* uh contains LucasU_s and vl contains LucasV_s */
	//   if ((mpz_cmp_ui(uh, 0) == 0) || (mpz_cmp_ui(vl, 0) == 0))
	if uh.Cmp(big.NewInt(0)) == 0 || vl.Cmp(big.NewInt(0)) == 0 {
		return true
		//     mpz_clear(zD);
		//     mpz_clear(s);
		//     mpz_clear(nmj);
		//     mpz_clear(res);
		//     mpz_clear(uh);
		//     mpz_clear(vl);
		//     mpz_clear(vh);
		//     mpz_clear(ql);
		//     mpz_clear(qh);
		//     mpz_clear(tmp);
		//     return PRP_PRP;
	}

	//   for (j = 1; j < r; j++)
	for j = 1; j < r; j++ {
		//     /* vl = vl*vl - 2*ql (mod n) */
		//     mpz_mul(vl, vl, vl);
		vl.Mul(vl, vl)
		//     mpz_mul_si(tmp, ql, 2);
		tmp.Mul(ql, big.NewInt(2))
		//     mpz_sub(vl, vl, tmp);
		vl.Sub(vl, tmp)
		//     mpz_mod(vl, vl, n);
		vl.Mod(vl, n)

		//     /* ql = ql*ql (mod n) */
		//     mpz_mul(ql, ql, ql);
		ql.Mul(ql, ql)
		//     mpz_mod(ql, ql, n);
		ql.Mod(ql, n)

		//     if (mpz_cmp_ui(vl, 0) == 0)
		    if vl.Cmp(big.NewInt(0)) == 0 {
		//       mpz_clear(zD);
		//       mpz_clear(s);
		//       mpz_clear(nmj);
		//       mpz_clear(res);
		//       mpz_clear(uh);
		//       mpz_clear(vl);
		//       mpz_clear(vh);
		//       mpz_clear(ql);
		//       mpz_clear(qh);
		//       mpz_clear(tmp);
		//       return PRP_PRP;
		return true
		}
	}

	//   mpz_clear(zD);
	//   mpz_clear(s);
	//   mpz_clear(nmj);
	//   mpz_clear(res);
	//   mpz_clear(uh);
	//   mpz_clear(vl);
	//   mpz_clear(vh);
	//   mpz_clear(ql);
	//   mpz_clear(qh);
	//   mpz_clear(tmp);
	//   return PRP_COMPOSITE;
	return false

} /* method mpz_stronglucas_prp */

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
			rightLimit.SetString("10000", 10)
			for ; num.Cmp(rightLimit) <= 0; num.Add(num, big.NewInt(1)) {

				// r := StrongLucasSelfridge(num)
				r := mpz_strongselfridge_prp(num)
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
		r := StrongLucasSelfridge(big.NewInt(105))
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
