package main

import (
	"strconv"
	"strings"
)

type instruction struct {
	rectInstruction
}

type rectInstruction struct {
	x, y    int
	command string
}

func newInstructionFromString(s string) (i instruction, err error) {
	splitted := strings.Split(s, " ")
	stringNums := strings.Split(splitted[1], "x")

	i = instruction{command: splitted[0]}

	i.x, err = strconv.Atoi(stringNums[0])
	if err != nil {
		return
	}
	i.y, err = strconv.Atoi(stringNums[1])
	if err != nil {
		return
	}
	return
}
