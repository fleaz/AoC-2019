package main

import "testing"

type testpair struct {
	Wire1    []string
	Wire2    []string
	Distance int
}

func Test(t *testing.T) {
	pairs := []testpair{
		{
			Wire1:    []string{"R8", "U5", "L5", "D3"},
			Wire2:    []string{"U7", "R6", "D4", "L4"},
			Distance: 6,
		},
		{
			Wire1:    []string{"R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72"},
			Wire2:    []string{"U62", "R66", "U55", "R34", "D71", "R55", "D58", "R83"},
			Distance: 159,
		},
		{
			Wire1:    []string{"R98", "U47", "R26", "D63", "R33", "U87", "L62", "D20", "R33", "U53", "R51"},
			Wire2:    []string{"U98", "R91", "D20", "R16", "D67", "R40", "U7", "R15", "U6", "R7"},
			Distance: 135,
		},
	}

	for _, p := range pairs {
		is := calculateDistance(p.Wire1, p.Wire2)
		if is != p.Distance {
			t.Errorf("Wanted %d but got %d", p.Distance, is)
		}

	}
}
