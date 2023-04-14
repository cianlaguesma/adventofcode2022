package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	filePath := "/Users/cianlaguesma/codes/adventofcode2022/Day 3/input.txt"
	readFile, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string
	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}
	readFile.Close()

	sumPrio := 0
	for i := 0; i < len(fileLines); i += 3 {
		firstLine := fileLines[i]
		secondLine := fileLines[i+1]
		thirdLine := fileLines[i+2]

		firstLineLetters := []string{}
		secondLineLetters := []string{}
		thirdLineLetters := []string{}

		for _, letter := range firstLine {
			convertedLetter := string(letter)
			firstLineLetters = append(firstLineLetters, convertedLetter)
		}
		for _, letter := range secondLine {
			convertedLetter := string(letter)
			secondLineLetters = append(secondLineLetters, convertedLetter)
		}
		for _, letter := range thirdLine {
			convertedLetter := string(letter)
			thirdLineLetters = append(thirdLineLetters, convertedLetter)
		}
		letterCheckerArr := []string{}
		sameLetters := []string{}
	out:
		for _, letter := range firstLineLetters {
			skip := false
			for _, letterChecker := range letterCheckerArr {
				if letterChecker == letter {
					skip = true
				}
			}
			for _, letter2 := range secondLineLetters {
				if (letter == letter2) && !skip {

					for _, letter3 := range thirdLineLetters {
						if (letter == letter3) && !skip {
							letterCheckerArr = append(letterCheckerArr, letter)
							sameLetters = append(sameLetters, letter)
							break out
						}

					}
				}

			}

		}
		for _, letter := range sameLetters {
			prioNumber, _ := utf8.DecodeRuneInString(letter)
			if unicode.IsLower(prioNumber) {

				prioNumber = prioNumber - 96
			} else {

				prioNumber = prioNumber - 38
			}
			sumPrio += int(prioNumber)

		}

	}
	fmt.Println(sumPrio)
}
