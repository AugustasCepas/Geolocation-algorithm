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
	Start      int
	Finish     int
	CountryTag string
	City       string
}

var database []GeolocationData

func main() {
	fmt.Println("READY")
	readInput()
}

func findGeoLocation(ipSum int) {
	for _, v := range database {
		if v.Start <= ipSum {
			if ipSum <= v.Finish {
				fmt.Println(v.CountryTag + "," + v.City)
				break
			}
		}
	}
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

		parsedLine.Start, _ = strconv.Atoi(rec[0])
		parsedLine.Finish, _ = strconv.Atoi(rec[1])
		parsedLine.CountryTag = rec[2]
		parsedLine.City = rec[5]

		database = append(database, parsedLine)
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
