package slice

import (
	"errors"
	"reflect"
	"testing"
)

func TestDelete(t *testing.T) {
	testCases := []struct {
		name    string
		s       []int
		idx     int
		want    []int
		wantErr error
	}{
		{"empty slice", []int{}, 0, []int{}, ErrIndexOutOfRange},
		{"negative index", []int{1, 2, 3, 4, 5}, -1, []int{1, 2, 3, 4, 5}, ErrIndexOutOfRange},
		{"index out of range", []int{1, 2, 3, 4, 5}, 5, []int{1, 2, 3, 4, 5}, ErrIndexOutOfRange},
		{"delete head", []int{1, 2, 3, 4, 5}, 0, []int{2, 3, 4, 5}, nil},
		{"delete tail", []int{1, 2, 3, 4, 5}, 4, []int{1, 2, 3, 4}, nil},
		{"delete middle", []int{1, 2, 3, 4, 5}, 2, []int{1, 2, 4, 5}, nil},
		{"single element", []int{1}, 0, []int{}, nil},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			slice := tc.s
			res, err := Delete(tc.s, tc.idx)
			if !errors.Is(err, tc.wantErr) {
				t.Errorf("Delete(%v, %v) = %v, want %v",
					slice, tc.idx, err, tc.wantErr)
			}
			if err != nil {
				return
			}
			if !reflect.DeepEqual(res, tc.want) {
				t.Errorf("Delete(%v, %v) = %v, want %v",
					slice, tc.idx, res, tc.want)
			}
		})
	}
}

func TestAdd(t *testing.T) {
	testCases := []struct {
		name    string
		s       []int
		idx     int
		val     int
		want    []int
		wantErr error
	}{
		{"add to empty slice", []int{}, 0, 1, []int{1}, nil},
		{"negative index", []int{1, 2, 3, 4, 5}, -1, 6, []int{1, 2, 3, 4, 5}, ErrIndexOutOfRange},
		{"index out of range", []int{1, 2, 3, 4, 5}, 6, 6, []int{1, 2, 3, 4, 5}, ErrIndexOutOfRange},
		{"add to head", []int{1, 2, 3, 4, 5}, 0, 6, []int{6, 1, 2, 3, 4, 5}, nil},
		{"add to tail", []int{1, 2, 3, 4, 5}, 5, 6, []int{1, 2, 3, 4, 5, 6}, nil},
		{"add to middle", []int{1, 2, 3, 4, 5}, 2, 6, []int{1, 2, 6, 3, 4, 5}, nil},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			slice := tc.s
			res, err := Add(tc.s, tc.idx, tc.val)
			if !errors.Is(err, tc.wantErr) {
				t.Errorf("Add(%v, %v, %v) = %v, want %v",
					slice, tc.idx, tc.val, err, tc.wantErr)
			}
			if err != nil {
				return
			}
			if !reflect.DeepEqual(res, tc.want) {
				t.Errorf("Add(%v, %v, %v) = %v, want %v",
					slice, tc.idx, tc.val, res, tc.want)
			}
		})
	}
}
