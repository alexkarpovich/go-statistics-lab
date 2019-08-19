package main

import (
	"fmt"
	"github.com/alexkarpovich/go-statistics-lab/library"
)

func main() {
	var count uint64
	var mean, stddev, p, zScore float64
	var a, b float64

	fmt.Sscanf("100\n500\n80\n.95\n1.96", "%d\n%f\n%f\n%f\n%f", &count, &mean, &stddev, &p, &zScore)
	_ = p

	a = library.DetermineCentralLimitA(zScore, count, mean, stddev)
	b = library.DetermineCentralLimitB(zScore, count, mean, stddev)

	fmt.Printf("%.2f\n%.2f", a, b)
}
