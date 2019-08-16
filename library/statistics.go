package library

import (
	"math"
	"sort"
)

func Mean(arr []int) (mean float64) {
	for i := 0; i < len(arr); i++ {
		mean = mean + float64(arr[i])
	}
	mean = mean / float64(len(arr))

	return
}

func WeightedMean(arr []int, weights []int) (mean float64) {
	weight_sum := 0

	for i := 0; i < len(arr); i++ {
		mean = mean + float64(arr[i]*weights[i])
		weight_sum = weight_sum + weights[i]
	}
	mean = mean / float64(weight_sum)

	return
}

func Median(arr []int, need_sort bool) float64 {
	mid := float64(len(arr)-1) / 2
	midl := int(mid)
	midr := int(math.Ceil(mid))

	if need_sort {
		sort.Sort(sort.IntSlice(arr))
	}

	if midl != midr {
		return float64(arr[midl]+arr[midr]) / 2
	}

	return float64(arr[midl])
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

func Quartiles(arr []int) (q1 float64, q2 float64, q3 float64) {
	var midl int
	var midr int
	mid := float64(len(arr)) / 2

	if len(arr)%2 == 0 {
		midr = int(mid)
		midl = midr
	} else {
		midl = int(mid)
		midr = midl + 1
	}

	sort.Sort(sort.IntSlice(arr))

	q1 = Median(arr[:midl], false)
	q2 = Median(arr, false)
	q3 = Median(arr[midr:], false)

	return
}

func Variance(arr []int) float64 {
	variance := 0.0
	mean := Mean(arr)

	for i := 0; i < len(arr); i++ {
		variance = variance + math.Pow(float64(arr[i])-mean, 2)
	}

	return variance / float64(len(arr))
}

func StdDeviation(arr []int) float64 {
	return math.Sqrt(Variance(arr))
}

func InterquartileRange(arr []int) float64 {
	q1, _, q3 := Quartiles(arr)

	return q3 - q1
}

func BinomialDistribution(x uint64, n uint64, p float64) float64 {
	return float64(Combinations(n, x)) * math.Pow(p, float64(x)) * math.Pow(1-p, float64(n-x))
}

func NegativeBinomialDistribution(x uint64, n uint64, p float64) float64 {
	return float64(Combinations(n-1, x-1)) * math.Pow(p, float64(x)) * math.Pow(1-p, float64(n-x))
}

func GeometricDistribution(n uint64, p float64) float64 {
	return math.Pow(1-p, float64(n-1)) * p
}
