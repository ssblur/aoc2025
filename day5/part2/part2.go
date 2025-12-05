package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	start uint64
	end   uint64
}

func newRange(start uint64, end uint64) *Range {
	return &Range{start: start, end: end}
}

func (rangeA *Range) inside(rangeB *Range) bool {
	return rangeA.start >= rangeB.start && rangeA.end <= rangeB.end
}

func (rangeA *Range) endsInside(rangeB *Range) bool {
	return rangeA.end >= rangeB.start && rangeA.end <= rangeB.end
}

func (rangeA *Range) size() uint64 {
	return rangeA.end - rangeA.start + 1
}

func (rangeA *Range) String() string {
	return fmt.Sprintf("%d-%d", rangeA.start, rangeA.end)
}

func in(key int, slice []int) bool {
	for _, k := range slice {
		if key == k {
			return true
		}
	}
	return false
}

func reduce(ranges *[]Range) []Range {
	out := make([]Range, 0)
	discard := make([]int, 0)

	for i, v := range *ranges {
		for j, w := range *ranges {
			if i >= j {
				continue
			}

			if v.inside(&w) {
				fmt.Printf("%v can be discarded because it fits into %v\n", v, w)
				discard = append(discard, i)
				break
			}

			if w.inside(&v) {
				fmt.Printf("%v can be discarded because it fits into %v\n", w, v)
				discard = append(discard, j)
				break
			}

			if v.endsInside(&w) || w.endsInside(&v) {
				fmt.Printf("%v and %v can be combined into %v\n", v, w, newRange(min(v.start, w.start), max(v.end, w.end)))
				discard = append(discard, i)
				discard = append(discard, j)
				out = append(out, *newRange(min(v.start, w.start), max(v.end, w.end)))
				break
			}
		}
	}

	for i, v := range *ranges {
		if !in(i, discard) {
			out = append(out, v)
		}
	}

	return out
}

func main() {
	file, fileErr := os.Open("day5/input.txt")
	if fileErr != nil {
		fmt.Println(fileErr)
		return
	}
	defer file.Close()

	var min uint64 = math.MaxInt
	var max uint64 = 0
	ranges := make([]Range, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		fields := strings.Split(line, "-")

		if len(fields) < 2 {
			break
		}

		start, convErr := strconv.ParseUint(fields[0], 10, 64)
		if convErr != nil {
			fmt.Println(convErr)
			continue
		}
		if start < min {
			min = start
		}

		end, convErr := strconv.ParseUint(fields[1], 10, 64)
		if convErr != nil {
			fmt.Println(convErr)
			continue
		}
		if end > max {
			max = end
		}
		ranges = append(ranges, *newRange(start, end))
	}

	r := len(ranges)
	for {
		ranges = reduce(&ranges)
		if r == len(ranges) {
			break
		}
		r = len(ranges)
		fmt.Println()
	}

	var fresh uint64 = 0
	for _, v := range ranges {
		fmt.Printf("Range %v has %d fresh ingredients\n", v, v.size())
		fresh += v.size()
	}

	fmt.Printf("There are %d fresh ingredients.\n", fresh)
}
