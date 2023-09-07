package embed

import "github.com/mikhalytch/eggs/constraints"

func String[T constraints.String](t string) T { return T(t) }

func Bool[T constraints.Bool](t bool) T { return T(t) }

func Complex64[T constraints.Complex64](t complex64) T    { return T(t) }
func Complex128[T constraints.Complex128](t complex128) T { return T(t) }

func Float64[T constraints.Float64](t float64) T { return T(t) }
func Float32[T constraints.Float32](t float32) T { return T(t) }

func Int[T constraints.Int](t int) T       { return T(t) }
func Int64[T constraints.Int64](t int64) T { return T(t) }
func Int32[T constraints.Int32](t int32) T { return T(t) }
func Int16[T constraints.Int16](t int16) T { return T(t) }
func Int8[T constraints.Int8](t int8) T    { return T(t) }

func Uint[T constraints.Uint](t uint) T          { return T(t) }
func Uint8[T constraints.Uint8](t uint8) T       { return T(t) }
func Uint16[T constraints.Uint16](t uint16) T    { return T(t) }
func Uint32[T constraints.Uint32](t uint32) T    { return T(t) }
func Uint64[T constraints.Uint64](t uint64) T    { return T(t) }
func Uintptr[T constraints.Uintptr](t uintptr) T { return T(t) }
