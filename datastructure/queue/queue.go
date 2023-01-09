package queue

// Queue is a generic queue data structure.
type Queue[T any] []T

// New returns a new queue.
func New[T any]() *Queue[T] {
	return &Queue[T]{}
}

// Enqueue adds a value to the queue.
func (q *Queue[T]) Enqueue(v T) {
	*q = append(*q, v)
}

// Dequeue removes a value from the queue.
func (q *Queue[T]) Dequeue() T {
	v := (*q)[0]
	*q = (*q)[1:]
	return v
}

// Peek returns the value at the front of the queue.
func (q *Queue[T]) Peek() T {
	return (*q)[0]
}

// IsEmpty returns true if the queue is empty.
func (q *Queue[T]) IsEmpty() bool {
	return len(*q) == 0
}

// Size returns the size of the queue.
func (q *Queue[T]) Size() int {
	return len(*q)
}
