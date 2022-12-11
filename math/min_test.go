package math_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/mikhalytch/eggs/math"
	"github.com/mikhalytch/eggs/slices"
)

func TestMin(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		tests := []struct {
			a, b int
			want int
		}{
			{1, 2, 1},
			{3, 2, 2},
			{-3, 2, -3},
		}
		for i, test := range tests {
			t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
				got := math.Min(test.a, test.b)
				require.Equal(t, test.want, got)
			})
		}
	})
	t.Run("string", func(t *testing.T) {
		tests := []struct {
			a, b string
			want string
		}{
			{"a", "b", "a"},
			{"string", "strings", "string"},
		}
		for i, test := range tests {
			t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
				got := math.Min(test.a, test.b)
				require.Equal(t, test.want, got)
			})
		}
	})
	t.Run("float with slices", func(t *testing.T) {
		tests := []struct {
			elems []float32
			want  float32
		}{
			{[]float32{0.1, 2, 3, 4, 2.3, 5.6}, 0.1},
		}
		for i, test := range tests {
			t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
				got := math.Min(slices.Head(test.elems), slices.Tail(test.elems)...)
				require.Equal(t, test.want, got)
			})
		}
	})
}
