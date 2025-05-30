package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Input your text to count the words: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	fmt.Println(countWords(text))
}

func countWords(text string) int {
	elements := strings.Fields(text)
	return len(elements)
}
