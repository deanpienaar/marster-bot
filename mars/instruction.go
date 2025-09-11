package mars

import "fmt"

type Instruction interface {
	String() string
	isInstruction()
}

type MovementInstruction struct {
	Distance int8
}

func (m MovementInstruction) String() string {
	return fmt.Sprintf("Move (%d)", m.Distance)
}

func (m MovementInstruction) isInstruction() {}

type RotationInstruction struct {
	Orientation Rotation
}

func (r RotationInstruction) String() string {
	return fmt.Sprintf("Rotate (%d)", r.Orientation)
}

func (r RotationInstruction) isInstruction() {}

func NewMovementInstruction(direction int8) *MovementInstruction {
	return &MovementInstruction{
		Distance: direction,
	}
}

func NewOrientationInstruction(orientation Rotation) *RotationInstruction {
	return &RotationInstruction{
		Orientation: orientation,
	}
}
