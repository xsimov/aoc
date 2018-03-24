package main

import (
	"testing"
)

func TestInstructionBuilderWithInvalidCommand(t *testing.T) {
	_, err := newInstructionFromString("gibberish 3x2")
	if err == nil {
		t.Fatalf("expected error not to be nil, but it was")
	}
	expectedError := `instruction first argument must be rect or rotate, got "gibberish" instead (from "gibberish 3x2")`
	if err.Error() != expectedError {
		t.Fatalf("error (%q) was different from expected error: %q", err.Error(), expectedError)
	}
}

func TestInstructionBuilderWithRectangle(t *testing.T) {
	i, _ := newInstructionFromString("rect 3x2")
	_, ok := i.(rectInstruction)
	if !ok {
		t.Fatalf("expected %v(%T) to be of type rectInstruction", i, i)
	}
}

func TestRectInstructionFieldsAreAssigned(t *testing.T) {
	i, err := newInstructionFromString("rect 3x2")
	r := i.(rectInstruction)
	if err != nil {
		t.Fatalf("expected error to be nil, but was: %v", err)
	}
	if r.Command() != "rect" {
		t.Fatalf("expected command to be 'rect' and it was %q", r.Command())
	}
	if r.X != 3 {
		t.Fatalf("expected x to be 3 and it was %v", r.X)
	}
	if r.Y != 2 {
		t.Fatalf("expected y to be 2 and it was %v", r.Y)
	}
}

func TestInstructionBuilderWithRectangleIncorrectX(t *testing.T) {
	_, err := newInstructionFromString("rect sx8")
	if err == nil {
		t.Fatalf("expected error not to be nil but it was: %v", err)
	}
	expectedError := `instruction first argument must be rect or rotate, got "gibberish" instead (from "gibberish 3x2")`
	if err.Error() != expectedError {
		t.Fatalf("error (%q) was different from expected error: %q", err.Error(), expectedError)
	}
}

func TestInstructionBuilderWithRectangleIncorrectY(t *testing.T) {
	_, err := newInstructionFromString("rect 2xd")
	if err == nil {
		t.Fatalf("expected error not to be nil but it was: %v", err)
	}
	expectedError := `instruction first argument must be rect or rotate, got "gibberish" instead (from "gibberish 3x2")`
	if err.Error() != expectedError {
		t.Fatalf("error (%q) was different from expected error: %q", err.Error(), expectedError)
	}
}

func TestInstructionBuilderWithRotate(t *testing.T) {
	i, err := newInstructionFromString("rotate row y=1 by 5")
	r := i.(rotateInstruction)
	if err != nil {
		t.Fatalf("expected error to be nil, but was: %v", err)
	}
	if r.Command() != "rotate" {
		t.Fatalf("expected command to be 'rotate' and it was %q", r.Command())
	}
	if r.Index != 1 {
		t.Fatalf("expected index to be 1 and it was %v", r.Index)
	}
	if r.Direction != "row" {
		t.Fatalf("expected direction to be row and it was %v", r.Direction)
	}
}
