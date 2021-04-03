package landscape

import (
	"reflect"
	"testing"
)

func TestInit(t *testing.T) {
	for _, testcase := range []struct{
		obstacles [][]int
	}{{obstacles: [][]int{{1,1}, {1,2}}}} {
		landscape := Init(testcase.obstacles);
		if !reflect.DeepEqual(landscape.Obstacles, testcase.obstacles) {
			t.Errorf("TC Failed. Obstacle not found in the List. TC: %+v, Landscape: %+v\n", testcase, landscape)
		} else {
			t.Logf("Works for testcase: %+v", testcase)
		}
	}
}

func TestLandscape_HasObstacleAt(t *testing.T) {
	for _, testcase := range []struct{
		obstacles [][]int
		testInput []int
		shouldExist bool
	}{{obstacles: [][]int{{1,1}, {1,2}}, testInput: []int{1,1}, shouldExist: true},
		{obstacles: [][]int{{1,1}, {1,2}}, testInput: []int{1,3}, shouldExist: false}} {
		landscape := Init(testcase.obstacles)
		if landscape.HasObstacleAt(testcase.testInput[0], testcase.testInput[1]) != testcase.shouldExist {
			t.Errorf("TC Failed. Obstacle not found in the List. TC: %+v, Landscape: %+v\n", testcase, landscape)
		} else {
			t.Logf("Works for testcase: %+v", testcase)
		}
	}
}