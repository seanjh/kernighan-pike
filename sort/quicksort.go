package sort

import "math/rand"

type QuickIntSorter []int

func quickSortInts(items QuickIntSorter, n int) {
	if n <= 1 {
		return
	}
	items = items[:n]

	r := rand.Intn(n)
	items.Swap(0, r)
	last := 0
	for i := 1; i < n; i++ {
		if items[i] < items[0] {
			last++
			items.Swap(i, last)
		}
	}
	items.Swap(0, last)
	quickSortInts(items, last)
	quickSortInts(items[last+1:], n-last-1)
}

func (l QuickIntSorter) Sort() {
	quickSortInts(l, len(l))
}

func (l QuickIntSorter) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}
