package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fresh := 0
	file, fileErr := os.Open("day5/input.txt")
	if fileErr != nil {
		fmt.Println(fileErr)
		return
	}
	defer file.Close()

	ranges := make([][2]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		fields := strings.Split(line, "-")

		if len(fields) < 2 {
			fmt.Println(fields)
			break
		}

		start, convErr := strconv.Atoi(fields[0])
		if convErr != nil {
			fmt.Println(convErr)
			continue
		}

		end, convErr := strconv.Atoi(fields[1])
		if convErr != nil {
			fmt.Println(convErr)
			continue
		}
		ranges = append(ranges, [2]int{start, end})
	}

scan:
	for scanner.Scan() {
		line := scanner.Text()

		id, convErr := strconv.Atoi(line)
		if convErr != nil {
			continue
		}
		for i := 0; i < len(ranges); i++ {
			fmt.Printf("%d in %d-%d?", id, ranges[i][0], ranges[i][1])
			if id >= ranges[i][0] && id <= ranges[i][1] {
				fmt.Print(" yes\n")
				fresh++
				continue scan
			} else {
				fmt.Print(" no\n")
			}
		}
	}

	fmt.Printf("There are %d fresh ingredients.\n", fresh)

}
