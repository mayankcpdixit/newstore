package rover

import (
	"newstore/src/landscape"
	"testing"
)

func TestRover_Move(t *testing.T) {
	testcases := []struct {
		inputX    int
		inputY    int
	}{{inputX: 20, inputY:-20}}

	for _, testcase := range testcases {
		rover := New(0, 0, "east", &landscape.Landscape{});
		rover.Move(testcase.inputX, testcase.inputY)

		if rover.X != testcase.inputX || rover.Y != testcase.inputY {
			t.Errorf("Testcase failed.\n Expected Coordinates: (%d, %d)\n " +
				"Received: Coordinates: (%d, %d)\n. Testcase: %+v\n",
				testcase.inputX, testcase.inputY, rover.X, rover.Y, testcase)
		} else {
			t.Logf("Works for testcase: %+v", testcase)
		}
	}
}

func TestRover_GetNextCoordinate(t *testing.T) {
	testcases := []struct {
		inputX    int
		inputY    int
		direction string
		movement  Movement
		expectedX int
		expectedY int
	}{{inputX: 0, inputY: 0, movement: MOVE_FORWARD, direction: "east", expectedX: 1, expectedY: 0},
		{inputX: 0, inputY: 0, movement: MOVE_FORWARD, direction: "west", expectedX: -1, expectedY: 0},
		{inputX: 0, inputY: 0, movement: MOVE_FORWARD, direction: "north", expectedX: 0, expectedY: 1},
		{inputX: 0, inputY: 0, movement: MOVE_FORWARD, direction: "south", expectedX: 0, expectedY: -1},
		{inputX: 0, inputY: 0, movement: MOVE_BACKWARD, direction: "east", expectedX: -1, expectedY: 0},
		{inputX: 0, inputY: 0, movement: MOVE_BACKWARD, direction: "west", expectedX: 1, expectedY: 0},
		{inputX: 0, inputY: 0, movement: MOVE_BACKWARD, direction: "north", expectedX: 0, expectedY: -1},
		{inputX: 0, inputY: 0, movement: MOVE_BACKWARD, direction: "south", expectedX: 0, expectedY: 1}}

	for _, testcase := range testcases {
		rover := New(0, 0, testcase.direction, &landscape.Landscape{});
		x,y := rover.GetNextCoordinate(testcase.movement)

		if x != testcase.expectedX || y != testcase.expectedY {
			t.Errorf("Testcase failed.\n Expected Coordinates: (%d, %d)\n " +
				"Received: Coordinates: (%d, %d)\n. Testcase: %+v\n",
				testcase.expectedX, testcase.expectedY, rover.X, rover.Y, testcase)
		} else {
			t.Logf("Works for testcase: %+v", testcase)
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
		rover := New(0, 0, testcase.inputDirection, &landscape.Landscape{});
		rover.Rotate(testcase.rotation)

		if rover.Direction != testcase.expectedDirection {
			t.Errorf("Testcase failed.\n Expected direction: %s\n " +
				"Received: %s\n. Testcase: %+v\n",
				testcase.expectedDirection, rover.Direction, testcase)
		} else {
			t.Logf("Works for testcase: %+v", testcase)
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
		Obstacles [][]int
		ExpectedState State
	}

	tcs := []testcase{{
		StartX: 0, StartY: 0, StartDirection: "east", InputCommand: "F",
		ExpectedX: 1, ExpectedY: 0, ExpectedDirection: "east", ExpectedState: MOBILE,
	}, {
		StartX: 0, StartY: 0, StartDirection: "east", InputCommand: "B",
		ExpectedX: -1, ExpectedY: 0, ExpectedDirection: "east", ExpectedState: MOBILE,
	}, {
		StartX: 0, StartY: 0, StartDirection: "east", InputCommand: "L",
		ExpectedX: 0, ExpectedY: 0, ExpectedDirection: "north", ExpectedState: MOBILE,
	}, {
		StartX: 0, StartY: 0, StartDirection: "east", InputCommand: "R",
		ExpectedX: 0, ExpectedY: 0, ExpectedDirection: "south", ExpectedState: MOBILE,
	}, {
		StartX: 0, StartY: 0, StartDirection: "west", InputCommand: "F",
		ExpectedX: -1, ExpectedY: 0, ExpectedDirection: "west", ExpectedState: MOBILE,
	}, {
		StartX: 0, StartY: 0, StartDirection: "west", InputCommand: "B",
		ExpectedX: 1, ExpectedY: 0, ExpectedDirection: "west", ExpectedState: MOBILE,
	}, {
		StartX: 0, StartY: 0, StartDirection: "west", InputCommand: "L",
		ExpectedX: 0, ExpectedY: 0, ExpectedDirection: "south", ExpectedState: MOBILE,
	}, {
		StartX: 0, StartY: 0, StartDirection: "west", InputCommand: "R",
		ExpectedX: 0, ExpectedY: 0, ExpectedDirection: "north", ExpectedState: MOBILE,
	}, {
		StartX: 0, StartY: 0, StartDirection: "north", InputCommand: "F",
		ExpectedX: 0, ExpectedY: 1, ExpectedDirection: "north", ExpectedState: MOBILE,
	}, {
		StartX: 0, StartY: 0, StartDirection: "north", InputCommand: "B",
		ExpectedX: 0, ExpectedY: -1, ExpectedDirection: "north", ExpectedState: MOBILE,
	}, {
		StartX: 0, StartY: 0, StartDirection: "north", InputCommand: "L",
		ExpectedX: 0, ExpectedY: 0, ExpectedDirection: "west", ExpectedState: MOBILE,
	}, {
		StartX: 0, StartY: 0, StartDirection: "north", InputCommand: "R",
		ExpectedX: 0, ExpectedY: 0, ExpectedDirection: "east", ExpectedState: MOBILE,
	}, {
		StartX: 0, StartY: 0, StartDirection: "south", InputCommand: "F",
		ExpectedX: 0, ExpectedY: -1, ExpectedDirection: "south", ExpectedState: MOBILE,
	}, {
		StartX: 0, StartY: 0, StartDirection: "south", InputCommand: "B",
		ExpectedX: 0, ExpectedY: 1, ExpectedDirection: "south", ExpectedState: MOBILE,
	}, {
		StartX: 0, StartY: 0, StartDirection: "south", InputCommand: "L",
		ExpectedX: 0, ExpectedY: 0, ExpectedDirection: "east", ExpectedState: MOBILE,
	}, {
		StartX: 0, StartY: 0, StartDirection: "south", InputCommand: "R",
		ExpectedX: 0, ExpectedY: 0, ExpectedDirection: "west", ExpectedState: MOBILE,
	}, {
		StartX: 4, StartY: 2, StartDirection: "east", InputCommand: "FLFFFRFLB",
		ExpectedX: 6, ExpectedY: 4, ExpectedDirection: "north", ExpectedState: MOBILE,
	}, {
		StartX: 4, StartY: 2, StartDirection: "east", InputCommand: "FLFFFRFLB",
		ExpectedX: 6, ExpectedY: 5, ExpectedDirection: "north",
		Obstacles: [][]int{{6,4}}, ExpectedState: STOPPED,
	}, {
		StartX: 0, StartY: 0, StartDirection: "east", InputCommand: "F",
		Obstacles: [][]int{{1,0}},
		ExpectedX: 0, ExpectedY: 0, ExpectedDirection: "east", ExpectedState: STOPPED,
	},}

	for _,tc := range tcs {
		land := landscape.Init(tc.Obstacles)
		rover := New(tc.StartX, tc.StartY, tc.StartDirection, land);
		returnedState := rover.Listen(tc.InputCommand);
		if returnedState != tc.ExpectedState {
			t.Errorf("TC failed. %+v", tc)
		} else if rover.X != tc.ExpectedX || rover.Y != tc.ExpectedY || rover.Direction != tc.ExpectedDirection {
			t.Errorf("TC Failed. \nExpected: (%d, %d) %s, \n" +
				"Received: (%d, %d) %s\nfor tc: %+v", tc.ExpectedX,
				tc.ExpectedY, tc.ExpectedDirection, rover.X, rover.Y, rover.Direction, tc);
		} else {
			t.Logf("Works for testcase: %+v", tc)
		}
	}

} 