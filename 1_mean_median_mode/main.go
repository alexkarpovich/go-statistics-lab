package main

import (
	"fmt"
	"github.com/alexkarpovich/go-statistics-lab/library"
)

func main() {
	n := 10
	//fmt.Scan(&n)

	arr := make([]int, n)
	fmt.Sscan("4978 11735 14216 14470 38120 51135 64630 67060 73429 99233", library.PackAddrs(arr)...)

	fmt.Print(library.Mean(arr), library.Median(arr, true), library.Mode(arr))
}
