package mars

type Direction struct {
	code  rune
	value int8
}

var (
	North = Direction{code: 'N', value: 0}
	East  = Direction{code: 'E', value: 1}
	South = Direction{code: 'S', value: 2}
	West  = Direction{code: 'W', value: 3}
)

func (d Direction) Rotate(rotation Rotation) Direction {
	switch rotation {
	case Right:
		newValue := (d.value + 1) % 4
		return DirectionFromValue(uint8(newValue))
	case Left:
		newValue := (d.value - 1 + 4) % 4
		return DirectionFromValue(uint8(newValue))
	default:
		return d
	}
}

func (d Direction) String() string {
	return string(d.code)
}

func (d Direction) Equals(other Direction) bool {
	return d.value == other.value
}

var valueToDirectionMap = []Direction{
	North,
	East,
	South,
	West,
}

func DirectionFromValue(value uint8) Direction {
	return valueToDirectionMap[value]
}

var codeToDirectionMap = map[rune]Direction{
	'N': North,
	'E': East,
	'S': South,
	'W': West,
}

func DirectionFromCode(code rune) Direction {
	return codeToDirectionMap[code]
}

type Rotation rune

const (
	Right Rotation = 'R'
	Left  Rotation = 'L'
)
