package btcec

import "testing"

func TestLessThanUint32(t *testing.T) {
	tests := []struct {
		x uint32
		y uint32
		a uint32
	}{
		{0, 1, 1},
		{2, 2, 0},
		{1 << 31, 1 << 31, 0},
		{17, 1 << 31, 1},
		{2 ^ 32, 0, 0},
	}

	t.Logf("Running %d tests", len(tests))
	for i, test := range tests {
		answer := LessThanUint32(test.x, test.y)
		if test.a != answer {
			t.Errorf("LessThanUint32 #%d wrong result\ngot: %v\n"+
				"want: %v", i, answer, test.a)
			continue
		}
	}
}

func TestLessOrEqUint32(t *testing.T) {
	tests := []struct {
		x uint32
		y uint32
		a uint32
	}{
		{0, 1, 1},
		{2, 2, 1},
		{1 << 31, 1 << 31, 1},
		{17, 1 << 31, 1},
		{2 ^ 32, 0, 0},
	}

	t.Logf("Running %d tests", len(tests))
	for i, test := range tests {
		answer := LessOrEqUint32(test.x, test.y)
		if test.a != answer {
			t.Errorf("LessOrEqUint32 #%d wrong result\ngot: %v\n"+
				"want: %v", i, answer, test.a)
			continue
		}
	}
}

func TestEqUint32(t *testing.T) {
	tests := []struct {
		x uint32
		y uint32
		a uint32
	}{
		{0, 1, 0},
		{2, 2, 1},
		{1 << 31, 1 << 31, 1},
		{17, 1 << 31, 0},
		{2 ^ 32, 0, 0},
	}

	t.Logf("Running %d tests", len(tests))
	for i, test := range tests {
		answer := EqUint32(test.x, test.y)
		if test.a != answer {
			t.Errorf("LessOrEqUint32 #%d wrong result\ngot: %v\n"+
				"want: %v", i, answer, test.a)
			continue
		}
	}
}
