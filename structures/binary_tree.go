package structures

type binaryNode[T comparable] struct {
	value T
	left  *binaryNode[T]
	right *binaryNode[T]
}

// ######
// Depth First Search
// ######

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
	return walkPostOrder(root, []T{})
}

func walkPostOrder[T comparable](curr *binaryNode[T], path []T) []T {
	if curr == nil {
		return path
	}

	path = walkPostOrder(curr.left, path)
	path = walkPostOrder(curr.right, path)

	path = append(path, curr.value)

	return path
}

func CompareTrees[T comparable](t1, t2 *binaryNode[T]) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}
	if t1.value != t2.value {
		return false
	}

	return CompareTrees(t1.left, t2.left) && CompareTrees(t1.right, t2.right)
}

func SearchBinarySearchTree[T int | float64](root *binaryNode[T], value T) bool {
	return search(root, value)
}

func search[T int | float64](curr *binaryNode[T], value T) bool {
	if curr == nil {
		return false
	}

	if curr.value == value {
		return true
	}
	if curr.value < value {
		return search(curr.right, value)
	}

	return search(curr.left, value)
}

// ######
// Breadth First Search
// ######

func BreadthFirstSearch[T comparable](root *binaryNode[T], value T) bool {
	queue := NewQueue[*binaryNode[T]]()
	queue.Enqueue(root)

	for !queue.IsEmpty() {
		curr, ok := queue.Dequeue()
		if !ok {
			// Should never happen since the IsEmpty() check is done before the Dequeue()
			panic("Queue is empty")
		}

		if curr == nil {
			continue
		}
		if curr.value == value {
			return true
		}

		queue.Enqueue(curr.left)
		queue.Enqueue(curr.right)
	}

	return false
}
