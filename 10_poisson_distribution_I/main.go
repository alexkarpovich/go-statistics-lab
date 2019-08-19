package main

import (
	"fmt"

	"github.com/alexkarpovich/go-statistics-lab/library"
)

func main() {
	var num uint64
	var mean float64
	var p float64

	fmt.Sscanf("2.5\n5", "%f\n%d", &mean, &num)

	p = library.PoissonDistribution(mean, num)

	fmt.Printf("%.3f", p)
}
