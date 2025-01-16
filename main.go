package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	input, err := getInput()
	if err != nil {
		log.Fatal(err)
	}

	hideMessageInImage("input.png", "out.png", input)

	message, err := extractMessageFromImage("out.png")
	if err != nil {
		log.Fatal(err)
	} else {
		log.Print("Extracted Message:", message)
	}
}

func getInput() (string, error) {
	var userChoice string = "data"
	var userData string
	var zeroesOffset int

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Insert data or pass file? (data/file):")
	fmt.Print("(default=data) ")
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return userData, fmt.Errorf("failed reading input: %w", err)
	}
	if scanner.Text() != "" {
		userChoice = scanner.Text()
	}

	fmt.Print("Input data/path: ")
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return userData, fmt.Errorf("failed reading input: %w", err)
	}

	if scanner.Text() != "" {
		if userChoice != "data" {
			file, err := os.Open(scanner.Text())
			if err != nil {
				return userData, fmt.Errorf("failed to open path: %w", err)
			}
			defer file.Close()

			fileInfo, err := file.Stat()
			if err != nil {
				return userData, fmt.Errorf("failed to read file properties: %w", err)
			}

			userData, err = getBytes(file, int(fileInfo.Size())*8)
			if err != nil {
				return userData, fmt.Errorf("failed converting to bytes slice: %w", err)
			}

			// TODO: Make this a toggle option
			// Removes trailing zeroes by crawling backwards through the userData slice until it finds a non-zero value
			for userDataIndex := range len(userData) {
				if userData[len(userData)-(userDataIndex+1)] != 0 {
					zeroesOffset = len(userData) - userDataIndex
					break
				}
			}
			userData = userData[:zeroesOffset]
			return userData, nil
		}
		userData = scanner.Text()
	}
	return userData, nil
}
