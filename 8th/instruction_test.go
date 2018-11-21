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

func TestInstructionBuilderWithInvalidCommandArguments(t *testing.T) {
	_, err := newInstructionFromString("rotate stupid stuff")
	if err == nil {
		t.Fatalf("expected error not to be nil, but it was")
	}
	expectedError := `rotate instruction could not be parsed from ["rotate" "stupid" "stuff"]`
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
		t.Fatalf("expected error not to be nil but it was!")
	}
	expectedError := `strconv.Atoi: parsing "s": invalid syntax`
	if err.Error() != expectedError {
		t.Fatalf("error (%q) was different from expected error: %q", err.Error(), expectedError)
	}
}

func TestInstructionBuilderWithRectangleIncorrectY(t *testing.T) {
	_, err := newInstructionFromString("rect 2xd")
	if err == nil {
		t.Fatalf("expected error not to be nil but it was: %v", err)
	}
	expectedError := `strconv.Atoi: parsing "d": invalid syntax`
	if err.Error() != expectedError {
		t.Fatalf("error (%q) was different from expected error: %q", err.Error(), expectedError)
	}
}

func TestInstructionBuilderWithRotate(t *testing.T) {
	i, _ := newInstructionFromString("rotate row y=1 by 5")
	_, ok := i.(rotateInstruction)
	if !ok {
		t.Fatalf("expected %v(%T) to be of type rotateInstruction", i, i)
	}
}

func TestInstructionBuilderWithRotateIncorrectIndex(t *testing.T) {
	_, err := newInstructionFromString("rotate row y=z by 5")
	if err == nil {
		t.Fatalf("expected error not to be nil but it was!")
	}
	expectedError := `strconv.Atoi: parsing "z": invalid syntax`
	if err.Error() != expectedError {
		t.Fatalf("error (%q) was different from expected error: %q", err.Error(), expectedError)
	}
}

func TestInstructionBuilderWithRotateIncorrectBy(t *testing.T) {
	_, err := newInstructionFromString("rotate row y=7 by o")
	if err == nil {
		t.Fatalf("expected error not to be nil but it was!")
	}
	expectedError := `strconv.Atoi: parsing "o": invalid syntax`
	if err.Error() != expectedError {
		t.Fatalf("error (%q) was different from expected error: %q", err.Error(), expectedError)
	}
}

func TestRotateInstructionFieldsAreAssigned(t *testing.T) {
	i, err := newInstructionFromString("rotate row y=1 by 5")
	r := i.(rotateInstruction)
	if err != nil {
		t.Fatalf("expected error to be nil, but was: %v", err)
	}
	if r.Index != 1 {
		t.Fatalf("expected index to be 1 and it was %v", r.Index)
	}
	if r.Direction != "row" {
		t.Fatalf("expected direction to be row and it was %v", r.Direction)
	}
}

func TestRectInstructionExecuteCallsRectUponDisplay(t *testing.T) {
	i, _ := newInstructionFromString("rect 3x2")
	d := NewDisplay(4, 4)
	if err := i.Execute(d); err != nil {
		t.Fatalf("expected error to be nil and it was: %v", err)
	}
}
