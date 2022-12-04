package strconv_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/mikhalytch/eggs/strconv"
)

func TestUtoA(t *testing.T) {
	t.Run("uint8", func(t *testing.T) {
		tests := []struct {
			in   uint8
			want string
		}{
			{2, "2"},
		}
		for i, test := range tests {
			t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
				got := strconv.UtoA(test.in)
				require.Equal(t, test.want, got)
			})
		}
	})
}
