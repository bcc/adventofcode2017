package main

import (
	"testing"
)

var test = []string{"s1", "x3/4", "pe/b"}

func TestRun16(t *testing.T) {
	i := ParseInstr(test)
	input := "abcde"
	v := Run16(input, i)
	if v != "baedc" {
		t.Error(
			"For", test,
			"expected baedc",
			"got", v,
		)
	}
}

func TestRun1M(t *testing.T) {
	i := ParseInstr(test)
	input := "abcdefghijlmnop"

	for x :=0; x<10000000; x++ {
		input = Run16(input, i)
	}

	if input != "fghdijlmnopabce" {
		t.Error(
			"For", test,
			"expected fghdijlmnopabce",
			"got", input,
		)
	}
}