package search

import (
	"testing"
)

func TestLinearSearch(t *testing.T) {
	tests := []struct {
		arr    []int64
		target int64
		want   int
	}{
		{
			arr:    []int64{1, 2, 3, 4, 5},
			target: 3,
			want:   2,
		},
		{
			arr:    []int64{10, 22, 3, 48, 5},
			target: 7,
			want:   -1,
		},
	}

	for _, tt := range tests {
		if got := LinearSearch(tt.arr, tt.target); got != tt.want {
			t.Errorf("LinearSearch() = %v, want %v", got, tt.want)
		}
	}
}
