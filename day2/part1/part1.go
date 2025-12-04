package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func checkForInvalid(start uint64, end uint64) uint64 {
	var sum uint64 = 0
	var i uint64
	fmt.Print("Invalid Input: ")
	for i = start; i <= end; i++ {
		number := strconv.FormatUint(i, 10)
		n := len(number)
		if (n % 2) == 1 {
			continue
		}

		if number[0:n/2] == number[n/2:n] {
			sum += i
			fmt.Printf("%d, ", i)
		}
	}
	fmt.Print("\n")

	return sum
}

func main() {
	var sum uint64 = 0

	file, fileErr := os.Open("day2/input.txt")
	if fileErr != nil {
		fmt.Println(fileErr)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	var start uint64 = 0
	var end uint64 = 0
	buffer := ""
	for {
		char, _, err := reader.ReadRune()
		if err != nil && err != io.EOF {
			break
		}

		if char == '-' {
			val, convErr := strconv.ParseUint(buffer, 10, 64)
			if convErr != nil {
				break
			}
			start = val
			buffer = ""
		} else if char == ',' || err == io.EOF {
			val, convErr := strconv.ParseUint(buffer, 10, 64)
			if convErr != nil {
				break
			}
			end = val
			fmt.Printf("\nChecking Range: %d-%d\n", start, end)
			sum += checkForInvalid(start, end)
			buffer = ""
			if err == io.EOF {
				break
			}
		} else if char >= '0' && char <= '9' {
			buffer += string(char)
		}
	}

	fmt.Printf("Sum: %d\n", sum)
}
