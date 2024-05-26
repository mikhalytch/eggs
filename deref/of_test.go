package deref_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/mikhalytch/eggs/deref"
	"github.com/mikhalytch/eggs/ptr"
)

func TestOf(t *testing.T) {
	t.Run("simple types", func(t *testing.T) {
		require.Equal(t, 2, deref.Of(ptr.Of(2)))
		require.Equal(t, 0, deref.Of[int](nil))
		require.Equal(t, uint(0), deref.Of[uint](nil))
		require.Equal(t, "", deref.Of[string](nil))
		require.Equal(t, false, deref.Of[bool](nil))
		require.NotEqual(t, true, deref.Of[bool](nil))
		require.Equal(t, false, deref.Of(ptr.Of(false)))
		require.Equal(t, true, deref.Of(ptr.Of(true)))
	})
	t.Run("struct", func(t *testing.T) {
		type A struct {
			v  int
			v1 string
		}

		a := A{1, "a"}
		require.Equal(t, a, deref.Of(&a))
		require.Equal(t, A{1, "a"}, deref.Of(&a))
		require.NotEqual(t, A{2, "a"}, deref.Of(&a))
	})
	t.Run("pointer", func(t *testing.T) {
		var unassigned *string

		require.Nil(t, deref.Of(&unassigned))
		require.Panics(t, func() {
			_ = *deref.Of(&unassigned)
		})

		assigned := ptr.Of("bad")
		require.Equal(t, "bad", *deref.OrDefault(&assigned))
	})
}
