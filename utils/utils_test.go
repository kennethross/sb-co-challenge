package utils

import "testing"

func TestValidateMove(t *testing.T) {
	tests := []struct {
		name     string
		prev     Tick
		next     Tick
		expected bool
	}{
		{
			name:     "Test 1: Valid Move X Position",
			prev:     Tick{VelX: 1, VelY: 0},
			next:     Tick{VelX: 0, VelY: 0},
			expected: true,
		},
		{
			name:     "Test 2: Invalid Move X Position",
			prev:     Tick{VelX: 1, VelY: 0},
			next:     Tick{VelX: -1, VelY: 0},
			expected: false,
		},
		{
			name:     "Test 3: Valid Move Y Position",
			prev:     Tick{VelX: 0, VelY: 0},
			next:     Tick{VelX: 0, VelY: 1},
			expected: true,
		},
		{
			name:     "Test 4: Invalid Move Y Position",
			prev:     Tick{VelX: 0, VelY: 1},
			next:     Tick{VelX: 0, VelY: -1},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateMove(tt.prev, tt.next); got != tt.expected {
				t.Errorf("validateMove() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestValidateInBound(t *testing.T) {
	tests := []struct {
		name            string
		currentPosition CurrentPosition
		width           int
		height          int
		expected        bool
	}{
		{
			name:            "Test 1: Valid InBound",
			currentPosition: CurrentPosition{X: 2, Y: 2},
			width:           2,
			height:          2,
			expected:        true,
		},
		{
			name:            "Test 2: Invalid out of bound",
			currentPosition: CurrentPosition{X: 1, Y: 3},
			width:           2,
			height:          2,
			expected:        false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateInBound(tt.currentPosition, tt.width, tt.height); got != tt.expected {
				t.Errorf("validateInBound() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestNewPosition(t *testing.T) {
	tests := []struct {
		name     string
		position CurrentPosition
		tick     Tick
		expected CurrentPosition
	}{
		{
			name:     "Test 1: Valid New Position X",
			position: CurrentPosition{X: 1, Y: 3},
			tick:     Tick{Down, Neutral},
			expected: CurrentPosition{X: 0, Y: 3},
		},
		{
			name:     "Test 2: Valid New Position X",
			position: CurrentPosition{X: 3, Y: 1},
			tick:     Tick{Up, Neutral},
			expected: CurrentPosition{X: 4, Y: 1},
		},
		{
			name:     "Test 3: Valid New Position Y",
			position: CurrentPosition{X: 0, Y: 3},
			tick:     Tick{Neutral, Left},
			expected: CurrentPosition{X: 0, Y: 2},
		},
		{
			name:     "Test 4: Valid New Position X",
			position: CurrentPosition{X: 0, Y: 0},
			tick:     Tick{Neutral, Right},
			expected: CurrentPosition{X: 0, Y: 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.position.NewPosition(tt.tick)
			if tt.position != tt.expected {
				t.Errorf("newPosition() = %v, want %v", tt.position, tt.expected)
			}
		})
	}
}
