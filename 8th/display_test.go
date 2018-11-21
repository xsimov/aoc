package main

import "testing"

func TestDisplay_String(t *testing.T) {
	type fields struct {
		x      int
		y      int
		Matrix [][]bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Displays columns",
			fields: fields{
				x: 2,
				y: 2,
			},
			want: "[ false false ]\n[ false false ]\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := NewDisplay(2, 2)
			if got := d.String(); got != tt.want {
				t.Errorf("Display.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
