// Experimenting with a different solution

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// A - Rock
// B - Paper
// C - Scizzors

// X - Rock - 1
// Y - Paper - 2
// Z - Scizzors - 3

// Win - 6
// Draw - 3
// Lose - 0

func main() {
	startTime := time.Now()
	fileName := "./2_input.txt"
	file, _ := os.Open(fileName)

	scanner := bufio.NewScanner(file)
	score1 := 0
	score2 := 0
	scoreMap1 := initMapPart1()
	scoreMap2 := initMapPart2()

	for scanner.Scan() {
		value := strings.Split(scanner.Text(), " ")
		p1, p2 := value[0], value[1]
		score1 += scoreMap1[p1+p2]
		score2 += scoreMap2[p1+p2]
	}

	fmt.Println("Answer part 1: ", score1)
	fmt.Println("Answer part2: ", score2)

	file.Close()

	elapsed := time.Since(startTime)

	fmt.Printf("Done in:  %s", elapsed)
}

func initMapPart1() map[string]int {
	scoreMap := make(map[string]int)
	scoreMap["AX"] = 4
	scoreMap["AY"] = 8
	scoreMap["AZ"] = 3
	scoreMap["BX"] = 1
	scoreMap["BY"] = 5
	scoreMap["BZ"] = 9
	scoreMap["CX"] = 7
	scoreMap["CY"] = 2
	scoreMap["CZ"] = 6

	return scoreMap
}

func initMapPart2() map[string]int {
	scoreMap := make(map[string]int)
	scoreMap["AX"] = 3
	scoreMap["AY"] = 4
	scoreMap["AZ"] = 8
	scoreMap["BX"] = 1
	scoreMap["BY"] = 5
	scoreMap["BZ"] = 9
	scoreMap["CX"] = 2
	scoreMap["CY"] = 6
	scoreMap["CZ"] = 7

	return scoreMap
}
