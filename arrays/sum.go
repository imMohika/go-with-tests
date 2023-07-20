package arrays

func Sum(numbers []int) int {
	sum := 0
	// range returns idx, value
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func SumAll(sets ...[]int) []int {
	sums := make([]int, len(sets))
	for i, set := range sets {
		sums[i] = Sum(set)
	}
	return sums
}
