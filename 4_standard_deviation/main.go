package main

import (
	"fmt"

	"github.com/alexkarpovich/go-statistics-lab/library"
)

func main() {
	n := 5
	//fmt.Scan(&n)

	arr := make([]float64, n)

	fmt.Sscan("10 40 30 50 20", library.PackFloatAddrs(arr)...)

	fmt.Printf("%.1f\n", library.StdDeviation(arr))
}
