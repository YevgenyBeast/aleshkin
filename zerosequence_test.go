package aleshkin

import "testing"

func TestSolutionBinaryGap(t *testing.T) {
	type testData struct {
		list int
		want int
	}
	tests := []testData{
		testData{list: 0, want: 1},
		testData{list: 7, want: 0},
		testData{list: -7, want: 2},
		testData{list: 8, want: 3},
		testData{list: 9223372036854775807, want: 0},
		testData{list: -9223372036854775808, want: 63},
	}
	for _, test := range tests {
		got := SolutionBinaryGap(test.list)
		if got != test.want {
			t.Errorf("SolutionBinaryGap(%#v) = %v, but want %v", test.list, got, test.want)
		}
	}
}
