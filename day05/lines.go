package main

import (
	"advent/utils"
	"fmt"
	"regexp"
)

func main() {
	part1()
	part2()
}

func part1() {
	board := make(map[string]int)
	lines := getLines()
	for _, l := range lines {
		if l.isStraight() {
			for _, p := range l.covers() {
				board[p]++
			}
		}
	}
	sum := 0
	for _, v := range board {
		if v >= 2 {
			sum++
		}
	}
	fmt.Println(sum)
}

func part2() {
	board := make(map[string]int)
	lines := getLines()
	for _, l := range lines {
		for _, p := range l.covers() {
			board[p]++
		}
	}
	sum := 0
	for _, v := range board {
		if v >= 2 {
			sum++
		}
	}
	fmt.Println(sum)
}

func (l line) covers() (result []string) {
	if l.x1 == l.x2 {
		smallest, largest := utils.SortNumbers(l.y1, l.y2)
		for i := smallest; i <= largest; i++ {
			result = append(result, utils.Position(i, l.x1))
		}
		return
	}
	if l.y1 == l.y2 {
		smallest, largest := utils.SortNumbers(l.x1, l.x2)
		for i := smallest; i <= largest; i++ {
			result = append(result, utils.Position(l.y1, i))
		}
		return
	}
	deltaX := 1
	deltaY := 1
	y := l.y1
	if l.x2 < l.x1 {
		deltaX = -1
	}
	if l.y2 < l.y1 {
		deltaY = -1
	}
	for i := l.x1; i != l.x2; i += deltaX {
		result = append(result, utils.Position(y, i))
		y += deltaY
	}
	result = append(result, utils.Position(l.y2, l.x2))
	return
}

func (l line) isStraight() bool {
	return l.x1 == l.x2 || l.y1 == l.y2
}

type line struct {
	x1, x2, y1, y2 int
}

func getLines() (result []line) {
	input := utils.ReadInput()
	lineRE := regexp.MustCompile(`^(\d+),(\d+) -> (\d+),(\d+)$`)
	for _, s := range input {
		l := line{}
		if lineRE.MatchString(s) {
			fields := lineRE.FindStringSubmatch(s)
			l.x1 = utils.ToInt(fields[1])
			l.y1 = utils.ToInt(fields[2])
			l.x2 = utils.ToInt(fields[3])
			l.y2 = utils.ToInt(fields[4])
			result = append(result, l)
		} else {
			panic("Could not parse input line " + s)
		}
	}
	return
}
