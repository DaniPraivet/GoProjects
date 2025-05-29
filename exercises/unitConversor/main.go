package main

import (
	"fmt"
)

func main() {
	fmt.Println(
		celsiusToFahrenheit(38),
		"\n",
		metersToFeet(50),
	)
}

func celsiusToFahrenheit(celsius float32) (fahrenheit float32) {
	return (celsius * 1.8) + 32
}

func metersToFeet(meters float32) (feet float32) {
	return meters * 3.28084
}
