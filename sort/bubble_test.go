package sort

import "testing"

func TestBubbleSort(t *testing.T) {
	tests := []struct {
		arr  []int64
		want []int64
	}{
		{
			arr:  []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			want: []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			arr:  []int64{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			want: []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			arr:  []int64{1, 6, 2, 8, 9, 5},
			want: []int64{1, 2, 5, 6, 8, 9},
		},
	}

	for _, tt := range tests {
		BubbleSort(tt.arr)

		for i := 0; i < len(tt.arr); i++ {
			if tt.arr[i] != tt.want[i] {
				t.Errorf("BubbleSort() = %v, want %v", tt.arr, tt.want)
			}
		}
	}

}
