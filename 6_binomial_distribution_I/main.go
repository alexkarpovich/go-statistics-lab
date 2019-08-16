package main

import (
	"fmt"

	"github.com/alexkarpovich/go-statistics-lab/library"
)

const n uint64 = 6

func main() {
	var i uint64
	var res float64
	var boysRation float64
	var girlsRation float64

	fmt.Sscanf("1.09 1", "%f %f", &boysRation, &girlsRation)

	fmt.Println(boysRation, girlsRation)

	totalRation := boysRation + girlsRation
	pBoys := boysRation / totalRation

	for i = 3; i <= n; i++ {
		res = res + library.BinomialDistribution(i, n, pBoys)
	}

	fmt.Printf("%.3f", res)
}
