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
)

type GeolocationData struct {
	Start      int
	Finish     int
	CountryTag string
	City       string
}

var database []GeolocationData

func main() {
	fmt.Println("> READY")
	ReadInput()
}

func FindGeoLocation(ipSum int) {

	for _, v := range database {
		if v.Start <= ipSum {
			if ipSum <= v.Finish {
				fmt.Println(v.CountryTag + "," + v.City)
				break
			}
		}
	}
}

func ReadCSV() {
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

func CalculateIPSum(input []int) int {

	result := 0
	for i := 0; i < 4; i++ {
		power := 3 - i
		sum := math.Pow(256, float64(power))
		result += input[i] * int(sum)
	}
	return result
}

func GetIntsArray(input string) ([]int, error) {
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

func ReadInput() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	inputStrings := strings.Split(text, " ")

	ipSum := 0
	if inputStrings[0] == "LOAD" {
		ReadCSV()
		fmt.Println("< OK")

	} else if inputStrings[0] == "EXIT" {
		fmt.Println("< OK")
		os.Exit(0)

	} else if inputStrings[0] == "LOOKUP" {
		inputSlice, err := GetIntsArray(inputStrings[1])

		if err != nil {
			fmt.Println(err.Error())
		} else {
			ipSum = CalculateIPSum(inputSlice)
		}

		FindGeoLocation(ipSum)
	}

	ReadInput()
}
