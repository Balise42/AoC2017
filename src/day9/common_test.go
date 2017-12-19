package day9

import (
	"reflect"
	"testing"
)

func TestCleanData1(t *testing.T) {
	testCases := make(map[string][]byte)

	testCases["{}"] = []byte("{}")
	testCases["{{{}}}"] = []byte("{{{}}}")
	testCases["{{},{}}"] = []byte("{{},{}}")
	testCases["{<{},{},{{}}>}"] = []byte("{}")
	testCases["{<a>,<a>,<a>,<a>}"] = []byte("{,,,}")
	testCases["{{<a>},{<a>},{<a>},{<a>}}"] = []byte("{{},{},{},{}}")
	testCases["{{<!>},{<!>},{<!>},{<a>}}"] = []byte("{{}}")

	for k, v := range testCases {
		res, _ := CleanData([]byte(k))
		if !reflect.DeepEqual(res, v) {
			t.Error("not equal: ", string(v), " ", string(res))
		}
	}
}

func TestCountGarbage(t *testing.T) {
	testCases := make(map[string]int)

	testCases["<>"] = 0
	testCases["<random characters>"] = 17
	testCases["<<<<>"] = 3
	testCases["<{!>}>"] = 2
	testCases["<!!>"] = 0
	testCases["<!!!>>"] = 0
	testCases["<{o\"i!a,<{i<a>"] = 10

	for k, v := range testCases {
		_, res := CleanData([]byte(k))
		if res != v {
			t.Error("not equal for ", k, " : ", v, " ", res)
		}
	}
}

func TestComputeScore(t *testing.T) {
	testCases := make(map[string]int)
	testCases["{}"] = 1
	testCases["{{{}}}"] = 6
	testCases["{{},{}}"] = 5
	testCases["{,,,}"] = 1
	testCases["{{},{},{},{}}"] = 9
	testCases["{{}}"] = 3

	for k, v := range testCases {
		res := ComputeScore([]byte(k))
		if res != v {
			t.Error("not equal for ", k, " : ", v, " ", res)
		}
	}

}
