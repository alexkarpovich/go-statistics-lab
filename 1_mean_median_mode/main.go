package main

import (
	"fmt"
	"math"
	"sort"
)

func packAddrs(n []int) []interface{} {
	p := make([]interface{}, len(n))
	for i := range n {
		p[i] = &n[i]
	}
	return p
}

func Mean(arr []int) (mean float32) {
	for i := 0; i < len(arr); i++ {
		mean = mean + float32(arr[i])
	}
	mean = mean / float32(len(arr))

	return
}

func Median(arr []int) float32 {
	mid := float64(len(arr)-1) / 2
	midl := int(mid)
	midr := int(math.Ceil(mid))

	sort.Sort(sort.IntSlice(arr))

	if midl != midr {
		return float32(arr[midl]+arr[midr]) / 2
	}

	return float32(arr[midl])
}

func Mode(arr []int) int {
	var cur int

	mid := 0
	max_count := 0
	counter := make(map[int]int)

	for i := 0; i < len(arr); i++ {
		cur = arr[i]
		counter[cur]++

		if max_count < counter[cur] || counter[cur] == counter[arr[mid]] && cur < arr[mid] {
			max_count = counter[cur]
			mid = i
		}
	}

	return arr[mid]
}

func main() {
	n := 10
	//fmt.Scan(&n)

	arr := make([]int, n)
	fmt.Sscan("4978 11735 14216 14470 38120 51135 64630 67060 73429 99233", packAddrs(arr)...)

	fmt.Print(Mean(arr), Median(arr), Mode(arr))
}
