package ptr_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/mikhalytch/eggs/ptr"
)

func TestOf(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		require.Equal(t, 1, *ptr.Of(1))
		require.Equal(t, "a", *ptr.Of("a"))

		type Name string

		require.Equal(t, Name("abc"), *ptr.Of(Name("abc")))
		require.Equal(t, Name("abc"), *ptr.Of[Name]("abc"))
	})
	t.Run("result is a link to a copy", func(t *testing.T) {
		t.Run("even with a pointer", func(t *testing.T) {
			v := "some value, pointer to which is under test" // string
			value := &v                                       // *string
			pointerToAValue := &value                         // **string
			pointerToACopyOfAValue := ptr.Of(value)           // **string
			copyOfAValue := *pointerToACopyOfAValue           // *string

			require.False(t, pointerToACopyOfAValue == pointerToAValue,
				"these are two different pointers, pointing onto two different pointers on v")
			require.True(t, value == copyOfAValue,
				"in go, pointers considered equal, if both nil OR point to the same variable")
			require.True(t, *value == *copyOfAValue,
				"same as above")
		})
	})
}
