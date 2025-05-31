package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	sc := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the path of the file to rename: ")
	oldPath, _ := sc.ReadString('\n')
	oldPath = strings.TrimSpace(oldPath)

	if _, err := os.Stat(oldPath); os.IsNotExist(err) {
		fmt.Println("File does not exist.")
		return
	}

	fmt.Print("Enter the new file name: ")
	newName, _ := sc.ReadString('n')
	newName = strings.TrimSpace(newName)

	dir := filepath.Dir(oldPath)
	newPath := filepath.Join(dir, newName)

	err := os.Rename(oldPath, newPath)

	if err != nil {
		fmt.Println("Error renaming the file:", err)
		return
	}

	fmt.Println("File renamed succesfully to: ", newPath)
}
