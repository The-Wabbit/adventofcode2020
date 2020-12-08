package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func checkContents(lines []string) (int,int) {
	validTask1 := 0
	validTask2 := 0
	for _, line := range lines {
		content := strings.Split(line, " ")
		//check if we have all 3 parameters
		if len(content) == 3 {
			minmax := strings.Split(content[0], "-")
			min, _ := strconv.Atoi(minmax[0])
			max, _ := strconv.Atoi(minmax[1])
			reference := strings.Trim(content[1], ":")
			counter := strings.Count(content[2], reference)
			if counter >= min && counter <= max {
				validTask1++
			}
			if string(content[2][min-1]) == reference || string(content[2][max-1]) == reference {
				if (content[2][min-1] != content[2][max-1]) {
					validTask2++
				}
			}

		}
	}
	return validTask1, validTask2
}

func main() {
	lines := loadFile("input")
	fmt.Println(checkContents(lines))
}
