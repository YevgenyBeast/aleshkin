package aleshkin

import "testing"

func TestSolutionSquareGenerator(t *testing.T) {
	type testData struct {
		list1 int
		list2 int
		want  []int
	}
	tests := []testData{
		testData{list1: 1, list2: 5, want: []int{1, 4, 9, 16, 25}},
		testData{list1: 100, list2: 2, want: []int{10000, 10201}},
	}
	for _, test := range tests {
		got := SolutionSquareGenerator(test.list1, test.list2)
		if !compareSliceSquare(test.want, got) {
			t.Errorf("SolutionSquareGenerator(%#v, %#v) = %v, but want %v", test.list1, test.list2, got, test.want)
		}
	}
}

func compareSliceSquare(w []int, g []int) bool {
	for i := 0; i < len(w); i++ {
		if w[i] != g[i] {
			return false
		}
	}
	return true
}
