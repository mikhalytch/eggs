package zap

import (
	"go.uber.org/zap"
)

// Int64 constructs a field with the given key and value.
func Int64[T ~int64](key string, val T) zap.Field { return zap.Int64(key, int64(val)) }

// String constructs a field with the given key and value.
func String[T ~string](key string, val T) zap.Field { return zap.String(key, string(val)) }
