package search

type comparison int

// Absent is returned when no match is located
const Absent = -1

const (
	lt	comparison = -1
	eq	comparison = 0
	gt	comparison = 1
)

type intSearcher interface {
	Search([]int) int
}

type stringSearcher interface {
	Search([]string) int
}

type comparator func(int) comparison

func makeIntComparator(items []int, target int) comparator {
	return func(i int) comparison {
		if val := items[i]; val == target {
			return eq
		} else if val < target {
			return lt 
		}
		return gt
	}
}

func makeStringComparator(items []string, target string) comparator {
	return func(i int) comparison {
		if val := items[i]; val == target {
			return eq
		} else if val < target {
			return lt 
		}
		return gt
	}
}
