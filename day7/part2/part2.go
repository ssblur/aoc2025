package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	beams := make([]uint64, len(line)+1)

	file.Seek(0, 0)
	scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()

		for i, v := range line {
			if v == 'S' {
				beams[i] = 1
				fmt.Print("      S")
			} else if v == '^' && beams[i] > 0 {
				beams[i-1] += beams[i]
				beams[i+1] += beams[i]
				beams[i] = 0
				fmt.Print("      ^")
			} else {
				if beams[i] > 0 {
					if beams[i] >= 100000 {
						d := strconv.FormatUint(beams[i], 10)
						fmt.Printf("%s.%se%.3d ", string(d[0]), string(d[1]), len(d))
					} else {
						fmt.Printf("%6.d ", beams[i])
					}
				} else {
					fmt.Print("       ")
				}
			}
		}
		fmt.Print(string('\n'))
	}

	timelines := uint64(0)
	for _, i := range beams {
		timelines += i
	}
	fmt.Printf("Split %d times\n", timelines)
}
