package mars

import (
	"fmt"
	"marster-bot/output"
)

type Rover struct {
	Position  Position
	Direction Direction
	Grid      *Grid
}

func NewRover(x, y int8, startingDirection Direction, grid *Grid) *Rover {
	return &Rover{
		Position: Position{
			X: x,
			Y: y,
		},
		Direction: startingDirection,
		Grid:      grid,
	}
}

// Move
// Moves the rover by the specified distance with the direction (forwards or backwards) dictated by the sign.
func (r *Rover) Move(console *output.Console, distance int8) error {
	switch r.Direction {
	case North:
		if !r.Grid.PositionWithinBoundsXY(r.Position.X, r.Position.Y+1) {
			if !r.CurrentPositionIsScented() {
				return r.OnGridExit()
			}

			return nil
		}
		r.Position.Y += distance
	case East:
		if !r.Grid.PositionWithinBoundsXY(r.Position.X+1, r.Position.Y) {
			if !r.CurrentPositionIsScented() {
				return r.OnGridExit()
			}

			return nil
		}
		r.Position.X += distance
	case South:
		if !r.Grid.PositionWithinBoundsXY(r.Position.X, r.Position.Y-1) {
			if !r.CurrentPositionIsScented() {
				return r.OnGridExit()
			}

			return nil
		}
		r.Position.Y -= distance
	case West:
		if !r.Grid.PositionWithinBoundsXY(r.Position.X-1, r.Position.Y) {
			if !r.CurrentPositionIsScented() {
				return r.OnGridExit()
			}

			return nil
		}
		r.Position.X -= distance
	}

	console.Debug("Rover moved to (%d, %d)", r.Position.X, r.Position.Y)

	return nil
}

func (r *Rover) Rotate(orientation Rotation) error {
	r.Direction = r.Direction.Rotate(orientation)
	return nil
}

func (r *Rover) Instruct(console *output.Console, instruction Instruction) error {
	switch instruction.(type) {
	case *MovementInstruction:
		err := r.Move(console, instruction.(*MovementInstruction).Distance)
		if err != nil {
			return err
		}
		return nil
	case *RotationInstruction:
		err := r.Rotate(instruction.(*RotationInstruction).Orientation)
		if err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("unknown instruction: %v", instruction)
	}
}

func (r *Rover) CurrentPositionIsScented() bool {
	return r.Grid.IsScented(r.Position)
}

func (r *Rover) OnGridExit() error {
	r.Grid.AddScent(r.Position)
	return fmt.Errorf("Your rover fell off the grid at (%d, %d)!\n", r.Position.X, r.Position.Y)
}

type Movement struct {
	Direction Direction
	Distance  int
}
