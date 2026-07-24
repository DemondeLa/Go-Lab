package slice

import (
	"reflect"
	"testing"
)

func TestIntersect(t *testing.T) {
	testCases := []struct {
		name string
		a    []int
		b    []int
		want []int
	}{
		{"both empty", []int{}, []int{}, []int{}},
		{"empty first slice", []int{}, []int{1, 2, 3}, []int{}},
		{"empty second slice", []int{1, 2, 3}, []int{}, []int{}},
		{"no intersection", []int{1, 2, 3}, []int{4, 5, 6}, []int{}},
		{"all elements intersect", []int{1, 2, 3}, []int{3, 2, 1}, []int{1, 2, 3}},
		{"remove duplicates and preserve first slice order", []int{3, 1, 2, 2, 4}, []int{2, 3, 3, 5}, []int{3, 2}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := Intersect(tc.a, tc.b)
			if !reflect.DeepEqual(res, tc.want) {
				t.Errorf("Intersect(%v, %v) = %v, want %v",
					tc.a, tc.b, res, tc.want)
			}
		})
	}
}

func TestUnion(t *testing.T) {
	testCases := []struct {
		name string
		a    []int
		b    []int
		want []int
	}{
		{"both empty", []int{}, []int{}, []int{}},
		{"empty first slice", []int{}, []int{1, 2, 2, 3}, []int{1, 2, 3}},
		{"empty second slice", []int{1, 1, 2, 3}, []int{}, []int{1, 2, 3}},
		{"disjoint slices", []int{1, 2, 3}, []int{4, 5, 6}, []int{1, 2, 3, 4, 5, 6}},
		{"remove duplicates and preserve first appearance order", []int{3, 1, 3, 2}, []int{2, 4, 1, 5}, []int{3, 1, 2, 4, 5}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := Union(tc.a, tc.b)
			if !reflect.DeepEqual(res, tc.want) {
				t.Errorf("Union(%v, %v) = %v, want %v",
					tc.a, tc.b, res, tc.want)
			}
		})
	}
}

func TestDiff(t *testing.T) {
	testCases := []struct {
		name string
		a    []int
		b    []int
		want []int
	}{
		{"both empty", []int{}, []int{}, []int{}},
		{"empty first slice", []int{}, []int{1, 2, 3}, []int{}},
		{"empty second slice", []int{1, 1, 2, 3}, []int{}, []int{1, 2, 3}},
		{"no difference", []int{1, 2, 3}, []int{3, 2, 1}, []int{}},
		{"all elements differ", []int{1, 2, 3}, []int{4, 5, 6}, []int{1, 2, 3}},
		{"remove excluded values and duplicates", []int{3, 1, 2, 1, 4, 3}, []int{2, 5}, []int{3, 1, 4}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := Diff(tc.a, tc.b)
			if !reflect.DeepEqual(res, tc.want) {
				t.Errorf("Diff(%v, %v) = %v, want %v",
					tc.a, tc.b, res, tc.want)
			}
		})
	}
}
