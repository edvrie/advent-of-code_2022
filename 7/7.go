package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

var commands []string

func main() {
	fileName := "7_input.txt"
	input, _ := ioutil.ReadFile(fileName)
	commands = strings.Split(string(input), "\n")
	run()
}

func run() {
	root := &Node{name: "/", size: 0}
	current := root

	for _, element := range commands[1:] {
		symbols := strings.Split(element, " ")
		switch symbols[0] {
		case "$":
			if (symbols[1] == "cd") {
				if (symbols[2] == ".."){
					if current.parent == nil {
						continue
					}
					current = current.parent
				} else {
					child := current.children[findIndex(symbols[2], current.children)]
					current = child
				}
				} else {
					continue
				}
			case "dir":
				current.children = append(current.children, &Node{name: symbols[1], size: 0, parent: current})

		default:
			val, _ := strconv.Atoi(symbols[0])
			addToParents(val, current)
			current.size += val
		}
	}

	fmt.Println("Answer part 1: ", findSum(root))

	fmt.Println("Answer part 2: ", findMin(root))
}

func addToParents(size int, node *Node) {
	if node.parent != nil {
		node.parent.size += size
		addToParents(size, node.parent)
	}
}

func findSum(node *Node) int {
	sum := 0
	if (node.size <= 100000) {
		sum += node.size
	}

	for _, el := range node.children {
		sum += findSum(el)
	}

	return sum
}

func findMin(node *Node) int {
	var validSizes []int
	lim :=  30000000 - (70000000 - node.size)

	validSizes = getValid(node, validSizes, lim)

	min := math.MaxInt64

	for _, el := range validSizes {
		if (el < min) {
			min = el
		}
	}

	return min
}

func getValid(node *Node, validSizes []int, lim int) []int {
	if (node.size >= lim) {
		validSizes = append(validSizes, node.size)
	}
	for _, el := range node.children {
		validSizes = getValid(el, validSizes, lim)
	}

	return validSizes
}

func findIndex(value string, children []*Node) int {
	for index := range children {
		if (string(children[index].name) == value) {
			return index
		}
	}

	return -1
}

//---------------TREE--------------

type Node struct {
	name string
	size int
	parent *Node
	children []*Node
}