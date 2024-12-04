package day02

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseDataFromFile(filePath string) [][]int {
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	var data [][]int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var allNumbers []int

		for _, value := range strings.Split(scanner.Text(), " ") {
			oneNumber, _ := strconv.Atoi(value)
			allNumbers = append(allNumbers, oneNumber)
		}

		data = append(data, allNumbers)
	}

	return data
}

func isSafe(data []int) bool {
	counter := 0
	directionFlag := 0

	for i := 0; i < len(data) - 1; i++ {
		if data[i + 1] - data[i] <= 3 && data[i + 1] - data[i] >= 1 && (directionFlag == 0 || directionFlag == 1) {
			counter++
			directionFlag = 1
		}

		if data[i] - data[i + 1] <= 3 && data[i] - data[i + 1] >= 1 && (directionFlag == 0 || directionFlag == -1) {
			counter++
			directionFlag = -1
		}
	}

	if counter == len(data) - 1 {
		return true
	}

	return false
}

func Part01() int {
	counter := 0

	data := parseDataFromFile("./inputs/day02.txt")

	for i := 0; i < len(data); i++ {
		if isSafe(data[i]) {
			counter++
		}
	}

	return counter
}

func Part02() int {
	counter := 0

	data := parseDataFromFile("./inputs/day02.txt")

	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			var newData []int

			for k := 0; k < len(data[i]); k++ {
				if j != k {
					newData = append(newData, data[i][k])
				}
			}

			if isSafe(newData) {
				counter++
				break
			}
		}
	}

	return counter
}
