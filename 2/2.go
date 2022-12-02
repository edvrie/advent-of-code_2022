package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fileName := "./2_input.txt"
	file, _ := os.Open(fileName)

	scanner := bufio.NewScanner(file)
	score := 0
	for scanner.Scan() {
		value := strings.Split(scanner.Text(), " ")
		p1, p2 := value[0], value[1]

		// score += part1(p1, p2)
		score += part2(p1, p2)
	}

	fmt.Println(score)
}

func part1(p1 string, p2 string) int {
	score := 0
	switch p1 {
	case "A":
		if (p2 == "X") {
			score += 4
		}
		if (p2 == "Y") {
			score += 8
		}
		if (p2 == "Z") {
			score += 3
		}
	case "B": 
		if (p2 == "X") {
			score += 1
		}
		if (p2 == "Y") {
			score += 5
		}
		if (p2 == "Z") {
			score += 9
		}
	case "C":
		if (p2 == "X") {
			score += 7
		}
		if (p2 == "Y") {
			score += 2
		}
		if (p2 == "Z") {
			score += 6
		}
	}

	return score
}

func part2(p1 string, outcome string) int {
	score := 0

	switch p1 {
	case "A":
		if (outcome == "X") {
			score += 3
		}
		if (outcome == "Y") {
			score += 4
		}
		if (outcome == "Z") {
			score += 8
		}
	case "B":
		if (outcome == "X") {
			score += 1
		}
		if (outcome == "Y") {
			score += 5
		}
		if (outcome == "Z") {
			score += 9
		}
	case "C":
		if (outcome == "X") {
			score += 2
		}
		if (outcome == "Y") {
			score += 6
		}
		if (outcome == "Z") {
			score += 7
		}
	}

	return score
}