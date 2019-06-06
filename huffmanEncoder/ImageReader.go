package main

import (
	"fmt"
	"image/jpeg"
	"os"
)

func readImage(filepath string) []uint32 {

	existingImageFile, _ := os.Open(filepath)
	defer existingImageFile.Close()

	// Alternatively, since we know it is a png already
	// we can call png.Decode() directly
	loadedImage, _ := jpeg.Decode(existingImageFile)

	width := loadedImage.Bounds().Size().X
	height := loadedImage.Bounds().Size().Y

	image := make([]uint32, width*height)
	fmt.Println("Loaded image size = (width, height) = (", width, ", ", height, ")")

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			r, _, _, _ := loadedImage.At(j, i).RGBA()
			image[i] = r
		}
	}

	return image
}
