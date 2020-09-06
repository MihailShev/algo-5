package sorting

import "algo-5/utils"

type Selection struct {
}

func (s Selection) Run(data []string) string {
	a := utils.ParseTestData(data)
	return utils.ToString(s.Sort(a))
}

func (s Selection) Sort(a []int) []int {
	var min int

	for i, _ := range a {
		min = i
		var l = len(a)
		for j := min + 1; j < l; j++ {
			if a[min] > a[j] {
				min = j
			}
		}

		if min != i {
			a[i], a[min] = a[min], a[i]
		}
	}

	return a
}
