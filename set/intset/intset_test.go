// set provides a very implementation of sets for a number of basic types.
package set

import (
	"sort"
	"testing"
)

func Test_set(t *testing.T) {

	a := []int{1, 2, 3, 4}
	aSet := NewIntSet()
	aSet.Add(a...)
	cSet := aSet.Clone()

	if !aSet.Equals(cSet) {
		t.Error("set equality does not work properly.")
	}

	b := []int{3, 4, 5, 6}
	bSet := NewIntSet()
	bSet.Add(b...)

	if !aSet.Contains(3) || aSet.Contains(10) {
		t.Error("set containment does not work properly.")
	}

	aList := sort.IntSlice(aSet.ToList())
	aList.Sort()
	for i, l := range a {
		if aList[i] != l {
			t.Error("list representation of aSet does not match a")
		}
	}

	bList := sort.IntSlice(bSet.ToList())
	bList.Sort()
	for i, l := range bList {
		if b[i] != l {
			t.Error("list representation of bSet does not match b")
		}
	}

	aSet.Union(bSet)
	aList = sort.IntSlice(aSet.ToList())
	aList.Sort()
	for i, l := range []int{1, 2, 3, 4, 5, 6} {
		if aList[i] != l {
			t.Error("union operation on aSet and bSet failed")
		}
	}

	aSet.Complement(bSet)
	aList = sort.IntSlice(aSet.ToList())
	aList.Sort()
	for i, l := range []int{1, 2} {
		if aList[i] != l {
			t.Error("complement operation on aSet and bSet failed")
		}
	}

	aSet.Intersect(bSet)
	aList = sort.IntSlice(aSet.ToList())
	aList.Sort()
	if len(aList) != 0 {
		t.Error("intersect operation on aSet and bSet failed")
	}
}
