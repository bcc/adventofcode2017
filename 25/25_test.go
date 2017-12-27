package main

import (
	"testing"
)

var test = []instruction{
	{"a", 0, 1, 1, "b"},
	{"a", 1, 0, -1, "b"},
	{"b", 0, 1, -1, "a"},
	{"b", 1, 1, 1, "a"}}

func TestInstructions(t *testing.T) {
	v := ProcessInstructions(test, 6)
	if v != 3 {
		t.Error(
			"For", test,
			"expected 3",
			"got", v,
		)
	}
}
