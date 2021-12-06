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

func part1(inputPath string) int64 {
	defer timeTrack(time.Now(), "Part One")
	return powerConsumption(initTwoDimBitArray(readStrings(inputPath)))
}

func part2(inputPath string) int64 {
	defer timeTrack(time.Now(), "Part Two")
	array := initTwoDimBitArray(readStrings(inputPath))
	return getOxygenRate(array) * getCO2Rate(array)
}

func powerConsumption(bitArray [][]string) int64 {
	var binGamma []string
	var binEpsilon []string

	for i := range bitArray[0] {
		if mostCommonBit(bitArray, i) == "0" {
			binGamma = append(binGamma, "0")
			binEpsilon = append(binEpsilon, "1")
		} else {
			binGamma = append(binGamma, "1")
			binEpsilon = append(binEpsilon, "0")
		}
	}

	gammaRate, _ := strconv.ParseInt(strings.Join(binGamma, ""), 2, 64)
	epsilonRate, _ := strconv.ParseInt(strings.Join(binEpsilon, ""), 2, 64)

	return gammaRate * epsilonRate
}

func getOxygenRate(bitArray [][]string) int64 {
	i := 0
	for len(bitArray) != 1 {
		if mostCommonBit(bitArray, i) == "0" {
			bitArray = filterArray(bitArray, i, "0")
		} else if mostCommonBit(bitArray, i) == "=" {
			bitArray = filterArray(bitArray, i, "1")
		} else {
			bitArray = filterArray(bitArray, i, "1")
		}
		i++
	}

	oxygenRate, _ := strconv.ParseInt(strings.Join(bitArray[0], ""), 2, 64)

	return oxygenRate
}

func getCO2Rate(bitArray [][]string) int64 {
	i := 0
	for len(bitArray) != 1 {
		if mostCommonBit(bitArray, i) == "0" {
			bitArray = filterArray(bitArray, i, "1")
		} else if mostCommonBit(bitArray, i) == "=" {
			bitArray = filterArray(bitArray, i, "0")
		} else {
			bitArray = filterArray(bitArray, i, "0")
		}
		i++
	}

	CO2Rate, _ := strconv.ParseInt(strings.Join(bitArray[0], ""), 2, 64)

	return CO2Rate
}

func filterArray(bitArray [][]string, column int, value string) [][]string {
	var filteredArray [][]string
	for _, bits := range bitArray {
		if bits[column] == value {
			filteredArray = append(filteredArray, bits)
		}
	}
	return filteredArray
}

func mostCommonBit(bitArray [][]string, column int) string {
	countZero := 0
	countOne := 0

	for i := range bitArray {
		if bitArray[i][column] == "0" {
			countZero++
		} else {
			countOne++
		}
	}

	if countZero > countOne {
		return "0"
	} else if countZero == countOne {
		return "="
	} else {
		return "1"
	}
}

func initTwoDimBitArray(rows []string) [][]string {
	x := len(strings.SplitAfter(rows[1], ""))
	y := len(rows)

	bitArray := make([][]string, y)
	for i := range bitArray {
		bitArray[i] = make([]string, x)
	}

	for i, row := range rows {
		for j, bit := range strings.SplitAfter(row, "") {
			bitArray[i][j] = bit
		}
	}
	return bitArray
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
