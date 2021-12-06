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

type number struct {
	value int
	drawn bool // Boolean to check if the value has been drawn
}

type grid struct {
	table     [][]number
	validated bool // Boolean to check if grid has been validated or no
}

func checkIfGridWin(g [][]number) bool {
	// Check on each lines
	for i := range g {
		// Reset to 0 for each line checks
		countPerLines := 0
		for _, num := range g[i] {
			if num.drawn == true {
				countPerLines++
			}
		}
		if countPerLines == 5 {
			return true
		}
	}

	// Check on each columns
	for i := range g[0] {
		// Reset to 0 for each columns checks
		countPerColumns := 0
		for _, line := range g {
			if line[i].drawn == true {
				countPerColumns++
			}
		}
		if countPerColumns == 5 {
			return true
		}
	}

	return false
}

func initBingoGrid(rows []string) [][]number {
	bingoGrid := make([][]number, 5)
	for i := range bingoGrid {
		bingoGrid[i] = make([]number, 5)
	}

	for i, row := range rows {
		for j, num := range strings.Fields(row) {
			bingoGrid[i][j] = number{
				value: toInt(num),
				drawn: false,
			}
		}
	}
	return bingoGrid
}

func initGridList(input []string) []grid {
	// Divided by 6 (5 lines with numbers, and one empty line)
	bingoGridLenght := len(input) / 6

	var bingoList []grid

	// Init range of first bingo grid to parse
	rangeStart := 0
	rangeEnd := 5

	for i := 0; i <= bingoGridLenght; i++ {
		bingoList = append(bingoList, grid{table: initBingoGrid(input[i+rangeStart : i+rangeEnd]), validated: false})
		rangeStart += 5
		rangeEnd += 5
	}
	return bingoList
}

func initNumberList(input string) []int {
	var intList []int
	strList := strings.Split(input, ",")
	for _, number := range strList {
		intList = append(intList, toInt(number))
	}
	return intList
}

func drawNumberInGrid(num int, grid [][]number) [][]number {
	for i, line := range grid {
		for j, n := range line {
			if n.value == num {
				grid[i][j] = number{
					value: n.value,
					drawn: true,
				}
			}
		}
	}
	return grid
}

func main() {
	inputPath := "./input"
	fmt.Println("--- Part One ---")
	fmt.Println(part1(inputPath))

	fmt.Println("--- Part Two ---")
	fmt.Println(part2(inputPath))
}

func part1(inputPath string) int {
	defer timeTrack(time.Now(), "Part One")
	input := readStrings(inputPath)
	numList := initNumberList(input[0])
	gridList := initGridList(input[2:])
	finalScore := 0
	finalDrawn := 0

mainloop:
	for _, num := range numList {
		for _, grid := range gridList {
			grid.table = drawNumberInGrid(num, grid.table)
			if checkIfGridWin(grid.table) {
				finalDrawn = num
				for i := range grid.table {
					for _, n := range grid.table[i] {
						if n.drawn == false {
							finalScore += n.value
						}
					}
				}
				break mainloop
			}
		}
	}
	return finalScore * finalDrawn
}

func part2(inputPath string) int {
	defer timeTrack(time.Now(), "Part Two")
	input := readStrings(inputPath)
	numList := initNumberList(input[0])
	gridList := initGridList(input[2:])
	finalScore := 0
	finalDrawn := 0
	var lastGrid grid

	for _, num := range numList {
		for i, grid := range gridList {
			if !grid.validated {
				gridList[i].table = drawNumberInGrid(num, gridList[i].table)
				if checkIfGridWin(gridList[i].table) {
					gridList[i].validated = true
					lastGrid = gridList[i]
					finalDrawn = num
				}
			}
		}
	}

	for i := range lastGrid.table {
		for _, n := range lastGrid.table[i] {
			if n.drawn == false {
				finalScore += n.value
			}
		}
	}

	return finalScore * finalDrawn
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
