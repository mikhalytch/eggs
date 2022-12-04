package math_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/mikhalytch/eggs/math"
)

func TestMax(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		tests := []struct {
			a, b int
			want int
		}{
			{1, 2, 2},
			{3, 2, 3},
			{-3, 2, 2},
		}
		for i, test := range tests {
			t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
				got := math.Max(test.a, test.b)
				require.Equal(t, test.want, got)
			})
		}
	})
	t.Run("string", func(t *testing.T) {
		tests := []struct {
			a, b string
			want string
		}{
			{"a", "b", "b"},
			{"string", "strings", "strings"},
		}
		for i, test := range tests {
			t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
				got := math.Max(test.a, test.b)
				require.Equal(t, test.want, got)
			})
		}
	})
}
