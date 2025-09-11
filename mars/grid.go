package mars

type Grid struct {
	XSize            uint8
	YSize            uint8
	scentedPositions *PositionSet
}

func NewGrid(xSize, ySize uint8) *Grid {
	return &Grid{
		XSize:            xSize,
		YSize:            ySize,
		scentedPositions: NewPositionSet(),
	}
}

func (m *Grid) IsScented(pos Position) bool {
	return m.scentedPositions.Has(pos)
}

func (m *Grid) IsScentedXY(x, y int8) bool {
	return m.scentedPositions.Has(NewPosition(x, y))
}

func (m *Grid) AddScent(pos Position) {
	m.scentedPositions.Add(pos)
}

func (m *Grid) PositionWithinBounds(pos Position) bool {
	if pos.X < 0 || pos.Y < 0 {
		return false
	}

	return uint8(pos.X) <= m.XSize && pos.X > 0 && uint8(pos.Y) <= m.YSize
}

func (m *Grid) PositionWithinBoundsXY(x, y int8) bool {
	if x < 0 || y < 0 {
		return false
	}

	return uint8(x) <= m.XSize && uint8(y) <= m.YSize
}
