package main

import (
	"testing"
)

var test = map[int]firewall{
	0: {0, 2, 0, true},
	1: {1, 1, 0, true},
	4: {4, 3, 0, true},
	6: {6, 3, 0, true},
}

func TestFindDelay(t *testing.T) {
	test = ResetHash(test)
	v := FindDelay(test)
	if v != 10 {
		t.Error(
			"For", test,
			"expected 10",
			"got", v,
		)
	}
}

func TestRunFirewall(t *testing.T) {
	test = ResetHash(test)
	v, _ := RunFirewall(test, 0)
	if v != 24 {
		t.Error(
			"For", test,
			"expected 24",
			"got", v,
		)
	}
}
