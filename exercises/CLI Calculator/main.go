package main

import (
	"fmt"
	"strings"
)

func main() {
	for {
		var x, y float64
		var option string

		fmt.Println("Input your type of operation (+, -, *, /) or type 'exit' to disconnect: ")
		fmt.Scan(&option)

		if strings.ToLower(option) == "exit" {
			fmt.Println("o7")
			break
		}

		fmt.Println("Input your first number: ")
		fmt.Scan(&x)

		fmt.Println("Input your second number: ")
		fmt.Scan(&y)

		switch option {
		case "+":
			fmt.Printf("%.2f %s %.2f = %.2f \n", x, option, y, x+y)
		case "-":
			fmt.Printf("%.2f %s %.2f = %.2f \n", x, option, y, x-y)
		case "*":
			fmt.Printf("%.2f %s %.2f = %.2f \n", x, option, y, x*y)
		case "/":
			if y != 0 {
				fmt.Printf("%.2f %s %.2f = %.2f \n", x, option, y, x/y)
			} else {
				fmt.Printf("%.2f %s %.2f = %s \n", x, option, y, "∞")
			}
		default:
			fmt.Println("Invalid option.")
		}
	}
}
