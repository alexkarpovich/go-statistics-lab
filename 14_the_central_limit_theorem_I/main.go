package main

import (
	"fmt"
	"github.com/alexkarpovich/go-statistics-lab/library"
)

func main() {
	var count uint64
	var mean, stddev, limit float64
	var r float64


	fmt.Sscanf("9800\n49\n205\n15", "%f\n%d\n%f\n%f", &limit, &count, &mean, &stddev)

	r = library.CentralLimit(limit, count, mean, stddev)

	fmt.Printf("%.4f", r)
}
