package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fileName := "8_input.txt"

	fileLines, _ := ioutil.ReadFile(fileName)

	rows := strings.Split(string(fileLines), "\n")
	treeMap := make([][]int, len(rows))

	for index, el := range rows {
		columns := strings.Split(el, "")
		for _, jel := range columns {
			value,  _ := strconv.Atoi(jel)

			treeMap[index] = append(treeMap[index], value)
		}
	}

	sum := len(treeMap)*4 - 4

	fmt.Println(sum)
	visitMap := make([][]int, len(treeMap))

	for i := range visitMap {
		for j:=0; j < len(treeMap); j++ {
			if (i == 0 || i == len(treeMap)-1) {

			visitMap[i] = append(visitMap[i], 1)
			continue
			}

			if (j == 0 || j == len(treeMap)-1) {
			visitMap[i] = append(visitMap[i], 1)
			continue
			}
			visitMap[i] = append(visitMap[i], 0)
		}
	}

	newSum, newSum2 := 0, 0;
	for i := 0; i < len(treeMap); i++ {
		newSum, visitMap = findRowCount(treeMap[i], i, visitMap)
		sum += newSum
		newSum2, visitMap = findRowCountRev(treeMap[i], i, visitMap)
		sum+= newSum2
	}

	fmt.Println(sum)
	

	for i := 0; i < len(treeMap); i++ {
		max := treeMap[0][i]
		for j := 0; j < len(treeMap); j++ {
			if (treeMap[j][i] <= max) {
				continue
			}
			max = treeMap[j][i]

			if (visitMap[j][i] == 1) {
				continue
			}
			visitMap[j][i] = 1
			sum += 1
		}
	}

	for i := 0; i < len(treeMap); i++ {
		max := treeMap[0][len(treeMap)-1]
		for j := len(treeMap)-1; j > 0; j-- {
			fmt.Println("AA: ", j, i)
			if (treeMap[j][i] <= max) {
				continue
			}

			max = treeMap[j][i]

			if (visitMap[j][i] == 1) {
				continue
			}

			visitMap[j][i] = 1
			sum += 1
		}
	}

	fmt.Println()

	fmt.Println(sum)

	//1837 is wrong

	// var columns []int
	for _, el := range visitMap {

		fmt.Println(el)

	}
}

func findRowCount(treeRow []int, index int, visitMap [][]int) (int, [][]int) {
	sum := 0
	max := treeRow[0]

	for i := range treeRow {
		// fmt.Println(treeRow[i])
		if (treeRow[i] <= max) {
			continue
		}
		max = treeRow[i]

		if (visitMap[index][i] == 1) {
			continue
		}

		visitMap[index][i] = 1
		// fmt.Println(treeRow[i])
		sum += 1
	}

	// fmt.Println("Row: ", visitMap)
	return sum, visitMap
}

func findRowCountRev(treeRow []int, index int, visitMap [][]int) (int, [][]int) {
	sum:=0
	max:=treeRow[len(treeRow) -1]

	for i:=len(treeRow)-1; i>0; i-- {
		if (treeRow[i] <= max) {
			continue
		}
		max = treeRow[i]

		if (visitMap[index][i] == 1) {
			continue
		}

		visitMap[index][i] = 1
		sum += 1
	}
	// fmt.Println("Reverse row: ", visitMap)
	return sum, visitMap
}
