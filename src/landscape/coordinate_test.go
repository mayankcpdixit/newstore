package landscape

import "testing"

func TestNewCoordinate(t *testing.T) {
	for _, testcase := range []struct{
		input Coordinate
	}{{input: Coordinate{X:1, Y:2}}}{
		coordinate := NewCoordinate(testcase.input.X, testcase.input.Y)
		if coordinate != testcase.input {
			t.Errorf("TC failed.\nExpected: %+v, Received: %+v", testcase.input, coordinate)
		} else {
			t.Logf("Works for testcase: %+v", testcase)
		}
	}
}
