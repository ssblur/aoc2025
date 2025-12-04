package main

import (
	"bufio"
	"fmt"
	"math"
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
		indices := [12]int{}
		for i := 0; i < len(indices); i++ {
			if i > 0 {
				indices[i] = indices[i-1] + 1
			}

			for x := indices[i]; x < len(line)-(len(indices)-i-1); x++ {
				if line[x] > line[indices[i]] {
					indices[i] = x
				}
			}
		}

		bank := 0
		for i := 0; i < len(indices); i++ {
			bank += int(math.Pow(10, float64(len(indices)-i-1))) * int(line[indices[i]]-'0')
		}
		fmt.Printf("\nBank Joltage: %d\n", bank)
		fmt.Println(indices)
		sum += bank
	}

	fmt.Printf("\nTotal Joltage: %d\n\n", sum)

}
