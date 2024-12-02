package day01

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
		firstNumber, _ := strconv.Atoi(strings.Split(scanner.Text(), "   ")[0])
		secondNumber, _ := strconv.Atoi(strings.Split(scanner.Text(), "   ")[1])

		data = append(data, []int{firstNumber, secondNumber})
	}

	return data
}

func sortData(data [][]int) {
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data) - 1; j++ {
			for k := 0; k < 2; k++ {
				if data[j][k] > data[j + 1][k] {
					temp := data[j + 1][k]

					data[j + 1][k] = data[j][k]
					data[j][k] = temp
				}
			}
		}
	}
}

func Part01() int {
	sum := 0

	data := parseDataFromFile("./inputs/day01.txt")

	sortData(data)

	for i := 0; i < len(data); i++ {
		if data[i][0] > data[i][1] {
			sum += data[i][0] - data[i][1]
		} else {
			sum += data[i][1] - data[i][0]
		}
	}

	return sum
}

func Part02() int {
	sum := 0

	data := parseDataFromFile("./inputs/day01.txt")

	for i := 0; i < len(data); i++ {
		multiplier := 0

		for j := 0; j < len(data); j++ {
			if data[i][0] == data[j][1] {
				multiplier++
			}
		}

		sum += data[i][0] * multiplier
	}

	return sum
}
