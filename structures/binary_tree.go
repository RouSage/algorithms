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

func InOrderSearch[T comparable](root *binaryNode[T]) []T {
	return walkInOrder(root, []T{})
}

func walkInOrder[T comparable](curr *binaryNode[T], path []T) []T {
	if curr == nil {
		return path
	}

	path = walkInOrder(curr.left, path)

	path = append(path, curr.value)

	path = walkInOrder(curr.right, path)

	return path
}

func PostOrderSearch[T comparable](root *binaryNode[T]) []T {
	return walPostOrder(root, []T{})
}

func walPostOrder[T comparable](curr *binaryNode[T], path []T) []T {
	if curr == nil {
		return path
	}

	path = walPostOrder(curr.left, path)
	path = walPostOrder(curr.right, path)

	path = append(path, curr.value)

	return path
}
