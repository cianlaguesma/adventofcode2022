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
	for _, lines := range fileLines {
		compartmentSize := len(lines)
		firstHalf := []string{}
		secondHalf := []string{}
		for count, letter := range lines {
			convertedLetter := string(letter)
			if count >= compartmentSize/2 {
				secondHalf = append(secondHalf, convertedLetter)
			} else {
				firstHalf = append(firstHalf, convertedLetter)
			}
		}
		letterCheckedArr := []string{}
		for _, letter := range firstHalf {
			skip := false
			for _, letterChecker := range letterCheckedArr {
				if letterChecker == letter {
					//fmt.Println(letter)
					skip = true
				}
			}
			for _, letter2 := range secondHalf {
				if (letter == letter2) && !skip {
					letterCheckedArr = append(letterCheckedArr, letter)
					//fmt.Println(letter, letter2)

					prioNumber, _ := utf8.DecodeRuneInString(letter)
					if unicode.IsLower(prioNumber) {

						prioNumber = prioNumber - 96
					} else {

						prioNumber = prioNumber - 38
					}
					sumPrio += int(prioNumber)
					break
				}
			}
		}
		fmt.Println(letterCheckedArr)

	}
	fmt.Println(sumPrio)
}
