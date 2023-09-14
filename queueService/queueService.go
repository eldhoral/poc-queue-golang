package queue

type SliceQueue[T any] []T

func NewSliceQueue[T any]() Queue[T] {
	return &SliceQueue[T]{}
}

func (q *SliceQueue[T]) Len() int {
	return len(*q)
}

func (q *SliceQueue[T]) Enqueue(value T) error {
	*q = append(*q, value)
	return nil
}

func (q *SliceQueue[T]) Dequeue() T {
	queue := *q
	if len(*q) > 0 {
		card := queue[0]
		*q = queue[1:]
		return card
	}

	var empty T
	return empty
}
