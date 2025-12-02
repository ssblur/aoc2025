package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	dial := 50
	password := 0
	file, fileErr := os.Open("day1/input.txt")
	if fileErr != nil {
		fmt.Println(fileErr)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		turns, convErr := strconv.Atoi(line[1:])
		if convErr != nil {
			continue
		}

		signum := 1
		if line[0] == 'L' {
			turns *= -1
			signum = -1
		}
		fmt.Printf("Turning %d from %d\n", turns, dial)

		for i := 0; i != turns; i += signum {
			dial += signum
			dial = (dial + 100) % 100
			if dial == 0 {
				password++
				fmt.Println("Passed 0")
			}
		}

	}

	fmt.Printf("\nPassword is %d\n\n", password)
}
