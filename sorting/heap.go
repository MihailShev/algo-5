package sorting

import "algo-5/utils"

type Heap struct {

}

func (h Heap) Run(data []string) string {
	a := utils.ParseTestData(data)
	return utils.ToString(h.Sort(a))
}

func (h Heap) Sort(a []int) []int {
	l := len(a)
	for root := l /2 - 1; root >= 0; root -- {
		moveMaxToRoot(a, root, l)
	}

	for i := l - 1; i >= 0; i-- {
		a[0], a[i] = a[i], a[0]
		moveMaxToRoot(a, 0, a[i])
	}

	return a
}

func moveMaxToRoot(a []int, root int, size int) {
	left := 2 * root + 1
	right := left + 1

	x := root

	if left < size && a[x] < a[left] {
		x = left
	}

	if right < size && a[x] < a[right] {
		x = right
	}

	if x == root {
		return
	}

	a[x], a[root] = a[root], a[x]
	moveMaxToRoot(a, x, size)
}
