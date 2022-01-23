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

	"github.com/tushar2708/altcsv"
)

var startValues []int
var endValues []int
var codes []string
var cities []string

var loadFileName string
var cityIndex int = 3

func main() {
	argsPassed := os.Args[1:]

	if len(argsPassed) != 1 {
		shortCSV()
		os.Exit(1)
	}
	loadFileName = argsPassed[0]

	if argsPassed[0] == "database.csv" {
		cityIndex = 5
	}

	fmt.Println("READY")
	readInput()
}

func findGeoLocation(ipSum int) int {
	r := -1 // not found
	start := 0
	end := len(endValues) - 1
	for start <= end {
		mid := (start + end) / 2
		if startValues[mid] <= ipSum {
			if ipSum <= endValues[mid] {
				r = mid // found
				fmt.Println(codes[r] + "," + cities[r])
				break
			}
		}
		if endValues[mid] < ipSum {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return r
}

func readCSV() {
	csvIn, err := os.Open(loadFileName)

	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(csvIn)

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
		countryTag := rec[2]
		city := rec[cityIndex]

		startValues = append(startValues, start)
		endValues = append(endValues, finish)
		codes = append(codes, countryTag)
		cities = append(cities, city)
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
		os.Exit(1)
	}

	ipSum := 0
	if command == "LOAD" {
		readCSV()
		fmt.Println("OK")

	} else if command == "EXIT" {
		fmt.Println("OK")
		os.Exit(0)

	} else if command == "LOOKUP" {
		ip := trimLastChar(inputStrings[1])
		inputSlice, err := getIntsArray(ip)

		if err != nil {
			fmt.Println(err.Error())
		} else {
			ipSum = calculateIPSum(inputSlice)
		}

		go findGeoLocation(ipSum)
	}

	readInput()
}

func shortCSV() {
	csvIn, err := os.Open("database.csv")

	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(csvIn)

	f, err := os.Create("shorterdb.csv")

	if err != nil {
		log.Fatalln("failed to open file", err)
	}

	for {
		rec, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		w := altcsv.NewWriter(f)
		w.AllQuotes = true
		err = w.Write([]string{rec[0], rec[1], rec[2], rec[5]})
		if err != nil {
			log.Fatal(err)
		}
		w.Flush()
	}
	fmt.Println("Database file shortening finished")
}
