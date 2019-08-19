package main

import (
	"fmt"
	"github.com/alexkarpovich/go-statistics-lab/library"
)

func main() {
	var mean, stddev float64
	var p0, p1, p2, r1, r2 float64


	fmt.Sscanf("20 2\n19.5\n20 22", "%f %f\n%f\n%f %f", &mean, &stddev, &p0, &p1, &p2)

	r1 = library.NormalProbability(p0, mean, stddev)
	r2 = library.NormalProbability(p2, mean, stddev) - library.NormalProbability(p1, mean, stddev)

	fmt.Printf("%.3f\n%.3f", r1, r2)
}
