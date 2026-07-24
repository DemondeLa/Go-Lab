package slice

import "testing"

func TestContains(t *testing.T) {
	testCases := []struct {
		name   string
		s      []int
		target int
		want   bool
	}{
		{"empty slice", []int{}, 1, false},
		{"contains head", []int{1, 2, 3, 4, 5}, 1, true},
		{"contains middle", []int{1, 2, 3, 4, 5}, 3, true},
		{"contains tail", []int{1, 2, 3, 4, 5}, 5, true},
		{"target appears more than once", []int{1, 2, 2, 3}, 2, true},
		{"does not contain target", []int{1, 2, 3, 4, 5}, 6, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := Contains(tc.s, tc.target)
			if res != tc.want {
				t.Errorf("Contains(%v, %v) = %v, want %v",
					tc.s, tc.target, res, tc.want)
			}
		})
	}
}

func TestContainsFunc(t *testing.T) {
	type item struct {
		name string
		tags []string
	}

	testCases := []struct {
		name  string
		s     []item
		match func(item) bool
		want  bool
	}{
		{
			"empty slice",
			[]item{},
			func(item) bool { return true },
			false,
		},
		{
			"matching element",
			[]item{{name: "first"}, {name: "second"}, {name: "third"}},
			func(v item) bool { return v.name == "second" },
			true,
		},
		{
			"no matching element",
			[]item{{name: "first"}, {name: "second"}, {name: "third"}},
			func(v item) bool { return v.name == "fourth" },
			false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := ContainsFunc(tc.s, tc.match)
			if res != tc.want {
				t.Errorf("ContainsFunc(%v, match) = %v, want %v",
					tc.s, res, tc.want)
			}
		})
	}
}

func TestIndex(t *testing.T) {
	testCases := []struct {
		name   string
		s      []int
		target int
		want   int
	}{
		{"empty slice", []int{}, 1, -1},
		{"find head", []int{1, 2, 3, 4, 5}, 1, 0},
		{"find middle", []int{1, 2, 3, 4, 5}, 3, 2},
		{"find tail", []int{1, 2, 3, 4, 5}, 5, 4},
		{"return first matching index", []int{1, 2, 2, 3}, 2, 1},
		{"target not found", []int{1, 2, 3, 4, 5}, 6, -1},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := Index(tc.s, tc.target)
			if res != tc.want {
				t.Errorf("Index(%v, %v) = %v, want %v",
					tc.s, tc.target, res, tc.want)
			}
		})
	}
}

func TestIndexFunc(t *testing.T) {
	type item struct {
		name string
		tags []string
	}

	testCases := []struct {
		name  string
		s     []item
		match func(item) bool
		want  int
	}{
		{
			"empty slice",
			[]item{},
			func(item) bool { return true },
			-1,
		},
		{
			"find matching element",
			[]item{{name: "first"}, {name: "second"}, {name: "third"}},
			func(v item) bool { return v.name == "second" },
			1,
		},
		{
			"return first matching index",
			[]item{{tags: []string{"go"}}, {tags: []string{"go"}}, {tags: []string{"solidity"}}},
			func(v item) bool { return len(v.tags) > 0 && v.tags[0] == "go" },
			0,
		},
		{
			"no matching element",
			[]item{{name: "first"}, {name: "second"}, {name: "third"}},
			func(v item) bool { return v.name == "fourth" },
			-1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := IndexFunc(tc.s, tc.match)
			if res != tc.want {
				t.Errorf("IndexFunc(%v, match) = %v, want %v",
					tc.s, res, tc.want)
			}
		})
	}
}
