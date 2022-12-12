package funcs_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/mikhalytch/eggs/funcs"
	"github.com/mikhalytch/eggs/strconv"
)

func TestApplier_Apply(t *testing.T) {
	app := funcs.Applier[uint64, string](strconv.UtoA[uint64])
	require.Equal(t, "42", app(42))
	require.Equal(t, app(42), app.Apply(42))
}

func TestPanic(t *testing.T) {
	require.Panics(t, func() { panic("abc") })
}
