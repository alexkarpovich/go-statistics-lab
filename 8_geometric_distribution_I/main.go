package main

import (
	"fmt"

	"github.com/alexkarpovich/go-statistics-lab/library"
)

func main() {
	var r float64
	var num, denom uint64
	var n uint64
	var pDefect float64

	fmt.Sscanf("1 3", "%d %d", &num, &denom)
	fmt.Sscanf("5", "%d", &n)

	pDefect = float64(num) / float64(denom)
	r = library.GeometricDistribution(n, pDefect)

	fmt.Printf("%.3f", r)
}
