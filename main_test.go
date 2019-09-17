package main

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	cases := []struct {
		name string
		got  []string
		want []string
	}{
		{
			"Sam's input",
			[]string{"WALK 14", "TURN 90", "WALK 13", "TURN 270", "WALK 12", "TURN 270", "WALK 8", "TURN 270", "WALK 7", "TURN 270", "WALK 11", "TURN 270", "WALK 9", "TURN 90", "WALK 12", "TURN 90", "WALK 15", "TURN 270", "WALK 15", "TURN 270", "WALK 10", "TURN 90", "WALK 9", "TURN 90", "WALK 14", "TURN 90", "WALK 6", "TURN 270", "WALK 7", "TURN 90", "WALK 6", "TURN 90", "WALK 10", "TURN 270", "WALK 15", "TURN 270", "WALK 15", "TURN 270", "WALK 14", "TURN 90", "WALK 6", "TURN 90", "WALK 6", "TURN 270",
				"WALK 6", "TURN 270", "WALK 10", "TURN 270", "WALK 15", "TURN 270", "WALK 13", "TURN 270", "WALK 13", "TURN 270", "WALK 6", "TURN 90", "WALK 13", "TURN 270", "WALK 10", "TURN 90", "WALK 14", "TURN 270", "WALK 12", "TURN 270", "WALK 10", "TURN 270", "WALK 6", "TURN 270", "WALK 8", "TURN 90", "WALK 11", "TURN 90", "WALK 10", "TURN 270", "WALK 14", "TURN 90", "WALK 15", "TURN 270", "WALK 12", "TURN 90", "WALK 14", "TURN 90", "WALK 12", "TURN 90", "WALK 12", "TURN 90", "WALK 9", "TURN 90", "WALK 14", "TURN 90", "WALK 7", "TURN 270", "WALK 10", "TURN 90", "WALK 13", "TURN 270", "WALK 8", "TURN 270", "WALK 14", "TURN 270", "WALK 9"},
			[]string{"TURN 63.434948822921996", "WALK 26.832815729997478"},
		},
		{
			"Return to (0,0)",
			[]string{"WALK 2", "TURN 90", "WALK 2", "TURN 90", "WALK 2", "TURN 90", "WALK 2"},
			[]string{"TURN 0", "WALK 0"},
		},
		{
			"First quadrant",
			[]string{"WALK 2", "TURN 270", "WALK 2"},
			[]string{"TURN 315", fmt.Sprintf("WALK %v", math.Sqrt(2*2 + 2*2))},
		},
		{
			"Second quadrant",
			[]string{"WALK 2", "TURN 90", "WALK 2"},
			[]string{"TURN 45", fmt.Sprintf("WALK %v", math.Sqrt(2*2 + 2*2))},
		},
		{
			"Third quadrant",
			[]string{"TURN 90", "WALK 2", "TURN 90", "WALK 2"},
			[]string{"TURN 135", fmt.Sprintf("WALK %v", math.Sqrt(2*2 + 2*2))},
		},
		{
			"Fourth quadrant",
			[]string{"TURN 270", "WALK 2", "TURN 270", "WALK 2"},
			[]string{"TURN 225", fmt.Sprintf("WALK %v", math.Sqrt(2*2 + 2*2))},
		},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			got, err := walk(test.got)
			assertNoError(t, err)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("got %v but wanted %v", got, test.want)
			}
		})
	}

	t.Run("empty list commands throws errEmptySlice", func(t *testing.T) {
		_, err := walk([]string{})
		if err.Error() != errEmptySlice.Error() {
			t.Errorf("Unexpected error: %v", err)
		}
	})
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("Got error: %v", err)
	}
}
