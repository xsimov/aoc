package main

import "fmt"

type display struct {
	x, y   int
	Matrix [][]bool
}

func main() {
	d := New(10, 6)
	fmt.Println(d)
	err := d.Rect(12, 3)
	if err != nil {
		fmt.Println("could not create rect:", err)
	}
	fmt.Println(d)
}

func New(columns, rows int) *display {
	d := display{x: columns, y: rows}
	d.Matrix = make([][]bool, 0)
	for i := 0; i < rows; i++ {
		d.Matrix = append(d.Matrix, make([]bool, columns))
	}
	return &d
}

func (d *display) String() string {
	var s, r string
	for i := 0; i < d.y; i++ {
		r = fmt.Sprintf("[")
		for j := 0; j < d.x; j++ {
			r = fmt.Sprintf("%s%6v", r, d.Matrix[i][j])
		}
		r = fmt.Sprintf("%s ]\n", r)
		s = fmt.Sprintf("%s\n%s", s, r)
	}
	return s
}

func (d *display) Rotate(direction string, index, by int) (err error) {
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

func (d *display) Rect(x, y int) error {
	if d.boundaryViolation(x, y) {
		return fmt.Errorf("could not execute Rect, there is a boundary violation: [%v %v] exceeds [%v %v]", x, y, d.x, d.y)
	}
	for i := 0; i <= x; i++ {
		for j := 0; j <= y; j++ {
			d.Matrix[i][j] = true
		}
	}
	return nil
}

func (d *display) boundaryViolation(x, y int) bool {
	return x >= d.x || y >= d.y
}
