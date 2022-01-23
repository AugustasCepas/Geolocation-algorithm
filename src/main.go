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

	"github.com/xitongsys/parquet-go-source/local"
	"github.com/xitongsys/parquet-go/parquet"
	"github.com/xitongsys/parquet-go/reader"
	"github.com/xitongsys/parquet-go/writer"
)

type GeolocationData struct {
	Start int    `parquet:"name=start, type=INT64"`
	End   int    `parquet:"name=end, type=INT64"`
	Code  string `parquet:"name=code, type=BYTE_ARRAY, convertedtype=UTF8"`
	City  string `parquet:"name=city, type=BYTE_ARRAY, convertedtype=UTF8"`
}

var geolocations []GeolocationData

var convertFileName string = "database"
var loadFileName string

func main() {
	argsPassed := os.Args[1:]

	if len(argsPassed) != 1 {
		csvToParquet()
		os.Exit(1)
	}
	loadFileName = argsPassed[0]

	fmt.Println("READY")
	readInput()
}

func readParquet() {
	fr, err := local.NewLocalFileReader(loadFileName)
	if err != nil {
		log.Println("Can't open file")
		return
	}

	pr, err := reader.NewParquetReader(fr, new(GeolocationData), 4)
	if err != nil {
		log.Println("Can't create parquet reader", err)
		return
	}

	rows := int(pr.GetNumRows())
	geolocations = make([]GeolocationData, rows)

	if err = pr.Read(&geolocations); err != nil {
		log.Println("Read error", err)
	}

	pr.ReadStop()
	fr.Close()
}

func findGeoLocation(ipSum int) int {
	r := -1 // not found
	start := 0
	end := len(geolocations) - 1
	for start <= end {
		mid := (start + end) / 2
		if geolocations[mid].Start <= ipSum {
			if ipSum <= geolocations[mid].End {
				r = mid // found
				fmt.Println(geolocations[r].Code + "," + geolocations[r].City)
				break
			}
		}

		if geolocations[mid].End < ipSum {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return r
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
		readParquet()
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

func csvToParquet() {
	var err error
	var start int = 0
	var end int = 0

	fw, err := local.NewLocalFileWriter(convertFileName + ".parquet")
	if err != nil {
		log.Println("Can't create local file", err)
		return
	}

	pw, err := writer.NewParquetWriter(fw, new(GeolocationData), 2)
	if err != nil {
		log.Println("Can't create parquet writer", err)
		return
	}

	pw.RowGroupSize = 128 * 1024 * 1024 //128M
	pw.CompressionType = parquet.CompressionCodec_SNAPPY

	csvFile, _ := os.Open(convertFileName + ".csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))

	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		start, _ = strconv.Atoi(line[0])
		end, _ = strconv.Atoi(line[1])

		shoe := GeolocationData{
			Start: start,
			End:   end,
			Code:  line[2],
			City:  line[5],
		}

		if err = pw.Write(shoe); err != nil {
			log.Println("Write error", err)
		}
	}

	if err = pw.WriteStop(); err != nil {
		log.Println("WriteStop error", err)
		return
	}

	log.Println("Write Finished")
	fw.Close()
}
