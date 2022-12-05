package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	// startTime := time.Now()

	fileName := "./5_input.txt"	

	file, _ := os.Open(fileName)

	scanner := bufio.NewScanner(file)

	var containers []stack

	for scanner.Scan() {
		for i:=0; i <= len(scanner.Text())/4; i = i + 1 {
				var tempStack stack
				containers = append(containers, tempStack)
		}
		break;
	}

	fmt.Println(containers)

file.Seek(0, io.SeekStart)

newScanner := bufio.NewScanner(file)

	for newScanner.Scan() {
		if (newScanner.Text() == "") {
			break
		}
		current := strings.Split(newScanner.Text(), "")
		for i:=0; i < len(newScanner.Text()) - 1; i = i + 4 {
			if current[i] == "[" {
				fmt.Println(i, (i/4)-1, current[i+1])
				if (i/4 - 1 < 0){
					containers[0].reversePush(current[i+1])
				} else {
					containers[(i/4)].reversePush(current[i+1])
				}
			}
		}
	}
	fmt.Println(containers)

	for newScanner.Scan() {
		instructions := strings.Split(newScanner.Text(), " ")
		qty, _ := strconv.Atoi(instructions[1])
		from, _ := strconv.Atoi(instructions[3])
		to, _ := strconv.Atoi(instructions[5])
		if (qty-1 == 0) {
			item := containers[from-1].pop()
			containers[to-1].push(item)
		} else {
			for i := 0; i < qty; i++ {
				item := containers[from-1].pop()
				containers[to-1].push(item)
			}
		}
	}

	var answers []string

	for _, el := range containers {
		answers = append(answers, el.pop())
	}

fmt.Println(answers)
}

// -----------------STACK------------------

type stack []string

func (s *stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *stack) push(item string) {
	*s = append(*s, item)
}

func (s *stack) reversePush(item string) {
	res := make([]string, len(*s)+1)
	copy(res[1:], *s)
	res[0] = item
	*s = res
}

func (s *stack) pop() (string) {
	if (s.isEmpty()) {
		return ""
	}

	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index]

	return element
}