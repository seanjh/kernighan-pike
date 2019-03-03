package search

import "testing"

func TestBinaryIntSearcher(t *testing.T) {
	cases := []struct {
		values BinaryIntSearcher
		target, expected int
	}{
		{make([]int, 0), 42, Absent},
		{[]int{99}, 98, Absent},
		{[]int{99}, 99, 0},
		{[]int{1, 99}, 99, 1},
		{[]int{1, 99}, 98, Absent},
		{[]int{1, 5, 99}, 98, Absent},
		{[]int{1, 5, 99}, 99, 2},
		{[]int{1, 5, 99}, 1, 0},
		{[]int{1, 5, 99}, -1, Absent},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}, 2, 1},
		{[]int{1, 1, 1, 2}, 2, 3},
		{[]int{1, 1, 1, 2}, 1, 1},
		{[]int{-42, 0, 42}, 5, Absent},
		{[]int{-42, 0, 42}, 42, 2},
		{[]int{-42, 0, 42}, -42, 0},
	}

	for _, c := range cases {
		if c.values.Search(c.values, c.target) != c.expected {
			t.Error("Expected to find target at index", c.expected)
		}
	}
}
