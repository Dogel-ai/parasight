package main

import (
	"bufio"
	"fmt"
	"os"
)

// Convert an 8-bit binary string to a byte
func binaryStringToByte(s string) byte {
	var b byte
	for i := 0; i < 8; i++ {
		b <<= 1
		if s[i] == '1' {
			b |= 1
		}
	}
	return b
}

// Convert binary string to readable text
func binaryToString(binary string) string {
	message := ""
	for i := 0; i < len(binary); i += 8 {
		char := binaryStringToByte(binary[i : i+8])
		message += string(char)
	}
	return message
}

func byteToBinaryString(b byte) string {
	s := ""
	for i := 7; i >= 0; i-- {
		if (b & (1 << i)) > 0 {
			s += "1"
		} else {
			s += "0"
		}
	}
	return s
}

func getBytes(image *os.File, size int) (string, error) {
	imageBytes := make([]uint8, size)

	reader := bufio.NewReader(image)
	_, err := reader.Read(imageBytes)
	if err != nil {
		return string(imageBytes), fmt.Errorf("failed to read image bytes: %w", err)
	}

	return string(imageBytes), nil
}
