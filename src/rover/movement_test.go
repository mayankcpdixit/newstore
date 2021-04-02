package rover

import (
	"fmt"
	"testing"
)

func TestRover_Rotate(t *testing.T) {
	testcases := []struct {
		InputDirection 		string
		Clockwise 			bool
		ExpectedDirection 	string
	}{{
		InputDirection: "east",
		Clockwise: false,
		ExpectedDirection: "north",
	},{
		InputDirection: "north",
		Clockwise: false,
		ExpectedDirection: "west",
	},{
		InputDirection: "west",
		Clockwise: false,
		ExpectedDirection: "south",
	},{
		InputDirection: "south",
		Clockwise: false,
		ExpectedDirection: "east",
	},{
		InputDirection: "east",
		Clockwise: true,
		ExpectedDirection: "south",
	},{
		InputDirection: "south",
		Clockwise: true,
		ExpectedDirection: "west",
	},{
		InputDirection: "west",
		Clockwise: true,
		ExpectedDirection: "north",
	},{
		InputDirection: "north",
		Clockwise: true,
		ExpectedDirection: "east",
	}}

	for _, testcase := range testcases{
		rover := New(0, 0, testcase.InputDirection);
		rover.Rotate(testcase.Clockwise)

		if rover.Direction != testcase.ExpectedDirection {
			t.Errorf("Testcase failed.\n Expected direction: %s\n Received: %s\n. Testcase: %+v\n", testcase.ExpectedDirection, rover.Direction, testcase)
		}
	}
}

func TestRover_Move(t *testing.T) {
	type testcase struct {
		StartX int
		StartY int
		StartDirection string
		InputCommand string
		ExpectedX int
		ExpectedY int
		ExpectedDirection string
	}

	tcs := []testcase{{
		StartX: 0, StartY: 0, StartDirection: "east",
		InputCommand: "F",
		ExpectedX: 1, ExpectedY: 0, ExpectedDirection: "east",
	}, {
		StartX: 0, StartY: 0, StartDirection: "east",
		InputCommand: "B",
		ExpectedX: -1, ExpectedY: 0, ExpectedDirection: "east",
	}, {
		StartX: 0, StartY: 0, StartDirection: "east",
		InputCommand: "L",
		ExpectedX: 0, ExpectedY: 0, ExpectedDirection: "north",
	}, {
		StartX: 0, StartY: 0, StartDirection: "east",
		InputCommand: "R",
		ExpectedX: 0, ExpectedY: 0, ExpectedDirection: "south",
	}, {
		StartX: 0, StartY: 0, StartDirection: "west",
		InputCommand: "F",
		ExpectedX: -1, ExpectedY: 0, ExpectedDirection: "west",
	}, {
		StartX: 0, StartY: 0, StartDirection: "west",
		InputCommand: "B",
		ExpectedX: 1, ExpectedY: 0, ExpectedDirection: "west",
	}, {
		StartX: 0, StartY: 0, StartDirection: "west",
		InputCommand: "L",
		ExpectedX: 0, ExpectedY: 0, ExpectedDirection: "south",
	}, {
		StartX: 0, StartY: 0, StartDirection: "west",
		InputCommand: "R",
		ExpectedX: 0, ExpectedY: 0, ExpectedDirection: "north",
	}, {
		StartX: 0, StartY: 0, StartDirection: "north",
		InputCommand: "F",
		ExpectedX: 0, ExpectedY: 1, ExpectedDirection: "north",
	}, {
		StartX: 0, StartY: 0, StartDirection: "north",
		InputCommand: "B",
		ExpectedX: 0, ExpectedY: -1, ExpectedDirection: "north",
	}, {
		StartX: 0, StartY: 0, StartDirection: "north",
		InputCommand: "L",
		ExpectedX: 0, ExpectedY: 0, ExpectedDirection: "west",
	}, {
		StartX: 0, StartY: 0, StartDirection: "north",
		InputCommand: "R",
		ExpectedX: 0, ExpectedY: 0, ExpectedDirection: "east",
	}, {
		StartX: 0, StartY: 0, StartDirection: "south",
		InputCommand: "F",
		ExpectedX: 0, ExpectedY: -1, ExpectedDirection: "south",
	}, {
		StartX: 0, StartY: 0, StartDirection: "south",
		InputCommand: "B",
		ExpectedX: 0, ExpectedY: 1, ExpectedDirection: "south",
	}, {
		StartX: 0, StartY: 0, StartDirection: "south",
		InputCommand: "L",
		ExpectedX: 0, ExpectedY: 0, ExpectedDirection: "east",
	}, {
		StartX: 0, StartY: 0, StartDirection: "south",
		InputCommand: "R",
		ExpectedX: 0, ExpectedY: 0, ExpectedDirection: "west",
	}, {
		StartX: 4, StartY: 2, StartDirection: "east",
		InputCommand: "FLFFFRFLB",
		ExpectedX: 6, ExpectedY: 4, ExpectedDirection: "north",
	}, {

	}}

	for _,tc := range tcs {
		rover := New(tc.StartX, tc.StartY, tc.StartDirection);
		rover.Move(tc.InputCommand);
		if rover.X != tc.ExpectedX || rover.Y != tc.ExpectedY || rover.Direction != tc.ExpectedDirection {
			t.Errorf("TC Failed. \nExpected: (%d, %d) %s, \nReceived: (%d, %d) %s\nfor tc: %+v", tc.ExpectedX,
			tc.ExpectedY, tc.ExpectedDirection, rover.X, rover.Y, rover.Direction, tc);
		} else {
			fmt.Printf("TC Passed. \nExpected: (%d, %d) %s, \nReceived: (%d, %d) %s\n", tc.ExpectedX, 
			tc.ExpectedY, tc.ExpectedDirection, rover.X, rover.Y, rover.Direction);
		}
	}

	
} 