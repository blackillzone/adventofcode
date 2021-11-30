package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type password struct {
	min, max         int
	letter, password string
}

func (p password) CheckValidityFirstPart() bool {
	result := strings.Count(p.password, p.letter)

	if (result >= p.min) && (result <= p.max) {
		return true
	} else {
		return false
	}
}

func (p password) CheckValiditySecondPart() bool {
	result := strings.SplitAfter(p.password, "")

	if ((result[p.min] == p.letter) && (result[p.max] != p.letter)) || ((result[p.min] != p.letter) && (result[p.max] == p.letter)) {
		return true
	} else {
		return false
	}
}

func main() {
	inputPath := "./input"
	fmt.Println("--- Part One ---")
	fmt.Println(part1(inputPath))

	fmt.Println("--- Part Two ---")
	fmt.Println(part2(inputPath))
}

func part1(inputPath string) int {
	input := readStrings(inputPath)
	return countValidPasswords(parseInput(input), true)
}

func part2(inputPath string) int {
	input := readStrings(inputPath)
	return countValidPasswords(parseInput(input), false)
}

func parseInput(input []string) []password {
	var splittedLine []string
	var rules []string
	var charRange []string
	var passwordList []password

	passwordList = nil

	for _, l := range input {
		splittedLine = strings.Split(l, ":")
		rules = strings.Split(splittedLine[0], " ")
		charRange = strings.Split(rules[0], "-")
		p := password{
			min:      toInt(charRange[0]),
			max:      toInt(charRange[1]),
			letter:   rules[1],
			password: splittedLine[1],
		}
		passwordList = append(passwordList, p)
	}
	return passwordList
}

func countValidPasswords(passwords []password, part1 bool) int {
	var Count int

	Count = 0

	for _, p := range passwords {
		if part1 {
			if p.CheckValidityFirstPart() {
				Count++
			}
		} else {
			if p.CheckValiditySecondPart() {
				Count++
			}
		}
	}
	return Count
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
