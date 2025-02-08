package structures

type binaryNode[T comparable] struct {
	value T
	left  *binaryNode[T]
	right *binaryNode[T]
}

func PreOrderSearch[T comparable](root *binaryNode[T]) []T {
	return walkPreOrder(root, []T{})
}

func walkPreOrder[T comparable](curr *binaryNode[T], path []T) []T {
	if curr == nil {
		return path
	}

	path = append(path, curr.value)

	path = walkPreOrder(curr.left, path)
	path = walkPreOrder(curr.right, path)

	return path
}
