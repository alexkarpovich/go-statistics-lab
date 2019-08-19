package main

import (
	"fmt"

	"github.com/alexkarpovich/go-statistics-lab/library"
)

func main() {
	n := 5
	//fmt.Scan(&n)

	arr := make([]float64, n)
	weights := make([]float64, n)
	fmt.Sscan("10 40 30 50 20", library.PackFloatAddrs(arr)...)
	fmt.Sscan("1 2 3 4 5", library.PackFloatAddrs(weights)...)

	fmt.Printf("%.1f", library.WeightedMean(arr, weights))
}
