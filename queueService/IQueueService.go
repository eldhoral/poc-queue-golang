package queue

import "errors"

var (
	ErrQueueEmpty = errors.New("queue empty")
	ErrQueueFull  = errors.New("queue full")
)

type Queue[T any] interface {
	Enqueue(T) error
	Dequeue() T
	Len() int
}
