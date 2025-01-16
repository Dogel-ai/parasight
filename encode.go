package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func hideMessageInImage(inputImagePath, outputImagePath, message string) error {
	imageFile, err := os.Open(inputImagePath)
	if err != nil {
		return err
	}
	defer imageFile.Close()

	img, _, err := image.Decode(imageFile)
	if err != nil {
		return err
	}

	message += string(rune(0))
	binaryMessage := ""
	for _, char := range message {
		binaryMessage += byteToBinaryString(byte(char))
	}

	bounds := img.Bounds()
	newImg := image.NewRGBA(bounds)

	bitIndex := 0
	messageLen := len(binaryMessage)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			r8, g8, b8 := uint8(r>>8), uint8(g>>8), uint8(b>>8)

			if bitIndex < messageLen {
				if binaryMessage[bitIndex] == '1' {
					r8 |= 1
				} else {
					r8 &= 0xFE
				}
				bitIndex++
			}
			newImg.Set(x, y, color.RGBA{r8, g8, b8, uint8(a >> 8)})
		}
	}

	outFile, err := os.Create(outputImagePath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	err = png.Encode(outFile, newImg)
	if err != nil {
		return err
	}

	return nil
}
