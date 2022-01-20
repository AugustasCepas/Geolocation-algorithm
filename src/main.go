package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

type GeolocationData struct {
	CountryTag string
	City       string
}

var database []GeolocationData
var finishValues []int
var startValues []int

func main() {
	fmt.Println("READY")
	readInput()
}

func findGeoLocation(ipSum int) int {
	r := -1 // not found
	start := 0
	end := len(finishValues) - 1
	for start <= end {
		mid := (start + end) / 2
		if startValues[mid] <= ipSum && ipSum <= finishValues[mid] {
			r = mid // found
			fmt.Println(database[r].CountryTag + "," + database[r].City)
			break
		} else if finishValues[mid] < ipSum {
			start = mid + 1
		} else if finishValues[mid] > ipSum {
			end = mid - 1
		}
	}
	return r
}

func readCSV() {
	csvIn, err := os.Open("./database.csv")

	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(csvIn)

	var parsedLine GeolocationData

	for {
		rec, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		start, _ := strconv.Atoi(rec[0])
		finish, _ := strconv.Atoi(rec[1])
		parsedLine.CountryTag = rec[2]
		parsedLine.City = rec[5]

		database = append(database, parsedLine)

		startValues = append(startValues, start)
		finishValues = append(finishValues, finish)
	}
}

func calculateIPSum(input []int) int {

	result := 0
	for i := 0; i < 4; i++ {
		power := 3 - i
		sum := math.Pow(256, float64(power))
		result += input[i] * int(sum)
	}
	return result
}

func getIntsArray(input string) ([]int, error) {
	inputStrings := strings.Split(input, ".")

	if len(inputStrings) != 4 {
		return nil, errors.New("input: invalid value entered")
	}

	inputSlice := make([]int, len(inputStrings))

	for i, v := range inputStrings {

		ipPart, err := strconv.Atoi(v)

		if err != nil {
			return nil, errors.New("input: ip contains invalid value")
		}
		inputSlice[i] = ipPart
	}
	return inputSlice, nil
}

func trimLastChar(s string) string {
	r, size := utf8.DecodeLastRuneInString(s)
	if r == utf8.RuneError && (size == 0 || size == 1) {
		size = 0
	}
	return s[:len(s)-size]
}

func readInput() {
	var command string

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSuffix(text, "\n")

	inputStrings := strings.Split(text, " ")
	inputLength := len(inputStrings)

	if inputLength == 1 {
		command = trimLastChar(inputStrings[0])
	} else if inputLength == 2 {
		command = inputStrings[0]
	} else {
		exit()
	}

	ipSum := 0
	if command == "LOAD" {
		readCSV()
		fmt.Println("OK")

	} else if command == "EXIT" {
		exit()

	} else if command == "LOOKUP" {
		ip := trimLastChar(inputStrings[1])
		inputSlice, err := getIntsArray(ip)

		if err != nil {
			fmt.Println(err.Error())
		} else {
			ipSum = calculateIPSum(inputSlice)
		}

		findGeoLocation(ipSum)
	}

	readInput()
}

func exit() {
	fmt.Println("OK")
	os.Exit(0)
}
