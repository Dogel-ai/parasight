package main

import (
	"fmt"
	"image"
	"os"
)

func extractMessageFromImage(imagePath string) (string, error) {
	// Open image file
	imgFile, err := os.Open(imagePath)
	if err != nil {
		return "", err
	}
	defer imgFile.Close()

	// Decode image
	img, _, err := image.Decode(imgFile)
	if err != nil {
		return "", err
	}

	bounds := img.Bounds()
	binaryMessage := ""

	// Iterate over pixels
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, _, _, _ := img.At(x, y).RGBA()
			r8 := uint8(r >> 8)

			// Extract LSB from red channel
			if r8&1 == 1 {
				binaryMessage += "1"
			} else {
				binaryMessage += "0"
			}

			// Stop if we detect the NULL character
			if len(binaryMessage)%8 == 0 {
				char := binaryStringToByte(binaryMessage[len(binaryMessage)-8:])
				if char == 0 {
					return binaryToString(binaryMessage[:len(binaryMessage)-8]), nil
				}
			}
		}
	}

	return "", fmt.Errorf("no message found")
}
