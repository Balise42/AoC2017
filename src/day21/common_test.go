package day21

import (
	"testing"
	"reflect"
)

func TestSplit(t *testing.T) {
	grid := []string{"abcd", "efgh", "ijkl", "mnop"}
	splitGrid := Split(grid)

	assert(t, len(splitGrid) == 2)
	assert(t, len(splitGrid[0]) == 2)
	assert(t, splitGrid[0][0][0] == "ab")
	assert(t, splitGrid[0][0][1] == "ef")
	assert(t, splitGrid[0][1][0] == "cd")
	assert(t, splitGrid[0][1][1] == "gh")
	assert(t, splitGrid[1][0][0] == "ij")
	assert(t, splitGrid[1][0][1] == "mn")
	assert(t, splitGrid[1][1][0] == "kl")
	assert(t, splitGrid[1][1][1] == "op")
}

func TestFlipH2x2(t *testing.T) {
	grid := []string{"ab", "cd"}
	flip := FlipH(grid)
	assert(t, len(flip) == 2)
	assert(t, len(flip[0]) == 2)
	assert(t, flip[0] == "ba")
	assert(t, flip[1] == "dc")
}

func TestFlipH3x3(t *testing.T) {
	grid := []string{"abc", "def", "ghi"}
	flip := FlipH(grid)
	assert(t, len(flip) == 3)
	assert(t, len(flip[0]) == 3)
	assert(t, flip[0] == "cba")
	assert(t, flip[1] == "fed")
	assert(t, flip[2] == "ihg")
}

func TestFlipV3x3(t *testing.T) {
	grid := []string{"abc", "def", "ghi"}
	flip := FlipV(grid)
	assert(t, len(flip) == 3)
	assert(t, len(flip[0]) == 3)
	assert(t, flip[0] == "ghi")
	assert(t, flip[1] == "def")
	assert(t, flip[2] == "abc")
}

func TestFlipV2x2(t *testing.T) {
	grid := []string{"ab", "cd"}
	flip := FlipV(grid)
	assert(t, len(flip) == 2)
	assert(t, len(flip[0]) == 2)
	assert(t, flip[0] == "cd")
	assert(t, flip[1] == "ab")

}

func TestRotate2x2(t *testing.T) {
	grid := []string{"ab", "cd"}
	rot := Rotate(grid)
	assert(t, rot[0] == "bd")
	assert(t, rot[1] == "ac")
	assert(t, reflect.DeepEqual(grid,Rotate(Rotate(Rotate(rot)))))
}

func TestRotate3x3(t *testing.T) {
	grid := []string{"abc", "def", "ghi"}
	rot := Rotate(grid)
	assert(t, len(rot) == 3)
	assert(t, len(rot[0]) == 3)
	assert(t, rot[0] == "cfi")
	assert(t, rot[1] == "beh")
	assert(t, rot[2] == "adg")
	assert(t, reflect.DeepEqual(grid,Rotate(Rotate(Rotate(rot)))))
}

func TestRebuild(t *testing.T) {
	grid := [][]string {
		{"ab/ef", "cd/gh"},
		{"ij/mn", "kl/op"},
	}
	reb := Rebuild(grid)
	assert(t, len(reb) == 4)
	assert(t, reb[0] == "abcd")
	assert(t, reb[1] == "efgh")
	assert(t, reb[2] == "ijkl")
	assert(t, reb[3] == "mnop")
}

func assert(t *testing.T, cond bool) {
	if !cond {
		t.Fail()
	}
}
