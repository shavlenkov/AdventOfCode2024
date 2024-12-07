package day03

import (
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func parseDataFromFile(filePath string) string {
	data, err := ioutil.ReadFile(filePath)

	if err != nil {
		log.Fatal(err)
	}

	return string(data)
}

func calcSumOfAllMultiplications(data string) int {
	var sum int

	matches := regexp.MustCompile(`mul\((\d{1,}),(\d{1,})\)`).FindAllStringSubmatch(data, -1)

	for _, match := range matches {
		firstNumber, _ := strconv.Atoi(match[1])
		secondNumber, _ := strconv.Atoi(match[2])

		sum += firstNumber * secondNumber
	}

	return sum
}

func Part01() int {
	return calcSumOfAllMultiplications(parseDataFromFile("./inputs/day03.txt"))
}

func Part02() int {
	var processedData string

	segments := strings.Split(parseDataFromFile("./inputs/day03.txt"), "do()")

	for i := 0; i < len(segments); i++ {
		if strings.Contains(segments[i], "don't()") {
			processedData += strings.Split(segments[i], "don't()")[0]
		} else {
			processedData += segments[i]
		}
	}

	return calcSumOfAllMultiplications(processedData)
}
