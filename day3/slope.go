package main

import (
	"bufio"
	"fmt"
	"os"
)

func loadFile(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		fmt.Errorf("Error loading input file")
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func countTrees(lines []string, horizontalIncrement int, verticalIncrement int) int {
	horizontalIndex := 0
	verticalIndex := 0
	trees := 0
	for {
		// check if end of end of file
		if verticalIndex >= len(lines) {
			break
		}
		line := lines[verticalIndex]
		//check if non-empty line
		if len(line) > 0 {
			//reset the cursor to emulate endless pattern
			if horizontalIndex >= len(line) {
				horizontalIndex = horizontalIndex - len(line)
			}
			//check for tree
			if string(line[horizontalIndex]) == "#" {
				trees++
			}
			//increment the horizontal cursor
			horizontalIndex = horizontalIndex + horizontalIncrement
		}
		verticalIndex = verticalIndex + verticalIncrement
	}
	return trees
}

func main() {
	lines := loadFile("input")
	fmt.Println("1, 1 : ", countTrees(lines, 1, 1))
	fmt.Println("3, 1 : ", countTrees(lines, 3, 1))
	fmt.Println("5, 1 : ", countTrees(lines, 5, 1))
	fmt.Println("7, 1 : ", countTrees(lines, 7, 1))
	fmt.Println("1, 2 : ", countTrees(lines, 1, 2))
    fmt.Println("Multiplied : ", countTrees(lines, 1, 1)*countTrees(lines, 3, 1)*countTrees(lines, 5, 1) * countTrees(lines, 7, 1)*countTrees(lines, 1, 2))
}
