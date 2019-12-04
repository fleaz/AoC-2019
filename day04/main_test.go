package main

import "testing"

type testpair struct {
	value int
	want  bool
}

func Test(t *testing.T) {
	pairs := []testpair{
		{111111, true},
		{223450, false},
		{123789, false},
		{777770, false},
	}

	for _, p := range pairs {
		is := isValid(p.value)
		if is != p.want {
			t.Errorf("%d: Wanted %v but got %v", p.value, p.want, is)
		}

	}
}
