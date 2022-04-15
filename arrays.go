package dataz

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
