package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	imagePath := "example-image.png"
	//inputData := []uint8{72, 101, 108, 108, 111}
	inputData := "According to all known laws of aviation, there is no way a bee should be able to fly. Its wings are too small to get its fat little body off the ground. The bee, of course, flies anyway because bees don't care what humans think is impossible. Yellow, black. Yellow, black. Yellow, black. Yellow, black. Ooh, black and yellow! Let's shake it up a little."

	_, _, err := getBytes(imagePath)
	if err != nil {
		log.Fatal("Reading image failed: ", err)
	}

	var dataMask [][]uint8
	if typeOf(inputData) == "string" {
		dataMask, _ = getDataMask(stringToUint8(inputData))
	} else {
		//dataMask, _ = getDataMask(inputData)
	}
	fmt.Println(dataMask)
}

// TODO: Make more performant. Reduce var assignments? Currently O(n)
func getDataMask(inputData []uint8) ([][]uint8, error) {
	dataBytes := make([][]uint8, len(inputData))
	// Each index indicates each value in provided array/slice of data.
	for dataBytesIndex := range len(dataBytes) {
		// Each index indicates each bit of each value provided in array/slice of data.
		dataBytes[dataBytesIndex] = make([]uint8, 8)
		for dataBitsIndex := range 8 {
			var currentBit uint8
			if dataBitsIndex == 0 {
				currentBit = (inputData[dataBytesIndex] & 1)
			} else {
				currentBit = (inputData[dataBytesIndex] & uint8(getIntPower(2, dataBitsIndex)))
			}
			// [7-dataBitsIndex] inverts the order bits are pushed onto the slice. Removing 7- will cause a flip: [0 64 0 0 8 0 0 0] -> [0 0 0 8 0 0 64 0]
			dataBytes[dataBytesIndex][7-dataBitsIndex] = currentBit
		}
	}
	return dataBytes, nil
}

func getBytes(imagePath string) ([]uint8, int, error) {
	var imageBytes []uint8
	var size int

	image, err := os.Open(imagePath)
	if err != nil {
		return imageBytes, size, fmt.Errorf("failed to open image: %w", err)
	}
	defer image.Close()

	fileInfo, err := image.Stat()
	if err != nil {
		return imageBytes, size, fmt.Errorf("failed to read image properties: %w", err)
	}
	size = int(fileInfo.Size())
	imageBytes = make([]uint8, size)

	reader := bufio.NewReader(image)
	_, err = reader.Read(imageBytes)
	if err != nil {
		return imageBytes, size, fmt.Errorf("failed to read image bytes: %w", err)
	}

	return imageBytes, size, nil
}
