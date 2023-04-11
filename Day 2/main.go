package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	filePath := "/Users/cianlaguesma/codes/adventofcode2022/Day 2/input.txt"
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

	fmt.Println(fileLines)
	score := 0
	secondScore := 0
	for i := 0; i < len(fileLines); i++ {
		pairArr := strings.Fields(fileLines[i])
		if pairArr[1] == "X" {
			score += 1
		} else if pairArr[1] == "Y" {
			score += 2
		} else {
			score += 3
		}

		if pairArr[0] == "A" && pairArr[1] == "X" {
			score += 3
		} else if pairArr[0] == "A" && pairArr[1] == "Y" {
			score += 6
		} else if pairArr[0] == "A" && pairArr[1] == "Z" {
			score += 0
		}

		if pairArr[0] == "B" && pairArr[1] == "Y" {
			score += 3
		} else if pairArr[0] == "B" && pairArr[1] == "Z" {
			score += 6
		} else if pairArr[0] == "B" && pairArr[1] == "X" {
			score += 0
		}

		if pairArr[0] == "C" && pairArr[1] == "Z" {
			score += 3
		} else if pairArr[0] == "C" && pairArr[1] == "X" {
			score += 6
		} else if pairArr[0] == "C" && pairArr[1] == "Y" {
			score += 0
		}

		if pairArr[1] == "X" && pairArr[0] == "A" {
			secondScore += 0
			secondScore += 3
		} else if pairArr[1] == "X" && pairArr[0] == "B" {
			secondScore += 0
			secondScore += 1

		} else if pairArr[1] == "X" && pairArr[0] == "C" {
			secondScore += 0
			secondScore += 2
		}

		if pairArr[1] == "Y" && pairArr[0] == "A" {
			secondScore += 3
			secondScore += 1
		} else if pairArr[1] == "Y" && pairArr[0] == "B" {
			secondScore += 3
			secondScore += 2

		} else if pairArr[1] == "Y" && pairArr[0] == "C" {
			secondScore += 3
			secondScore += 3
		}

		if pairArr[1] == "Z" && pairArr[0] == "A" {
			secondScore += 6
			secondScore += 2
		} else if pairArr[1] == "Z" && pairArr[0] == "B" {
			secondScore += 6
			secondScore += 3

		} else if pairArr[1] == "Z" && pairArr[0] == "C" {
			secondScore += 6
			secondScore += 1
		}
		fmt.Println(pairArr)
		fmt.Println(score)
		fmt.Println(secondScore)

	}
}
