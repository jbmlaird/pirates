package main

import (
	"fmt"
	"math"
)

func generateFinalInstructions(directionsMap map[string]float64) []string {
	x := directionsMap["E"] - directionsMap["W"]
	y := directionsMap["N"] - directionsMap["S"]
	distance, angle := generateAngleAndDistance(x, y)
	return buildInstructionList(distance, angle)
}

func buildInstructionList(distance, angle float64) []string {
	if math.IsNaN(angle) {
		angle = 0
	}
	return []string{
		fmt.Sprintf("TURN %v", angle),
		fmt.Sprintf("WALK %v", distance),
	}
}

// https://en.wikipedia.org/wiki/Polar_coordinate_system#Converting_between_polar_and_Cartesian_coordinates
// using arccos as we have the values of both x and y
func generateAngleAndDistance(x, y float64) (distance, angleDegrees float64) {
	distance = math.Sqrt(x*x + y*y)
	var angle float64
	if y >= 0 {
		angle = math.Acos(x / distance)
	} else {
		angle = -math.Acos(x / distance)
	}
	angleDegrees = (angle * (180 / math.Pi)) - 90
	if angleDegrees < 0 {
		angleDegrees = 360 - angleDegrees
	}
	return distance, angleDegrees
}
