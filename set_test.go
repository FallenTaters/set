package set_test

import (
	"sort"
	"strings"
	"testing"

	"github.com/FallenTaters/set"
)

func TestSet(t *testing.T) {
	intSet := set.New(1, 2)

	intSet.Add(3)

	intSet.AddAll(2, 3, 4)

	if intSet.Has(5) {
		t.Error(`contains 5 but shouldn't`)
	}
	if !intSet.Has(1) {
		t.Error(`doesn't contain 1 but should`)
	}
	intSet.Remove(1)
	if intSet.Has(1) {
		t.Error(`contains 1 but shouldn't`)
	}

	vals := intSet.Values()
	sort.Ints(vals)
	if vals[0] != 2 || vals[1] != 3 || vals[2] != 4 {
		t.Errorf(`%#v should be []int{2, 3, 4}`, vals)
	}

	intSet2 := make(set.Set[int])
	intSet2.Add(3)
	intSet.Add(5)

	intersection := set.Intersect(intSet, intSet2)
	if len(intersection) != 1 || !intersection.Has(3) {
		t.Errorf(`%#v should be set.New[int](3)`, intersection)
	}

	union := set.Union(intSet, intSet2)
	if len(union) != 4 {
		t.Errorf(`%#v should be set.New[int](2, 3, 4, 5)`, union)
	}

	if !union.HasAll(2, 3, 4, 5) || union.HasAll(1, 2, 3, 4, 5) {
		t.Error(`HasAll is not working`)
	}

	if !union.HasAny(4) || union.HasAny(10) {
		t.Error(`HasAny is not working`)
	}

	goString := intSet.GoString()
	if !strings.HasPrefix(goString, `set.New[int](`) || !strings.HasSuffix(goString, `)`) || len(goString) != 24 {
		t.Errorf(`%q should have been "set.New[int](2, 3, 4, 5) or similar"`, intSet.GoString())
	}
}
