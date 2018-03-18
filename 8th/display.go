package main

import "fmt"

type display struct {
	columns, rows int
	Matrix        [][]bool
}

func main() {
	d := New(10, 6)
	fmt.Println(d)
	d.Rect(2, 3)
	fmt.Println(d)
}

func New(columns, rows int) *display {
	d := display{columns: columns, rows: rows}
	d.Matrix = make([][]bool, 0)
	for i := 0; i < rows; i++ {
		d.Matrix = append(d.Matrix, make([]bool, columns))
	}
	return &d
}

func (d *display) String() string {
	var s, r string
	for i := 0; i < d.rows; i++ {
		r = fmt.Sprintf("[")
		for j := 0; j < d.columns; j++ {
			r = fmt.Sprintf("%s%6v", r, d.Matrix[i][j])
		}
		r = fmt.Sprintf("%s ]\n", r)
		s = fmt.Sprintf("%s\n%s", s, r)
	}
	return s
}

func (d *display) Rect(x, y int) {
	for i := 0; i <= x; i++ {
		for j := 0; j <= y; j++ {
			d.Matrix[i][j] = true
		}
	}
}
