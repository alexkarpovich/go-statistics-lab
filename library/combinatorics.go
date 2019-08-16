package library

func Factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}

func Combinations(n uint64, r uint64) uint64 {
	return Factorial(n) / (Factorial(r) * Factorial(n-r))
}
