package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	startTime := time.Now()

	points := "abcdefghijklmnopqrstuvwxyz"

	fileName := "./3_input.txt"
	file, _ := os.Open(fileName)

	scanner := bufio.NewScanner(file)
	score1 := 0
	score2 := 0
	var groupArray []string

	for scanner.Scan() {
		score1 += part1(scanner.Text(), points)
		groupArray = append(groupArray, scanner.Text())
		score2 += part2(points, groupArray)
		if len(groupArray) == 3 {
			groupArray = nil
		}
	}

	fmt.Println("Answer part 1: ", score1)
	fmt.Println("Answer part 2: ", score2)
	elapsed := time.Since(startTime)

	fmt.Printf("Done in:  %s", elapsed)
}

func part1(fileLine string, points string) int {
	score := 0
	compartment1 := fileLine[0 : len(fileLine)/2]
	compartment2 := fileLine[len(fileLine)/2:]

	duplicates := findMatching(compartment1, compartment2)

	for _, el := range duplicates {
		score += getPoints(el, points)
	}

	return score
}

func part2(points string, groupArray []string) int {
	if len(groupArray) < 3 {
		return 0
	}

	var duplicates []string
	var badge string

	duplicates = findMatching(groupArray[0], groupArray[1])
	badge = findMatching(strings.Join(duplicates, ""), groupArray[2])[0]

	return getPoints(badge, points)
}

func getPoints(el string, points string) int {
	score := 0
	letter := strings.ToLower(el)

	if letter != el {
		score += 26
	}
	score += strings.Index(points, letter) + 1

	return score
}

func findMatching(comp1 string, comp2 string) []string {
	var duplicates []string
	for index := range comp1 {
		if strings.Contains(comp2, string(comp1[index])) && !arrayContains(duplicates, string(comp1[index])) {
			duplicates = append(duplicates, string(comp1[index]))
		}
	}

	return duplicates
}

func arrayContains(arr []string, val string) bool {
	for _, el := range arr {
		if el == val {
			return true
		}
	}

	return false
}
