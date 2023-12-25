package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	var lineNumber int
	var sum *int = new(int)
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fileScanner := bufio.NewScanner(file)
	patternNumbers := regexp.MustCompile("[0-9]+")
	patternAsterisk := regexp.MustCompile(`\*`)
	var numbersInLine [142][]Item
	var asterisksInLine [142][]Item

	for fileScanner.Scan() {
		currentLine := fileScanner.Text()
		numbersInLine[lineNumber] = searchPattern(currentLine, patternNumbers)
		asterisksInLine[lineNumber] = searchPattern(currentLine, patternAsterisk)
		result := identifyValidMultiplication(lineNumber, sameLine, &numbersInLine, &asterisksInLine[lineNumber])
		add(sum, result)
		stdout(lineNumber, "(sameLine): ", result, *sum)
		result = identifyValidMultiplication(lineNumber, numbersAbove, &numbersInLine, &asterisksInLine[lineNumber])
		add(sum, result)
		stdout(lineNumber, "(numberAbove): ", result, *sum)
		result = identifyValidMultiplication(lineNumber, numbersOnTwoLines, &numbersInLine,
			&asterisksInLine[lineNumber])
		add(sum, result)
		stdout(lineNumber, "(numberOnTwoLines, * on currentLine): ", result, *sum)
		if lineNumber > 0 {
			result = identifyValidMultiplication(lineNumber, numbersBelow, &numbersInLine,
				&asterisksInLine[lineNumber-1])
			add(sum, result)
			stdout(lineNumber, "(numberBelow): ", result, *sum)
		}
		if lineNumber > 0 {
			result = identifyValidMultiplication(lineNumber, numbersOnTwoLines, &numbersInLine,
				&asterisksInLine[lineNumber-1])
			add(sum, result)
			stdout(lineNumber, "(numberOnTwoLines, * on previous line): ", result, *sum)
		}
		if lineNumber > 1 {
			result = identifyValidMultiplication(lineNumber, numbersAboveAndBelow, &numbersInLine,
				&asterisksInLine[lineNumber-1])
			add(sum, result)
			stdout(lineNumber, "(numbersAboveAndBelow): ", result, *sum)
		}
		lineNumber++
		// if lineNumber > 106 {
		// 	break
		// }
	}
	fmt.Println("Sum is: ", *sum)
}

type Item struct {
	item     string
	startpos int
	endpos   int
}

func searchPattern(currentLine string, pattern *regexp.Regexp) []Item {
	var items []Item
	posItems := pattern.FindAllIndex([]byte(currentLine), -1)
	for _, s := range posItems {
		if len(s) == 0 {
			break
		}
		items = append(items, Item{
			item:     currentLine[s[0]:s[1]],
			startpos: s[0],
			endpos:   s[1],
		})
	}
	return items
}

func (asterisk *Item) numberNearBy(number Item) int {
	var numberInt int
	distanceFromStart := asterisk.startpos - number.startpos
	distanceFromEnd := asterisk.startpos - number.endpos + 1
	if distanceFromStart >= -1 && distanceFromEnd <= 1 {
		numberInt, _ = strconv.Atoi(number.item) // Very sure that it is a number due to regex
	}
	return numberInt
}

func identifyValidMultiplication(lineNumber int, call func(int, *Item, *[142][]Item) []int,
	numbersInLines *[142][]Item, asterisks *[]Item) []int {

	var result []int
	var numbers []int
	for _, asterisk := range *asterisks {
		numbers = call(lineNumber, &asterisk, numbersInLines)
		if len(numbers) == 2 {
			result = numbers
		}
		numbers = numbers[:0]
	}
	return result
}

func sameLine(lineNumber int, asterisk *Item, numbersInLines *[142][]Item) []int {
	var numbers []int
	for _, number := range numbersInLines[lineNumber] {
		numberInt := asterisk.numberNearBy(number)
		if numberInt != 0 {
			numbers = append(numbers, numberInt)
		}
	}
	return numbers
}

func numbersAbove(lineNumber int, asterisk *Item, numbersInLines *[142][]Item) []int {
	var numbers []int
	if lineNumber == 0 {
		return numbers
	}
	for _, number := range numbersInLines[lineNumber-1] {
		numberInt := asterisk.numberNearBy(number)
		if numberInt != 0 {
			numbers = append(numbers, numberInt)
		}
	}
	return numbers
}

func numbersBelow(lineNumber int, asterisk *Item, numbersInLines *[142][]Item) []int {
	var numbers []int
	if lineNumber == 0 {
		return numbers
	}
	for _, number := range numbersInLines[lineNumber] {
		numberInt := asterisk.numberNearBy(number)
		if numberInt != 0 {
			numbers = append(numbers, numberInt)
		}
	}
	return numbers
}

func numbersOnTwoLines(lineNumber int, asterisk *Item, numbersInLines *[142][]Item) []int {
	var numbers []int
	if lineNumber == 0 {
		return numbers
	}
	for _, number := range numbersInLines[lineNumber] {
		numberInt := asterisk.numberNearBy(number)
		if numberInt != 0 {
			numbers = append(numbers, numberInt)
			break
		}
	}
	for _, number := range numbersInLines[lineNumber-1] {
		numberInt := asterisk.numberNearBy(number)
		if numberInt != 0 {
			numbers = append(numbers, numberInt)
			break
		}
	}
	return numbers
}

func numbersAboveAndBelow(lineNumber int, asterisk *Item, numbersInLines *[142][]Item) []int {
	var numbers []int
	if lineNumber == 0 {
		return numbers
	}
	for _, number := range numbersInLines[lineNumber-2] {
		numberInt := asterisk.numberNearBy(number)
		if numberInt != 0 {
			numbers = append(numbers, numberInt)
			break
		}
	}
	for _, number := range numbersInLines[lineNumber] {
		numberInt := asterisk.numberNearBy(number)
		if numberInt != 0 {
			numbers = append(numbers, numberInt)
			break
		}
	}
	return numbers
}

func stdout(lineNumber int, text string, result []int, sum int) {
	if len(result) == 0 {
		return
	}
fmt.Println("Line:", lineNumber+1, text, result, "=", result[0] * result[1], "-> ", sum)
}

func add(sum *int, numbers []int) {
	if len(numbers) == 0 {
		return
	}
	*sum += numbers[0] * numbers[1]
}
