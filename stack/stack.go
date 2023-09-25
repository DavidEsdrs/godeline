package stack

type Stack[T any] struct {
	_arr   []T
	Length int
}

func NewStack[T any]() Stack[T] {
	return Stack[T]{
		_arr: []T{},
	}
}

func (stack *Stack[T]) Push(item T) {
	stack._arr = append(stack._arr, item)
	stack.Length++
}

func (stack *Stack[T]) Pop() T {
	last := stack._arr[stack.Length-1]
	stack._arr = stack._arr[:stack.Length-1]
	stack.Length--
	return last
}
