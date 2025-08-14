package main

import (
	"fmt"
	"math/big"
)

// https://eprint.iacr.org/2009/173 椭圆曲线标量乘法的快速多基方法和其他几种优化
// https://eprint.iacr.org/2009/173.pdf

type point struct {
	x *big.Int
	y *big.Int
}

const POLY_SIZE = 8192
const SIEVE_LIMIT = 1073741824

const LOG2 = 0.301029995

const PRECISION = 10000
const ERROR_SHIFT = 1000

const BMAX = 2000
const DMAX = 20

var Bmax = BMAX
var Dmax = DMAX

var one = 0
var ISPRINTSTEP = false
var Quiet = false
var verbose = false

func MillerRabbin(num *big.Int) bool {
	a := big.NewInt(2)
	if num.Cmp(big.NewInt(1)) <= 0 {
		return false
	}
	if num.Cmp(big.NewInt(2)) == 0 {
		return true
	}
	if num.Bit(0) == 0 {
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

/*
Algorithm 1.5.3 (Modified Cornacchia). See "A Course
in Computational Algebraic Number Theory" by Henri
Cohen page 36. Let p be a prime number and D be a
negative number such that D = 0 or 1 modulo 4 and
| D | < 4 * p. This algorithm either outputs an
integer solution (x, y) to the Diophantine equation
x * x + | D | * y * y = 4 * p, or says that such a
solution does not exist.
*/
func modified_Cornacchia(D, p, x, y *big.Int) *big.Int {
	// int value = 0;
	value := big.NewInt(0)
	// mpz_class dd, xx;
	dd := big.NewInt(0)
	xx := big.NewInt(0)
	// mpz_class a = 0, b = 0, c = 0, d = 0, e = 0;
	a := big.NewInt(0)
	b := big.NewInt(0)
	c := big.NewInt(0)
	d := big.NewInt(0)
	e := big.NewInt(0)
	// mpz_class l = 0, r = 0, x0 = 0;
	l := big.NewInt(0)
	r := big.NewInt(0)
	x0 := big.NewInt(0)

	// if (JACOBI(D, p) != -1) {
	if Jacobi(D, p).Cmp(big.NewInt(-1)) != 0 {
		// x0 = square_root_mod(D, p);
		x0.ModSqrt(D, p)

		// dd = modpos(D, 2);
		dd.Mod(D, big.NewInt(2))
		// xx = modpos(x0, 2);
		xx.Mod(x0, big.NewInt(2))

		// if (dd != xx)
		// x0 = p - x0;
		if dd.Cmp(xx) != 0 {
			x0.Sub(p, x0)
		}

		// a = p * 2;
		a.Mul(p, big.NewInt(2))
		// b = x0;
		b.Set(x0)
		// c = sqrt(p);
		c.Sqrt(p)

		// l = c * 2;
		l.Mul(c, big.NewInt(2))
		// while (b > l) {
		for b.Cmp(l) > 0 {
			// r = modpos(a, b);
			r.Mod(a, b)
			// a = b;
			a.Set(b)
			// b = r;
			b.Set(r)
		}

		// c = p * 4;
		c.Mul(p, big.NewInt(4))
		// a = b * b;
		a.Mul(b, b)
		// e = c - a;
		e.Sub(c, a)
		// d = abs(D);
		d.Abs(D)

		// c = e / d;
		c.Div(e, d)
		// r = modpos(e, d);
		r.Mod(e, d)

		// if ((r == 0) && square_test(c, y)) {
		sq := big.NewInt(0).Sqrt(c)
		square_test := false
		if c.Cmp(big.NewInt(0).Mul(sq, sq)) == 0 {
			square_test = true
			y.Set(sq)
		}
		if r.Cmp(big.NewInt(0)) == 0 && square_test {
			// *x = b;
			x.Set(b)
			// value = 1
			value.SetInt64(1)

		}
	}

	// return value;
	return value
}

func Atkin(N *big.Int) int {
	/* returns 2 if N is composite, 1 if probably prime, 0 if proven prime */
	// 	int value1, value2;
	value1 := big.NewInt(0)
	value2 := big.NewInt(0)
	// 	long p;
	p := big.NewInt(0)
	// 	mpz_class Ni = N, a, b, d, i = 0, m, q, t, x, y;
	Ni := big.NewInt(0).Set(N)
	a := big.NewInt(0)
	b := big.NewInt(0)
	d := big.NewInt(0)
	i := big.NewInt(0)
	m := big.NewInt(0)
	q := big.NewInt(0)
	t := big.NewInt(0)
	x := big.NewInt(0)
	y := big.NewInt(0)

	// 	struct point P, P1, P2;
	P := point{big.NewInt(0), big.NewInt(0)}
	P1 := point{big.NewInt(0), big.NewInt(0)}
	P2 := point{big.NewInt(0), big.NewInt(0)}
	// 	long n = 0, k = 0;
	n := big.NewInt(0)
	k := big.NewInt(0)
	// 	long D;
	D := big.NewInt(0)
	// 	bool found = false, found2 = false;
	found := false
	found2 := false
	// 	mpz_class u, v;
	u := big.NewInt(0)
	v := big.NewInt(0)
	// 	mpz_class g;
	g := big.NewInt(0)
	// 	long points_tried = 0;
	points_tried := big.NewInt(0)
	// 	mpz_class root[POLY_SIZE];
	var root [POLY_SIZE]*big.Int

	// 	long rootSize = 0;
	rootSize := big.NewInt(0)
	// 	mpz_class j, w;
	j := big.NewInt(0)
	w := big.NewInt(0)

	for {

		//  cout << "Step 2\n";
		if ISPRINTSTEP {
			fmt.Println("Step 2")
		}

		if Ni.Cmp(big.NewInt(SIEVE_LIMIT)) <= 0 {
			// p = 2;
			p.SetInt64(2)
			// for p <= sqrt(Ni) {
			for p.Cmp(big.NewInt(0).Sqrt(Ni)) <= 0 {
				// if (Ni % p == 0) {
				if big.NewInt(0).Mod(Ni, p).Cmp(big.NewInt(0)) == 0 {
					// cout << "1 factor = " << p << "\n";
					fmt.Println("1 factor = ", p)
					// fflush(stdout);
					return 2
				}
				// p++;
				p.Add(p, big.NewInt(1))
			}
			return 0
		}

		// 		if (!Rabin_Miller(Ni)) return 2;
		if !MillerRabbin(Ni) {
			return 2
		}

		// for Bmax <= 2000000000 && Dmax <= 312500 {
		for Bmax <= 2000000000 && Dmax <= 312500 {
			if !Quiet {
				// cout << "Bmax = " << Bmax << "\n"
				// cout << "Dmax = " << Dmax << "\n"
				// fflush(stdout)
				fmt.Println("Bmax = ", Bmax)
				fmt.Println("Dmax = ", Dmax)
			}

			// 			n = 1;
			n.Set(big.NewInt(1))
			// for ((D = -n++) > -Dmax) {
			D.Neg(n)
			n.Add(n, big.NewInt(1))
			for D.Cmp(big.NewInt(0).Neg(big.NewInt(0).SetInt64(int64(Dmax)))) > 0 {

				// 				if (D%4 != 0 && D%4 != -3 && D%4 != 1)
				// 					continue;
				if D.Mod(D, big.NewInt(4)).Cmp(big.NewInt(0)) != 0 && D.Mod(D, big.NewInt(4)).Cmp(big.NewInt(1)) != 0 {
					continue
				}

				found = false
				found2 = false

				if verbose {
					// cout << D << '\r'
					fmt.Println(D)
					// fflush(stdout)
				}

				//  cout << "Step 3\n";
				if ISPRINTSTEP {
					fmt.Println("Step 3")
				}

				// 				if (JACOBI(D + Ni, Ni) != 1)
				// 					continue;

				if Jacobi(big.NewInt(0).Add(D, Ni), Ni).Cmp(big.NewInt(1)) != 0 {
					continue
				}

				// 				if (!modified_Cornacchia(D, Ni, &u, &v))
				// 					continue;
				if modified_Cornacchia(D, Ni, u, v).Cmp(big.NewInt(0)) == 0 {
					continue
				}

				//  cout << "Step 4\n";
				if ISPRINTSTEP {
					fmt.Println("Step 4")
				}

				// 				t = sqrt(sqrt(Ni)) + 1;
				t.Sqrt(Ni).Sqrt(t).Add(t, big.NewInt(1))
				// 				t = t * t;
				t.Mul(t, t)

				// 				if (check_for_factor(&q, m = Ni + 1 + u, t))
				// 					found = true;
				// 				else if (check_for_factor(&q, m = Ni + 1 - u, t))
				// 					found = true;
				// 				else if (D == -4) {
				// 					if (check_for_factor(&q, m = Ni + 1 + 2*v, t))
				// 						found = true;
				// 					else if (check_for_factor(&q, m = Ni + 1 - 2*v, t))
				// 						found = true;
				// 				}
				// 				else if (D == -3) {
				// 					if (check_for_factor(&q, m = Ni + 1 + (u + 3*v)/2, t))
				// 						found = true;
				// 					else if (check_for_factor(&q, m = Ni + 1 - (u + 3*v)/2, t))
				// 						found = true;
				// 					else if (check_for_factor(&q, m = Ni + 1 + (u - 3*v)/2, t))
				// 						found = true;
				// 					else if (check_for_factor(&q, m = Ni + 1 - (u - 3*v)/2, t))
				// 						found = true;
				// 				}

				// 				if (!found)
				// 					continue;

				// //  cout << "Step 6\n";

				// 				rootSize = 0;

				// 				for (int type = 0; type <= 2; type++) {

				// 					if (!find_curve(type, &a, &b, D, Ni, root, &rootSize))
				// 						continue;

				// 					for (long roots_tried = 0; roots_tried < rootSize; roots_tried++) {
				// 						if (type == 0) {
				// 							a = modpos(a, Ni);
				// 							b = modpos(b, Ni);
				// 						}

				// 						else if (hilbert && type == 1) {
				// 							j = modpos(root[roots_tried], Ni);

				// 							if (!Quiet) {
				// 								cout << "j = " << j << "\n";
				// 								fflush(stdout);
				// 							}

				// 							mpz_class c;
				// 							c = modpos(j * inverse(j - 1728, Ni), Ni);
				// 							a = modpos(-3 * c, Ni);
				// 							b = modpos(2 * c, Ni);
				// 						}

				// 						else if (weber && type == 2) {
				// 							w = modpos(root[roots_tried], Ni);

				// 							if (!Quiet) {
				// 								cout << "u = " << w << "\n";
				// 								fflush(stdout);
				// 							}

				// 							if (D % 4 == 0)
				// 								j = Vegas(w, Ni, D / 4);
				// 							else
				// 								j = Vegas(w, Ni, D);

				// 							if (!Quiet) {
				// 								cout << "j = " << j << "\n";
				// 								fflush(stdout);
				// 							}

				// 							mpz_class c;
				// 							c = modpos(j * inverse(j - 1728, Ni), Ni);
				// 							a = modpos(-3 * c, Ni);
				// 							b = modpos(2 * c, Ni);
				// 						}

				// //  cout << "Step 7\n";

				// 						do {
				// 							do g = modpos(rand2(), Ni); while (g == 0);
				// 							if (JACOBI(g, Ni) != -1)
				// 								continue;
				// 							if (D == -3)	// Cohen
				// //							if (Ni % 3 == 1) // Studholme
				// 								if (exp_mod(g, (Ni - 1)/3, Ni) == 1)
				// 									continue;
				// 							break;
				// 						} while (true);

				// //  cout << "Step 8\n";

				// 						points_tried = 0;

				// 						do {

				// 							do {
				// 								do {
				// 									do x = modpos(rand2(), Ni); while (x == 0);
				// 									y = modpos(((x * x) % Ni * x) % Ni + a * x + b, Ni);
				// 								} while (JACOBI(y, Ni) == -1);
				// 								y = square_root_mod(y, Ni);
				// 							} while (y == 0);

				// //  cout << "Step 9\n";

				// 							P.x = x, P.y = y;
				// 							points_tried++;
				// 							k = 0;

				// 							do {
				// //  cout << "Step 12\n";

				// 								value2 = multiply(a, m/q, Ni, P, &P2, &d);

				// 								if (value2 == 1) {
				// 									cout << "3 factor = " << d << "\n";
				// 									fflush(stdout);
				// 									return 2;
				// 								}

				// 								value1 = multiply(a, q, Ni, P2, &P1, &d);

				// 								if (value1 == 1) {
				// 									cout << "2 factor = " << d << "\n";
				// 									fflush(stdout);
				// 									return 2;
				// 								}

				// 								if ((value1 == -1) && (value2 == 0)) {
				// //  cout << "Step 13\n";

				// 									found2 = true;
				// 									break;
				// 								}

				// //  cout << "Step 10\n";

				// 								++k;

				// 								if (D == -3) {
				// 									if (k >= 6)
				// 										break;
				// 									b *= g;
				// 								}
				// 								else if (D == -4) {
				// 									if (k >= 4)
				// 										break;
				// 									a *= g;
				// 								}
				// 								else {
				// 									if (k >= 2)
				// 										break;
				// 									a *= (g * g);
				// 									b *= (g * g * g);
				// 								}

				// 								a = modpos(a, Ni);
				// 								b = modpos(b, Ni);

				// 							} while (true);

				// 						} while (!found2 && points_tried < 100);

				// 						if (found2)
				// 							break;
				// 					}

				// 					if (found2)
				// 						break;
				// 				}

				// 				if (found2)
				// 					break;
				D.Neg(n)
				n.Add(n, big.NewInt(1))
			}

			// 			if (found2) {
			// 				if (!staticBmax)
			// 					Bmax = BMAX;
			// 				if (!staticDmax)
			// 					Dmax = DMAX;
			// 				break;
			// 			}
			// 			else {
			// 				if (!staticBmax)
			// 					Bmax *= 10;
			// 				if (!staticDmax)
			// 					Dmax *= 5;
			// 			}
		}

		// 		if (Dmax > 312500) {
		// 			if (!Quiet) {
		// 				cout << "ProvePrime: ran out of discriminants\n";
		// 				fflush(stdout);
		// 			}
		// 			return 2;
		// 		}
		// 		if (Bmax > 2000000000) {
		// 			if (!Quiet) {
		// 				cout << "ProvePrime: exceeded maximum factoring bounds\n";
		// 				fflush(stdout);
		// 			}
		// 			return 2;
		// 		}

		// 		cout << "N[" << i << "] = " << Ni << "\n";
		// 		cout << "a = " << a << "\n";
		// 		cout << "b = " << b << "\n";
		// 		cout << "m = " << m << "\n";
		// 		cout << "q = " << q << "\n";
		// 		cout << "P = (" << P.x << ", " << P.y << ")\n";
		// 		cout << "P1 = (" << P1.x << ", " << P1.y << ")\n";
		// 		cout << "P2 = (" << P2.x << ", " << P2.y << ")\n";

		// 		fflush(stdout);

		// 		i++;
		// 		Ni = q;
		if one == 0 {
			break
		}
	}
	// 	} while (!one);

	// 	if (!one) return 0;
	// 	if (one) return 1;
	if one == 0 {
		return 0
	} else {
		return 1
	}
}

func main() {
	fmt.Println("Hello World")
}
