package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	inputPath := "./input"
	fmt.Println("--- Part One ---")
	fmt.Println(part1(inputPath))

	fmt.Println("--- Part Two ---")
	fmt.Println(part2(inputPath))
}

func part1(inputPath string) int {
	defer timeTrack(time.Now(), "Part One")
	commands := readStrings(inputPath)
	start := [2]int{0, 0}

	for _, command := range commands {
		start = move01(start, command)
	}

	return start[0] * start[1]
}

func part2(inputPath string) int {
	defer timeTrack(time.Now(), "Part Two")
	commands := readStrings(inputPath)
	start := [3]int{0, 0, 0}

	for _, command := range commands {
		start = move02(start, command)
	}

	return start[0] * start[1]
}

func move01(origin [2]int, command string) [2]int {
	var target [2]int

	splittedCommand := strings.Split(command, " ")
	action := splittedCommand[0]
	move := toInt(splittedCommand[1])

	if action == "forward" {
		target[0] = origin[0] + move
		target[1] = origin[1]
	} else if action == "down" {
		target[0] = origin[0]
		target[1] = origin[1] + move
	} else if action == "up" {
		target[0] = origin[0]
		target[1] = origin[1] - move
	}
	return target
}

func move02(origin [3]int, command string) [3]int {
	var target [3]int

	splittedCommand := strings.Split(command, " ")
	action := splittedCommand[0]
	move := toInt(splittedCommand[1])

	if action == "forward" {
		target[0] = origin[0] + move
		target[1] = origin[1] + origin[2]*move
		target[2] = origin[2]
	} else if action == "down" {
		target[0] = origin[0]
		target[1] = origin[1]
		target[2] = origin[2] + move
	} else if action == "up" {
		target[0] = origin[0]
		target[1] = origin[1]
		target[2] = origin[2] - move
	}
	return target
}

func readStrings(filename string) []string {
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var text []string
	for scanner.Scan() {
		text = append(text, strings.TrimRight(scanner.Text(), "\n"))
	}
	return text
}

func readNumbers(filename string) []int {
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	Scanner := bufio.NewScanner(file)

	var numbers []int
	for Scanner.Scan() {
		numbers = append(numbers, toInt(Scanner.Text()))
	}
	return numbers
}

func readRaw(filename string) string {
	content, err := ioutil.ReadFile(filename)
	check(err)
	return strings.TrimRight(string(content), "\n")
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func toInt(s string) int {
	result, err := strconv.Atoi(s)
	check(err)
	return result
}

func max(numbers []int) int {
	currMax := numbers[0]
	for _, val := range numbers {
		if val > currMax {
			currMax = val
		}
	}
	return currMax
}

func min(numbers []int) int {
	currMin := numbers[0]
	for _, val := range numbers {
		if val < currMin {
			currMin = val
		}
	}
	return currMin
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Println(name, "took", elapsed)
}
