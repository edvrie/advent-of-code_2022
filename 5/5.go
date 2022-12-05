package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	startTime := time.Now()

	fileName := "./5_input.txt"

	file, _ := os.Open(fileName)

	scanner := bufio.NewScanner(file)

	var containers []stack

	for scanner.Scan() {
		for i := 0; i <= len(scanner.Text())/4; i++ {
			var tempStack stack
			containers = append(containers, tempStack)
		}
		break
	}

	file.Seek(0, io.SeekStart)

	newScanner := bufio.NewScanner(file)

	for newScanner.Scan() {
		if newScanner.Text() == "" {
			break
		}
		current := strings.Split(newScanner.Text(), "")
		for i := 0; i < len(newScanner.Text())-1; i = i + 4 {
			if current[i] == "[" {
				if (i/4)-1 < 0 {
					containers[0].reversePush(current[i+1])
				} else {
					containers[(i / 4)].reversePush(current[i+1])
				}
			}
		}
	}
	fmt.Println(containers)

	containersPart1 := append([]stack(nil), containers...)
	copy(containersPart1, containers)
	containersPart2 := append([]stack(nil), containers...)
	copy(containersPart2, containers)

	for newScanner.Scan() {
		instructions := strings.Split(newScanner.Text(), " ")
		qty, _ := strconv.Atoi(instructions[1])
		from, _ := strconv.Atoi(instructions[3])
		to, _ := strconv.Atoi(instructions[5])
		part1(qty, from, to, containersPart1)
		part2(qty, from, to, containersPart2)
	}
	fmt.Println("Answer part 1: ")
	for _, el := range containersPart1 {
		fmt.Print(el.pop())
	}
	fmt.Println()
	fmt.Println("Answer part 2: ")
	for _, el := range containersPart2 {
		fmt.Print(el.pop())
	}
	fmt.Println()
	elapsed := time.Since(startTime)

	fmt.Printf("Done in:  %s", elapsed)
}

func part1(qty, from, to int, containers []stack) {
	for i := 0; i < qty; i++ {
		item := containers[from-1].pop()
		containers[to-1].push(item)
	}
}

func part2(qty, from, to int, containers []stack) {
	if qty-1 == 0 {
		item := containers[from-1].pop()
		containers[to-1].push(item)
	} else {
		items := containers[from-1].popMany(qty)
		containers[to-1].pushMany(items)
	}
}

// -----------------STACK------------------

type stack []string

func (s *stack) push(item string) {
	*s = append(*s, item)
}

func (s *stack) reversePush(item string) {
	res := make([]string, len(*s)+1)
	copy(res[1:], *s)
	res[0] = item
	*s = res
}

func (s *stack) pushMany(items []string) {
	revItems := make([]string, len(items))
	for i, j := 0, len(items)-1; i <= j; i, j = i+1, j-1 {
		revItems[i], revItems[j] = items[j], items[i]
	}

	*s = append(*s, revItems...)
}

func (s *stack) pop() string {
	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index]

	return element
}

func (s *stack) popMany(count int) []string {
	var items []string
	for i := 0; i < count; i++ {
		items = append(items, s.pop())
	}

	return items
}
