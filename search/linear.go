package search

func linearSearch(length int, cmp comparator) int {
	for i := 0; i < length; i++ {
		if cmp(i) == eq {
			return i
		}
	}
	return Absent
}

type LinearIntSearcher []int

func (l LinearIntSearcher) Search(items []int, target int) int {
	return linearSearch(len(items), makeIntComparator(l, target))
}

type LinearStringSearcher []string

func (l LinearStringSearcher) Search(items []string, target string) int {
	return linearSearch(len(items), makeStringComparator(l, target))
}
