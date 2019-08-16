package main

import (
	"fmt"

	"github.com/alexkarpovich/go-statistics-lab/library"
)

func main() {
	var i uint64
	var r1, r2 float64
	var rejectPercent uint64
	var count uint64
	var pReject float64

	fmt.Sscanf("12 10", "%d %d", &rejectPercent, &count)

	pReject = float64(rejectPercent) / 100

	for i = 0; i <= 2; i++ {
		r1 = r1 + library.BinomialDistribution(i, count, pReject)
	}

	for i = 2; i <= count; i++ {
		r2 = r2 + library.BinomialDistribution(i, count, pReject)
	}

	fmt.Printf("%.3f\n%.3f", r1, r2)
}
