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
