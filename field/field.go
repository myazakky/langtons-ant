package field

import "fmt"

type field struct {
	Width  int
	Height int
	square [][]int
}

func New(width, height int) *field {
	f := new(field)

	f.Width = width
	f.Height = height
	f.square = make([][]int, height)
	for i := range f.square {
		f.square[i] = make([]int, height)
	}

	return f
}

func (f *field) GetStatusAt(x, y int) int {
	return f.square[y][x]
}

func (f *field) SwitchStatusAt(x, y int) {
	if f.square[y][x] == 0 {
		f.square[y][x] = 1
	} else {
		f.square[y][x] = 0
	}
}

func (f *field) Draw(point_x, point_y int) {
	for y := range f.square {
		for x := range f.square[y] {
			switch {
			case x == point_x && y == point_y:
				fmt.Print("蟻")
			case f.GetStatusAt(x, y) == 0:
				fmt.Print("■")
			case f.GetStatusAt(x, y) == 1:
				fmt.Print("◇")
			}
		}
		fmt.Println("")
	}
}
