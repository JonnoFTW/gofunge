package gofunge

import (
	"testing"
)

func TestNewInterpreter(t *testing.T) {
	source := `>987v
 v456<
 >321^`
	width := 6
	height := 3
	interpreter, err := NewInterpreter(width, height, source)

	if err != nil {
		t.Fatalf("Error creating interpreter: %v", err)
	}

	if interpreter.px != 0 || interpreter.py != 0 {
		t.Errorf("Expected program counter at (0, 0), got (%d, %d)", interpreter.px, interpreter.py)
	}

	if len(interpreter.board) != height || len(interpreter.board[0]) != width {
		t.Errorf("Expected board dimensions %dx%d, got %dx%d", height, width, len(interpreter.board), len(interpreter.board[0]))
	}
}

func TestInterpreter_Step(t *testing.T) {
	source := `>987v 
	 v456<
	 >321^`

	interpreter, err := NewInterpreter(7, 3, source)
	if err != nil {
		t.Fatalf("Error creating interpreter: %v", err)
	}

	// Initial state before stepping
	initialPX, initialPY := interpreter.px, interpreter.py
	initialStackLen := len(interpreter.stack.stack)

	// Perform a step
	interpreter.Step()

	// Validate the state after stepping
	if interpreter.px == initialPX && interpreter.py == initialPY {
		t.Error("Expected program counter to move after a step")
	}
	if len(interpreter.stack.stack) != initialStackLen {
		t.Error("Expected stack length to change after a step")
	}
	interpreter.Step()
	// After 1 step 9 should be on the stack
	if interpreter.stack.Peek() != 9 {
		t.Error("Expected 9 to be on the stack after 2nd step")
	}
}

func TestInterpreter_Show(t *testing.T) {
	source := `>987v
	 v456<
	 >321^`

	interpreter, err := NewInterpreter(7, 3, source)
	if err != nil {
		t.Fatalf("Error creating interpreter: %v", err)
	}

	// Test Show function
	interpreter.Show()
}
