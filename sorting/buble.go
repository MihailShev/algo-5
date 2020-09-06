package sorting

import (
	"algo-5/utils"
)

type BubbleSort struct {
}

func (b BubbleSort) Run(data []string) string {
	s := utils.ParseTestData(data)

	return utils.ToString(b.Sort(s))
}

func (b BubbleSort) Sort(a []int) []int {
	var l int
	for i := 0; i < len(a); i++ {
		l = len(a) - 1 - i
		for j := 0; j < l; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	return a
}
