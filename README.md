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

## Technical Decisions
### Choice of Go
We chose Go for its ease in setting up the repo, as well as compiling the program. It also had nice lightweight utils
for building CLI apps. It's super easy to get started, and simple enough for the project timelines.

Other options included:
* Python + Commander lib: Python is very finicky to set up if one does not have it running already, but Commander is awesome!
* TypeScript + oclif: A little too familiar in terms of challenges, can be harder to read and maintain than Go's simplicity.
* Rust + clap: Rust would've been great with longer time constraints.
* .NET: A little too heavy

### Input/Output
The intention is for these modules to have new and other mechanisms for input/output. Right now
Console is way too strongly coupled.

### Parser
Validation currently takes place within the parser but could be moved into a separate module for mixing and matching
input mechanisms.

### Running the Simulation
Concurrency: The simulation is currently single-threaded and would require thread-safe maps if we were to allow processing
multiple rovers at once. This does not seem necessary yet; especially given the only existing input mechanism is
console.

Entrypoint: The entrypoint is doing a lot of work that has to be repeated in tests. This should be refactored so
that the engine is available across whichever input we are using.

### Extensibility
Capability has been added for instructions like moving backwards (B), turning around (T), and moving at longer distances (F3/B2), but these
are not yet implemented.
