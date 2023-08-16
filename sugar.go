package sugar

import "golang.org/x/exp/constraints"

// Contains returns true if array of K contains e
func Contains[K constraints.Ordered](s []K, e K) bool {
	for _, elem := range s {
		if elem == e {
			return true
		}
	}
	return false
}

// IndexesOf returns []int with indexes of occurrences e in array s
func IndexesOf[K constraints.Ordered](s []K, e K) (indexes []int) {
	for i, elem := range s {
		if elem == e {
			indexes = append(indexes, i)
		}
	}
	return
}

// Map returns array with every element of s passed through the mapFunc
func Map[K any](s []K, mapFunc func(K) K) (result []K) {
	for _, elem := range s {
		result = append(result, mapFunc(elem))
	}
	return
}

// Filter returns array with elements of s if filterFunc returned true for according element
func Filter[K any](s []K, filterFunc func(K) bool) (result []K) {
	for _, elem := range s {
		if filterFunc(elem) {
			result = append(result, elem)
		}
	}
	return
}

// Reduce returns one K var that is result of cumulative execution of reduceFunc on every element of s
func Reduce[K any](s []K, reduceFunc func(K, K) K) (result K) {
	if len(s) == 0 {
		return
	}
	result = s[0]
	if len(s) == 1 {
		return
	}
	for i := 1; i < len(s); i++ {
		result = reduceFunc(result, s[i])
	}
	return
}

// Flatten returns slice based on matrix values
func Flatten[K any](s [][]K) (result []K) {
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s[i]); j++ {
			result = append(result, s[i][j])
		}
	}
	return
}

// Min returns min element of array
func Min[K constraints.Ordered](s []K) (result K) {
	if len(s) == 0 {
		return
	}
	result = s[0]
	for _, elem := range s {
		if elem < result {
			result = elem
		}
	}
	return
}

// Max returns max element of array
func Max[K constraints.Ordered](s []K) (result K) {
	if len(s) == 0 {
		return
	}
	result = s[0]
	for _, elem := range s {
		if elem > result {
			result = elem
		}
	}
	return
}

// Keys returns array of T containing keys of map[T]K
func Keys[T constraints.Ordered, K any](m map[T]K) (result []T) {
	result = make([]T, len(m))
	i := 0
	for key, _ := range m {
		result[i] = key
		i++
	}
	return
}

// Values returns array of T containing values of map[T]K
func Values[T constraints.Ordered, K any](m map[T]K) (result []K) {
	result = make([]K, len(m))
	i := 0
	for _, value := range m {
		result[i] = value
		i++
	}
	return
}

// SplitByN splits a slice s into sub-slices of maximum size n.
// Returns nil if n is less than or equal to 0.
func SplitByN[T any](s []T, n int) [][]T {
	if n <= 0 || len(s) == 0 {
		return nil // Invalid chunk size or empty input slice
	}

	// Calculate the number of chunks needed, rounded up
	count := (len(s) + n - 1) / n
	result := make([][]T, count)

	for i := 0; i < count; i++ {
		start := i * n
		end := start + n

		// Ensure the last chunk doesn't exceed slice bounds
		if end > len(s) {
			end = len(s)
		}

		result[i] = s[start:end]
	}

	return result
}
