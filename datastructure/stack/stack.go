package stack

// Stack is a generic stack
type Stack[T any] []T

// New returns a new stack
func New[T any]() *Stack[T] {
	return &Stack[T]{}
}

// Push adds a value to the stack
func (s *Stack[T]) Push(v T) {
	*s = append(*s, v)
}

// Pop removes a value from the stack
func (s *Stack[T]) Pop() T {
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v
}

// Peek returns the value at the top of the stack
func (s *Stack[T]) Peek() T {
	return (*s)[len(*s)-1]
}

// IsEmpty returns true if the stack is empty
func (s *Stack[T]) IsEmpty() bool {
	return len(*s) == 0
}

// Size returns the size of the stack
func (s *Stack[T]) Size() int {
	return len(*s)
}
