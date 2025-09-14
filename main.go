package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"marster-bot/input"
	"marster-bot/mars"
	"marster-bot/output"
	"os"

	"github.com/urfave/cli/v3"
)

func processRover(console *output.Console, grid *mars.Grid, roverNum int) error {
	console.Blank()
	console.Header(fmt.Sprintf("Rover #%d", roverNum))
	console.Divider()

	rover, err := input.CollectRoverFromInput(console, grid)
	if err != nil {
		return err
	}

	instructions, err := input.CollectInstructionsFromInput(console)
	if err != nil {
		return err
	}

	// Process the rover (placeholder for now)
	console.Debug("No. instructions: %d", len(*instructions))
	console.Blank()
	console.Info("Processing rover movements...")

	for _, instruction := range *instructions {
		console.Debug("Current position: %d %d %s", rover.Position.X, rover.Position.Y, rover.Direction)
		console.Debug("Processing instruction: %v", instruction)
		err := rover.Instruct(console, instruction)

		if err != nil {
			return err
		}
	}

	console.Success("Final position:  %d %d %s", rover.Position.X, rover.Position.Y, rover.Direction)

	return nil
}

func runRoverSimulation(console *output.Console) error {
	console.HeaderWithBorder("Mars Rover Explorer")
	console.Blank()

	grid, err := input.CollectGridFromInput(console)

	if err != nil {
		return err
	}

	roverNum := 1
	for {
		err := processRover(console, grid, roverNum)
		if err != nil {
			if err.Error() == "exit" {
				console.Blank()
				console.Success("Thank you for using Mars Rover Explorer!")
				break
			}
			console.Error("Error processing rover #%d: %v", roverNum, err)

			// If there was an error with this rover, ask if they want to try again
			response, _ := console.Prompt("Would you like to add another rover? (Y/n): ")
			if response != "" && response != "y" {
				console.Success("Thank you for using Mars Rover Explorer!")
				break
			}
		}

		roverNum++
		console.Blank()
		console.Divider()
	}

	return nil
}

func main() {
	app := &cli.Command{
		Name:  "Marster Bot",
		Usage: "A Mars rover navigation simulator",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "debug",
				Usage: "Enable debug output",
				Value: false,
			},
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			debugMode := c.Bool("debug")
			reader := bufio.NewReader(os.Stdin)
			console := output.NewConsole(*reader, debugMode)
			return runRoverSimulation(console)
		},
	}

	err := app.Run(context.Background(), os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
