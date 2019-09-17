package main

import (
	"github.com/pkg/errors"
)

// Add the movements to a dictionary
type Walker struct {
	currentDirection float64
	directions       []string
	directionsMap    map[string]float64
}

func NewWalker() *Walker {
	return &Walker{
		0,
		[]string{"N", "E", "S", "W"},
		map[string]float64{},
	}
}

func (w *Walker) executeCommand(command *Command) error {
	switch command.commandString {
	case string(WALK):
		var mappedIndex int
		if w.currentDirection >= 0 {
			mappedIndex = int(w.currentDirection) % 4
		} else {
			mappedIndex = -3 * int(w.currentDirection) % 4
		}
		w.directionsMap[w.directions[mappedIndex]] += command.amountString
	case string(TURN):
		switch command.amountString {
		case float64(LEFT):
			w.currentDirection -= 1
		case float64(RIGHT):
			w.currentDirection += 1
		default:
			return errors.Errorf("Expected 90 or 270 but got %v", command.amountString)
		}
	default:
		return errors.Errorf("Unable to read command from String %v", command)
	}
	return nil
}
