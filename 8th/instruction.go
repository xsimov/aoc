package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Instruction must be executable upon a Display
type Instruction interface {
	Execute(Display) error
}

type rectInstruction struct {
	X, Y int
}

type rotateInstruction struct {
	Index, By int
	Direction string
}

func newInstructionFromString(s string) (i Instruction, err error) {
	splitted := strings.Split(s, " ")

	switch splitted[0] {
	case "rect":
		i, err = newRectInstruction(splitted)
	case "rotate":
		i, err = newRotateInstruction(splitted)
	default:
		err = fmt.Errorf("instruction first argument must be rect or rotate, got %q instead (from %q)", splitted[0], s)
	}
	return
}

func newRotateInstruction(splitted []string) (rotateInstruction, error) {
	if len(splitted) < 5 {
		return rotateInstruction{}, fmt.Errorf("rotate instruction could not be parsed from %q", splitted)
	}
	strIndex := strings.Split(splitted[2], "=")[1]
	index, err := strconv.Atoi(strIndex)
	if err != nil {
		return rotateInstruction{}, err
	}

	by, err := strconv.Atoi(splitted[4])
	if err != nil {
		return rotateInstruction{}, err
	}
	return rotateInstruction{Direction: splitted[1], Index: index, By: by}, nil
}

func newRectInstruction(splitted []string) (rectInstruction, error) {
	stringNums := strings.Split(splitted[1], "x")
	x, err := strconv.Atoi(stringNums[0])
	if err != nil {
		return rectInstruction{}, err
	}
	y, err := strconv.Atoi(stringNums[1])
	if err != nil {
		return rectInstruction{}, err
	}
	return rectInstruction{X: x, Y: y}, nil
}

func (i rectInstruction) Execute(d Display) error {
	return d.rect(i.X, i.Y)
}

func (i rotateInstruction) Execute(d Display) error {
	return d.rotate(i.Direction, i.Index, i.By)
}
