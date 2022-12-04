package maps_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/mikhalytch/eggs/funcs/mapper"
	"github.com/mikhalytch/eggs/maps"
	"github.com/mikhalytch/eggs/strconv"
)

func TestMapKeys(t *testing.T) {
	t.Run("Identity", func(t *testing.T) {
		tests := []map[int]string{
			{},
			{1: "a", 2: "b"},
		}
		for i, test := range tests {
			t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
				require.Equal(t, test, maps.MapKeys(test, mapper.Identity[int]))
			})
		}
	})
	t.Run("StoA", func(t *testing.T) {
		tests := []struct {
			in   map[int]string
			want map[string]string
		}{
			{map[int]string{}, map[string]string{}},
			{map[int]string{1: "a", 2: "b"}, map[string]string{"1": "a", "2": "b"}},
		}
		for i, test := range tests {
			t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
				require.Equal(t, test.want, maps.MapKeys(test.in, strconv.StoA[int]))
			})
		}
	})
}

func TestMapValues(t *testing.T) {
	require.Equal(t, map[string]string{}, maps.MapValues(map[string]string{}, mapper.Identity[string]))
	require.Equal(t, map[string]string{"a": "1", "b": "20"},
		maps.MapValues(map[string]int8{"a": 1, "b": 20}, strconv.StoA[int8]))
}
