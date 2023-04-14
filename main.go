package main

import (
	"bufio"
	"fmt"
	"github.com/uia-worker/minyr/yr"
	"log"
	"os"
	"strings"
)

func GetInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	//Format input
	input = strings.TrimSpace(input)
	input = strings.ToUpper(input)
	return input
}

func main() {

	inputFileName := "kjevik-temp-celsius-20220318-20230318.csv"
	outputFileName := "kjevik-temp-fahr-20220318-20230318.csv"
	studentName := "GABRIEL MOLINSKI"

	for {

		fmt.Println("Write a option")
		fmt.Println("---------------------")

		input := GetInput()

		switch input {

		case "Q":
			fmt.Println("QUITTING")
			os.Exit(0)

		case "EXIT":
			fmt.Println("EXITING")
			os.Exit(0)

		case "CONVERT":

			if _, err := os.Stat(outputFileName); os.IsNotExist(err) {
				err := yr.ConvertTemperatureAndAddStudent(inputFileName, outputFileName, studentName)
				if err != nil {
					log.Fatalf("Error converting temperature and replacing last line: %v", err)
				}
				os.Exit(0)

			} else {
				fmt.Println("The file already exists. Want to create it anyways, Y/N")
				fmt.Println("---------------------------------------------------------------")
				input = GetInput()

				switch input {
				case "Y":
					err := yr.ConvertTemperatureAndAddStudent(inputFileName, outputFileName, studentName)
					if err != nil {
						log.Fatalf("Error converting temperature and replacing last line: %v", err)
					}
					os.Exit(0)
				default:
					break
				}
			}

			break

		case "AVERAGE":

			filePath := "kjevik-temp-celsius-20220318-20230318.csv"

			fmt.Println("What temperature unit would you like to get the average in, C/F")
			fmt.Println("---------------------------------------------------------------")
			input = GetInput()

			switch input {
			case "C":
				avgCelsius := yr.CalculateAverageCelsius(filePath)
				fmt.Printf("Average temperature in Celsius: %.2f\n", avgCelsius)
				os.Exit(0)

			case "F":
				avgFahrenheit := yr.CalculateAverageFahrenheit(filePath)
				fmt.Printf("Average temperature in Fahrenheit: %.2f\n", avgFahrenheit)
				os.Exit(0)

			default:
				break
			}

		default:
			fmt.Println("Please write a valid option")
			fmt.Println("-----------------------------------")
			break

		}
	}
}
