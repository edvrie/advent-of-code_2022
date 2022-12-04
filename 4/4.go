package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	startTime := time.Now()
	fileName := "./4_input.txt"
	file, _ := os.Open(fileName)

	scanner := bufio.NewScanner(file)

	count1 := 0
	count2 := 0

	for scanner.Scan() {
		intervals := strings.Split(scanner.Text(), ",")
		firstPairInt := strings.Split(intervals[0], "-")
		secondPairInt := strings.Split(intervals[1], "-")

		if isFullyOverlapping(parseFloat(firstPairInt[0]), parseFloat(secondPairInt[0]), parseFloat(firstPairInt[1]), parseFloat(secondPairInt[1])) {
			count1 += 1
		}

		if isOverlapping(parseFloat(firstPairInt[0]), parseFloat(secondPairInt[0]), parseFloat(firstPairInt[1]), parseFloat(secondPairInt[1])) {
			count2 += 1
		}
	}
	fmt.Println("Answer part 1: ", count1)
	fmt.Println("Answer part 2: ", count2)

	file.Close()

	elapsed := time.Since(startTime)

	fmt.Printf("Done in:  %s", elapsed)
}

func parseFloat(input string) float64 {
	parsed, _ := strconv.ParseFloat(input, 64)
	return parsed
}

func isFullyOverlapping(x1, x2, y1, y2 float64) bool {
	return (x1 <= x2 && y2 <= y1) || (x2 <= x1 && y1 <= y2)
}

func isOverlapping(x1, x2, y1, y2 float64) bool {
	return math.Max(x1, x2) <= math.Min(y1, y2)
}
