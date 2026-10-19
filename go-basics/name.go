package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Create a reader to read from standard input
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("What is your name? ")

	// Read string input until the user presses Enter (\n)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occurred while reading input:", err)
		return
	}

	// Clean up the input by removing whitespace and newline characters
	name := strings.TrimSpace(input)

	// Fallback to "stranger" if the user hits Enter without typing a name
	if name == "" {
		name = "stranger"
	}

	// Print the customized greeting
	fmt.Printf("Hello, %s! Welcome to your first Go CLI.\n", name)
}
