package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input, _ := os.Open("/mnt/c/Cdrive/github/adventofcode/2022/day02/input/input.txt")
	defer input.Close()

	var lines []string
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	var totalscore int
	for _, value := range lines {
		hands := strings.Fields(value)
		result, handscore := part1(hands[0], hands[1])
		totalscore += (result + handscore)
	}
	fmt.Println(totalscore)
}

func part1(opponent string, mine string) (int, int) {
	switch {
	case opponent == "A":
		if mine == "X" {
			return 3, 1
		} else if mine == "Y" {
			return 6, 2
		} else if mine == "Z" {
			return 0, 3
		}
	case opponent == "B":
		if mine == "X" {
			return 0, 1
		} else if mine == "Y" {
			return 3, 2
		} else if mine == "Z" {
			return 6, 3
		}
	case opponent == "C":
		if mine == "X" {
			return 6, 1
		} else if mine == "Y" {
			return 0, 2
		} else if mine == "Z" {
			return 3, 3
		}
	}
	return -1, -1
}

func part2() {

}
