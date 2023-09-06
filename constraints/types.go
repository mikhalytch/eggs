package constraints

type (
	String interface{ ~string }

	Complex64  interface{ ~complex64 }
	Complex128 interface{ ~complex128 }

	Float64 interface{ ~float64 }
	Float32 interface{ ~float32 }

	Int   interface{ ~int }
	Int64 interface{ ~int64 }
	Int32 interface{ ~int32 }
	Int16 interface{ ~int16 }
	Int8  interface{ ~int8 }

	Uint    interface{ ~uint }
	Uint8   interface{ ~uint8 }
	Uint16  interface{ ~uint16 }
	Uint32  interface{ ~uint32 }
	Uint64  interface{ ~uint64 }
	Uintptr interface{ ~uintptr }
)
