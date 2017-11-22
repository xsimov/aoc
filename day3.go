// Triangler solves the puzzle for AOC day3: count how many of the input triplets of values could make a valid triangle.

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
)

type Side struct {
	v   int
	set bool
}

func (s *Side) IsNotFilled() bool {
	return !s.set
}

type Triangle struct {
	A, B, C Side
}

func (t *Triangle) AddSide(s int) error {
	if t.A.IsNotFilled() {
		t.A = Side{s, true}
		return nil
	}
	if t.B.IsNotFilled() {
		t.B = Side{s, true}
		return nil
	}
	if t.C.IsNotFilled() {
		t.C = Side{s, true}
		return nil
	}
	return fmt.Errorf("could not add side %v to t(%v)", s, t)
}

func (t Triangle) IsLegal() bool {
	return t.A.v+t.B.v > t.C.v && t.B.v+t.C.v > t.A.v && t.C.v+t.A.v > t.B.v
}

var triangleRE = regexp.MustCompile(`\d+`)
var count int

func main() {
	tfile, err := ioutil.ReadFile("assets/day3.txt")
	if err != nil {
		log.Fatalf("could not read file %v:", tfile)
	}
	for _, triangle := range getTriangles(tfile) {
		if triangle.IsLegal() {
			count += 1
		}
	}

	fmt.Println(count)
}

func getTriangles(tfile []byte) []Triangle {
	allSides := triangleRE.FindAllString(string(tfile), -1)
	triangles := make([]Triangle, len(allSides)/3)

	for i, side := range allSides {
		position := verticalPositioned(i)
		intSide, _ := strconv.Atoi(side)
		triangles[position].AddSide(intSide)
	}
	return triangles
}

func horizontalPositioned(i int) int {
	return i / 3
}

func verticalPositioned(i int) int {
	return i%3 + (i / 9 * 3)
}
