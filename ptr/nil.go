package ptr

func Nil[T any]() *T  { return nil }
func None[T any]() *T { return Nil[T]() } // alias
