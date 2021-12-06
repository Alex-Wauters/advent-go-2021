package main

import (
	"advent/utils"
	"fmt"
	"strings"
)

const BOARD_SIZE = 5

func main() {
	part1()
	part2()
}

func part1() {
	numbers := draws()
	boards := getBoards()
	for _, d := range numbers {
		for _, b := range boards {
			b.mark(d)
			if b.hasWon() {
				fmt.Println(b.score(d))
				return
			}
		}
	}
}

func part2() {
	numbers := draws()
	boards := getBoards()
	remaining := len(boards)
	finishedBoards := make(map[int]bool)
	for _, d := range numbers {
		for i, b := range boards {
			b.mark(d)
			_, finished := finishedBoards[i]
			if !finished && b.hasWon() {
				finishedBoards[i] = true
				remaining--
				if remaining == 0 {
					fmt.Println(b.score(d))
					return
				}
			}
		}
	}
}

func draws() []int {
	return []int{4, 75, 74, 31, 76, 79, 27, 19, 69, 46, 98, 59, 83, 23, 90, 52, 87, 6, 11, 92, 80, 51, 43, 5, 94, 17, 15, 67, 25, 30, 48, 47, 62, 71, 85, 58, 60, 1, 72, 99, 3, 35, 42, 10, 96, 49, 37, 36, 8, 44, 70, 40, 45, 39, 0, 63, 2, 78, 68, 53, 50, 77, 20, 55, 38, 86, 54, 93, 26, 88, 12, 91, 95, 34, 9, 14, 33, 66, 41, 13, 28, 57, 29, 73, 56, 22, 89, 21, 64, 61, 32, 65, 97, 84, 18, 82, 81, 7, 16, 24}
}

type board map[string]*cell
type cell struct {
	number int
	marked bool
}
type cellList []*cell

func (l cellList) isComplete() bool {
	for _, c := range l {
		if !c.marked {
			return false
		}
	}
	return true
}

func (b board) hasWon() bool {
	for i := 0; i < BOARD_SIZE; i++ {
		if b.getRow(i).isComplete() {
			return true
		}
	}
	for i := 0; i < BOARD_SIZE; i++ {
		if b.getColumn(i).isComplete() {
			return true
		}
	}
	return false
}

func (b board) getColumn(c int) (result cellList) {
	for i := 0; i < BOARD_SIZE; i++ {
		result = append(result, b[pos(i, c)])
	}
	return
}

func (b board) getRow(r int) (result cellList) {
	for i := 0; i < BOARD_SIZE; i++ {
		result = append(result, b[pos(r, i)])
	}
	return
}

func (b board) mark(n int) {
	for _, v := range b {
		if v.number == n {
			v.marked = true
			return
		}
	}
}

func (b board) score(lastNumber int) int {
	s := 0
	for _, v := range b {
		if !v.marked {
			s += v.number
		}
	}
	return lastNumber * s
}

func getBoards() (boards []board) {
	lines := utils.ReadInput()
	row := 0
	b := board{}
	for _, l := range lines {
		if l == "" {
			if row != 0 {
				boards = append(boards, b)
			}
			row = 0
			b = board{}
			continue
		}
		numbers := strings.Fields(l)
		for i, s := range numbers {
			n := utils.ToInt(s)
			b[pos(row, i)] = &cell{number: n, marked: false}
		}
		row++
	}
	if row != 0 {
		boards = append(boards, b)
	}
	return
}

func pos(row, col int) string {
	return fmt.Sprintf("%v,%v", row, col)
}
