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

func process(filename string, outFilename string) int {
	rolls := 0
	file, fileErr := os.Open(filename)
	if fileErr != nil {
		fmt.Println(fileErr)
		return 0
	}
	defer file.Close()
	outFile, fileErr := os.Create(outFilename)
	if fileErr != nil {
		fmt.Println(fileErr)
		return 0
	}
	defer outFile.Close()

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
					outFile.WriteString(string('x'))
					rolls++
				} else {
					outFile.WriteString(string(lines[1][i]))
				}
			}
			outFile.WriteString("\n")
		}
	}

	for i := 1; i < len(lines); i++ {
		lines[i-1] = lines[i]
	}
	lines[len(lines)-1] = ""
	for i := 0; i < len(lines[1]); i++ {
		if lines[1][i] == '@' && adjacent(lines, i) < 4 {
			outFile.WriteString("x")
			rolls++
		} else {
			outFile.WriteString(string(lines[1][i]))
		}
	}

	fmt.Printf("%d rolls of paper are accessible.\n", rolls)
	return rolls
}

func main() {
	rolls := process("day4/input.txt", "day4/dataA.txt")
	whichFile := true
	var r int
	for {
		if whichFile {
			r = process("day4/dataA.txt", "day4/dataB.txt")

		} else {
			r = process("day4/dataB.txt", "day4/dataA.txt")
		}
		whichFile = !whichFile
		if r == 0 {
			break
		}
		rolls += r
	}

	fmt.Printf("\nBy the end, %d rolls of paper will be accessible.\n", rolls)
}
