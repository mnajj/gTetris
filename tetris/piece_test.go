package tetris

import (
	"testing"
)

func TestClockwiseRotate(t *testing.T) {
	table := []struct {
		from, to shape
	}{
		{
			[][]bool{
				{true, true},
				{true, true},
			},
			[][]bool{
				{true, true},
				{true, true},
			},
		},
		{
			[][]bool{
				{false, true},
				{true, true},
				{true, false},
			}, [][]bool{
				{true, true, false},
				{false, true, true},
			},
		},
		{
			[][]bool{
				{true, true, false},
				{false, true, true},
			},
			[][]bool{
				{false, true},
				{true, true},
				{true, false},
			},
		},
	}

	for _, s := range table {
		s.from = s.from.clockwiseRotate()
		if !equal2DArrays(s.from, s.to) {
			t.Errorf("get %v, expected %v", s.from, s.to)
		}
	}
}

func equal2DArrays(a, b shape) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := 0; j < len(a[i]); j++ {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}
