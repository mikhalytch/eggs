package str_test

import (
	"reflect"
	"testing"

	"github.com/mikhalytch/eggs/funcs/mapper/str"
	"github.com/mikhalytch/eggs/slices"

	"github.com/stretchr/testify/require"
)

func TestPrepend(t *testing.T) {
	type testType string

	in := []testType{"a", "b"}
	got := slices.Map(str.Prepend(testType("1")))(in)
	require.True(t, reflect.DeepEqual([]testType{"1a", "1b"}, got))
}
