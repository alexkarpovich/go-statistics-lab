package main

import (
	"fmt"

	"github.com/alexkarpovich/go-statistics-lab/library"
)

func main() {
	n := 6
	//fmt.Scan(&n)

	items := make([]float64, n)
	freqs := make([]int, n)
	arr := make([]float64, 0)

	fmt.Sscan("6 12 8 10 20 16", library.PackFloatAddrs(items)...)
	fmt.Sscan("5 4 3 2 1 5", library.PackAddrs(freqs)...)

	for i := 0; i < len(items); i++ {
		for j := 0; j < freqs[i]; j++ {
			arr = append(arr, items[i])
		}
	}

	fmt.Printf("%.1f\n", library.InterquartileRange(arr))
}
