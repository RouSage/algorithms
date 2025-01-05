package search

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		arr    []int64
		target int64
		want   int
	}{
		{
			arr:    []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			target: 5,
			want:   4,
		},
		{
			arr:    []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			target: 11,
			want:   -1,
		},
		{
			arr:    []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 20, 20, 30},
			target: 20,
			want:   10,
		},
		{
			arr:    []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			target: 3,
			want:   2,
		},
		{
			arr:    []int64{},
			target: 3,
			want:   -1,
		},
	}

	for _, tt := range tests {
		if got := BinarySearch(tt.arr, tt.target); got != tt.want {
			t.Errorf("BinarySearch() = %v, want %v", got, tt.want)
		}
	}
}
