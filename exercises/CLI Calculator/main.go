package main

import (
	"fmt"
)

func main() {
	var x, y float64
	var option string

	fmt.Println("Input your first number: ")
	fmt.Scan(&x)

	fmt.Println("Input your second number: ")
	fmt.Scan(&y)

	fmt.Println("Input your type of operation (+, -, *, /): ")
	fmt.Scan(&option)

	switch option {
	case "+":
		fmt.Printf("%.2f %s %.2f = %.2f", x, option, y, x+y)
	case "-":
		fmt.Printf("%.2f %s %.2f = %.2f", x, option, y, x-y)
	case "*":
		fmt.Printf("%.2f %s %.2f = %.2f", x, option, y, x*y)
	case "/":
		if y != 0 {
			fmt.Printf("%.2f %s %.2f = %.2f", x, option, y, x/y)
		} else {
			fmt.Printf("%.2f %s %.2f = %s", x, option, y, "∞")
		}
	default:
		fmt.Println("Invalid option.")
	}
}
