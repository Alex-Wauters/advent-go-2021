package main

import (
	utils "advent/utils"
	"fmt"
	"regexp"
)

func main() {
	part1()
	part2()
}

func part1() {
	horizontal, vertical := 0, 0
	instructions := readInput()
	for _, instruction := range instructions {
		horizontal += instruction.horizontal
		vertical += instruction.vertical
	}
	fmt.Println(horizontal * vertical)
}

func part2() {
	horizontal, depth, aim := 0, 0, 0
	instructions := readInput()
	for _, instruction := range instructions {
		horizontal += instruction.horizontal
		aim += instruction.vertical
		depth += instruction.horizontal * aim
	}
	fmt.Println(horizontal * depth)
}

type instruction struct {
	horizontal, vertical int
}

func readInput() (result []instruction) {
	instructionRE := regexp.MustCompile(`^(down|up|forward) (\d)+$`)
	lines := utils.ReadInput()
	for _, l := range lines {
		i := instruction{}
		if instructionRE.MatchString(l) {
			fields := instructionRE.FindStringSubmatch(l)
			delta := utils.ToInt(fields[2])
			switch fields[1] {
			case "down":
				i.vertical = delta
			case "up":
				i.vertical = -delta
			case "forward":
				i.horizontal = delta
			default:
				panic("Did not recognize " + fields[1])
			}
		} else {
			panic("Could not parse " + l)
		}
		result = append(result, i)
	}

	return
}
