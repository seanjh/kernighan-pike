package search

import "testing"

func TestLinearIntSearcher(t *testing.T) {
	cases := []struct {
		values LinearIntSearcher
		target, expected int
	}{
		{make([]int, 0), 5, Absent},
		{[]int{1, 2, 3, 4, 5}, 5, 4},
		{[]int{5, 4, 3, 2, 1}, 3, 2},
		{[]int{5, 5, 5, 5, 5}, 5, 0},
		{[]int{4, 5, 5, 5, 5}, 5, 1},
		{[]int{0}, 0, 0},
		{[]int{1}, 1, 0},
		{[]int{-1, 1}, -1, 0},
	}

	for _, c := range cases {
		if c.values.Search(c.values, c.target) != c.expected {
			t.Error("Expected to find target at index", c.expected)
		}
	}
}

func TestLinearStringSearcher(t *testing.T) {
	cases := []struct {
		values LinearStringSearcher
		target string
		expected int
	}{
		{make([]string, 0), "foo", -1},
		{[]string{"foo", "bar", "baz", "yab", "gab"}, "", -1},
		{[]string{"foo", "bar", "baz", "yab", "gab"}, "gab", 4},
		{[]string{"foo", "bar", "baz", "yab", "gab"}, "yab", 3},
		{[]string{"foo"}, "foo", 0},
		{[]string{"bar"}, "bar", 0},
		{[]string{""}, "", 0},
	}

	for _, c := range cases {
		if c.values.Search(c.values, c.target) != c.expected {
			t.Error("Expected to find target at index", c.expected)
		}
	}
}
