package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// Display simulator
type Display struct {
	columns, rows int
	Matrix        [][]bool
}

func main() {
	f, err := os.Open("../assets/day8.txt")
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}
	content, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalf("could not read contents: %v", err)
	}
	instructions := strings.Split(string(content), "\n")

	d := NewDisplay(50, 16)

	for index, s := range instructions {
		if s == "" {
			continue
		}
		i, _ := newInstructionFromString(s)
		fmt.Println(d, i, index)
		i.Execute(d)
	}
	fmt.Println(d, d.countPixelsOn())
}

// NewDisplay returns a new instance of Display
func NewDisplay(columns, rows int) Display {
	d := Display{columns: columns, rows: rows}
	d.Matrix = make([][]bool, 0)
	for i := 0; i < rows; i++ {
		d.Matrix = append(d.Matrix, make([]bool, columns))
	}
	return d
}
func (d Display) countPixelsOn() int {
	var count int
	for i := 0; i < d.rows; i++ {
		for j := 0; j < d.columns; j++ {
			if d.Matrix[i][j] {
				count++
			}
		}
	}
	return count
}

func (d Display) String() string {
	var s, r string
	var c rune
	for i := 0; i < 6; i++ {
		r = fmt.Sprintf("[")
		for j := 0; j < d.columns; j++ {
			if d.Matrix[i][j] {
				c = '#'
			} else {
				c = '-'
			}
			r = fmt.Sprintf("%s%2c", r, c)
		}
		r = fmt.Sprintf("%s ]\n", r)
		s = fmt.Sprintf("%s%s", s, r)
	}
	return s
}

func (d Display) rotate(direction string, index, by int) (err error) {
	switch direction {
	case "row":
		err = d.rotateRow(index, by)
	case "column":
		err = d.rotateColumn(index, by)
	default:
		return fmt.Errorf("direction must be \"column\" or \"row\", and got: %v", direction)
	}
	return err
}

func (d Display) rect(x, y int) error {
	if d.boundaryViolation(x, y) {
		return fmt.Errorf("could not execute Rect, there is a boundary violation: [%v %v] exceeds [%v %v]", x, y, d.columns, d.rows)
	}
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			d.Matrix[i][j] = true
		}
	}
	return nil
}

func (d Display) boundaryViolation(x, y int) bool {
	return x >= d.rows || y >= d.columns
}

func (d Display) rotateRow(index, by int) error {
	cpy := make([]bool, d.columns)
	copy(cpy[by:len(d.Matrix[index])], d.Matrix[index][:(d.columns-by)])
	copy(cpy[:by], d.Matrix[index][d.columns-by:])
	copy(d.Matrix[index], cpy)
	return nil
}

func (d Display) rotateColumn(index, by int) error {
	cpy := make([]bool, d.rows)
	for c := 0; c < d.rows; c++ {
		cpy[(c+by)%d.rows] = d.Matrix[c][index]
	}
	for c := 0; c < d.rows; c++ {
		d.Matrix[c][index] = cpy[c]
	}
	return nil
}
