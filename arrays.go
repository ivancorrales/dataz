package dataz

import "sort"

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
