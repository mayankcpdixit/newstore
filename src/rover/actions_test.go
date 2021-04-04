package rover

import (
	"newstore/src/landscape"
	"testing"
)

var ORIGIN = landscape.Coordinate{X:0, Y:0}

func TestRover_Move(t *testing.T) {
	testcases := []struct {
		inputCoordinate landscape.Coordinate
	}{{inputCoordinate: landscape.Coordinate{X: -20, Y: 20}}}

	for _, testcase := range testcases {
		rover := New(ORIGIN, "east", &landscape.Landscape{})
		rover.Move(testcase.inputCoordinate)

		if rover.Coordinate != testcase.inputCoordinate {
			t.Errorf("Testcase failed.\n Expected Coordinates: %+v\n " +
				"Received: Coordinates: %+v\n. Testcase: %+v\n",
				testcase.inputCoordinate, rover.Coordinate, testcase)
		} else {
			t.Logf("Works for testcase: %+v", testcase)
		}
	}
}

func TestRover_GetNextCoordinateFromOrigin(t *testing.T) {
	testcases := []struct {
		direction          string
		movement           Movement
		expectedCoordinate landscape.Coordinate
	}{{movement: MOVE_FORWARD, direction: "east", expectedCoordinate: landscape.Coordinate{X: 1, Y:0}},
		{movement: MOVE_FORWARD, direction: "west", expectedCoordinate: landscape.Coordinate{X: -1, Y:0}},
		{movement: MOVE_FORWARD, direction: "north", expectedCoordinate: landscape.Coordinate{X: 0, Y:1}},
		{movement: MOVE_FORWARD, direction: "south", expectedCoordinate: landscape.Coordinate{X: 0, Y:-1}},
		{movement: MOVE_BACKWARD, direction: "east", expectedCoordinate: landscape.Coordinate{X: -1, Y:0}},
		{movement: MOVE_BACKWARD, direction: "west", expectedCoordinate: landscape.Coordinate{X: 1, Y:0}},
		{movement: MOVE_BACKWARD, direction: "north", expectedCoordinate: landscape.Coordinate{X: 0, Y:-1}},
		{movement: MOVE_BACKWARD, direction: "south", expectedCoordinate: landscape.Coordinate{X: 0, Y:1}},
	}

	for _, testcase := range testcases {
		rover := New(ORIGIN, testcase.direction, &landscape.Landscape{})
		nextCoordinate := rover.GetNextCoordinate(testcase.movement)

		if nextCoordinate != testcase.expectedCoordinate {
			t.Errorf("Testcase failed.\n Expected Coordinates: %+v\n " +
				"Received: Coordinates: %+v\n. Testcase: %+v\n",
				testcase.expectedCoordinate, rover.Coordinate, testcase)
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
		rover := New(ORIGIN, testcase.inputDirection, &landscape.Landscape{})
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
		coordinate         landscape.Coordinate
		direction          string
		command            string
		expectedCoordinate landscape.Coordinate
		expectedDirection  string
		obstacles          []landscape.Coordinate
		expectedState      State
	}

	tcs := []testcase{{
		coordinate: ORIGIN, direction: "east", command: "F",
		expectedCoordinate: landscape.Coordinate{ X: 1, Y: 0}, expectedDirection: "east", expectedState: MOBILE,
	}, {
		coordinate: ORIGIN, direction: "east", command: "B",
		expectedCoordinate: landscape.Coordinate{ X: -1, Y: 0}, expectedDirection: "east", expectedState: MOBILE,
	}, {
		coordinate: ORIGIN, direction: "east", command: "L",
		expectedCoordinate: ORIGIN, expectedDirection: "north", expectedState: MOBILE,
	}, {
		coordinate: ORIGIN, direction: "east", command: "R",
		expectedCoordinate: ORIGIN, expectedDirection: "south", expectedState: MOBILE,
	}, {
		coordinate: ORIGIN, direction: "west", command: "F",
		expectedCoordinate: landscape.Coordinate{ X: -1, Y: 0}, expectedDirection: "west", expectedState: MOBILE,
	}, {
		coordinate: ORIGIN, direction: "west", command: "B",
		expectedCoordinate: landscape.Coordinate{ X: 1, Y: 0}, expectedDirection: "west", expectedState: MOBILE,
	}, {
		coordinate: ORIGIN, direction: "west", command: "L",
		expectedCoordinate: ORIGIN, expectedDirection: "south", expectedState: MOBILE,
	}, {
		coordinate: ORIGIN, direction: "west", command: "R",
		expectedCoordinate: ORIGIN, expectedDirection: "north", expectedState: MOBILE,
	}, {
		coordinate: ORIGIN, direction: "north", command: "F",
		expectedCoordinate: landscape.Coordinate{ X: 0, Y: 1}, expectedDirection: "north", expectedState: MOBILE,
	}, {
		coordinate: ORIGIN, direction: "north", command: "B",
		expectedCoordinate: landscape.Coordinate{ X: 0, Y: -1}, expectedDirection: "north", expectedState: MOBILE,
	}, {
		coordinate: ORIGIN, direction: "north", command: "L",
		expectedCoordinate: ORIGIN, expectedDirection: "west", expectedState: MOBILE,
	}, {
		coordinate: ORIGIN, direction: "north", command: "R",
		expectedCoordinate: ORIGIN, expectedDirection: "east", expectedState: MOBILE,
	}, {
		coordinate: ORIGIN, direction: "south", command: "F",
		expectedCoordinate: landscape.Coordinate{ X: 0, Y: -1}, expectedDirection: "south", expectedState: MOBILE,
	}, {
		coordinate: ORIGIN, direction: "south", command: "B",
		expectedCoordinate: landscape.Coordinate{ X: 0, Y: 1}, expectedDirection: "south", expectedState: MOBILE,
	}, {
		coordinate: ORIGIN, direction: "south", command: "L",
		expectedCoordinate: ORIGIN, expectedDirection: "east", expectedState: MOBILE,
	}, {
		coordinate: ORIGIN, direction: "south", command: "R",
		expectedCoordinate: ORIGIN, expectedDirection: "west", expectedState: MOBILE,
	}, {
		coordinate: landscape.Coordinate{ X: 4, Y: 2}, direction: "east", command: "FLFFFRFLB",
		expectedCoordinate: landscape.Coordinate{ X: 6, Y: 4}, expectedDirection: "north", expectedState: MOBILE,
	}, {
		coordinate: landscape.Coordinate{ X: 4, Y: 2}, direction: "east", command: "FLFFFRFLB",
		expectedCoordinate: landscape.Coordinate{ X: 6, Y: 5}, expectedDirection: "north",
		obstacles: []landscape.Coordinate{{X: 6, Y: 4}}, expectedState: STOPPED,
	}, {
		coordinate: ORIGIN, direction: "east", command: "F",
		obstacles:          []landscape.Coordinate{ {X: 1, Y: 0}},
		expectedCoordinate: ORIGIN, expectedDirection: "east", expectedState: STOPPED,
	}}

	for _,tc := range tcs {
		land := landscape.NewLandscape(tc.obstacles)
		rover := New(tc.coordinate, tc.direction, land)
		returnedState := rover.Listen(tc.command)
		if returnedState != tc.expectedState {
			t.Errorf("TC failed. %+v", tc)
		} else if rover.Coordinate != tc.expectedCoordinate || rover.Direction != tc.expectedDirection {
			t.Errorf("TC Failed. \nExpected: %+v %s, \n" +
				"Received: %+v %s\nfor tc: %+v", tc.expectedCoordinate, tc.expectedDirection, rover.Coordinate, rover.Direction, tc)
		} else {
			t.Logf("Works for testcase: %+v", tc)
		}
	}
}

// TODO: Part 3
func TestRover_PathFinder(t *testing.T) {
	for _, testcase := range []struct {
		source landscape.Coordinate
		destination landscape.Coordinate
		obstacles []landscape.Coordinate
	}{{
		source: landscape.NewCoordinate(0, 0),
		destination: landscape.NewCoordinate(5,5),
	}}{
		land := landscape.NewLandscape(testcase.obstacles)
		rover := New(ORIGIN, "east", land)
		rover.PathFinder(testcase.destination)
		// TODO: Test Part 3
	}
}