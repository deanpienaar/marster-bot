package input

import (
	"fmt"
	"marster-bot/mars"
	"marster-bot/output"
	"strconv"
	"strings"
	"unicode/utf8"
)

func CollectGridFromInput(console *output.Console) (*mars.Grid, error) {
	// Get grid boundaries (only once)
	gridInput, err := console.Prompt("Enter grid upper-right coordinates (x,y): ")
	if err != nil {
		console.Error("Failed to read grid boundaries: %v", err)
		return nil, err
	}
	parts := strings.Split(gridInput, ",")

	if len(parts) != 2 {
		return nil, fmt.Errorf("expected format 'x,y' (e.g., '5,5'), got '%s'", gridInput)
	}

	maxX, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		return nil, fmt.Errorf("invalid x boundary: '%s' is not a number", parts[0])
	}

	maxY, err := strconv.Atoi(strings.TrimSpace(parts[1]))
	if err != nil {
		return nil, fmt.Errorf("invalid y boundary: '%s' is not a number", parts[1])
	}

	if maxX < 0 || maxY < 0 {
		return nil, fmt.Errorf("grid boundaries must be positive (got %d,%d)", maxX, maxY)
	}

	if maxX == 0 || maxY == 0 {
		return nil, fmt.Errorf("grid must have non-zero dimensions (got %d,%d)", maxX, maxY)
	}

	grid := mars.NewGrid(uint8(maxX), uint8(maxY))

	console.Success("Grid established: %dx%d", grid.XSize, grid.YSize)

	return grid, nil
}

func CollectRoverFromInput(console *output.Console, grid *mars.Grid) (*mars.Rover, error) {
	positionInput, err := console.Prompt("Enter rover position and direction (x y D) or 'exit' to quit: ")
	if err != nil {
		console.Error("Failed to read rover position: %v", err)
		return nil, err
	}

	if strings.ToLower(positionInput) == "exit" {
		return nil, fmt.Errorf("exit")
	}

	parts := strings.Fields(positionInput)

	if len(parts) != 3 {
		return nil, fmt.Errorf("expected format 'x y D' where D is N/S/E/W (e.g., '1 2 N'), got '%s'", positionInput)
	}

	x, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, fmt.Errorf("invalid x position: '%s' is not a number", parts[0])
	}

	y, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, fmt.Errorf("invalid y position: '%s' is not a number", parts[1])
	}

	directionCode, _ := utf8.DecodeRuneInString(parts[2])
	if directionCode != 'N' && directionCode != 'S' && directionCode != 'E' && directionCode != 'W' {
		return nil, fmt.Errorf("invalid direction '%s': must be N, S, E, or W", string(directionCode))
	}

	if x < 0 || uint8(x) > grid.XSize {
		return nil, fmt.Errorf("x position %d is outside grid bounds (0-%d)", x, grid.XSize)
	}

	if y < 0 || uint8(y) > grid.YSize {
		return nil, fmt.Errorf("y position %d is outside grid bounds (0-%d)", y, grid.YSize)
	}

	direction := mars.DirectionFromCode(directionCode)

	rover := mars.NewRover(int8(x), int8(y), direction, grid)

	console.Success("Rover positioned at (%d, %d) facing %s", rover.Position.X, rover.Position.Y, rover.Direction)

	return rover, nil
}

func CollectInstructionsFromInput(console *output.Console) (*[]mars.Instruction, error) {
	instructionInput, err := console.Prompt("Enter movement instructions (R=Right, L=Left, F=Forward): ")
	if err != nil {
		console.Error("Failed to read instructions: %v", err)
		return nil, err
	}
	instructionInput = strings.TrimSpace(strings.ToUpper(instructionInput))

	if len(instructionInput) == 0 {
		return nil, fmt.Errorf("instructions cannot be empty")
	}

	var instructions []mars.Instruction

	for i, char := range instructionInput {
		var instruction mars.Instruction

		switch char {
		case 'R':
			instruction = mars.NewOrientationInstruction(mars.Right)
		case 'L':
			instruction = mars.NewOrientationInstruction(mars.Left)
		// Case is always true, but if we add a
		case 'F':
			instruction = mars.NewMovementInstruction(1)
		default:
			return nil, fmt.Errorf("invalid instruction '%c' at position %d: only R, L, F are allowed", char, i)
		}

		console.Debug("New instruction: %v", instruction)
		instructions = append(instructions, instruction)
	}

	console.Success("Instructions received: %s", instructionInput)

	return &instructions, nil
}
