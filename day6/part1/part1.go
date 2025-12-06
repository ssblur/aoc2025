package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, fileErr := os.Open("day6/input.txt")
	if fileErr != nil {
		fmt.Println(fileErr)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var line string
	for scanner.Scan() {
		line = scanner.Text()
	}

	operators := make([]rune, 0)
	for _, c := range line {
		if c == '*' || c == '+' {
			operators = append(operators, c)
		}
	}
	fmt.Printf("Operators are %s\n", string(operators))

	file.Seek(0, 0)
	scanner = bufio.NewScanner(file)
	var results []uint64 = make([]uint64, len(operators))

	for i, v := range operators {
		if v == '*' {
			results[i] = 1
		} else {
			results[i] = 0
		}
	}

	var buffer string
	for scanner.Scan() {
		line = scanner.Text() + " "
		i := 0
		buffer = ""
		for _, c := range line {
			if c >= '0' && c <= '9' {
				buffer += string(c)
			} else {
				if len(buffer) > 0 {
					val, err := strconv.ParseUint(buffer, 10, 64)
					buffer = ""
					if err != nil {
						fmt.Println(err)
						continue
					}

					if operators[i] == '*' {
						results[i] *= val
					} else {
						results[i] += val
					}
					i++
				}
			}
		}
	}

	fmt.Printf("Results are %v\n", results)

	var sum uint64 = 0
	for _, v := range results {
		sum += v
	}
	fmt.Printf("Total is %d\n\n", sum)
}
