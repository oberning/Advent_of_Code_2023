package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	cachedLineNumber := 0
	var numbersInLine [][]Item
	var charsInLine [][]Item
	sum := 0
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fileScanner := bufio.NewScanner(file)
	patternNumbers := regexp.MustCompile("[0-9]+")
	patternSpecialChars := regexp.MustCompile(`[^0-9,\.\n\r]+`)

	for fileScanner.Scan() {
		currentLine := fileScanner.Text()
		fmt.Print(cachedLineNumber, currentLine)

		numbers := searchPattern(currentLine, patternNumbers)
		numbersInLine = append(numbersInLine, numbers)
		characters := searchPattern(currentLine, patternSpecialChars)
		charsInLine = append(charsInLine, characters)

		// Check on same line
		for _, character := range characters {
			for _, number := range numbers {
				distanceFromStart := character.startpos - number.startpos
				distanceFromEnd := character.startpos - (number.endpos - 1)
				if distanceFromStart >= -1 && distanceFromEnd <= 1 {
					fmt.Println("To sum up inner line: ", number.item)
					numberInt, _ := strconv.Atoi(number.item) // Regex ensure that it is a number
					sum += numberInt
				}
			}
		}
		// Check on the previous line (first with searching for characters and then numbers)
		if cachedLineNumber > 0 {
			for _, character := range charsInLine[cachedLineNumber-1] {
				for _, number := range numbers {
					distanceFromStart := character.startpos - number.startpos
					distanceFromEnd := character.startpos - (number.endpos - 1)
					if distanceFromStart >= -1 && distanceFromEnd <= 1 {
						fmt.Println("To sum up: ", number.item)
						numberInt, _ := strconv.Atoi(number.item) // Regex ensure that it is a number
						sum += numberInt
					}
				}
			}
			for _, number := range numbersInLine[cachedLineNumber-1] {
				for _, character := range characters {
					distanceFromStart := character.startpos - number.startpos
					distanceFromEnd := character.startpos - (number.endpos - 1)
					if distanceFromStart >= -1 && distanceFromEnd <= 1 {
						fmt.Println("To sum up: ", number.item)
						numberInt, _ := strconv.Atoi(number.item) // Regex ensure that it is a number
						sum += numberInt
					}
				}
			}
		}
		cachedLineNumber++
	}
	fmt.Println("The sum is: ", sum)
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
