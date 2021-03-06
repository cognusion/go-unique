package unique

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUniqueStrings(t *testing.T) {

	unique := New()
	strings := []string{"four", "one", "two", "three", "three", "four"}
	ustrings := []interface{}{"four", "one", "two", "three"}
	bools := []bool{true, true, true, true, false, false}

	Convey("Given new Unique, when checking strings the values are as expected", t, func() {
		for c, s := range strings {
			v := unique.IsUnique(s)
			So(v, ShouldEqual, bools[c])
		}
		So(len(unique.Things()), ShouldEqual, len(ustrings))
	})
}

func TestUniqueInts(t *testing.T) {

	unique := New()
	ints := []int{4, 1, 2, 3, 3, 4}
	uints := []int{4, 1, 2, 3}
	bools := []bool{true, true, true, true, false, false}

	Convey("Given new Unique, when checking ints the values are as expected", t, func() {
		for c, s := range ints {
			v := unique.IsUnique(s)
			So(v, ShouldEqual, bools[c])
		}
		So(len(unique.Things()), ShouldEqual, len(uints))
	})
}

func TestUniqueMixed(t *testing.T) {

	unique := New()
	strings := []string{"four", "one", "two", "three", "three", "four"}
	ustrings := []interface{}{"four", "one", "two", "three"}
	ints := []int{4, 1, 2, 3, 3, 4}
	uints := []int{4, 1, 2, 3}
	bools := []bool{true, true, true, true, false, false, true, true, true, true, false, false}

	Convey("Given new Unique, when checking ints the values are as expected", t, func() {
		for c, s := range ints {
			v := unique.IsUnique(s)
			So(v, ShouldEqual, bools[c])
		}
		for c, s := range strings {
			v := unique.IsUnique(s)
			So(v, ShouldEqual, bools[c])
		}
		So(len(unique.Things()), ShouldEqual, len(uints)+len(ustrings))
	})
}

func BenchmarkIsUniqueTrue(b *testing.B) {
	unique := New()

	for i := 0; i < b.N; i++ {
		unique.IsUnique((i + 1) * b.N)
	}
}

func BenchmarkIsUniqueFalse(b *testing.B) {
	unique := New()
	unique.IsUnique("blah")
	for i := 0; i < b.N; i++ {
		unique.IsUnique("blah")
	}
}

/*
func BenchmarkIsUniqueTrueM(b *testing.B) {
	unique := NewMap()

	for i := 0; i < b.N; i++ {
		unique.IsUnique((i + 1) * b.N)
	}
}

func BenchmarkIsUniqueFalseM(b *testing.B) {
	unique := NewMap()
	unique.IsUnique("blah")
	for i := 0; i < b.N; i++ {
		unique.IsUnique("blah")
	}
}
*/
