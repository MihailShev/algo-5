package sorting

import "algo-5/utils"

type Insertion struct {
}

func (in Insertion) Run(data []string) string {
	a := utils.ParseTestData(data)
	return utils.ToString(in.Sort(a))
}

func (in Insertion) Sort(a []int) []int {
	for i := 1; i < len(a); i++ {
		var x = a[i]
		var j = i - 1

		for j >= 0 && a[j] > x {
			a[j + 1] = a[j]
			j--
		}
		a[j + 1] = x
	}

	return a
}
