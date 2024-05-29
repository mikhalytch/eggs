package deref_test

import (
	"testing"

	"github.com/mikhalytch/eggs/deref"
	"github.com/mikhalytch/eggs/ptr"

	"github.com/stretchr/testify/require"
)

func Test_Or(t *testing.T) {
	type s string

	var sVal *s

	require.Equal(t, s(""), deref.Or(sVal, ""))
	require.Equal(t, s("123"), deref.Or(sVal, "123"))

	sVal = ptr.Of(s("abc"))
	require.Equal(t, s("abc"), deref.Or(sVal, "123"))
}
