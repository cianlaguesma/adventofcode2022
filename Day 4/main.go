package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	filePath := "/Users/cianlaguesma/codes/adventofcode2022/Day 4/input.txt"
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
	ans := 0
	for i := 0; i < len(fileLines); i += 1 {
		line := []rune(fileLines[i])
		firstMin := ""
		firstMax := ""
		secondMin := ""
		secondMax := ""

		switcher := 0

		//0 = first first, 1 = second first, 2 = first second, 3 = second second
		for _, char := range line {
			newChar := string(char)
			if newChar == "-" {
				switcher++
				continue
			} else if newChar == "," {
				switcher++
				continue
			}
			if switcher == 0 {
				firstMin += newChar
			} else if switcher == 1 {
				firstMax += newChar
			} else if switcher == 2 {
				secondMin += newChar
			} else {
				secondMax += newChar
			}
		}
		firstMinInt, _ := strconv.Atoi(firstMin)
		firstMaxInt, _ := strconv.Atoi(firstMax)
		secondMinInt, _ := strconv.Atoi(secondMin)
		secondMaxInt, _ := strconv.Atoi(secondMax)

		// if firstMinInt > secondMinInt || firstMaxInt < secondMaxInt {
		// } else {
		// 	fmt.Println("Yes to: ")
		// 	fmt.Println(firstMin, firstMax, secondMin, secondMax)
		// 	ans++
		// 	continue
		// }

		// if firstMinInt < secondMinInt || firstMaxInt > secondMaxInt {
		// } else {

		// 	fmt.Println("Yes to: ")
		// 	fmt.Println(firstMin, firstMax, secondMin, secondMax)
		// 	ans++
		// 	continue
		// }
		// IS SECOND PAIR IN FIRST PAIR?
		//2-4, 5-6
		//3-5, 1-2
		if firstMaxInt < secondMinInt {

		} else if firstMinInt > secondMaxInt {

		} else {
			ans++
		}

	}
	fmt.Println(ans)
}
