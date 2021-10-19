package unique

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUniqueStrings(t *testing.T) {

	unique := NewUnique()
	strings := []string{"four", "one", "two", "three", "three", "four"}
	bools := []bool{true, true, true, true, false, false}

	Convey("Given new Unique, when checking strings the values are as expected", t, func() {
		for c, s := range strings {
			v := unique.IsUnique(s)
			So(v, ShouldEqual, bools[c])
		}

	})
}

func TestUniqueInts(t *testing.T) {

	unique := NewUnique()
	ints := []int{4, 1, 2, 3, 3, 4}
	bools := []bool{true, true, true, true, false, false}

	Convey("Given new Unique, when checking ints the values are as expected", t, func() {
		for c, s := range ints {
			v := unique.IsUnique(s)
			So(v, ShouldEqual, bools[c])
		}

	})
}

func TestUniqueMixed(t *testing.T) {

	unique := NewUnique()
	strings := []string{"four", "one", "two", "three", "three", "four"}
	ints := []int{4, 1, 2, 3, 3, 4}
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

	})
}
