package landscape

import (
	"reflect"
	"testing"
)

func TestInit(t *testing.T) {
	for _, testcase := range []struct{
		obstacles []Coordinate
	}{{obstacles: []Coordinate{{X:1, Y:1}, {X:1, Y:2}}}} {
		landscape := NewLandscape(testcase.obstacles)
		if !reflect.DeepEqual(landscape.Obstacles, testcase.obstacles) {
			t.Errorf("TC Failed. Obstacle not found in the List. TC: %+v, Landscape: %+v\n", testcase, landscape)
		} else {
			t.Logf("Works for testcase: %+v", testcase)
		}
	}
}

func TestLandscape_HasObstacleAt(t *testing.T) {
	for _, testcase := range []struct{
		obstacles   []Coordinate
		testInput   Coordinate
		shouldExist bool
	}{{obstacles: []Coordinate{{X:1, Y:1}, {X:1, Y:2}}, testInput: Coordinate{X:1, Y:1}, shouldExist: true},
		{obstacles: []Coordinate{{X:1, Y:1}, {X:1, Y:2}}, testInput: Coordinate{X:1, Y:3}, shouldExist: false}} {
		landscape := NewLandscape(testcase.obstacles)
		if landscape.HasObstacleAt(testcase.testInput) != testcase.shouldExist {
			t.Errorf("TC Failed. Obstacle not found in the List. TC: %+v, Landscape: %+v\n", testcase, landscape)
		} else {
			t.Logf("Works for testcase: %+v", testcase)
		}
	}
}