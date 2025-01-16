package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const BitsInByte = 8

func main() {
	var dataMask [][]uint8
	imagePath := "example-image.png"
	//inputData := []uint8{72, 101, 108, 108, 111}
	inputData := "Hello"

	// Opens actual image.
	image, err := os.Open(imagePath)
	if err != nil {
		log.Fatal("Failed to open image: ", err)
	}
	defer image.Close()

	// Gets bytes of image and dumps into single array with totaled numbers
	imageBytes, err := getBytes(image, len(inputData)*BitsInByte)
	if err != nil {
		log.Fatal("Reading image failed: ", err)
	}

	// Converts the bytes array into 2D slice containing separate bits
	imageBits, err := getBitSlice(imageBytes)
	if err != nil {
		log.Fatal("Converting to bit slice failed: ", err)
	}

	// Converts inputted data into 2D slice of bits
	if typeOf(inputData) == "string" {
		dataMask, _ = getBitSlice(stringToUint8(inputData))
	} else {
		//dataMask, _ = getDataMask(inputData)
	}

	fmt.Println(dataMask)
	fmt.Println(imageBits)
}

// TODO: Make more performant. Currently O(n)
func getBitSlice(inputData []uint8) ([][]uint8, error) {
	dataBytes := make([][]uint8, len(inputData))
	// Each index indicates each value in provided array/slice of data.
	for dataBytesIndex := range len(dataBytes) {
		// Each index indicates each bit of each value provided in array/slice of data.
		dataBytes[dataBytesIndex] = make([]uint8, BitsInByte)
		for dataBitsIndex := range BitsInByte {
			currentBit := (inputData[dataBytesIndex] & uint8(getIntPower(2, dataBitsIndex)))

			// [7-dataBitsIndex] inverts the order bits are pushed onto the slice. Removing 7- will cause a flip: [0 64 0 0 8 0 0 0] -> [0 0 0 8 0 0 64 0]
			dataBytes[dataBytesIndex][7-dataBitsIndex] = currentBit
		}
	}
	return dataBytes, nil
}

func getBytes(image *os.File, size int) ([]uint8, error) {
	imageBytes := make([]uint8, size)

	reader := bufio.NewReader(image)
	_, err := reader.Read(imageBytes)
	if err != nil {
		return imageBytes, fmt.Errorf("failed to read image bytes: %w", err)
	}

	return imageBytes, nil
}
