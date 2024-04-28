package game

type Enum int

func (enum Enum) EnumIndex() int {
	return int(enum)
}

type Axis Enum

func (axis Axis) String() string {
	return [...]string{"X", "Y"}[axis-1]
}

const (
	X Axis = iota + 1
	Y
)
