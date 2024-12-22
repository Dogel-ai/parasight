package main

import "fmt"

func getIntPower(base, power int) int {
	calculatedInt := 1
	for power != 0 {
		calculatedInt *= base
		power -= 1
	}
	return calculatedInt
}

func typeOf(variable interface{}) string {
	return fmt.Sprintf("%T", variable)
}

func stringToUint8(str string) []uint8 {
	convertedSlice := make([]uint8, len(str))
	for characterIndex := range len(str) {
		convertedSlice[characterIndex] = str[characterIndex]
	}
	return convertedSlice
}
