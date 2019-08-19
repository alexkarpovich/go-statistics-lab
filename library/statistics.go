package library

import (
	"math"
	"sort"
)

func Mean(arr []float64) (mean float64) {
	for i := 0; i < len(arr); i++ {
		mean = mean + arr[i]
	}
	mean = mean / float64(len(arr))

	return
}

func WeightedMean(arr []float64, weights []float64) (mean float64) {
	var weight_sum float64

	for i := 0; i < len(arr); i++ {
		mean = mean + arr[i]*weights[i]
		weight_sum = weight_sum + weights[i]
	}
	mean = mean / weight_sum

	return
}

func Median(arr []float64, need_sort bool) float64 {
	mid := float64(len(arr)-1) / 2
	midl := int(mid)
	midr := int(math.Ceil(mid))

	if need_sort {
		sort.Sort(sort.Float64Slice(arr))
	}

	if midl != midr {
		return (arr[midl] + arr[midr]) / 2
	}

	return arr[midl]
}

func Mode(arr []float64) float64 {
	var cur float64

	mid := 0
	max_count := 0
	counter := make(map[float64]int)

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

func Quartiles(arr []float64) (q1 float64, q2 float64, q3 float64) {
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

	sort.Sort(sort.Float64Slice(arr))

	q1 = Median(arr[:midl], false)
	q2 = Median(arr, false)
	q3 = Median(arr[midr:], false)

	return
}

func Variance(arr []float64) float64 {
	variance := 0.0
	mean := Mean(arr)

	for i := 0; i < len(arr); i++ {
		variance = variance + math.Pow(arr[i]-mean, 2)
	}

	return variance / float64(len(arr))
}

func StdDeviation(arr []float64) float64 {
	return math.Sqrt(Variance(arr))
}

func InterquartileRange(arr []float64) float64 {
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

func PoissonDistribution(lambda float64, k uint64) float64 {
	return math.Pow(lambda, float64(k)) * math.Pow(math.E, -lambda) / float64(Factorial(k))
}

func NormalProbability(x float64, mean float64, stddev float64) float64 {
	return (1 + math.Erf((x-mean)/(stddev*math.Sqrt(2)))) / 2
}

func CentralLimit(limit float64, count uint64, mean float64, stddev float64) float64 {
	return NormalProbability(limit, float64(count)*mean, math.Sqrt(float64(count))*stddev)
}

func DetermineCentralLimitA(zScore float64, count uint64, mean float64, stddev float64) float64 {
	return (mean*float64(count) - zScore*stddev*math.Sqrt(float64(count))) / float64(count)
}

func DetermineCentralLimitB(zScore float64, count uint64, mean float64, stddev float64) float64 {
	return (mean*float64(count) + zScore*stddev*math.Sqrt(float64(count))) / float64(count)
}

func Correlation(x []float64, y []float64) float64 {
	var numerator float64

	ln := len(x)
	xMean := Mean(x)
	xStdDev := StdDeviation(x)
	yMean := Mean(y)
	yStdDev := StdDeviation(y)

	for i := 0; i < ln; i++ {
		numerator = numerator + (x[i]-xMean)*(y[i]-yMean)
	}

	return numerator / (float64(ln) * xStdDev * yStdDev)
}

func RankCorrelation(x []float64, y []float64) float64 {
	xrid := 1
	yrid := 1
	xRanks := make([]float64, len(x))
	yRanks := make([]float64, len(y))
	xRank := make(map[float64]int)
	yRank := make(map[float64]int)
	copy(xRanks, x)
	copy(yRanks, y)
	sort.Sort(sort.Float64Slice(xRanks))
	sort.Sort(sort.Float64Slice(yRanks))
	for i := 0; i < len(x); i++ {
		_, okx := xRank[xRanks[i]]
		if !okx {
			xRank[xRanks[i]] = xrid
			xrid++
		}

		_, oky := yRank[yRanks[i]]
		if !oky {
			yRank[yRanks[i]] = yrid
			yrid++
		}
	}

	for i := 0; i < len(x); i++ {
		xRanks[i] = float64(xRank[x[i]])
		yRanks[i] = float64(yRank[y[i]])
	}

	return Correlation(xRanks, yRanks)
}
