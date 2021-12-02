package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	day1()
	day2()
}

func day1() {
	horizontal, vertical := 0, 0
	instructions := readInput()
	for _, instruction := range instructions {
		horizontal += instruction.horizontal
		vertical += instruction.vertical
	}
	fmt.Println(horizontal * vertical)
}

func day2() {
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

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i := instruction{}
		line := scanner.Text()
		if instructionRE.MatchString(line) {
			fields := instructionRE.FindStringSubmatch(line)
			delta := ToInt(fields[2])
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
			if line != "" {
				panic("Could not parse " + line)
			}
		}
		result = append(result, i)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return
}

func ToInt(s string) (i int) {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return
}
