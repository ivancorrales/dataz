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
		input    string{}
		expected string{}
		fn       func(string) bool
	}{
		{
			input:    []string{"home", "port", "table"},
			expected: []string{"home"},
			fn:       func(item string) bool { return !strings.Contains(item, "t") },
		},
	}

	for _, c := range cases {
		output := DropWhile(c.input, c.fn)
		assert.Equal(t, output, c.expected)
	}
}
