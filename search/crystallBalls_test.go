package search

import "testing"

func TestTwoCrystallBalls(t *testing.T) {
	tests := []struct {
		breaks   []bool
		expected int
	}{
		{
			breaks:   []bool{false, false, false, false, false, false, false, true, true, true, true, true, true, true, true, true},
			expected: 7,
		},
		{
			breaks:   []bool{false, true, true},
			expected: 1,
		},
		{
			breaks:   []bool{false, false, false, false, false, false, false, false, true, true, true, true, true, true, true, true, true},
			expected: 8,
		},
		{
			breaks:   []bool{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true},
			expected: 16,
		},
		{
			breaks:   []bool{},
			expected: -1,
		},
	}

	for _, test := range tests {
		if actual := TwoCrystallBalls(test.breaks); actual != test.expected {
			t.Errorf("expected %d, got %d", test.expected, actual)
		}
	}
}
