package aleshkin

import "testing"

func TestSolutionRotate(t *testing.T) {
	type testData struct {
		list1 []int
		list2 int
		want  []int
	}
	tests := []testData{
		testData{list1: []int{1, 2, 3, 4, 5}, list2: 3, want: []int{3, 4, 5, 1, 2}},
		testData{list1: []int{1, 2, 3, 4, 5}, list2: 6, want: []int{5, 1, 2, 3, 4}},
	}
	for _, test := range tests {
		got := SolutionRotate(test.list1, test.list2)
		if !compareSliceRotate(test.want, got) {
			t.Errorf("SolutionRotate(%#v, %#v) = %v, but want %v", test.list1, test.list2, got, test.want)
		}
	}
}

func compareSliceRotate(w []int, g []int) bool {
	for i := 0; i < len(w); i++ {
		if w[i] != g[i] {
			return false
		}
	}
	return true
}
