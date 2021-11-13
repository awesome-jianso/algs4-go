package sorting_test

import (
	"testing"

	"github.com/youngzhu/algs4-go/sorting"
	. "github.com/youngzhu/algs4-go/sorting"
)

func TestHeapsortInts(t *testing.T) {
	data := ints
	x := sorting.IntSortSlice(data[0:])
	Heapsort(x)
	if !sorting.IsSorted(x) {
		t.Errorf("sorting %v", ints)
		t.Errorf("    got %v", data)
	}
}

func TestHeapsortFloat64s(t *testing.T) {
	data := float64s
	x := sorting.Float64SortSlice(data[0:])
	Heapsort(x)
	if !sorting.IsSorted(x) {
		t.Errorf("sorting %v", float64s)
		t.Errorf("    got %v", data)
	}
}

func TestHeapsortStrings(t *testing.T) {
	data := strings
	x := sorting.StringSortSlice(data[0:])
	Heapsort(x)
	if !sorting.IsSorted(x) {
		t.Errorf("sorting %v", strings)
		t.Errorf("    got %v", data)
	}
}
