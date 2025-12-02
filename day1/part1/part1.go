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

		if line[0] == 'L' {
			turns *= -1
		}
		dial += turns
		dial %= 100
		fmt.Printf("Turned %d to land at %d\n", turns, dial)

		if dial == 0 {
			fmt.Println("Dial at 0, incrementing password")
			password++
		}
	}

	fmt.Printf("\nPassword is %d\n\n", password)

}
