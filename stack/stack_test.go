package stack

import "testing"

func TestStackNew(t *testing.T) {
	var stack Stack[int]
	const MAXLEN = 5
	const MAXCAP = 10

	{
		stack = New[int]()
		for range stack.values {
			t.Errorf("expect empty slice")
		}
		if stack.len != 0 {
			t.Errorf("expect slice of zero length, got %d", stack.len)
		}
		if stack.cap != 0 {
			t.Errorf("expect slice of zero capacity, got %d", stack.cap)
		}
	}

	{
		stack = New[int](MAXCAP)
		if len_ := len(stack.values); len_ != MAXCAP {
			t.Errorf("expect slice of length %d, got %d", MAXCAP, len_)
		}
		if cap_ := cap(stack.values); cap_ != MAXCAP {
			t.Errorf("expect stack of capacity %d, got %d", MAXCAP, cap_)
		}
		for i, value := range stack.values {
			if value != 0 {
				t.Errorf("expect zero at position %d, got %v", i, value)
			}
		}
		if stack.len != MAXCAP {
			t.Errorf("expect stack of length %d, got %d", MAXCAP, stack.len)
		}
		if stack.cap != MAXCAP {
			t.Errorf("expect stack of capacity %d, got %d", MAXCAP, stack.len)
		}
	}

	{
		stack = New[int](MAXLEN, MAXCAP)
		if len_ := len(stack.values); len_ != MAXLEN {
			t.Errorf("expect slice of length %d, got %d", MAXLEN, len_)
		}
		if cap_ := cap(stack.values); cap_ != MAXCAP {
			t.Errorf("expect stack of capacity %d, got %d", MAXCAP, cap_)
		}
		for i, value := range stack.values {
			if value != 0 {
				t.Errorf("expect zero at position %d, got %v", i, value)
			}
		}
		if len_ := stack.len; len_ != MAXLEN {
			t.Errorf("expect stack of length %d, got %d", MAXLEN, len_)
		}
		if cap_ := stack.cap; cap_ != MAXCAP {
			t.Errorf("expect stack of capacity %d, got %d", MAXCAP, cap_)
		}
	}
}

func TestStackLen(t *testing.T) {
	stack := New[int](3)

	stack.Push(1)
	stack.Push(-2)
	stack.Push(30)

	stack.Pop(0)

	if len_ := stack.Len(); len_ != 2 {
		t.Errorf("expect length %v, got %v", 2, len_)
	}

	stack.Push(3)

	if len_ := stack.Len(); len_ != 3 {
		t.Errorf("expect length %v, got %v", 3, len_)
	}
}

func TestStackCap(t *testing.T) {
	stack := New[int](3)

	stack.Push(1)
	stack.Push(-2)
	stack.Push(30)

	stack.Pop(0)

	if cap_ := stack.Cap(); cap_ != 3 {
		t.Errorf("expect capacity %v, got %v", 3, cap_)
	}

	stack.Push(100)
	stack.Push(10)

	if cap_ := stack.Cap(); cap_ != 5 {
		t.Errorf("expect capacity %v, got %v", 5, cap_)
	}
}

func TestStackPush(t *testing.T) {
	stack := New[int]()

	stack.Push(1)
	stack.Push(2)
	stack.Push(-3)

	correct := []int{1, 2, -3}
	for i := 0; i < len(correct); i++ {
		if stack.values[i] != correct[i] {
			t.Errorf("expect %v at position %d, got %v", stack.values[i], i, correct[i])
		}
	}
	if stack.len != 3 {
		t.Errorf("expect %v, got %v", 3, stack.len)
	}
	if stack.cap != 3 {
		t.Errorf("expect %v, got %v", 3, stack.cap)
	}
}

func TestStackPop(t *testing.T) {
	stack := New[int]()
	const FILLER = -1

	if value, ok := stack.Pop(FILLER); ok {
		t.Errorf("expect fail, got %v", value)
	} else if value != FILLER {
		t.Errorf("expect filler value %v, got %v", -1, value)
	}

	values := []int{1, 2, -3}
	for _, value := range values {
		stack.Push(value)
	}

	for i := len(values) - 1; i >= 0; i-- {
		if value, ok := stack.Pop(FILLER); ok {
			if value != values[i] {
				t.Errorf("expect %v, got %v", values[i], value)
			}
		} else {
			if value != FILLER {
				t.Errorf("expect filler value %v, got %v", FILLER, value)
			}
			t.Errorf("empty stack")
		}
	}
}

func TestStackPeek(t *testing.T) {
	stack := New[int]()

	if value, ok := stack.Peek(-1); ok {
		t.Errorf("expect fail, got %v", value)
	} else if value != -1 {
		t.Errorf("expect FILLER value %v, got %v", -1, value)
	}

	values := []int{1, 2, -3, -10, 0}
	for _, value := range values {
		stack.Push(value)
	}

	const FILLER = -1
	if value, ok := stack.Peek(FILLER); ok {
		okValue := values[len(values)-1]
		if value != okValue {
			t.Errorf("expect %v, got %v", okValue, value)
		}
	} else {
		if value != FILLER {
			t.Errorf("expect filler value %v, got %v", FILLER, value)
		}
		t.Errorf("expect value, got filler")
	}
}
