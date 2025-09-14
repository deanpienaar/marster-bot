package mars

import (
	"testing"
)

func TestPositionSet(t *testing.T) {
	t.Run("NewPositionSet creates empty set", func(t *testing.T) {
		set := NewPositionSet()
		if set.Len() != 0 {
			t.Errorf("Expected empty set, got length %d", set.Len())
		}
	})

	t.Run("Add and Has", func(t *testing.T) {
		set := NewPositionSet()
		pos1 := NewPosition(3, 4)
		pos2 := NewPosition(-1, 5)

		set.Add(pos1)
		if !set.Has(pos1) {
			t.Errorf("Expected set to contain position %v", pos1)
		}
		if set.Has(pos2) {
			t.Errorf("Expected set to not contain position %v", pos2)
		}

		set.Add(pos2)
		if !set.Has(pos2) {
			t.Errorf("Expected set to contain position %v", pos2)
		}
		if set.Len() != 2 {
			t.Errorf("Expected set length 2, got %d", set.Len())
		}
	})

	t.Run("Add duplicate position", func(t *testing.T) {
		set := NewPositionSet()
		pos := NewPosition(1, 2)

		set.Add(pos)
		set.Add(pos)
		if set.Len() != 1 {
			t.Errorf("Expected set length 1 after adding duplicate, got %d", set.Len())
		}
	})

	t.Run("Del removes position", func(t *testing.T) {
		set := NewPositionSet()
		pos1 := NewPosition(1, 2)
		pos2 := NewPosition(3, 4)

		set.Add(pos1)
		set.Add(pos2)

		set.Del(pos1)
		if set.Has(pos1) {
			t.Errorf("Expected position %v to be removed", pos1)
		}
		if !set.Has(pos2) {
			t.Errorf("Expected position %v to still be in set", pos2)
		}
		if set.Len() != 1 {
			t.Errorf("Expected set length 1, got %d", set.Len())
		}
	})

	t.Run("Del non-existent position", func(t *testing.T) {
		set := NewPositionSet()
		pos1 := NewPosition(1, 2)
		pos2 := NewPosition(3, 4)

		set.Add(pos1)
		set.Del(pos2) // Delete position that was never added

		if set.Len() != 1 {
			t.Errorf("Expected set length 1, got %d", set.Len())
		}
	})

	t.Run("Keys returns all positions", func(t *testing.T) {
		set := NewPositionSet()
		pos1 := NewPosition(1, 2)
		pos2 := NewPosition(3, 4)
		pos3 := NewPosition(-1, -2)

		set.Add(pos1)
		set.Add(pos2)
		set.Add(pos3)

		keys := set.Keys()
		if len(keys) != 3 {
			t.Errorf("Expected 3 keys, got %d", len(keys))
		}

		// Check that all positions are in the keys
		foundPos1, foundPos2, foundPos3 := false, false, false
		for _, key := range keys {
			if key.Equals(pos1) {
				foundPos1 = true
			}
			if key.Equals(pos2) {
				foundPos2 = true
			}
			if key.Equals(pos3) {
				foundPos3 = true
			}
		}

		if !foundPos1 || !foundPos2 || !foundPos3 {
			t.Errorf("Not all positions found in keys: pos1=%v, pos2=%v, pos3=%v", foundPos1, foundPos2, foundPos3)
		}
	})

	t.Run("Positions with same coordinates are treated as same", func(t *testing.T) {
		set := NewPositionSet()
		pos1 := NewPosition(5, 7)
		pos2 := NewPosition(5, 7) // Same coordinates as pos1

		set.Add(pos1)
		if !set.Has(pos2) {
			t.Errorf("Expected set to recognize position with same coordinates")
		}

		set.Add(pos2)
		if set.Len() != 1 {
			t.Errorf("Expected set length 1, got %d", set.Len())
		}
	})

	t.Run("Handles negative coordinates", func(t *testing.T) {
		set := NewPositionSet()
		pos1 := NewPosition(-5, -10)
		pos2 := NewPosition(-5, 10)
		pos3 := NewPosition(5, -10)

		set.Add(pos1)
		set.Add(pos2)
		set.Add(pos3)

		if !set.Has(pos1) || !set.Has(pos2) || !set.Has(pos3) {
			t.Errorf("Set should contain all positions with negative coordinates")
		}
		if set.Len() != 3 {
			t.Errorf("Expected set length 3, got %d", set.Len())
		}
	})

	t.Run("Empty set Keys returns empty slice", func(t *testing.T) {
		set := NewPositionSet()
		keys := set.Keys()
		if len(keys) != 0 {
			t.Errorf("Expected empty keys slice, got length %d", len(keys))
		}
	})
}

func TestPosition(t *testing.T) {
	t.Run("NewPosition creates correct position", func(t *testing.T) {
		pos := NewPosition(3, 5)
		if pos.X != 3 || pos.Y != 5 {
			t.Errorf("Expected position (3,5), got (%d,%d)", pos.X, pos.Y)
		}
	})

	t.Run("String formats correctly", func(t *testing.T) {
		pos := NewPosition(2, -4)
		expected := "(2,-4)"
		if pos.String() != expected {
			t.Errorf("Expected string %s, got %s", expected, pos.String())
		}
	})

	t.Run("Equals compares correctly", func(t *testing.T) {
		pos1 := NewPosition(1, 2)
		pos2 := NewPosition(1, 2)
		pos3 := NewPosition(1, 3)
		pos4 := NewPosition(2, 2)

		if !pos1.Equals(pos2) {
			t.Errorf("Expected positions with same coordinates to be equal")
		}
		if pos1.Equals(pos3) {
			t.Errorf("Expected positions with different Y to not be equal")
		}
		if pos1.Equals(pos4) {
			t.Errorf("Expected positions with different X to not be equal")
		}
	})

	t.Run("Copy creates independent copy", func(t *testing.T) {
		pos1 := NewPosition(5, 7)
		pos2 := pos1.Copy()

		if !pos1.Equals(pos2) {
			t.Errorf("Expected copy to have same values")
		}

		// Modify original
		pos1.X = 10
		if pos2.X == 10 {
			t.Errorf("Expected copy to be independent, but X changed")
		}
	})
}
