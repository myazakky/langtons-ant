package main

import (
	"./ant"
	"./field"
	"fmt"
)

func main() {
	size := 100
	sandbox := field.New(size, size)
	antman := ant.New(size/2, size/2)
	var p int
	var count int

	for {
		p = sandbox.GetStatusAt(antman.X, antman.Y)
		if p == 0 {
			antman.Rotate("right")
		} else {
			antman.Rotate("left")
		}
		sandbox.SwitchStatusAt(antman.X, antman.Y)
		antman.Move(sandbox.Width-1, sandbox.Height-1)

		if count%100 == 0 {
			fmt.Println(count)
			sandbox.Draw(antman.X, antman.Y)
		}

		count += 1
	}
}
