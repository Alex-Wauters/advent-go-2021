package main

import (
	utils "advent/utils"
	"fmt"
	"strconv"
)

func main() {
	part1()
	part2()
}

func part1() {
	lines := utils.ReadInput()
	gamma, epsilon := "", ""
	for i := 0; i < len(lines[0]); i++ {
		most, least := mostAndLeastCommonRune(lines, i)
		gamma += string(most)
		epsilon += string(least)
	}
	fmt.Println(fromBinary(gamma) * fromBinary(epsilon))
}

func mostAndLeastCommonRune(lines []string, i int) (uint8, uint8) {
	count0, count1 := 0, 0
	for _, line := range lines {
		if line[i] == '0' {
			count0++
		} else {
			count1++
		}
	}
	if count0 > count1 {
		return '0', '1'
	}
	return '1', '0'
}

func filter(lines []string, i int, mostCommon bool) string {
	if len(lines) == 1 {
		return lines[0]
	}
	most, least := mostAndLeastCommonRune(lines, i)
	comparator := least
	if mostCommon {
		comparator = most
	}
	filtered := make([]string, 0)
	for _, l := range lines {
		if l[i] == comparator {
			filtered = append(filtered, l)
		}
	}
	return filter(filtered, i+1, mostCommon)
}

func part2() {
	oxygen := filter(utils.ReadInput(), 0, true)
	scrubber := filter(utils.ReadInput(), 0, false)
	fmt.Println(fromBinary(oxygen) * fromBinary(scrubber))
}

func fromBinary(s string) (r int) {
	i, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		panic(err)
	}
	return int(i)
}
