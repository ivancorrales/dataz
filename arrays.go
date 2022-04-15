package dataz

import (
	"math"
	"sort"
)

func Reverse[T any](input []T) []T {
	for first, last := 0, len(input)-1; first < last; first++ {
		input[first], input[last] = input[last], input[first]
		last--
	}
	return input
}

func DropWhile[T any](input []T, fn func(T) bool) []T {
	for i, element := range input {
		if !fn(element) {
			input = input[:i]
			return input
		}
	}
	return input
}

func Filter[T any](input []T, fn func(T) bool) []T {
	filtered := make([]T, 0)
	for _, element := range input {
		if fn(element) {
			filtered = append(filtered, element)
		}
	}
	return filtered
}

func Drop[T any](input []T, indexes ...int) []T {
	sort.Ints(indexes)
	limit := len(input) - 1
	for i, index := range indexes {
		if index < 0 || index > limit {
			indexes = Drop(indexes, i)
			continue
		}
		if i > 0 && index <= indexes[i-1] {
			indexes = Drop(indexes, i)
		}
	}
	for rIndex, index := range indexes {
		input = append(input[:index-rIndex], input[index-rIndex+1:]...)
	}
	return input
}

func Duplicates[T comparable](input []T) []T {
	duplicates := make([]T, 0)
	inputSet := make(map[T]struct{})
	for _, element := range input {
		if _, ok := inputSet[element]; ok {
			duplicates = append(duplicates, element)
			continue
		}
		inputSet[element] = struct{}{}
	}
	return duplicates
}

func ForEach[T any](input []T, fn func(item T)) {
	for _, element := range input {
		fn(element)
	}
}

func Concat[T any](input [][]T) []T {
	output := make([]T, 0)
	for i := range input {
		output = append(output, input[i]...)
	}
	return output
}

func Chunk[T any](input []T, size int) [][]T {
	totalGroups := len(input) / size
	if len(input)%size != 0 {
		totalGroups++
	}
	groups := make([][]T, totalGroups)
	for i := 0; i < totalGroups; i++ {
		groups[i] = input[i*size : int(math.Min(float64((i+1)*size), float64(len(input))))]
	}
	return groups
}

func Difference[T comparable](input, input2 []T) []T {
	output := make([]T, 0)
	set, set2 := make(map[T]struct{}), make(map[T]struct{})
	for _, element := range input {
		set[element] = struct{}{}
	}
	for _, element := range input2 {
		set2[element] = struct{}{}
	}
	for i := range input2 {
		val := input2[i]
		if _, ok := set[val]; !ok {
			output = append(output, val)
		}
	}
	for i := range input {
		val := input[i]
		if _, ok := set2[val]; !ok {
			output = append(output, val)
		}
	}
	return output
}

func Flatten[T any](input [][]T) []T {
	out := make([]T, 0)
	for i := range input {
		out = append(out, input[i]...)
	}
	return out
}
