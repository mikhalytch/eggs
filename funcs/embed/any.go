package embed

import baseConstraints "github.com/mikhalytch/eggs/constraints/base"

func String[T baseConstraints.String](t string) T { return T(t) }

func Bool[T baseConstraints.Bool](t bool) T { return T(t) }

func Complex64[T baseConstraints.Complex64](t complex64) T    { return T(t) }
func Complex128[T baseConstraints.Complex128](t complex128) T { return T(t) }

func Float64[T baseConstraints.Float64](t float64) T { return T(t) }
func Float32[T baseConstraints.Float32](t float32) T { return T(t) }

func Int[T baseConstraints.Int](t int) T       { return T(t) }
func Int64[T baseConstraints.Int64](t int64) T { return T(t) }
func Int32[T baseConstraints.Int32](t int32) T { return T(t) }
func Int16[T baseConstraints.Int16](t int16) T { return T(t) }
func Int8[T baseConstraints.Int8](t int8) T    { return T(t) }

func Uint[T baseConstraints.Uint](t uint) T          { return T(t) }
func Uint8[T baseConstraints.Uint8](t uint8) T       { return T(t) }
func Uint16[T baseConstraints.Uint16](t uint16) T    { return T(t) }
func Uint32[T baseConstraints.Uint32](t uint32) T    { return T(t) }
func Uint64[T baseConstraints.Uint64](t uint64) T    { return T(t) }
func Uintptr[T baseConstraints.Uintptr](t uintptr) T { return T(t) }
