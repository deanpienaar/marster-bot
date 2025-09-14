package mars

import (
	"bufio"
	"marster-bot/output"
	"strings"
	"testing"
)

type UnknownInstruction struct{}

func (u UnknownInstruction) String() string {
	return "Unknown"
}

func (u UnknownInstruction) isInstruction() {}

func TestRoverInstruct(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader(""))
	console := output.NewConsole(*reader, false)
	grid := NewGrid(10, 10)
	rover := NewRover(5, 5, North, grid)

	t.Run("MovementInstruction", func(t *testing.T) {
		instruction := NewMovementInstruction(2)
		err := rover.Instruct(console, instruction)
		if err != nil {
			t.Errorf("Expected no error for MovementInstruction, got: %v", err)
		}
		if rover.Position.Y != 7 {
			t.Errorf("Expected Y position to be 7, got: %d", rover.Position.Y)
		}
	})

	t.Run("RotationInstruction", func(t *testing.T) {
		instruction := NewOrientationInstruction(Right)
		initialDirection := rover.Direction
		err := rover.Instruct(console, instruction)
		if err != nil {
			t.Errorf("Expected no error for RotationInstruction, got: %v", err)
		}
		if rover.Direction == initialDirection {
			t.Errorf("Expected direction to change, but it remained: %v", rover.Direction)
		}
	})

	t.Run("UnknownInstruction", func(t *testing.T) {
		instruction := UnknownInstruction{}
		err := rover.Instruct(console, instruction)
		if err == nil {
			t.Errorf("Expected error for unknown instruction, got nil")
		}
	})
}
