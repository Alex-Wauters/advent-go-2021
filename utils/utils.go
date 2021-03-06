package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func ToInt(s string) (i int) {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return
}

func ReadInput() (result []string) {
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
		line := scanner.Text()
		result = append(result, line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return result
}

func Position(row, col int) string {
	return fmt.Sprintf("%v,%v", row, col)
}

func SortNumbers(a, b int) (c, d int) {
	if a < b {
		return a, b
	}
	return b, a
}

func Abs(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}
