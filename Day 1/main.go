package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {

	filePath := "/Users/cianlaguesma/codes/adventofcode2022/Day 1/mainInput.txt"
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
	testArr := []int{}
	for _, line := range fileLines {
		if line == "" {
			personNum += 1
			personSum = 0
		}
		num, _ := strconv.Atoi(line)
		personSum += num
		testArr = append(testArr, personSum)
	}

	sort.Slice(testArr, func(i, j int) bool {
		return testArr[i] < testArr[j]
	})
	top3 := 0
	top1 := testArr[len(testArr)-1]
	for i := len(testArr) - 1; i >= len(testArr)-3; i-- {
		top3 += testArr[i]
	}
	fmt.Println(top1)
	fmt.Println(top3)

}
