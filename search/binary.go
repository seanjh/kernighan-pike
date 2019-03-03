package search

type BinaryIntSearcher []int

func (l BinaryIntSearcher) Search(items []int, target int) int {
	low, high := 0, len(items) - 1
	for low <= high {
		mid := (low + high) / 2
		if r := items[mid]; r == target {
			return mid
		} else if target < r {
			high = mid - 1 // shift left
		} else { // target > r
			low = mid + 1 // shift right
		}
	}
	return Absent
}

type BinaryStringSearcher []string

func (l BinaryStringSearcher) Search(items []int) int {
	return -99
}
