package ptr

// ToSlice returns unary slices of optional value, otherwise the empty slice.
func ToSlice[T any](t *T) []T {
	if t == nil {
		return nil
	}

	return []T{*t}
}
