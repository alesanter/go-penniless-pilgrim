package set

import "testing"

func TestNewSet(t *testing.T) {
	{
		set := New[int]()
		for range set.values {
			t.Errorf("expect empty map, got %v", set.values)
		}
	}

	{
		set := New[int](300)
		for range set.values {
			t.Errorf("expect empty map, got %v", set.values)
		}
	}

	{
		defer func() { _ = recover() }()
		set := New[int](1, 2, 3)
		t.Errorf("expect to panic, got %v", set.values)
	}
}

func TestAddSet(t *testing.T) {
	set := New[int]()

	{
		values := []int{1, 5, 100}
		for _, value := range values {
			ok := set.Add(value)
			if !ok {
				t.Errorf("expect %d to be a new element", value)
			}
		}
	}

	{
		count := 0
		for value := range set.values {
			count += 1
			if value != 1 && value != 5 && value != 100 {
				t.Errorf("expect 1 or 5 or 100, got %v", value)
			}
		}
		if count != 3 {
			t.Errorf("expect map to hold %d values, got %d", 3, count)
		}
	}
}

func TestSetRemove(t *testing.T) {
	set := New[int]()

	{
		ok := set.Remove(1)
		if ok {
			t.Errorf("expect %v to not be removed", 1)
		}
	}

	values := []int{1, 5, 100}
	for _, value := range values {
		set.Add(value)
	}

	for _, value := range values {
		ok := set.Remove(value)
		if !ok {
			t.Errorf("expect %v to be removed", value)
		}
	}
}

func TestSetContains(t *testing.T) {
	set := New[int]()

	{
		value := 10
		if set.Contains(value) {
			t.Errorf("expect set not to contain %v", value)
		}
	}

	values := []int{1, 5, 100}
	for _, value := range values {
		set.Add(value)
	}

	testValues := []int{1, 3, -4, 100, 499}
	testResult := []bool{true, false, false, true, false}
	for i, value := range testValues {
		ok := set.Contains(value)
		result := testResult[i]
		if ok != result {
			if result {
				t.Errorf("expect set to contain %v", value)
			} else {
				t.Errorf("expect set not to contain %v", value)
			}
		}
	}
}

func TestSetEmpty(t *testing.T) {
	set := New[int]()

	if !set.Empty() {
		t.Errorf("expect set to be empty")
	}

	values := []int{1, 5, 100}
	for _, value := range values {
		set.Add(value)
	}

	if set.Empty() {
		t.Errorf("expect set not to be empty")
	}

	for _, value := range values {
		set.Remove(value)
	}

	if !set.Empty() {
		t.Errorf("expect set to be empty")
	}
}
