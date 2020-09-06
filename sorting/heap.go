package sorting

import "algo-5/utils"

type Heap struct {
}

func (h Heap) Run(data []string) string {
	a := utils.ParseTestData(data)
	return utils.ToString(h.Sort(a))
}

func (h Heap) Sort(a []int) []int {
	size := len(a)
	for i := size/2 - 1; i >= 0; i-- {
		heapify(a, size, i)
	}

	for i := size - 1; i >= 0; i-- {
		a[0], a[i] = a[i], a[0]
		heapify(a, i, 0)
	}

	return a
}

func heapify(a []int, size, i int) {
	large := i
	left := 2*i + 1
	right := left + 1

	if left < size && a[left] > a[large] {
		large = left
	}

	if right < size && a[right] > a[large] {
		large = right
	}

	if large != i {
		a[large], a[i] = a[i], a[large]
		heapify(a, size, large)
	}

}
