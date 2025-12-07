package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, fileErr := os.Open("day7/input.txt")
	if fileErr != nil {
		fmt.Println(fileErr)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()

	beams := make([]bool, len(line)+1)
	splits := 0

	file.Seek(0, 0)
	scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()

		for i, v := range line {
			if v == 'S' {
				beams[i] = true
				fmt.Print("S")
			} else if v == '^' && beams[i] {
				beams[i-1] = true
				beams[i+1] = true
				beams[i] = false
				splits++
				fmt.Print("^")
			} else {
				if beams[i] {
					fmt.Print("|")
				} else {
					fmt.Print(" ")
				}
			}
		}
		fmt.Print(string('\n'))
	}

	fmt.Printf("Split %d times\n", splits)

}
