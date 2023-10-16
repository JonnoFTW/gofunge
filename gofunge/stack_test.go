package gofunge

import (
	"testing"
)

func TestStack_Pop(t *testing.T) {
	s := Stack{}

	// Test popping from an empty stack
	assertPanic(t, func() { s.Pop() }, "Stack is empty!")

	// Test popping from a non-empty stack
	s.Push(42)
	s.Push(56)

	// Check if values are popped in the correct order
	pop1 := s.Pop()
	if pop1 != 56 {
		t.Errorf("Expected popped value 56, got %d", pop1)
	}

	pop2 := s.Pop()
	if pop2 != 42 {
		t.Errorf("Expected popped value 42, got %d", pop2)
	}

	// Check if the stack is empty after popping all elements
	assertPanic(t, func() { s.Pop() }, "Stack is empty!")
}

func TestStack_Push(t *testing.T) {
	s := Stack{}

	// Test pushing values onto the stack
	push1 := s.Push(10)
	if push1 != 10 {
		t.Errorf("Expected pushed value 10, got %d", push1)
	}

	push2 := s.Push(20)
	if push2 != 20 {
		t.Errorf("Expected pushed value 20, got %d", push2)
	}

	// Test if the stack grows as expected
	if len(s.stack) != 2 {
		t.Errorf("Expected stack length to be 2, got %d", len(s.stack))
	}
}

func TestStack_Peek(t *testing.T) {
	s := Stack{}

	// Test peeking at an empty stack
	assertPanic(t, func() { s.Peek() }, "Stack is empty!")

	// Test peeking at a non-empty stack
	s.Push(5)
	s.Push(8)

	peekedValue := s.Peek()
	if peekedValue != 8 {
		t.Errorf("Expected peeked value 8, got %d", peekedValue)
	}

	// Ensure the stack is still the same after peeking
	if len(s.stack) != 2 {
		t.Errorf("Expected stack length to be 2, got %d", len(s.stack))
	}
}

// assertPanic is a helper function to check if the provided function panics with the expected message.
func assertPanic(t *testing.T, f func(), expectedMsg string) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic, but none occurred.")
		} else {
			errMsg, ok := r.(string)
			if !ok || errMsg != expectedMsg {
				t.Errorf("Expected panic with message '%s', got '%v'", expectedMsg, errMsg)
			}
		}
	}()

	f()
}
