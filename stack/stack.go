package stack

type Stack[T any] struct {
	values []T
	len    int
	cap    int
}

func New[T any](size ...int) Stack[T] {
	var values []T
	switch len(size) {
	case 0:
		values = make([]T, 0)
	case 1:
		values = make([]T, size[0])
	case 2:
		values = make([]T, size[0])
	default:
		panic(size)
	}

	return Stack[T]{
		values: values,
		len:    len(values),
		cap:    cap(values),
	}
}

func (stack Stack[T]) Len() int {
	return stack.len
}

func (stack Stack[T]) Cap() int {
	return stack.cap
}

func (stack *Stack[T]) Push(value T) {
	if stack.len < stack.cap {
		stack.values[stack.len] = value
		stack.len += 1
	} else {
		stack.values = append(stack.values, value)
		stack.len += 1
		stack.cap += 1
	}
}

func (stack *Stack[T]) Pop(filler T) (value T, ok bool) {
	if stack.len == 0 {
		return filler, false
	} else {
		value = stack.values[stack.len-1]
		stack.len -= 1
		return value, true
	}
}

func (stack Stack[T]) Peek(filler T) (value T, ok bool) {
	if stack.len == 0 {
		return filler, false
	} else {
		return stack.values[stack.len-1], true
	}
}
