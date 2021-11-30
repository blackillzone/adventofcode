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
	expenses := readNumbers(inputPath)
	return sumOfTwoEqualsTo2020(expenses, 2020)
}

func part2(inputPath string) int {
	expenses := readNumbers(inputPath)
	return sumOfThreeEqualsTo2020(expenses)
}

func sumOfTwoEqualsTo2020(numbers []int, target int) int {
	for _, number := range numbers {
		diff := target - number
		if isElementExist(numbers, diff) {
			return diff * number
		}
	}
	return 0
}

func sumOfThreeEqualsTo2020(numbers []int) int {
	for i, number := range numbers {
		diff := 2020 - number
		// Call previous function, and exclude already tested values
		third := sumOfTwoEqualsTo2020(numbers[i:], diff)
		if third != 0 {
			return third * number
		}
	}
	return 0
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

func isElementExist(i []int, num int) bool {
  for _, v := range i {
    if v == num {
      return true
    }
  }
  return false
}
