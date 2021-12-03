package main

import "fmt"

func main() {
	day1()
	day2()
}

func day1() {
	m := measurements()
	previous := 0
	inc := 0
	for _, i := range m {
		if previous != 0 && i > previous {
			inc++
		}
		previous = i
	}
	fmt.Println(inc)
}

func day2() {
	m := measurements()
	inc := 0
	for i := 1; i+2 < len(m); i = i + 1 {
		previous := m[i-1] + m[i] + m[i+1]
		current := m[i] + m[i+1] + m[i+2]
		if current > previous {
			inc++
		}
	}
	fmt.Println(inc)
}
