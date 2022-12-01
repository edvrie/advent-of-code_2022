package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fileName := "./1_input.txt";
	file, _ := os.Open(fileName);

	var lines []int
	scanner := bufio.NewScanner(file)
	var sum int

	for scanner.Scan() {
		if (scanner.Text() != "") {
			value, _ := strconv.Atoi(scanner.Text())
			sum += value
			} else {
				lines = append(lines, sum)
				sum = 0
		}
	}

	part1(lines)

	part2(lines)

	file.Close()
}

func part1(fileLines []int) {
	max := fileLines[0]

	for _, element := range fileLines {
		fmt.Println(element, max)
		if (max < element) {
			max = element
		}
	}

	fmt.Println("Highest calories", max)
}

func part2(fileLines []int) {
	first, second, third := 0, 0, 0

	for _, element := range fileLines {
		if (element > first) {
			third = second
			second = first
			first = element
		} else if (element > second) {
			third = second
			second = element
		} else if (element > third) {
			third = element
		}
	}

	fmt.Println("Sum of top 3", first + second + third)
}