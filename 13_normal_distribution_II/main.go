package main

import (
	"fmt"
	"github.com/alexkarpovich/go-statistics-lab/library"
)

func main() {
	var mean, stddev float64
	var p0, p1, r1, r2, r3 float64


	fmt.Sscanf("70 10\n80\n60", "%f %f\n%f\n%f %f", &mean, &stddev, &p0, &p1)

	fmt.Print(mean, stddev, p0, p1)

	r1 = 1 - library.NormalProbability(p0, mean, stddev)
	r2 = 1 - library.NormalProbability(p1, mean, stddev)
	r3 = library.NormalProbability(p1, mean, stddev)

	fmt.Printf("%.3f\n%.3f\n%.3f", r1, r2, r3)
}
