package predicate

func NonNil[T any](t *T) bool { return t != nil }
func Nil[T any](t *T) bool    { return t == nil }
