package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type Triangle struct {
	A, B, C int
	IsLegal bool
}

var triangleRE = regexp.MustCompile(`\d+`)
var count int

func main() {
	tfile, err := ioutil.ReadFile("assets/day3.txt")
	if err != nil {
		log.Fatalf("could not read file %v:", tfile)
	}
	triangles := strings.Split(string(tfile), "\n")
	for _, strTriangle := range triangles {
		sides := make([]int, 3)
		for i, side := range triangleRE.FindAllString(strTriangle, -1) {
			intSide, err := strconv.Atoi(side)
			if err != nil {
				log.Fatalf("could not convert %v to int:", side)
			}
			sides[i] = intSide
		}
		triangle := Triangle{A: sides[0], B: sides[1], C: sides[2]}
		if triangleIsLegal(triangle) {
			triangle.IsLegal = true
			count += 1
		}
		fmt.Println(triangle)
	}
	fmt.Println(count)
}

func triangleIsLegal(t Triangle) bool {
	return t.A+t.B > t.C && t.B+t.C > t.A && t.C+t.A > t.B
}
