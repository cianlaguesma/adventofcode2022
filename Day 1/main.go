package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	// reader := bufio.NewReader(os.Stdin)
	filePath := "/Users/cianlaguesma/codes/adventofcode2022/Day 1/input.txt"
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

	personNum := 1
	personSum := 0
	outArr := make(map[int]int)
	for _, line := range fileLines {
		if line == "" {
			fmt.Println("LOL")
			personNum += 1
			personSum = 0
		}
		num, _ := strconv.Atoi(line)
		personSum += num
		outArr[personNum] = personSum
	}
	fmt.Println(outArr)
	for maxnumber := range outArr {
		fmt.Println(outArr)
	}
	fmt.Println(fileLines)

}
