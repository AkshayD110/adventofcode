package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("/mnt/c/Cdrive/github/adventofcode/2022/day03/input/input.txt")
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	sum_of_priorities := part1(lines)
	fmt.Println(sum_of_priorities)
}

func part1(rucksacks []string) int {
	var asci_repeatItem rune
	var priority int
	for _, rucksack := range rucksacks {
		items_in_rucksack := len(rucksack)
		for i := 0; i < items_in_rucksack/2; i++ {
			for j := items_in_rucksack / 2; j < items_in_rucksack; j++ {
				if string(rucksack[i]) == string(rucksack[j]) {
					asci_repeatItem = rune(rucksack[i])
				}
			}
		}
		// fmt.Println("In the rucksack ", rucksack, " the repeat item is: ", string(asci_repeatItem), " and it has the value: ", asci_repeatItem)
		if asci_repeatItem > 91 {
			priority += int(asci_repeatItem) - 96
		} else {
			priority += int(asci_repeatItem) - 38
		}

	}
	return priority
}

func part2() {

}
