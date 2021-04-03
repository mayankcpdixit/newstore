package landscape

import (
	"testing"
)

func TestLandscape_HasObstacleAt(t *testing.T) {
	for _, testcase := range []struct{
		obstacles [][]int
		testInput []int
	}{{obstacles: [][]int{{1,1}, {1,2}}, testInput: []int{1,1}}} {
		landscape := Init(testcase.obstacles);
		if !landscape.HasObstacleAt(testcase.testInput[0], testcase.testInput[1]) {
			t.Errorf("TC Failed. Obstacle not found in the List. TC: %+v, Landscape: %+v\n", testcase, landscape)
		} else {
			t.Logf("Works for testcase: %+v", testcase)
		}
	}
}