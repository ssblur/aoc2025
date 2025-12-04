package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sum := 0
	file, fileErr := os.Open("day3/input.txt")
	if fileErr != nil {
		fmt.Println(fileErr)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		index := 0
		for i := 1; i < len(line)-1; i++ {
			if line[index] < line[i] {
				index = i
			}
		}

		indexB := index + 1
		for i := index + 1; i < len(line); i++ {
			if line[indexB] < line[i] {
				indexB = i
			}
		}

		fmt.Printf("Bank joltage: %d\n", int((10*(line[index]-'0'))+(line[indexB]-'0')))
		sum += int((10 * (line[index] - '0')) + (line[indexB] - '0'))
	}

	fmt.Printf("\nTotal Joltage: %d\n\n", sum)

}
