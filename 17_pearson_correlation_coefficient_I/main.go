package main

import (
	"fmt"
	"os"

	"github.com/alexkarpovich/go-statistics-lab/library"
)

func main() {
	var r float64
	var n uint64

	f, err := os.Open("./test.txt")
	if err != nil {
		panic(err)
	}

	fmt.Fscanf(f, "%d", &n)

	x := make([]float64, n)
	y := make([]float64, n)

	fmt.Fscan(f, library.PackFloatAddrs(x)...)
	fmt.Fscan(f, library.PackFloatAddrs(y)...)

	r = library.Correlation(x, y)

	fmt.Printf("%.3f", r)
}
