package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputPath := "./input"
	fmt.Println("--- Part One ---")
	fmt.Println(part1(inputPath))

	fmt.Println("--- Part Two ---")
	fmt.Println(part2(inputPath))
}

func part1(inputPath string) int {
	count := 0
	depthList := readNumbers(inputPath)

	for i := 0; i < len(depthList)-1; i++ {
		if isDepthIncrease(depthList[i], depthList[i+1]) {
			count++
		}
	}
	return count
}

func part2(inputPath string) int {
	count := 0
	depthList := readNumbers(inputPath)

	for i := 0; i < len(depthList)-3; i++ {
		firstSumOfThree := depthList[i] + depthList[i+1] + depthList[i+2]
		secondSumOfThree := depthList[i+1] + depthList[i+2] + depthList[i+3]
		if isDepthIncrease(firstSumOfThree, secondSumOfThree) {
			count++
		}
	}
	return count
}

func isDepthIncrease(depth1 int, depth2 int) bool {
	if depth1 < depth2 {
		return true
	} else {
		return false
	}
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