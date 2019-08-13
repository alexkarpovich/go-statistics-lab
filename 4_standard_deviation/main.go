package main

import (
	"fmt"
	"github.com/alexkarpovich/go-statistics-lab/library"
)

func main() {
	n := 5
	//fmt.Scan(&n)

	arr := make([]int, n)

	fmt.Sscan("10 40 30 50 20", library.PackAddrs(arr)...)

	fmt.Printf("%.1f\n", library.StdDeviation(arr))
}
