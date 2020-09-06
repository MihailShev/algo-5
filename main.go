package main

import (
	"algo-5/sorting"
	"fmt"
	tester "github.com/MihailShev/algo-tester"
)

const pathToRandom = "test-data/0.random"
const pathToDigits = "test-data/1.digits"
const pathToSorted = "test-data/2.sorted"
const pathToRevers = "test-data/3.revers"

type TestData struct {
	path        string
	description string
}

func main() {
	tests := []TestData{{
		path:        pathToRandom,
		description: "array of random numbers",
	}, {
		path:        pathToDigits,
		description: "array of random digits",
	}, {
		path:        pathToSorted,
		description: "99% sorted array",
	}, {
		path:        pathToRevers,
		description: "revers sorted array",
	}}

	var sort tester.Tester

	//fmt.Printf("\n*** Test buble sort ***\n")
	//for _, v := range tests {
	//	fmt.Printf("\nType: %s:\n\n", v.description)
	//	sort = tester.NewTester(sorting.BubbleSort{}, v.path)
	//	sort.RunTestNum(6)
	//}
	//
	//fmt.Printf("\n*** Test selection sort ***\n")
	//for _, v := range tests {
	//	fmt.Printf("\nType: %s:\n\n", v.description)
	//	sort = tester.NewTester(sorting.Selection{}, v.path)
	//	sort.RunTestNum(6)
	//}

	//fmt.Printf("\n*** Test insertion sort ***\n")
	//for _, v := range tests {
	//	fmt.Printf("\nType: %s:\n\n", v.description)
	//	sort = tester.NewTester(sorting.Insertion{}, v.path)
	//	sort.RunTestNum(6)
	//}
	//
	//fmt.Printf("\n*** Test shell sort ***\n")
	//for _, v := range tests {
	//	fmt.Printf("\nType: %s:\n\n", v.description)
	//	sort = tester.NewTester(sorting.Shell{}, v.path)
	//	sort.RunTestWithCount(8)
	//}
	//
	//fmt.Printf("\n*** Test shell sort with Knut steps ***\n")
	//for _, v := range tests {
	//	fmt.Printf("\nType: %s:\n\n", v.description)
	//	sort = tester.NewTester(sorting.Shell{StepType: sorting.KnutSteps}, v.path)
	//	sort.RunTestWithCount(8)
	//}
	//
	//fmt.Printf("\n*** Test shell sort with Sedgewick steps ***\n")
	//for _, v := range tests {
	//	fmt.Printf("\nType: %s:\n\n", v.description)
	//	sort = tester.NewTester(sorting.Shell{StepType: sorting.SedgewickSteps}, v.path)
	//	sort.RunTestWithCount(8)
	//}
	//
	fmt.Printf("\n*** Test heap sort ***\n")
	for _, v := range tests {
		fmt.Printf("\nType: %s:\n\n", v.description)
		sort = tester.NewTester(sorting.Heap{}, v.path)
		sort.RunTestWithCount(8)
	}
}
