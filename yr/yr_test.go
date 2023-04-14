package yr

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"testing"
)

// Helper functions
func lineInFile(fileName, targetLine string) (bool, error) {
	file, err := os.Open("../" + fileName)
	if err != nil {
		return false, fmt.Errorf("Error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	_ = ""
	for scanner.Scan() {
		line := scanner.Text()

		if line == targetLine {
			return true, nil
		}

	}

	if err := scanner.Err(); err != nil {
		return false, fmt.Errorf("Error scanning file: %v", err)
	}

	return false, nil
}

func countLines(filePath string) int {
	file, err := os.Open("../" + filePath)
	if err != nil {
		os.Exit(0)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}

	if err := scanner.Err(); err != nil {
		os.Exit(0)
	}

	return lineCount
}

func TestTotalLines(t *testing.T) {
	type test struct {
		input string
		want  int
	}

	tests := []test{
		{input: "kjevik-temp-celsius-20220318-20230318.csv", want: 16756},
		{input: "kjevik-temp-fahr-20220318-20230318.csv", want: 16756},
	}

	for _, tc := range tests {
		got := countLines(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func TestConversion(t *testing.T) {
	type test struct {
		input    string
		filename string
		want     bool
	}

	tests := []test{
		{input: "Kjevik;SN39040;18.03.2022 01:50;42.8", filename: "kjevik-temp-fahr-20220318-20230318.csv", want: true},
		{input: "Kjevik;SN39040;07.03.2023 18:20;32.0", filename: "kjevik-temp-fahr-20220318-20230318.csv", want: true},
		{input: "Kjevik;SN39040;08.03.2023 02:20;12.2", filename: "kjevik-temp-fahr-20220318-20230318.csv", want: true},
	}

	for _, tc := range tests {
		got, _ := lineInFile(tc.filename, tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func TestMark(t *testing.T) {
	type test struct {
		input    string
		filename string
		want     bool
	}

	tests := []test{
		{input: "Data er gyldig per 18.03.2023 (CC BY 4.0), Meteorologisk institutt (MET);endringer gjort av GABRIEL MOLINSKI", filename: "kjevik-temp-fahr-20220318-20230318.csv", want: true},
	}

	for _, tc := range tests {
		got, _ := lineInFile(tc.filename, tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func TestAverage(t *testing.T) {
	type test struct {
		input string
		unit  string
		want  string
	}

	tests := []test{
		{input: "../kjevik-temp-celsius-20220318-20230318.csv", unit: "C", want: "8.56"},
		{input: "../kjevik-temp-celsius-20220318-20230318.csv", unit: "F", want: "47.41"},
	}

	for _, tc := range tests {

		var got float64
		var gotString string

		if tc.unit == "C" {
			got = CalculateAverageCelsius(tc.input)
			gotString = fmt.Sprintf("%.2f", got)
		}

		if tc.unit == "F" {
			got = CalculateAverageFahrenheit(tc.input)
			gotString = fmt.Sprintf("%.2f", got)
		}

		if !reflect.DeepEqual(tc.want, gotString) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}
