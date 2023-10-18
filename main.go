// This simple program will convert from Fahrenheit to Celcius or vice versa

package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("Welcome to Fahrenheit/Celcius temperature converter.")
	fmt.Println("Enter 'F' to convert a Fahrenheit temperature or 'C' to convert a Celcius temperature: ")

	//user enters F or C. Make sure the input here is one of those two values
	var inputType string
	for {
		fmt.Scan(&inputType)

		inputType = strings.ToUpper(inputType)

		if len(inputType) == 1 && (inputType == "F" || inputType == "C") {
			break
		} else {
			fmt.Println("Error: please enter either 'F' or 'C': ")
		}
	}

	fmt.Println("Now please enter the temperature: ")

	//user enters a decimal number. Make sure input is correct
	var inputValue float64
	for {
		_, err := fmt.Scan(&inputValue) //Scan returns num of items and error. We dont care about first one so we store it in _
		fmt.Scanln()
		if err == nil {
			break
		} else {
			fmt.Println("Error: Please enter a valid number.")
		}
	}

	//Display proper conversion
	if inputType == "F" {
		fmt.Print(inputValue)
		fmt.Print(" F is equal to ")
		fmt.Print(fahToCel(inputValue))
		fmt.Print(" C.")
	} else {
		fmt.Print(inputValue)
		fmt.Print(" C is equal to ")
		fmt.Print(celTofah(inputValue))
		fmt.Print(" F.")
	}

}

func fahToCel(f float64) float64 {
	return (f - 32.00) * (5.00 / 9.00)
}

func celTofah(c float64) float64 {
	return c*(9.00/5.00) + 32.00
}
