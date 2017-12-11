package main

import (
	"testing"
)

type testpair struct {
	values []string
	result int
  }
  
  var tests = []testpair{
	{[]string{"ne","ne","ne"}, 3},
	{[]string{"ne","ne","sw","sw"}, 0},
	{[]string{"ne","ne","s","s"}, 2},
	{[]string{"se","sw","se","sw","sw"}, 3},
  }

func TestCalcDistance(t *testing.T) {
	for _, pair := range tests {
		v := CalcDistance(pair.values)
		if v[0] != pair.result {
		  t.Error(
			"For", pair.values,
			"expected", pair.result,
			"got", v[0],
		  )
		}
	  }
}