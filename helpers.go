package main

func getIntPower(base, power int) int {
	calculatedInt := 1
	for power != 0 {
		calculatedInt *= base
		power -= 1
	}
	return calculatedInt
}

func stringToUint8(data string) []uint8 {
	convertedSlice := make([]uint8, len(data))
	for characterIndex := range len(data) {
		convertedSlice[characterIndex] = data[characterIndex]
	}
	return convertedSlice
}
