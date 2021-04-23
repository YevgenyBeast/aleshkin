package aleshkin

import "testing"

func TestSolutionUniq(t *testing.T) {
	type testData struct {
		list []int
		want int
	}
	tests := []testData{
		testData{list: []int{1}, want: 1},
		testData{list: []int{1, 1, 2}, want: 2},
		testData{list: []int{-1, -1, 0, 1, 1, 2, -2}, want: 5},
		testData{list: []int{1, 1}, want: 1},
	}
	for _, test := range tests {
		got := SolutionUniq(test.list)
		if got != test.want {
			t.Errorf("SolutionUniq(%#v) = %v, but want %v", test.list, got, test.want)
		}
	}
}
