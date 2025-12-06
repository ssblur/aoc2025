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

	out, fileErr := os.Create("day6/data.txt")
	if fileErr != nil {
		fmt.Println(fileErr)
		return
	}

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	size := len(line)

	for i := size - 1; i >= 0; i-- {
		file.Seek(0, 0)
		scanner = bufio.NewScanner(file)
		for scanner.Scan() {
			line = scanner.Text()
			out.WriteString(string(line[i]))
			if line[i] == '*' || line[i] == '+' {
				out.WriteString("\n")
			}
		}
	}

	out.Close()
	file.Close()

	file, fileErr = os.Open("day6/data.txt")
	if fileErr != nil {
		fmt.Println(fileErr)
		return
	}
	defer file.Close()

	sum := uint64(0)
	scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()
		buffer := ""
		numbers := make([]uint64, 0)
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

					numbers = append(numbers, val)
				}

				switch c {
				case '*':
					x := uint64(1)
					for _, y := range numbers {
						x *= y
					}
					sum += x
				case '+':
					x := uint64(0)
					for _, y := range numbers {
						x += y
					}
					sum += x
				}
			}
		}
	}
	fmt.Printf("Sum is %d\n", sum)
}
