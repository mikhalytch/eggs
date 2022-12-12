package mapkeys_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/mikhalytch/eggs/funcs/mapper"
	"github.com/mikhalytch/eggs/maps/mapkeys"
)

func TestMap(t *testing.T) {
	t.Run("String mapper", func(t *testing.T) {
		type Name string
		require.ElementsMatch(t, []string{"k1", "k2"}, mapkeys.Map(map[Name]int{"k1": 1, "k2": 2}, mapper.String[Name]))
		require.ElementsMatch(t, []string{}, mapkeys.Map(map[Name]int{}, mapper.String[Name]))
	})
	t.Run("Identity mapper", func(t *testing.T) {
		require.ElementsMatch(t, []string{"k1", "k2"}, mapkeys.Map(map[string]int{"k1": 1, "k2": 2}, mapper.Identity[string]))
		require.ElementsMatch(t, []string{}, mapkeys.Map(map[string]int{}, mapper.Identity[string]))
	})
}
