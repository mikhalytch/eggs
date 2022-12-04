package zap_test

import (
	"bytes"
	"testing"

	"github.com/fgrosse/zaptest"
	"github.com/stretchr/testify/require"

	gzap "github.com/mikhalytch/eggs/zap"
)

func TestString(t *testing.T) {
	type Name string

	bb := bytes.NewBuffer(make([]byte, 0))
	logger := zaptest.LoggerWriter(bb)

	logger.Info("test msg", gzap.String("key", Name("name")))
	require.Equal(t, `INFO	test msg	{"key": "name"}`+"\n", bb.String())
}

func TestInt64(t *testing.T) {
	type Value int64

	bb := bytes.NewBuffer(make([]byte, 0))
	logger := zaptest.LoggerWriter(bb)

	logger.Info("test msg", gzap.Int64("key", Value(2)))
	require.Equal(t, `INFO	test msg	{"key": 2}`+"\n", bb.String())
}
