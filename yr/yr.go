package yr

import (
	"encoding/csv"
	"fmt"
	"github.com/B4KO/funtemps/conv"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func ConvertTemperatureAndAddStudent(inputFile, outputFile, studentName string) error {
	inFile, err := os.Open(inputFile)
	if err != nil {
		return fmt.Errorf("Error opening input file: %v", err)
	}
	defer inFile.Close()

	outFile, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("Error creating output file: %v", err)
	}
	defer outFile.Close()

	inputReader := csv.NewReader(inFile)
	inputReader.Comma = ';'
	outputWriter := csv.NewWriter(outFile)
	outputWriter.Comma = ';'
	defer outputWriter.Flush()

	lastRecord := []string{}
	for {
		record, err := inputReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("Error reading input file: %v", err)
		}

		if len(record) >= 4 && record[3] != "" && record[3] != "Lufttemperatur" {
			celsius, err := strconv.ParseFloat(record[3], 64)
			if err != nil {
				return fmt.Errorf("Error parsing Celsius temperature: %v", err)
			}
			record[3] = fmt.Sprintf("%.1f", conv.CelciusToFarenheit(celsius))
		}

		if len(lastRecord) > 0 {
			err = outputWriter.Write(lastRecord)
			if err != nil {
				return fmt.Errorf("Error writing to output file: %v", err)
			}
		}
		lastRecord = record
	}

	if strings.HasPrefix(lastRecord[0], "Data er gyldig per") {
		lastRecord = append(lastRecord[:len(lastRecord)-3], "endringer gjort av "+studentName)
	}

	err = outputWriter.Write(lastRecord)
	if err != nil {
		return fmt.Errorf("Error writing to output file: %v", err)
	}

	return nil
}

func CalculateAverageCelsius(filePath string) float64 {
	// Open the input CSV file
	inputFile, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening input file: %v", err)
	}
	defer inputFile.Close()

	inputReader := csv.NewReader(inputFile)
	inputReader.Comma = ';'

	var sumCelsius float64
	var count int


	headerProcessed := false
	for {
		record, err := inputReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error reading record: %v", err)
		}

		if len(record) < 4 {
			continue
		}

		if !headerProcessed {
			headerProcessed = true
			continue
		}

		if record[3] != "" {
			lufttemperatur, err := strconv.ParseFloat(record[3], 64)
			if err != nil {
				log.Fatalf("Error parsing Lufttemperatur: %v", err)
			}
			sumCelsius += lufttemperatur

			count++
		}
	}


	avgCelsius := sumCelsius / float64(count)

	return avgCelsius
}

func CalculateAverageFahrenheit(filePath string) float64 {
	// Open the input CSV file
	inputFile, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening input file: %v", err)
	}
	defer inputFile.Close()

	// Create a new CSV reader with a custom delimiter (semicolon)
	inputReader := csv.NewReader(inputFile)
	inputReader.Comma = ';'

	var sumFahrenheit float64
	var count int

	// Read and process records from the input CSV file
	headerProcessed := false
	for {
		record, err := inputReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error reading record: %v", err)
		}

		if len(record) < 4 {
			continue
		}

		if !headerProcessed {
			headerProcessed = true
			continue
		}

		// Calculate the sum of Celsius and Fahrenheit temperatures
		if record[3] != "" {
			lufttemperatur, err := strconv.ParseFloat(record[3], 64)
			if err != nil {
				log.Fatalf("Error parsing Lufttemperatur: %v", err)
			}
			sumFahrenheit += conv.CelciusToFarenheit(lufttemperatur)
			count++
		}
	}

	// Calculate the average temperatures
	avgFahrenheit := sumFahrenheit / float64(count)

	return avgFahrenheit
}
