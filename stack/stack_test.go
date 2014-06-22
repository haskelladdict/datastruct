// some sanity tests to make sure our queue works ok
package stack

import (
	"testing"
)

func Test_stack(t *testing.T) {

	as := []int{1, 2, 3, 4}
	aStack := NewStack()
	bStack := NewStack()
	cStack := NewStack()
	for _, a := range as {
		aStack.Push(a)
		bStack.Push(a)
		cStack.Push(a)
	}

	if aStack.Len() != 4 {
		t.Error("stack has incorrect length")
	}

	for i := 3; i >= 0; i-- {
		val, err := aStack.Pop()
		if err != nil {
			t.Error("failed to pop element from stack")
		}
		if as[i] != val.(int) {
			t.Error("elements in stack appear to be out of order")
		}
	}

	bStack.Clear()
	_, err := bStack.Pop()
	if bStack.Len() != 0 || err == nil {
		t.Error("Failed to properly clear stack")
	}

	cStack.Pop()
	cStack.Pop()
	cStack.Push(5)
	cStack.Push(6)
	comp := []int{6, 5, 2, 1}
	for _, c := range comp {
		val, err := cStack.Pop()
		if err != nil {
			t.Error("failed to pop element from stack")
		}
		if c != val.(int) {
			t.Error("elements in stack appear to be out of order")
		}
	}
}
