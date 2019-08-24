package ant

type direction int

const (
	North direction = iota
	East
	South
	West
)

type ant struct {
	X         int
	Y         int
	direction direction
}

func New(X, Y int) *ant {
	a := new(ant)
	a.X = X
	a.Y = Y

	return a
}

func (a *ant) Rotate(vector string) {
	switch vector {
	case "right":
		a.direction += 1
	case "left":
		a.direction -= 1
	}

	switch {
	case a.direction < North:
		a.direction = West
	case a.direction > West:
		a.direction = North
	}
}

func (a *ant) Move(maxX, maxY int) {
	switch a.direction {
	case North:
		a.Y -= 1
	case South:
		a.Y += 1
	case East:
		a.X += 1
	case West:
		a.X -= 1
	}

	switch {
	case a.Y > maxY:
		a.Y = 0
	case a.Y < 0:
		a.Y = maxY
	case a.X > maxX:
		a.X = 0
	case a.X < 0:
		a.X = maxX
	}
}
