package structures

import (
	"testing"

	. "github.com/go-playground/assert/v2"
)

func newTree() *binaryNode[int] {
	return &binaryNode[int]{
		value: 1,
		left: &binaryNode[int]{
			value: 2,
			left: &binaryNode[int]{
				value: 4,
				left: &binaryNode[int]{
					value: 5,
				},
				right: &binaryNode[int]{
					value: 6,
				},
			},
			right: &binaryNode[int]{
				value: 3,
			},
		},
		right: &binaryNode[int]{
			value: 7,
			left: &binaryNode[int]{
				value: 8,
			},
		},
	}

}

func TestPreOrderSearch(t *testing.T) {
	root := newTree()
	expected := []int{1, 2, 4, 5, 6, 3, 7, 8}
	actual := PreOrderSearch(root)

	Equal(t, len(actual), len(expected))

	for i, val := range actual {
		Equal(t, val, expected[i])
	}
}

func TestInOrderSearch(t *testing.T) {
	root := newTree()
	expected := []int{5, 4, 6, 2, 3, 1, 8, 7}
	actual := InOrderSearch(root)

	Equal(t, len(actual), len(expected))

	for i, val := range actual {
		Equal(t, val, expected[i])
	}
}

func TestPostOrderSearch(t *testing.T) {
	root := newTree()
	expected := []int{5, 6, 4, 3, 2, 8, 7, 1}
	actual := PostOrderSearch(root)

	Equal(t, len(actual), len(expected))

	for i, val := range actual {
		Equal(t, val, expected[i])
	}
}

func TestCompareTrees(t *testing.T) {
	t1 := newTree()
	t2 := newTree()
	actual := CompareTrees(t1, t2)

	Equal(t, actual, true)

	t2.right = nil
	actual = CompareTrees(t1, t2)

	Equal(t, actual, false)

	t1.right = nil
	actual = CompareTrees(t1, t2)

	Equal(t, actual, true)

	t1.left.value = 120
	actual = CompareTrees(t1, t2)
	Equal(t, actual, false)
}

func TestBreadthFirstSearch(t *testing.T) {
	root := newTree()
	tests := []struct {
		value int
		want  bool
	}{
		{value: 6, want: true},
		{value: 7, want: true},
		{value: 12, want: false},
		{value: 100, want: false},
	}

	for _, tt := range tests {
		got := BreadthFirstSearch(root, tt.value)
		Equal(t, got, tt.want)
	}
}

func newBinarySearchTree() *binaryNode[int] {
	return &binaryNode[int]{
		value: 10,
		left: &binaryNode[int]{
			value: 5,
			left: &binaryNode[int]{
				value: 2,
				left: &binaryNode[int]{
					value: 1,
				},
			},
			right: &binaryNode[int]{
				value: 6,
			},
		},
		right: &binaryNode[int]{
			value: 15,
			left: &binaryNode[int]{
				value: 14,
			},
			right: &binaryNode[int]{
				value: 20,
			},
		},
	}
}

func TestSearchBinarySearchTree(t *testing.T) {
	root := newBinarySearchTree()

	tests := []struct {
		value int
		want  bool
	}{
		{value: 10, want: true},
		{value: 5, want: true},
		{value: 2, want: true},
		{value: 1, want: true},
		{value: 6, want: true},
		{value: 15, want: true},
		{value: 14, want: true},
		{value: 20, want: true},
		{value: 12, want: false},
		{value: 11, want: false},
	}

	for _, tt := range tests {
		got := SearchBinarySearchTree(root, tt.value)
		Equal(t, got, tt.want)
	}
}
