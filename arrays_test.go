package dataz

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	cases := []struct {
		input    []any
		expected []any
	}{
		{
			input:    []any{"a", "b", "c"},
			expected: []any{"c", "b", "a"},
		},
		{
			input:    []any{20, 1, .3, true, "a"},
			expected: []any{"a", true, .3, 1, 20},
		},
	}
	for _, c := range cases {
		output := Reverse(c.input)
		assert.Equal(t, c.expected, output)
	}
}

func TestDropWhile_string(t *testing.T) {
	cases := []struct {
		input    []string
		expected []string
		fn       func(string) bool
	}{
		{
			input:    []string{"home", "port", "table", "window"},
			expected: []string{"home"},
			fn:       func(item string) bool { return !strings.Contains(item, "t") },
		},
	}

	for _, c := range cases {
		output := DropWhile(c.input, c.fn)
		assert.Equal(t, output, c.expected)
	}
}

func TestFilter_string(t *testing.T) {
	cases := []struct {
		input    []string
		expected []string
		fn       func(string) bool
	}{
		{
			input:    []string{"home", "port", "table", "window"},
			expected: []string{"home", "window"},
			fn:       func(item string) bool { return !strings.Contains(item, "t") },
		},
	}

	for _, c := range cases {
		output := Filter(c.input, c.fn)
		assert.Equal(t, output, c.expected)
	}
}

func TestDrop(t *testing.T) {
	cases := []struct {
		input    []any
		expected []any
		indexes  []int
	}{
		{
			input:    []any{"home", "port", "table", "window"},
			expected: []any{"home", "window"},
			indexes:  []int{1, 2},
		},
		{
			input:    []any{"home", "port", "table", "window"},
			expected: []any{"home", "port", "table", "window"},
			indexes:  []int{},
		},
		{
			input:    []any{"home", "port", "table", "window"},
			expected: []any{"home", "port", "table", "window"},
			indexes:  []int{-2},
		},
		{
			input:    []any{"home", "port", "table", "window"},
			expected: []any{"home", "port", "table", "window"},
			indexes:  []int{20},
		},
		{
			input:    []any{"home", "port", "table", "window"},
			expected: []any{"home", "port", "window"},
			indexes:  []int{20, 2},
		},
		{
			input:    []any{"home", "port", "table", "window"},
			expected: []any{"home", "window"},
			indexes:  []int{2, 1},
		},
	}

	for _, c := range cases {
		output := Drop(c.input, c.indexes...)
		assert.Equal(t, output, c.expected)
	}
}

func TestDuplicates_string(t *testing.T) {
	cases := []struct {
		input    []string
		expected []string
	}{
		{
			input:    []string{"home", "port", "window", "home", "table", "window"},
			expected: []string{"home", "window"},
		},
		{
			input:    []string{"home", "port", "table", "window", "window", "table", "port", "home"},
			expected: []string{"window", "table", "port", "home"},
		},
		{
			input:    []string{"home", "port", "table", "window"},
			expected: []string{},
		},
	}

	for _, c := range cases {
		output := Duplicates(c.input)
		assert.EqualValues(t, output, c.expected)
	}
}
