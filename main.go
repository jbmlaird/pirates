package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	commands := []string{"WALK 14", "TURN 90", "WALK 13", "TURN 270", "WALK 12", "TURN 270", "WALK 8", "TURN 270", "WALK 7", "TURN 270", "WALK 11", "TURN 270", "WALK 9", "TURN 90", "WALK 12", "TURN 90", "WALK 15", "TURN 270", "WALK 15", "TURN 270", "WALK 10", "TURN 90", "WALK 9", "TURN 90", "WALK 14", "TURN 90", "WALK 6", "TURN 270", "WALK 7", "TURN 90", "WALK 6", "TURN 90", "WALK 10", "TURN 270", "WALK 15", "TURN 270", "WALK 15", "TURN 270", "WALK 14", "TURN 90", "WALK 6", "TURN 90", "WALK 6", "TURN 270",
		"WALK 6", "TURN 270", "WALK 10", "TURN 270", "WALK 15", "TURN 270", "WALK 13", "TURN 270", "WALK 13", "TURN 270", "WALK 6", "TURN 90", "WALK 13", "TURN 270", "WALK 10", "TURN 90", "WALK 14", "TURN 270", "WALK 12", "TURN 270", "WALK 10", "TURN 270", "WALK 6", "TURN 270", "WALK 8", "TURN 90", "WALK 11", "TURN 90", "WALK 10", "TURN 270", "WALK 14", "TURN 90", "WALK 15", "TURN 270", "WALK 12", "TURN 90", "WALK 14", "TURN 90", "WALK 12", "TURN 90", "WALK 12", "TURN 90", "WALK 9", "TURN 90", "WALK 14", "TURN 90", "WALK 7", "TURN 270", "WALK 10", "TURN 90", "WALK 13", "TURN 270", "WALK 8", "TURN 270", "WALK 14", "TURN 270", "WALK 9"}

	finalInstruction, err := walk(commands)
	if err != nil {
		panic(err)
	}
	fmt.Println(finalInstruction)
	fmt.Printf("Execution time: %v", time.Since(start))
}

var errEmptySlice = errors.New("awkward. You got the parchment wet and all the instructions seem to have disappeared 🤔")

func walk(commands []string) ([]string, error) {
	if len(commands) < 1 {
		return nil, errEmptySlice
	}
	walker := NewWalker()

	for _, command := range commands {
		parsedCommand, err := parseCommand(command)
		if err != nil {
			fmt.Printf("Unable to parse command with String: %v", command)
		}
		if err := walker.executeCommand(parsedCommand); err != nil {
			fmt.Printf("Invalid parsed command: %v", parsedCommand)
		}
	}

	return generateFinalInstructions(walker.directionsMap), nil
}
