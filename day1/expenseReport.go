package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		fmt.Errorf("Could not load input file")
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	var lines []int
	for scanner.Scan() {
		intValue, _ := strconv.Atoi(scanner.Text())
		lines = append(lines, intValue)
	}

	for _, i := range lines {
		for _, k := range lines {
			for _, j := range lines {
				if i+k+j == 2020 {
					fmt.Println(i*k*j)
					os.Exit(0)
				}
			}
		}
	}
}
