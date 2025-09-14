package input

import (
	"bufio"
	"marster-bot/mars"
	"marster-bot/output"
	"strings"
	"testing"
)

func TestCollectGridFromInput(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantErr   bool
		wantX     uint8
		wantY     uint8
		errMsg    string
	}{
		{
			name:    "Valid grid input",
			input:   "5,5\n",
			wantErr: false,
			wantX:   5,
			wantY:   5,
		},
		{
			name:    "Valid grid with spaces",
			input:   " 10 , 8 \n",
			wantErr: false,
			wantX:   10,
			wantY:   8,
		},
		{
			name:    "Invalid format - no comma",
			input:   "5 5\n",
			wantErr: true,
			errMsg:  "expected format 'x,y'",
		},
		{
			name:    "Invalid format - too many values",
			input:   "5,5,5\n",
			wantErr: true,
			errMsg:  "expected format 'x,y'",
		},
		{
			name:    "Invalid x - not a number",
			input:   "abc,5\n",
			wantErr: true,
			errMsg:  "invalid x boundary",
		},
		{
			name:    "Invalid y - not a number",
			input:   "5,xyz\n",
			wantErr: true,
			errMsg:  "invalid y boundary",
		},
		{
			name:    "Negative x value",
			input:   "-1,5\n",
			wantErr: true,
			errMsg:  "grid boundaries must be positive",
		},
		{
			name:    "Negative y value",
			input:   "5,-1\n",
			wantErr: true,
			errMsg:  "grid boundaries must be positive",
		},
		{
			name:    "Zero x dimension",
			input:   "0,5\n",
			wantErr: true,
			errMsg:  "grid must have non-zero dimensions",
		},
		{
			name:    "Zero y dimension",
			input:   "5,0\n",
			wantErr: true,
			errMsg:  "grid must have non-zero dimensions",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.input))
			console := output.NewConsole(*reader, false)

			grid, err := CollectGridFromInput(console)

			if tt.wantErr {
				if err == nil {
					t.Errorf("Expected error containing '%s', got nil", tt.errMsg)
				} else if !strings.Contains(err.Error(), tt.errMsg) {
					t.Errorf("Expected error containing '%s', got '%s'", tt.errMsg, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if grid == nil {
					t.Fatal("Expected non-nil grid")
				}
				if grid.XSize != tt.wantX || grid.YSize != tt.wantY {
					t.Errorf("Expected grid size (%d,%d), got (%d,%d)", tt.wantX, tt.wantY, grid.XSize, grid.YSize)
				}
			}
		})
	}
}

func TestCollectRoverFromInput(t *testing.T) {
	grid := mars.NewGrid(5, 5)

	tests := []struct {
		name      string
		input     string
		wantErr   bool
		wantX     int8
		wantY     int8
		wantDir   mars.Direction
		errMsg    string
	}{
		{
			name:    "Valid rover input North",
			input:   "1 2 N\n",
			wantErr: false,
			wantX:   1,
			wantY:   2,
			wantDir: mars.North,
		},
		{
			name:    "Valid rover input South",
			input:   "3 3 S\n",
			wantErr: false,
			wantX:   3,
			wantY:   3,
			wantDir: mars.South,
		},
		{
			name:    "Valid rover input East",
			input:   "0 0 E\n",
			wantErr: false,
			wantX:   0,
			wantY:   0,
			wantDir: mars.East,
		},
		{
			name:    "Valid rover input West",
			input:   "5 5 W\n",
			wantErr: false,
			wantX:   5,
			wantY:   5,
			wantDir: mars.West,
		},
		{
			name:    "Exit command",
			input:   "exit\n",
			wantErr: true,
			errMsg:  "exit",
		},
		{
			name:    "Exit command uppercase",
			input:   "EXIT\n",
			wantErr: true,
			errMsg:  "exit",
		},
		{
			name:    "Invalid format - too few parts",
			input:   "1 2\n",
			wantErr: true,
			errMsg:  "expected format 'x y D'",
		},
		{
			name:    "Invalid format - too many parts",
			input:   "1 2 N E\n",
			wantErr: true,
			errMsg:  "expected format 'x y D'",
		},
		{
			name:    "Invalid x - not a number",
			input:   "abc 2 N\n",
			wantErr: true,
			errMsg:  "invalid x position",
		},
		{
			name:    "Invalid y - not a number",
			input:   "1 xyz N\n",
			wantErr: true,
			errMsg:  "invalid y position",
		},
		{
			name:    "Invalid direction",
			input:   "1 2 X\n",
			wantErr: true,
			errMsg:  "invalid direction",
		},
		{
			name:    "X position out of bounds - negative",
			input:   "-1 2 N\n",
			wantErr: true,
			errMsg:  "x position -1 is outside grid bounds",
		},
		{
			name:    "X position out of bounds - too large",
			input:   "6 2 N\n",
			wantErr: true,
			errMsg:  "x position 6 is outside grid bounds",
		},
		{
			name:    "Y position out of bounds - negative",
			input:   "1 -1 N\n",
			wantErr: true,
			errMsg:  "y position -1 is outside grid bounds",
		},
		{
			name:    "Y position out of bounds - too large",
			input:   "1 6 N\n",
			wantErr: true,
			errMsg:  "y position 6 is outside grid bounds",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.input))
			console := output.NewConsole(*reader, false)

			rover, err := CollectRoverFromInput(console, grid)

			if tt.wantErr {
				if err == nil {
					t.Errorf("Expected error containing '%s', got nil", tt.errMsg)
				} else if !strings.Contains(err.Error(), tt.errMsg) {
					t.Errorf("Expected error containing '%s', got '%s'", tt.errMsg, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if rover == nil {
					t.Fatal("Expected non-nil rover")
				}
				if rover.Position.X != tt.wantX || rover.Position.Y != tt.wantY {
					t.Errorf("Expected position (%d,%d), got (%d,%d)", tt.wantX, tt.wantY, rover.Position.X, rover.Position.Y)
				}
				if !rover.Direction.Equals(tt.wantDir) {
					t.Errorf("Expected direction %v, got %v", tt.wantDir, rover.Direction)
				}
			}
		})
	}
}

func TestCollectInstructionsFromInput(t *testing.T) {
	tests := []struct {
		name         string
		input        string
		wantErr      bool
		wantCount    int
		errMsg       string
		validateInst func([]mars.Instruction) bool
	}{
		{
			name:      "Valid instructions - RLF",
			input:     "RLF\n",
			wantErr:   false,
			wantCount: 3,
			validateInst: func(inst []mars.Instruction) bool {
				// Check first is rotation right
				if _, ok := inst[0].(*mars.RotationInstruction); !ok {
					return false
				}
				// Check second is rotation left
				if _, ok := inst[1].(*mars.RotationInstruction); !ok {
					return false
				}
				// Check third is movement
				if _, ok := inst[2].(*mars.MovementInstruction); !ok {
					return false
				}
				return true
			},
		},
		{
			name:      "Valid instructions - lowercase",
			input:     "rlf\n",
			wantErr:   false,
			wantCount: 3,
		},
		{
			name:      "Valid instructions - mixed case",
			input:     "RlF\n",
			wantErr:   false,
			wantCount: 3,
		},
		{
			name:      "Valid instructions - with spaces",
			input:     "  RLF  \n",
			wantErr:   false,
			wantCount: 3,
		},
		{
			name:      "Valid instructions - long sequence",
			input:     "FFRFFRFFRFF\n",
			wantErr:   false,
			wantCount: 11,
		},
		{
			name:      "Single valid instruction",
			input:     "F\n",
			wantErr:   false,
			wantCount: 1,
		},
		{
			name:    "Empty instructions",
			input:   "\n",
			wantErr: true,
			errMsg:  "instructions cannot be empty",
		},
		{
			name:    "Only whitespace",
			input:   "   \n",
			wantErr: true,
			errMsg:  "instructions cannot be empty",
		},
		{
			name:    "Invalid instruction character",
			input:   "RLX\n",
			wantErr: true,
			errMsg:  "invalid instruction 'X' at position 2",
		},
		{
			name:    "Invalid instruction - number",
			input:   "RL5\n",
			wantErr: true,
			errMsg:  "invalid instruction '5' at position 2",
		},
		{
			name:    "Invalid instruction - special character",
			input:   "RL@\n",
			wantErr: true,
			errMsg:  "invalid instruction '@' at position 2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.input))
			console := output.NewConsole(*reader, false)

			instructions, err := CollectInstructionsFromInput(console)

			if tt.wantErr {
				if err == nil {
					t.Errorf("Expected error containing '%s', got nil", tt.errMsg)
				} else if !strings.Contains(err.Error(), tt.errMsg) {
					t.Errorf("Expected error containing '%s', got '%s'", tt.errMsg, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if instructions == nil {
					t.Fatal("Expected non-nil instructions")
				}
				if len(*instructions) != tt.wantCount {
					t.Errorf("Expected %d instructions, got %d", tt.wantCount, len(*instructions))
				}
				if tt.validateInst != nil && !tt.validateInst(*instructions) {
					t.Errorf("Instructions validation failed")
				}
			}
		})
	}
}