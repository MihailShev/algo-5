package utils

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func ParseTestData(data []string) []int {
	str := strings.Split(data[1], " ")
	res := make([]int, 0)

	for _, v := range str {
		number, err := strconv.Atoi(v)

		if err != nil {
			log.Fatal("failed to parse test data")
		}

		res = append(res, number)
	}

	return res
}

func ToString(a []int) string {
	s := strings.Builder{}

	for _, v := range a {
		s.WriteString(fmt.Sprint(v))
		s.WriteString(" ")
	}

	return strings.TrimRight(s.String(), " ")
}