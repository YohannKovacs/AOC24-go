package solution

import (
	"strconv"
	"strings"
)

func SolutionDay2(input string) (int, error, int, error) {
	p1, err1 := D2part1(input)
	p2, err2 := D2part2(input)

	return p1, err1, p2, err2
}

func D2part1(input string) (int, error) {
	var answer int

	reports := strings.Split(input, "\n")
	for _, report := range reports {
		isIncreasing := true
		isDecreasing := true
		var levels []int
		a := strings.Split(report, " ")
		for _, v := range a {
			level, _ := strconv.Atoi(v)
			levels = append(levels, level)
		}
		for i := 1; i < len(levels); i++ {
			diff := levels[i] - levels[i-1]
			if diff < 1 || diff > 3 {
				isIncreasing = false
				break
			}
		}
		if isIncreasing {
			answer++
		} else {
			for i := 0; i < len(levels)-1; i++ {
				diff := levels[i] - levels[i+1]
				if diff < 1 || diff > 3 {
					isDecreasing = false
					break
				}
			}
			if isDecreasing {
				answer++
			}
		}
	}
	return answer, nil
}

func D2part2(input string) (int, error) {
	var answer int

	reports := strings.Split(input, "\n")
	for _, report := range reports {
		var levels []int
		a := strings.Split(report, " ")
		for _, v := range a {
			level, _ := strconv.Atoi(v)
			levels = append(levels, level)
		}
		if checkSafety(levels) {
			answer++
		} else {
			for i := range levels {
				newSlice := make([]int, len(levels))
				copy(newSlice, levels)
				if checkSafety(append(newSlice[:i], newSlice[i+1:]...)) {
					answer++
					break
				}
			}
		}
	}
	return answer, nil
}

func checkSafety(levels []int) bool {
	isIncreasing := true
	isDecreasing := true
	for i := 1; i < len(levels); i++ {
		diff := levels[i] - levels[i-1]
		if diff < 1 || diff > 3 {
			isIncreasing = false
			break
		}
	}
	if isIncreasing {
		return true
	} else {
		for i := 0; i < len(levels)-1; i++ {
			diff := levels[i] - levels[i+1]
			if diff < 1 || diff > 3 {
				isDecreasing = false
				break
			}
		}
		if isDecreasing {
			return true
		}
	}
	return false
}
