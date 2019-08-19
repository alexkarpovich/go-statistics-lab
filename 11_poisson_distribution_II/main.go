package main

import (
	"fmt"
)

func main() {
	var m1, m2 float64
	var c1, c2 float64

	fmt.Sscanf("0.88 1.55", "%f %f", &m1, &m2)

	c1 = 160+40*(m1+m1*m1)
	c2 = 128+40*(m2+m2*m2)

	fmt.Printf("%.3f\n%.3f", c1, c2)
}
