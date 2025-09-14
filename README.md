# Mars Rover Explorer

A command-line Mars rover navigation simulator written in Go.

## Overview

Control rovers on a rectangular grid, issuing movement and rotation commands while avoiding falling off the edges. When a rover falls off, it leaves a "scent" that prevents future rovers from falling off at the same location.

## Usage

```bash
# Run the program
./marster-bot

# Run with debug output
./marster-bot --debug
```

## Input Format

1. **Grid size**: `x,y` (e.g., `5,5` creates a 5x5 grid)
2. **Rover position**: `x y D` where D is direction (N/S/E/W)
3. **Instructions**: String of commands:
   - `F` - Move forward one space
   - `L` - Rotate 90° left
   - `R` - Rotate 90° right

## Example

```
Grid: 5,5
Rover: 1 2 N
Instructions: RFRFFRFRF
Result: 1 3 N
```

## Running Tests

```bash
# Run all tests
go test ./...

# Run with verbose output
go test -v ./...
```

## Building

```bash
go build -o marster-bot
```