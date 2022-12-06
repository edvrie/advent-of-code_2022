package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fileName := "./6_input.txt"
	file, _ := os.Open(fileName)

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	numb1 := part1(scanner.Text())
	numb2 := part2(scanner.Text())
	fmt.Print("Answer part 1: ", numb1)
	fmt.Println()
	fmt.Print("Answer part 2: ", numb2)
	fmt.Println()


	file.Close()
}

func part1(line string) int {
	for index := range line {
		uniqMap := make(map[string]bool)
		subStr := line[index:index+4]

		for _, el := range subStr {
			if !uniqMap[string(el)] {
				uniqMap[string(el)] = true
			} else {
				break
			}
		}

		if (len(uniqMap) == 4) {
			return index + 4
		}
	}

	return -1
}

func part2(line string) int {
	for index := range line {
		uniqMap := make(map[string]bool)
		subStr := line[index:index+14]

		for _, el := range subStr {
			if !uniqMap[string(el)] {
				uniqMap[string(el)] = true
			} else {
				break
			}
		}

		if (len(uniqMap) == 14) {
			return index + 14
		}
	}

	return -1
}