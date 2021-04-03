package rover

import (
	"testing"
)

func TestRover_Move(t *testing.T) {
	testcases := []struct {
		inputX    int
		inputY    int
		direction string
		movement  Movement
		expectedX int
		expectedY int
	}{{inputX: 0, inputY: 0, movement: FORWARD, direction: "east", expectedX: 1, expectedY: 0},
		{inputX: 0, inputY: 0, movement: FORWARD, direction: "west", expectedX: -1, expectedY: 0},
		{inputX: 0, inputY: 0, movement: FORWARD, direction: "north", expectedX: 0, expectedY: 1},
		{inputX: 0, inputY: 0, movement: FORWARD, direction: "south", expectedX: 0, expectedY: -1},
		{inputX: 0, inputY: 0, movement: BACKWARD, direction: "east", expectedX: -1, expectedY: 0},
		{inputX: 0, inputY: 0, movement: BACKWARD, direction: "west", expectedX: 1, expectedY: 0},
		{inputX: 0, inputY: 0, movement: BACKWARD, direction: "north", expectedX: 0, expectedY: -1},
		{inputX: 0, inputY: 0, movement: BACKWARD, direction: "south", expectedX: 0, expectedY: 1}}

	for _, testcase := range testcases {
		rover := New(0, 0, testcase.direction);
		x,y := rover.GetNextCoordinate(testcase.movement)
		rover.Move(x,y)

		if rover.X != testcase.expectedX || rover.Y != testcase.expectedY {
			t.Errorf("Testcase failed.\n Expected Coordinates: (%d, %d)\n " +
				"Received: Coordinates: (%d, %d)\n. Testcase: %+v\n",
				testcase.expectedX, testcase.expectedY, rover.X, rover.Y, testcase)
		} else {
			t.Log("Works for testcase: ", testcase)
		}
	}
}

func TestRover_Rotate(t *testing.T) {
	testcases := []struct {
		inputDirection    string
		rotation          Rotation
		expectedDirection string
	}{{ inputDirection: "east", rotation: ANTICLOCKWISE, expectedDirection: "north"},
		{ inputDirection: "north", rotation: ANTICLOCKWISE, expectedDirection: "west"},
		{ inputDirection: "west", rotation: ANTICLOCKWISE, expectedDirection: "south"},
		{ inputDirection: "south", rotation: ANTICLOCKWISE, expectedDirection: "east"},
		{ inputDirection: "east", rotation: CLOCKWISE, expectedDirection: "south"},
		{ inputDirection: "south", rotation: CLOCKWISE, expectedDirection: "west"},
		{ inputDirection: "west", rotation: CLOCKWISE, expectedDirection: "north"},
		{ inputDirection: "north", rotation: CLOCKWISE, expectedDirection: "east"}}

	for _, testcase := range testcases{
		rover := New(0, 0, testcase.inputDirection);
		rover.Rotate(testcase.rotation)

		if rover.Direction != testcase.expectedDirection {
			t.Errorf("Testcase failed.\n Expected direction: %s\n " +
				"Received: %s\n. Testcase: %+v\n",
				testcase.expectedDirection, rover.Direction, testcase)
		} else {
			t.Log("Works for testcase: ", testcase)
		}
	}
}

func TestRover_Listen(t *testing.T) {
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
		StartX: 0, StartY: 0, StartDirection: "east", InputCommand: "F",
		ExpectedX: 1, ExpectedY: 0, ExpectedDirection: "east",
	}, {
		StartX: 0, StartY: 0, StartDirection: "east", InputCommand: "B",
		ExpectedX: -1, ExpectedY: 0, ExpectedDirection: "east",
	}, {
		StartX: 0, StartY: 0, StartDirection: "east", InputCommand: "L",
		ExpectedX: 0, ExpectedY: 0, ExpectedDirection: "north",
	}, {
		StartX: 0, StartY: 0, StartDirection: "east", InputCommand: "R",
		ExpectedX: 0, ExpectedY: 0, ExpectedDirection: "south",
	}, {
		StartX: 0, StartY: 0, StartDirection: "west", InputCommand: "F",
		ExpectedX: -1, ExpectedY: 0, ExpectedDirection: "west",
	}, {
		StartX: 0, StartY: 0, StartDirection: "west", InputCommand: "B",
		ExpectedX: 1, ExpectedY: 0, ExpectedDirection: "west",
	}, {
		StartX: 0, StartY: 0, StartDirection: "west", InputCommand: "L",
		ExpectedX: 0, ExpectedY: 0, ExpectedDirection: "south",
	}, {
		StartX: 0, StartY: 0, StartDirection: "west", InputCommand: "R",
		ExpectedX: 0, ExpectedY: 0, ExpectedDirection: "north",
	}, {
		StartX: 0, StartY: 0, StartDirection: "north", InputCommand: "F",
		ExpectedX: 0, ExpectedY: 1, ExpectedDirection: "north",
	}, {
		StartX: 0, StartY: 0, StartDirection: "north", InputCommand: "B",
		ExpectedX: 0, ExpectedY: -1, ExpectedDirection: "north",
	}, {
		StartX: 0, StartY: 0, StartDirection: "north", InputCommand: "L",
		ExpectedX: 0, ExpectedY: 0, ExpectedDirection: "west",
	}, {
		StartX: 0, StartY: 0, StartDirection: "north", InputCommand: "R",
		ExpectedX: 0, ExpectedY: 0, ExpectedDirection: "east",
	}, {
		StartX: 0, StartY: 0, StartDirection: "south", InputCommand: "F",
		ExpectedX: 0, ExpectedY: -1, ExpectedDirection: "south",
	}, {
		StartX: 0, StartY: 0, StartDirection: "south", InputCommand: "B",
		ExpectedX: 0, ExpectedY: 1, ExpectedDirection: "south",
	}, {
		StartX: 0, StartY: 0, StartDirection: "south", InputCommand: "L",
		ExpectedX: 0, ExpectedY: 0, ExpectedDirection: "east",
	}, {
		StartX: 0, StartY: 0, StartDirection: "south", InputCommand: "R",
		ExpectedX: 0, ExpectedY: 0, ExpectedDirection: "west",
	}, {
		StartX: 4, StartY: 2, StartDirection: "east", InputCommand: "FLFFFRFLB",
		ExpectedX: 6, ExpectedY: 4, ExpectedDirection: "north",
	}}

	for _,tc := range tcs {
		rover := New(tc.StartX, tc.StartY, tc.StartDirection);
		rover.Listen(tc.InputCommand);
		if rover.X != tc.ExpectedX || rover.Y != tc.ExpectedY || rover.Direction != tc.ExpectedDirection {
			t.Errorf("TC Failed. \nExpected: (%d, %d) %s, \n" +
				"Received: (%d, %d) %s\nfor tc: %+v", tc.ExpectedX,
				tc.ExpectedY, tc.ExpectedDirection, rover.X, rover.Y, rover.Direction, tc);
		} else {
			t.Logf("TC Passed. \nExpected: (%d, %d) %s, \n" +
				"Received: (%d, %d) %s\n", tc.ExpectedX,
				tc.ExpectedY, tc.ExpectedDirection, rover.X, rover.Y, rover.Direction);
		}
	}

} 