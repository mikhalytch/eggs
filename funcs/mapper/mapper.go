package mapper

func String[A ~string](a A) string { return string(a) }
func Identity[A any](a A) A        { return a }
