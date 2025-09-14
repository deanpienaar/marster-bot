package mars

import "testing"

func TestGridBounds(t *testing.T) {
	grid := NewGrid(3, 3)
	
	tests := []struct {
		x, y     int8
		expected bool
		desc     string
	}{
		{0, 0, true, "origin should be valid"},
		{3, 3, true, "max corner should be valid"},
		{4, 3, false, "x=4 should be out of bounds"},
		{3, 4, false, "y=4 should be out of bounds"},
		{-1, 0, false, "negative x should be out of bounds"},
		{0, -1, false, "negative y should be out of bounds"},
		{1, 1, true, "middle position should be valid"},
		{2, 2, true, "position (2,2) should be valid"},
	}
	
	for _, tt := range tests {
		result := grid.PositionWithinBoundsXY(tt.x, tt.y)
		if result != tt.expected {
			t.Errorf("%s: PositionWithinBoundsXY(%d, %d) = %v, expected %v", 
				tt.desc, tt.x, tt.y, result, tt.expected)
		}
	}
}