package solution

import (
	"math"
	"slices"
	"strconv"
	"strings"
)

func SolutionDay1(input string) (int, error, int, error) {
	a1, err1 := part1(input)
	a2, err2 := part2(input)

	return a1, err1, a2, err2
}

func part1(input string) (int, error) {
	var list1 []int
	var list2 []int
	var answer int

	lines := strings.Split(input, "\n")
	for _, v := range lines {
		line := strings.Split(v, "   ")
		i1, _ := strconv.Atoi(line[0])
		i2, _ := strconv.Atoi(line[1])
		list1 = append(list1, i1)
		list2 = append(list2, i2)
	}
	slices.Sort(list1)
	slices.Sort(list2)
	for i, _ := range list1 {
		answer += int(math.Abs(float64(list1[i] - list2[i])))
	}

	return answer, nil
}

func part2(input string) (int, error) {
	var answer int
	var list1 []int
	var list2 []int

	lines := strings.Split(input, "\n")
	for _, v := range lines {
		line := strings.Split(v, "   ")
		i1, _ := strconv.Atoi(line[0])
		i2, _ := strconv.Atoi(line[1])
		list1 = append(list1, i1)
		list2 = append(list2, i2)
	}

	for i := range list1 {
		var counter int
		for j := range list2 {
			if list1[i] == list2[j] {
				counter += 1
			}
		}

		answer += list1[i] * counter
	}

	return answer, nil
}
