package utils

type Tick struct {
	VelX Direction `json:"velX"`
	VelY Direction `json:"velY"`
}

type Direction int

const (
	Neutral Direction = iota
	Left              = -1
	Right             = 1
	Down              = 1
	Up                = -1
)

type CurrentPosition struct {
	X, Y int
}

func (position *CurrentPosition) NewPosition(tick Tick) {
	if tick.VelX == Left {
		position.X -= 1
	}
	if tick.VelX == Right {
		position.X += 1
	}
	if tick.VelY == Up {
		position.Y += 1
	}
	if tick.VelY == Down {
		position.Y -= 1
	}
}

func ValidateMove(prev, next Tick) bool {
	if prev.VelX != 0 && next.VelX == -prev.VelX {
		return false
	}

	if prev.VelY != 0 && next.VelY == -prev.VelY {
		return false
	}
	return true
}

func ValidateInBound(
	CurrentPosition CurrentPosition,
	width int,
	height int,
) bool {
	if CurrentPosition.X < 0 || CurrentPosition.X > width {
		return false
	}

	if CurrentPosition.Y < 0 || CurrentPosition.Y > height {
		return false
	}

	return true
}
