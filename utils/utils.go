package utils

import (
	"bufio"
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
		if line != "" {
			result = append(result, line)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return result
}
