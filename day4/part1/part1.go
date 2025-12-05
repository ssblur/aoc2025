package main

import (
	"bufio"
	"fmt"
	"os"
)

func adjacent(lines [3]string, column int) int {
	sum := 0
	for y := 0; y < len(lines); y++ {
		for x := -1; x <= 1; x++ {
			if len(lines[y]) > column+x && column+x >= 0 && lines[y][column+x] == '@' {
				sum++
			}
		}
	}
	return sum - 1
}

func main() {
	rolls := 0
	file, fileErr := os.Open("day4/input.txt")
	if fileErr != nil {
		fmt.Println(fileErr)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := [3]string{}
	for scanner.Scan() {
		for i := 1; i < len(lines); i++ {
			lines[i-1] = lines[i]
		}
		lines[len(lines)-1] = scanner.Text()

		if len(lines[1]) > 0 {
			for i := 0; i < len(lines[1]); i++ {
				if lines[1][i] == '@' && adjacent(lines, i) < 4 {
					fmt.Print("x")
					rolls++
				} else {
					fmt.Print(string(lines[1][i]))
				}
			}
		}

		fmt.Print("\n")
	}

	for i := 1; i < len(lines); i++ {
		lines[i-1] = lines[i]
	}
	lines[len(lines)-1] = ""
	for i := 0; i < len(lines[1]); i++ {
		if lines[1][i] == '@' && adjacent(lines, i) < 4 {
			fmt.Print("x")
			rolls++
		} else {
			fmt.Print(string(lines[1][i]))
		}
	}

	fmt.Printf("\n%d rolls of paper are accessible.\n", rolls)

}
