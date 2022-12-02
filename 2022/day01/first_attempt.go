//change package to main when running later
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("/mnt/c/Cdrive/github/adventofcode/2022/day01/input/input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var caloriesum int
	cal_array := []int{}

	for _, val := range lines {
		if val == "" {
			cal_array = append(cal_array, caloriesum)
			caloriesum = 0
		} else {
			val_int, err := strconv.Atoi(val)
			if err != nil {
				log.Fatal(err)
			}
			caloriesum += val_int
		}

	}
	sort.Ints(cal_array)
	cal_len := len(cal_array)

	//getting the the largest. Golang doesn't have builtin func to get max !
	fmt.Println(cal_array[cal_len-1])

}
