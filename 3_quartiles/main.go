package main

import (
	"fmt"

	"github.com/alexkarpovich/go-statistics-lab/library"
)

func main() {
	n := 10
	//fmt.Scan(&n)

	arr := make([]float64, n)

	fmt.Sscan("3 7 8 5 12 14 21 15 18 14", library.PackFloatAddrs(arr)...)
	q1, q2, q3 := library.Quartiles(arr)

	fmt.Printf("%.0f\n%.0f\n%.0f\n", q1, q2, q3)
}
