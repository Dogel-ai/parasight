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

	inputData, err := getConvertedInput()
	if err != nil {
		log.Fatal("Failed getting input: ", err)
	}

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
	dataMask, _ = getBitSlice(inputData)

	fmt.Println("Data Mask:", dataMask)
	fmt.Println("Image Bits Slices:", imageBits)
}

// Unsure if getConvertedInput works correctly. I need to write to new file to figure out
func writeToFile([]uint8) {

}

func getConvertedInput() ([]uint8, error) {
	// TODO: Backup original file. Add confirmation. Add unsupported prompt
	//		 This function needs some solid cleanup
	//		 Write tests for this
	//		 Really needs comments as well
	//		 Add option for string/text. Use stringToUint8 in execution

	var userChoice string = "data"
	var userData []uint8
	var zeroesOffset int

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Insert data or pass file? (data/file):")
	fmt.Print("(default=data) ")
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed reading input: %w", err)
	}
	if scanner.Text() != "" {
		userChoice = scanner.Text()
	}

	fmt.Print("Input data/path: ")
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed reading input: %w", err)
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

			userData, err = getBytes(file, int(fileInfo.Size())*BitsInByte)
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
		userData = scanner.Bytes()
	}
	return userData, nil
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
