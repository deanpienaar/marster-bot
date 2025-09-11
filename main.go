package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"marster-bot/internal/input"
	"marster-bot/internal/output"
	"marster-bot/internal/rover"
	"os"

	"github.com/urfave/cli/v3"
)

func calculateEndPositionByCollapse(instructions string) {
	currentDirection := rover.North
	movement := map[rover.Direction]int{
		rover.North: 0,
		rover.East:  0,
		rover.South: 0,
		rover.West:  0,
	}

	for _, char := range instructions {
		switch char {
		case 'F':
			movement[currentDirection] += 1
			break
		case 'R':
			currentDirection = rover.North
			break
		case 'L':
			currentDirection = rover.South
			break
		}
	}

}

func runRoverSimulation(reader *bufio.Reader, console *output.Console) error {
	console.HeaderWithBorder("Mars Rover Explorer")
	console.Blank()

	// Get grid boundaries
	console.Prompt("Enter grid upper-right coordinates (x,y): ")
	gridInput, err := reader.ReadString('\n')
	if err != nil {
		console.Error("Failed to read grid boundaries: %v", err)
		return err
	}

	gridBounds, err := input.ParseGridBounds(gridInput)
	if err != nil {
		console.Error("Invalid grid boundaries: %v", err)
		return err
	}
	console.Success("Grid established: %dx%d", gridBounds.MaxX, gridBounds.MaxY)

	// Get rover starting position
	console.Prompt("Enter rover position and direction (x y D): ")
	positionInput, err := reader.ReadString('\n')
	if err != nil {
		console.Error("Failed to read rover position: %v", err)
		return err
	}

	roverPos, err := input.ParseRoverPosition(positionInput, gridBounds)
	if err != nil {
		console.Error("Invalid rover position: %v", err)
		return err
	}
	console.Success("Rover positioned at (%d, %d) facing %s", roverPos.X, roverPos.Y, roverPos.Direction)

	// Get movement instructions
	console.Prompt("Enter movement instructions (R=Right, L=Left, F=Forward): ")
	instructionInput, err := reader.ReadString('\n')
	if err != nil {
		console.Error("Failed to read instructions: %v", err)
		return err
	}

	instructions, err := input.ValidateInstructions(instructionInput)
	if err != nil {
		console.Error("Invalid instructions: %v", err)
		return err
	}
	console.Success("Instructions validated: %s", instructions.Commands)

	marsRover := rover.NewRover(roverPos.X, roverPos.Y, roverPos.Direction)

	console.Blank()
	console.Divider()
	console.Info("Processing rover movements...")
	console.Data("Grid size", fmt.Sprintf("%dx%d", gridBounds.MaxX, gridBounds.MaxY))
	console.Data("Starting position", fmt.Sprintf("(%d, %d) %s", marsRover.Position.X, marsRover.Position.Y, marsRover.Direction))
	console.Data("Instructions", instructions.Commands)
	console.Divider()
	console.Blank()
	console.Warning("Movement execution not yet implemented")

	return nil
}

func main() {
	app := &cli.Command{
		Name:  "Marster Bot",
		Usage: "A Mars rover navigation simulator",
		Action: func(ctx context.Context, c *cli.Command) error {
			reader := bufio.NewReader(os.Stdin)
			console := output.NewConsole()
			return runRoverSimulation(reader, console)
		},
	}

	err := app.Run(context.Background(), os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
