package mars

import (
	"bufio"
	"marster-bot/output"
	"strings"
	"testing"
)

func TestRoverIntegration(t *testing.T) {
	t.Run("Happy path - single rover moves correctly", func(t *testing.T) {
		// Setup
		reader := bufio.NewReader(strings.NewReader(""))
		console := output.NewConsole(*reader, false)
		grid := NewGrid(5, 5)
		rover := NewRover(1, 2, North, grid)

		// Instructions: RFRFFRFRF (should end at 1,3,N)
		instructions := []Instruction{
			NewOrientationInstruction(Right), // facing East
			NewMovementInstruction(1),         // move to (2,2)
			NewOrientationInstruction(Right), // facing South
			NewMovementInstruction(1),         // move to (2,1)
			NewMovementInstruction(1),         // move to (2,0)
			NewOrientationInstruction(Right), // facing West
			NewMovementInstruction(1),         // move to (1,0)
			NewOrientationInstruction(Right), // facing North
			NewMovementInstruction(1),         // move to (1,1)
		}

		// Execute instructions
		for _, inst := range instructions {
			err := rover.Instruct(console, inst)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
		}

		// Verify final position
		if rover.Position.X != 1 || rover.Position.Y != 1 {
			t.Errorf("Expected position (1,1), got (%d,%d)", rover.Position.X, rover.Position.Y)
		}
		if !rover.Direction.Equals(North) {
			t.Errorf("Expected direction North, got %v", rover.Direction)
		}
	})

	t.Run("Rover falls off and leaves scent", func(t *testing.T) {
		// Setup
		reader := bufio.NewReader(strings.NewReader(""))
		console := output.NewConsole(*reader, false)
		grid := NewGrid(5, 5)
		rover := NewRover(3, 3, North, grid)

		// Instructions to make rover fall off at (3,5)
		instructions := []Instruction{
			NewMovementInstruction(1), // move to (3,4)
			NewMovementInstruction(1), // move to (3,5)
			NewMovementInstruction(1), // try to move to (3,6) - should fall off
		}

		// Execute instructions
		var fallOffError error
		for _, inst := range instructions {
			err := rover.Instruct(console, inst)
			if err != nil {
				fallOffError = err
				break
			}
		}

		// Verify rover fell off
		if fallOffError == nil {
			t.Error("Expected rover to fall off grid")
		}
		if !strings.Contains(fallOffError.Error(), "fell off the grid") {
			t.Errorf("Expected fall off error, got: %v", fallOffError)
		}

		// Verify scent was left at the edge position
		edgePos := Position{X: 3, Y: 5}
		if !grid.IsScented(edgePos) {
			t.Errorf("Expected scent at position (%d,%d)", edgePos.X, edgePos.Y)
		}

		// Verify rover stayed at last valid position
		if rover.Position.X != 3 || rover.Position.Y != 5 {
			t.Errorf("Expected rover to stay at (3,5), got (%d,%d)", rover.Position.X, rover.Position.Y)
		}
	})

	t.Run("Second rover ignores instruction at scented position", func(t *testing.T) {
		// Setup
		reader := bufio.NewReader(strings.NewReader(""))
		console := output.NewConsole(*reader, false)
		grid := NewGrid(5, 5)
		
		// First rover falls off at (3,5) facing North
		rover1 := NewRover(3, 3, North, grid)
		instructions1 := []Instruction{
			NewMovementInstruction(1), // move to (3,4)
			NewMovementInstruction(1), // move to (3,5)
			NewMovementInstruction(1), // try to move to (3,6) - should fall off
		}
		
		for _, inst := range instructions1 {
			rover1.Instruct(console, inst)
		}

		// Verify scent exists
		if !grid.IsScented(Position{X: 3, Y: 5}) {
			t.Fatal("Expected scent at (3,5) after first rover fell off")
		}

		// Second rover at same position trying same move
		rover2 := NewRover(3, 5, North, grid)
		
		// Try to move forward (which would fall off)
		err := rover2.Instruct(console, NewMovementInstruction(1))
		
		// Should not error (instruction ignored due to scent)
		if err != nil {
			t.Errorf("Expected no error for scented position, got: %v", err)
		}

		// Rover should stay at (3,5)
		if rover2.Position.X != 3 || rover2.Position.Y != 5 {
			t.Errorf("Expected rover to stay at (3,5), got (%d,%d)", rover2.Position.X, rover2.Position.Y)
		}

		// Rover should be able to turn and move elsewhere
		err = rover2.Instruct(console, NewOrientationInstruction(Right)) // Face East
		if err != nil {
			t.Errorf("Unexpected error turning: %v", err)
		}
		
		err = rover2.Instruct(console, NewMovementInstruction(1)) // Move to (4,5)
		if err != nil {
			t.Errorf("Unexpected error moving east: %v", err)
		}

		if rover2.Position.X != 4 || rover2.Position.Y != 5 {
			t.Errorf("Expected rover at (4,5), got (%d,%d)", rover2.Position.X, rover2.Position.Y)
		}
	})

	t.Run("Multiple scents at different positions", func(t *testing.T) {
		// Setup
		reader := bufio.NewReader(strings.NewReader(""))
		console := output.NewConsole(*reader, false)
		grid := NewGrid(3, 3)
		
		// First rover falls off North edge
		rover1 := NewRover(1, 3, North, grid)
		rover1.Instruct(console, NewMovementInstruction(1))
		
		// Second rover falls off East edge
		rover2 := NewRover(3, 1, East, grid)
		rover2.Instruct(console, NewMovementInstruction(1))
		
		// Verify both scents exist
		if !grid.IsScented(Position{X: 1, Y: 3}) {
			t.Error("Expected scent at (1,3)")
		}
		if !grid.IsScented(Position{X: 3, Y: 1}) {
			t.Error("Expected scent at (3,1)")
		}

		// Third rover should be protected at both scented positions
		rover3 := NewRover(1, 3, North, grid)
		err := rover3.Instruct(console, NewMovementInstruction(1))
		if err != nil {
			t.Errorf("Expected no error at scented (1,3), got: %v", err)
		}
		
		// Move to other scented position
		rover3.Position = Position{X: 3, Y: 1}
		rover3.Direction = East
		err = rover3.Instruct(console, NewMovementInstruction(1))
		if err != nil {
			t.Errorf("Expected no error at scented (3,1), got: %v", err)
		}
	})

	t.Run("Rover at corner can fall off in two directions", func(t *testing.T) {
		// Setup
		reader := bufio.NewReader(strings.NewReader(""))
		console := output.NewConsole(*reader, false)
		
		// Test North direction with fresh grid
		grid1 := NewGrid(3, 3)
		rover1 := NewRover(3, 3, North, grid1)
		
		// Try to move North (should fall off)
		err := rover1.Instruct(console, NewMovementInstruction(1))
		if err == nil || !strings.Contains(err.Error(), "fell off") {
			t.Error("Expected rover to fall off moving North from corner")
		}
		
		// Test East direction with fresh grid
		grid2 := NewGrid(3, 3)
		rover2 := NewRover(3, 3, East, grid2)
		err = rover2.Instruct(console, NewMovementInstruction(1))
		if err == nil || !strings.Contains(err.Error(), "fell off") {
			t.Errorf("Expected rover to fall off moving East from corner, got err=%v, pos=(%d,%d)", err, rover2.Position.X, rover2.Position.Y)
		}
		
		// Verify scent protects from both directions on same grid
		grid3 := NewGrid(3, 3)
		rover3a := NewRover(3, 3, North, grid3)
		rover3a.Instruct(console, NewMovementInstruction(1)) // Falls off, leaves scent
		
		rover3b := NewRover(3, 3, North, grid3)
		err = rover3b.Instruct(console, NewMovementInstruction(1))
		if err != nil {
			t.Errorf("Expected no error at scented corner position facing North: %v", err)
		}
		
		rover3b.Direction = East
		err = rover3b.Instruct(console, NewMovementInstruction(1))
		if err != nil {
			t.Errorf("Expected no error at scented corner position facing East: %v", err)
		}
	})

	t.Run("Rover at origin boundary conditions", func(t *testing.T) {
		// Setup
		reader := bufio.NewReader(strings.NewReader(""))
		console := output.NewConsole(*reader, false)
		
		// Test South direction with fresh grid
		grid1 := NewGrid(5, 5)
		rover1 := NewRover(0, 0, South, grid1)
		err := rover1.Instruct(console, NewMovementInstruction(1))
		
		// Should fall off
		if err == nil || !strings.Contains(err.Error(), "fell off") {
			t.Error("Expected rover to fall off at origin moving South")
		}
		
		// Test West direction with fresh grid
		grid2 := NewGrid(5, 5)
		rover2 := NewRover(0, 0, West, grid2)
		err = rover2.Instruct(console, NewMovementInstruction(1))
		
		// Should fall off
		if err == nil || !strings.Contains(err.Error(), "fell off") {
			t.Error("Expected rover to fall off at origin moving West")
		}
		
		// Verify scents protect at origin on same grid
		grid3 := NewGrid(5, 5)
		rover3a := NewRover(0, 0, South, grid3)
		rover3a.Instruct(console, NewMovementInstruction(1)) // Falls off, leaves scent
		
		if !grid3.IsScented(Position{X: 0, Y: 0}) {
			t.Error("Expected scent at origin after fall")
		}
		
		// New rover at origin should be protected
		rover3b := NewRover(0, 0, West, grid3)
		err = rover3b.Instruct(console, NewMovementInstruction(1))
		if err != nil {
			t.Errorf("Expected no error at scented origin, got: %v", err)
		}
	})

	t.Run("Complex path with rotations and movements", func(t *testing.T) {
		// Setup
		reader := bufio.NewReader(strings.NewReader(""))
		console := output.NewConsole(*reader, false)
		grid := NewGrid(5, 5)
		rover := NewRover(1, 1, East, grid)

		// FFRFLFLFRFF - complex path
		instructions := []Instruction{
			NewMovementInstruction(1),         // (2,1) E
			NewMovementInstruction(1),         // (3,1) E
			NewOrientationInstruction(Right),  // (3,1) S
			NewMovementInstruction(1),         // (3,0) S
			NewOrientationInstruction(Left),   // (3,0) E
			NewMovementInstruction(1),         // (4,0) E
			NewOrientationInstruction(Left),   // (4,0) N
			NewMovementInstruction(1),         // (4,1) N
			NewOrientationInstruction(Right),  // (4,1) E
			NewMovementInstruction(1),         // (5,1) E
			NewMovementInstruction(1),         // Try (6,1) - should fall off
		}

		var lastErr error
		for _, inst := range instructions {
			err := rover.Instruct(console, inst)
			if err != nil {
				lastErr = err
				break
			}
		}

		// Should have fallen off at East boundary
		if lastErr == nil || !strings.Contains(lastErr.Error(), "fell off") {
			t.Error("Expected rover to fall off East boundary")
		}

		// Should be at (5,1) facing East
		if rover.Position.X != 5 || rover.Position.Y != 1 {
			t.Errorf("Expected position (5,1), got (%d,%d)", rover.Position.X, rover.Position.Y)
		}
		if !rover.Direction.Equals(East) {
			t.Errorf("Expected direction East, got %v", rover.Direction)
		}
	})
}