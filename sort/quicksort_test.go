package sort

import "testing"

func TestQuickIntSorter(t *testing.T) {
	cases := []struct {
		input QuickIntSorter
		expected []int
	}{
		{QuickIntSorter{}, []int{}},
		{QuickIntSorter{1}, []int{1}},
		{QuickIntSorter{2, 1}, []int{1, 2}},
		{QuickIntSorter{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{
			QuickIntSorter{
				19, 79, 75, 28, 88, 32, 92, 92, 29, 52, 53, 70, 25, 23, 13, 74,
				84, 60, 67, 59, 36, 91, 24, 31, 61, 76, 71, 95, 35, 31, 88, 22,
				11, 90, 30, 23, 9, 50, 73, 76, 0, 64, 75, 44, 65, 55, 78, 86,
				25, 99,
			},
			[]int{
				0, 9, 11, 13, 19, 22, 23, 23, 24, 25, 25, 28, 29, 30, 31, 31,
				32, 35, 36, 44, 50, 52, 53, 55, 59, 60, 61, 64, 65, 67, 70, 71,
				73, 74, 75, 75, 76, 76, 78, 79, 84, 86, 88, 88, 90, 91, 92, 92,
				95, 99,
			},
		},
		{QuickIntSorter{1, 1, 1}, []int{1, 1, 1}},
		{QuickIntSorter{1, 0, -1}, []int{-1, 0, 1}},
	}

	for _, c := range cases {
		c.input.Sort()
		for i := range c.input {
			if c.input[i] != c.expected[i] {
				t.Errorf(
					"Expected to find sorted result.\n" +
					"Actual:   %v\n" +
					"Expected: %v\n",
					c.input, c.expected,
				)
				break
			}
		}
	}
}
