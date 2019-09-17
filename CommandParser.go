package main

import (
	"github.com/pkg/errors"
	"strconv"
	"strings"
)

func parseCommand(command string) (*Command, error) {
	split := strings.Split(command, " ")
	if len(split) != 2 {
		return nil, errors.Errorf("Invalid command and amount supplied from String %v", command)
	}
	command, amount := split[0], split[1]
	amountInt, err := parseFloat(amount)
	if err != nil {
		return nil, errors.Wrapf(err, "Unable to create command with string %v", command)
	}
	return NewCommand(command, amountInt), nil
}

func parseFloat(amount string) (float64, error) {
	if leftOrRight, err := strconv.ParseFloat(amount, 64); err == nil {
		return leftOrRight, nil
	}
	return 0, errors.Errorf("Unable to parse number from String %v", amount)
}
