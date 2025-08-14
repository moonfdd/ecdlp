package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

// 佩兰失败
// https://en.wikipedia.org/wiki/Perrin_number#Perrin_pseudoprimes
func main() {
	a := mat.NewDense(3, 3, []float64{
		1, 2, 3,
		4, 5, 6,
		7, 8, 9})
	b := mat.NewDense(3, 3, []float64{
		8, 8, 8,
		8, 8, 8,
		8, 8, 8})
	var c mat.Dense
	c.Add(a, b)
	fmt.Printf("%v\n\n", mat.Formatted(&c))

	c.Sub(a, b)
	fmt.Printf("%v\n\n", mat.Formatted(&c))
}
