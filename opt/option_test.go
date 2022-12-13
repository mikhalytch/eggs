package opt_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/mikhalytch/eggs/funcs/predicate"
	"github.com/mikhalytch/eggs/opt"
	"github.com/mikhalytch/eggs/ptr"
	strconv2 "github.com/mikhalytch/eggs/strconv"
	"github.com/mikhalytch/eggs/tuple"
)

func TestOfPtr(t *testing.T) {
	require.False(t, opt.OfPtr[any](nil).IsDefined())
	require.True(t, opt.OfPtr[string](nil).IsEmpty())
	require.True(t, opt.OfPtr(ptr.Of(2)).IsDefined())
	require.False(t, opt.OfPtr(ptr.Of("string")).IsEmpty())
}

func TestSomething(t *testing.T) {
	type Number int

	require.Equal(t, opt.Some(1), opt.Some(1))
	require.Equal(t, opt.Some("string"), opt.Some("string"))
	require.Equal(t, opt.Some(Number(100)), opt.Some(Number(100)))
}

func TestNothing(t *testing.T) {
	require.Equal(t, opt.None[any](), opt.None[any]())
	require.Equal(t, opt.None[int](), opt.None[int]())
	require.Equal(t, opt.None[string](), opt.None[string]())
}

func TestToSlice(t *testing.T) {
	require.Equal(t, []int{1}, opt.Some(1).ToSlice())
	require.Equal(t, []int{}, opt.None[int]().ToSlice())
}

func TestGet(t *testing.T) {
	require.Equal(t, tuple.Err(1, nil), tuple.Err(opt.Some(1).Get()))
	require.Error(t, tuple.Err(opt.None[int]().Get()).T2())
}

func TestOrElse(t *testing.T) {
	require.Equal(t, 1, opt.Some(1).OrElse(2))
	require.Equal(t, 1, opt.None[int]().OrElse(1))
}

func TestNone_OrDefault(t *testing.T) {
	require.Equal(t, "", opt.None[string]().OrDefault())
	require.Equal(t, 0.0, opt.None[float64]().OrDefault())
	require.Equal(t, float32(0.0), opt.None[float32]().OrDefault())

	type Name struct{ s string }

	require.Equal(t, Name{s: ""}, opt.None[Name]().OrDefault())
}

func TestSome_OrDefault(t *testing.T) {
	require.Equal(t, "abc", opt.Some("abc").OrDefault())
	require.Equal(t, 42.1, opt.Some(42.1).OrDefault())

	type Name struct{ s string }

	require.Equal(t, Name{"eggs"}, opt.Some(Name{s: "eggs"}).OrDefault())
}

func TestFilter(t *testing.T) {
	require.Equal(t, opt.None[int](), opt.Some(1).Filter(func(i int) bool { return i > 1 }))
	require.Equal(t, opt.None[int](), opt.Some(1).Filter(predicate.Never[int]))
	require.Equal(t, opt.None[int](), opt.None[int]().Filter(predicate.Always[int]))
	require.Equal(t, opt.Some(1), opt.Some(1).Filter(func(i int) bool { return i == 1 }))
	require.Equal(t, opt.Some(1), opt.Some(1).Filter(predicate.Always[int]))
}

func TestSome_Map(t *testing.T) {
	require.Equal(t, opt.Some(2), opt.Some(1).Map(func(i int) int { return i * 2 }))
}

func TestNone_Map(t *testing.T) {
	require.Equal(t, opt.None[int](), opt.None[int]().Map(func(i int) int { return i * 2 }))
}

func TestSome_FlatMap(t *testing.T) {
	require.Equal(t, opt.Some("abc"), opt.Some("abc").FlatMap(opt.Some[string]))
	require.Equal(t, opt.Some("1"),
		opt.FlatMap(opt.Some(1), func(i int) opt.Option[string] { return opt.Some(strconv.Itoa(i)) }))
}

func TestNone_FlatMap(t *testing.T) {
	require.Equal(t, opt.None[string](), opt.None[string]().FlatMap(opt.Some[string]))
	require.Equal(t, opt.None[string](),
		opt.FlatMap(opt.None[int](), func(i int) opt.Option[string] { return opt.Some(strconv.Itoa(i)) }))
}

func TestNone_Foreach(t *testing.T) {
	ctr := 0

	opt.None[string]().Foreach(func(s string) { ctr++ })
	require.Equal(t, 0, ctr)
}

func TestSome_Foreach(t *testing.T) {
	ctr := 0

	opt.Some[string]("abc").Foreach(func(s string) { ctr++ })
	require.Equal(t, 1, ctr)
}

func TestFold(t *testing.T) {
	require.Equal(t, "1", opt.Fold(opt.Some(1), strconv2.StoA[int], func() string { return "" }))
	require.Equal(t, "", opt.Fold(opt.None[int](), strconv2.StoA[int], func() string { return "" }))
	require.Equal(t, "1", opt.Fold(opt.None[int](), strconv2.StoA[int], func() string { return "1" }))
}
