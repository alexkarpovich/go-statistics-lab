package main

import (
	"fmt"
)

func packAddrs(n []int) []interface{} {
	p := make([]interface{}, len(n))
	for i := range n {
		p[i] = &n[i]
	}
	return p
}

func WeightedMean(arr []int, weights []int) (mean float32) {
	weight_sum := 0

	for i := 0; i < len(arr); i++ {
		mean = mean + float32(arr[i]*weights[i])
		weight_sum = weight_sum + weights[i]
	}
	mean = mean / float32(weight_sum)

	return
}

func main() {
	n := 5
	//fmt.Scan(&n)

	arr := make([]int, n)
	weights := make([]int, n)
	fmt.Sscan("10 40 30 50 20", packAddrs(arr)...)
	fmt.Sscan("1 2 3 4 5", packAddrs(weights)...)

	fmt.Printf("%.1f", WeightedMean(arr, weights))
}
