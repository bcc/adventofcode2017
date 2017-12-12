package main

import (
	"testing"
)

var tests = []pipedata{
	{0, []int{2}},
	{1, []int{1}},
	{2, []int{0, 3, 4}},
	{3, []int{2, 4}},
	{4, []int{2, 3, 6}},
	{5, []int{6}},
	{6, []int{4, 5}},
}

func TestCountPipes(t *testing.T) {
	v := CountPipes(tests)
	if v != 6 {
		t.Error(
			"For", tests,
			"expected 6",
			"got", v,
		)
	}
}


func TestCountGroups(t *testing.T) {
	seen := make(map[int]int)
	v,_ := CountGroups(seen, tests)
	if v != 2 {
		t.Error(
			"For", tests,
			"expected 2",
			"got", v,
		)
	}
}