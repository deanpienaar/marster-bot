package mars

import "fmt"

type Position struct {
	X int8
	Y int8
}

func NewPosition(x, y int8) Position {
	return Position{X: x, Y: y}
}

func (p *Position) String() string {
	return fmt.Sprintf("(%d,%d)", p.X, p.Y)
}

func (p *Position) Equals(other Position) bool {
	return p.X == other.X && p.Y == other.Y
}

func (p *Position) Copy() Position {
	return Position{X: p.X, Y: p.Y}
}

type PositionSet struct {
	m map[string]Position
}

func NewPositionSet() *PositionSet {
	return &PositionSet{
		m: make(map[string]Position),
	}
}

func (s *PositionSet) getKey(pos Position) string {
	return fmt.Sprintf("%d,%d", pos.X, pos.Y)
}

func (s *PositionSet) Add(pos Position) {
	key := s.getKey(pos)
	s.m[key] = pos
}

func (s *PositionSet) Has(pos Position) bool {
	key := s.getKey(pos)
	_, ok := s.m[key]
	return ok
}

func (s *PositionSet) Del(pos Position) {
	key := s.getKey(pos)
	delete(s.m, key)
}

func (s *PositionSet) Keys() []Position {
	positions := make([]Position, 0, len(s.m))
	for _, pos := range s.m {
		positions = append(positions, pos)
	}
	return positions
}

func (s *PositionSet) Len() int {
	return len(s.m)
}
