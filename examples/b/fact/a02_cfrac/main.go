package main

// func ecm(N *big.Int, arange int64, plimit int64) *big.Int {
// 	if N.ProbablyPrime(0) {
// 		return N
// 	}

// 	if pow := big.NewInt(0).Log(N, big.NewInt(plimit)); pow.Cmp(big.NewInt(1)) == 0 {
// 		return N.Sqrt(nil)
// 	}

// 	primes := big.NewInt(plimit).PrimeRange(nil)

// 	for a := -arange; a <= arange; a++ {
// 		x := big.NewInt(0)
// 		y := big.NewInt(1)

// 		for _, B := range primes {
// 			t := big.NewInt(B).Exp(Big.NewInt(plimit), big.NewInt(B), nil)

// 			first := true
// 			sx := big.NewInt(0).Set(x)
// 			sy := big.NewInt(0).Set(y)

// 			for t.Cmp(big.Zero) > 0 {

// 				if t.BitLen()&1 == 1 { // if t is odd
// 					var xn, yn = new(big.Int), new(big.Int)
// 					if first {
// 						xn.Set(sx)
// 						yn.Set(sy)
// 						first = false
// 					} else {
// 						u := new(big.Int).Sub(sx, xn)
// 						u.ModInverse(u, N)

// 						if u == nil {
// 							d := new(big.Int).GCD(nil, nil, new(big.Int).Sub(sx, xn), N)
// 							if d.Cmp(N) != 0 {
// 								return d
// 							}
// 							break
// 						}

// 						L := new(big.Int).Mul(u, new(big.Int).Sub(sy, yn))
// 						L.Mod(L, N)
// 						xs := new(big.Int).Sub(new(big.Int).Mul(L, L), new(big.Int).Add(new(big.Int).Add(xn, sx), N))
// 						xs.Mod(xs, N)

// 						yn.Mul(L, new(big.Int).Sub(xn, xs))
// 						yn.Sub(yn, yn)
// 						yn.Mod(yn, N)

// 						xn.Set(xs)
// 					}
// 				}

// 				u := new(big.Int).Mul(sy, big.NewInt(2))
// 				u.ModInverse(u, N)

// 				if u == nil {
// 					d := new(big.Int).GCD(nil, nil, big.NewInt(2), N)
// 					if d.Cmp(N) != 0 {
// 						return d
// 					}
// 					break
// 				}

// 				L := new(big.Int).Mul(u, big.NewInt(3))
// 				L.Mul(L, new(big.Int).Mul(sx, sx))
// 				L.Add(L, big.NewInt(a))
// 				L.Mod(L, N)

// 				x2 := new(big.Int).Sub(new(big.Int).Mul(L, L), big.NewInt(2*int64(sx)))
// 				x2.Mod(x2, N)

// 				sy = new(big.Int).Sub(sy, new(big.Int).Mul(new(big.Int).Mul(L, new(big.Int).Sub(sx, x2)), new(big.Float))).Mod(sy, N)
// 			}
// 			x.Set(xn)
// 			y.Set(yn)
// 		}
// 	}

// 	return N
// }

func main() {
	// N := big.NewInt(14304849576137459)
	// result := ecm(N, 100, 10000)
	// fmt.Println(result.String())
}