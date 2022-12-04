package ptr

// Of is useful when one needs a one-liner for pointer of expression result;
// one needs to understand, the t will be passed as a copy, and `ptr.Of(x)` will result in pointer to copy of x.
func Of[T any](t T) *T { return &t }
