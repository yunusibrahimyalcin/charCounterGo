package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	letters := []string{}
	exit := "exit"
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Input char: ")
		char, _ := reader.ReadString('\n')
		trimmedChar := strings.TrimSpace(char)

		if trimmedChar == exit {
			break
		}
		letters = append(letters, trimmedChar)
		fmt.Printf("You entered: %s\n", trimmedChar)
		fmt.Printf("Letters slice length is: %v \n", len(letters))
	}

	fmt.Println("Collected letters:", letters)

	charCount := make(map[string]int)

	for _, char := range letters {
		charCount[char]++
	}

	for char, count := range charCount {
		fmt.Printf("%d%s", count, char)
	}
}
