package heapq

type Data[T any] struct {
	Priority int
	Value    T
}

// NewData returns a new Data
func NewData[T any](priority int, value T) Data[T] {
	return Data[T]{Priority: priority, Value: value}
}

type PriorityQueue[T any] []Data[T]

// New returns a new priority queue.
func New[T any]() *PriorityQueue[T] {
	return &PriorityQueue[T]{}
}

// Push adds a value to the queue.
func (h *PriorityQueue[T]) Push(v Data[T]) {
	*h = append(*h, v)
	h.up(len(*h) - 1)
}

// Pop removes a value from the queue.
func (h *PriorityQueue[T]) Pop() Data[T] {
	v := (*h)[0]
	(*h)[0] = (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	h.down(0)
	return v
}

// Peek returns the value at the top of the queue.
func (h *PriorityQueue[T]) Peek() Data[T] {
	return (*h)[0]
}

// IsEmpty returns true if the queue is empty.
func (h *PriorityQueue[T]) IsEmpty() bool {
	return len(*h) == 0
}

// Size returns the size of the queue.
func (h *PriorityQueue[T]) Size() int {
	return len(*h)
}

// up moves the element at index i up the heap.
func (h *PriorityQueue[T]) up(i int) {
	for {
		parent := (i - 1) / 2
		if parent == i || (*h)[parent].Priority <= (*h)[i].Priority {
			break
		}
		(*h)[parent], (*h)[i] = (*h)[i], (*h)[parent]
		i = parent
	}
}

// down moves the element at index i down the heap.
func (h *PriorityQueue[T]) down(i int) {
	for {
		a := 2*i + 1
		b := 2*i + 2
		if a >= len(*h) {
			break
		}
		c := a
		if b < len(*h) && (*h)[b].Priority < (*h)[a].Priority {
			c = b
		}
		if (*h)[i].Priority <= (*h)[c].Priority {
			break
		}
		(*h)[i], (*h)[c] = (*h)[c], (*h)[i]
		i = c
	}
}
