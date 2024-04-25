package set

type Set[T comparable] struct {
	values map[T]bool
}

func New[T comparable](size ...int) Set[T] {
	var values map[T]bool

	switch len(size) {
	case 0:
		values = make(map[T]bool)
	case 1:
		values = make(map[T]bool, size[0])
	default:
		panic(size)
	}

	return Set[T]{
		values: values,
	}
}

func (set Set[T]) Empty() bool {
	for range set.values {
		return false
	}
	return true
}

// Returns `true` if the map didn't contain the added element, otherwise `false`.
func (set *Set[T]) Add(value T) (ok bool) {
	_, ok = set.values[value]
	set.values[value] = true
	return !ok
}

// Returns `false` if the map didn't contain the element, otherwise `true`.
func (set *Set[T]) Remove(value T) (ok bool) {
	_, ok = set.values[value]
	delete(set.values, value)
	return ok
}

func (set Set[T]) Contains(value T) bool {
	_, ok := set.values[value]
	return ok
}
