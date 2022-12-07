package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var commands []string

func main() {
	fileName := "7_input.txt"
	input, _ := ioutil.ReadFile(fileName)
	commands = strings.Split(string(input), "\n")
	// fmt.Println(commands)
	run()

	fmt.Println()
}

func run() int {
	root := &Node{name: "/", size: 0}
	current := root
	// size := 0
	// fmt.Println(commands[1:])
	for _, element := range commands[1:] {
		// fmt.Println(current)
		// for _, t := range current.children {
		// 	fmt.Println(string(t.name))
		// }
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
			for _, t := range root.children {
			fmt.Println(string(t.name), t.size)
		}

	fmt.Println(root.children)
	return 0
}

func addToParents(size int, node *Node) {
	if node.parent != nil {
		node.parent.size += size
		addToParents(size, node.parent)
	}
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