package main

import "testing"

func TestInstructionBuilderWithRectangle(t *testing.T) {
	i, err := newInstructionFromString("rect 3x2")
	if err != nil {
		t.Fatalf("expected error to be nil, but was: %v", err)
	}
	if i.command != "rect" {
		t.Fatalf("expected command to be 'rect' and it was %q", i.command)
	}
	if i.x != 3 {
		t.Fatalf("expected x to be 3 and it was %v", i.x)
	}
	if i.y != 2 {
		t.Fatalf("expected y to be 2 and it was %v", i.y)
	}
}

func TestInstructionBuilderWithRectangleIncorrectX(t *testing.T) {
	_, err := newInstructionFromString("rect sx8")
	if err == nil {
		t.Fatalf("expected error not to be nil but it was: %v", err)
	}
}

func TestInstructionBuilderWithRectangleIncorrectY(t *testing.T) {
	_, err := newInstructionFromString("rect 2xd")
	if err == nil {
		t.Fatalf("expected error not to be nil but it was: %v", err)
	}
}

func TestInstructionBuilderWithRotate(t *testing.T) {
	i, err := newInstructionFromString("rect 3x2")
	if err != nil {
		t.Fatalf("expected error to be nil, but was: %v", err)
	}
	if i.command != "rect" {
		t.Fatalf("expected command to be 'rect' and it was %q", i.command)
	}
	if i.x != 3 {
		t.Fatalf("expected x to be 3 and it was %v", i.x)
	}
	if i.y != 2 {
		t.Fatalf("expected y to be 2 and it was %v", i.y)
	}
}
