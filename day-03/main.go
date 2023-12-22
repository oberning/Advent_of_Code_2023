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
	var numbersInLine [][]Number
	var charsInLine [][]Character
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

		numbers := searchNumbers(cachedLineNumber, currentLine, patternNumbers)
		numbersInLine = append(numbersInLine, numbers)
		characters := searchCharacters(cachedLineNumber, currentLine, patternSpecialChars)
		charsInLine = append(charsInLine, characters)

		// Check on same line
		for _, character := range characters {
			for _, number := range numbers {
				distanceFromStart := character.pos - number.startpos
				distanceFromEnd := character.pos - (number.endpos - 1)
				if distanceFromStart >= -1 && distanceFromEnd <= 1 {
					fmt.Println("To sum up inner line: ", number.number)
					sum += number.number
				}
			}
		}
		// Check on previous line - one time with numbers, one time with chars
		if cachedLineNumber > 0 {
			for _, character := range charsInLine[cachedLineNumber-1] {
				for _, number := range numbers {
					distanceFromStart := character.pos - number.startpos
					distanceFromEnd := character.pos - (number.endpos - 1)
					if distanceFromStart >= -1 && distanceFromEnd <= 1 {
						fmt.Println("To sum up: ", number.number)
						sum += number.number
					}
				}
			}
			for _, number := range numbersInLine[cachedLineNumber-1] {
				for _, character := range characters {
					distanceFromStart := character.pos - number.startpos
					distanceFromEnd := character.pos - (number.endpos - 1)
					if distanceFromStart >= -1 && distanceFromEnd <= 1 {
						fmt.Println("To sum up: ", number.number)
						sum += number.number
					}
				}
			}
		}
		cachedLineNumber++
	}
	fmt.Println("The sum is: ", sum)
}

type Number struct {
	number   int
	startpos int
	endpos   int
}

type Character struct {
	char string
	pos  int
}

func searchNumbers(linenumber int, currentLine string, pattern *regexp.Regexp) []Number {
	var numbers []Number
	posNumbers := pattern.FindAllIndex([]byte(currentLine), -1)
	for _, s := range posNumbers {
		if len(s) == 0 {
			break
		}
		number, _ := strconv.Atoi(currentLine[s[0]:s[1]])
		numbers = append(numbers, Number{
			number:   number,
			startpos: s[0],
			endpos:   s[1],
		})
	}
	return numbers
}

func searchCharacters(linenumber int, currentLine string, pattern *regexp.Regexp) []Character {
	var characters []Character
	posSpecialChars := pattern.FindAllIndex([]byte(currentLine), -1)
	for _, s := range posSpecialChars {
		if len(s) == 0 {
			break
		}
		characters = append(characters, Character{
			char: currentLine[s[0]:s[1]],
			pos:  s[0],
		})
	}
	return characters
}
