package main

type CommandString string

const (
	WALK CommandString = "WALK"
	TURN CommandString = "TURN"
)

type LeftOrRight int

const (
	LEFT  LeftOrRight = 90
	RIGHT LeftOrRight = 270
)

type Command struct {
	commandString string
	amountString  float64
}

func NewCommand(command string, amount float64) *Command {
	return &Command{
		command,
		amount,
	}
}
