package strconv_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/mikhalytch/eggs/strconv"
)

func TestStoA(t *testing.T) {
	req := require.New(t)
	req.Equal("1", strconv.StoA(1))
	req.Equal("1", strconv.StoA[int16](1))
	req.Equal("-2", strconv.StoA[int32](-2))
}
