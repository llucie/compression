package main

import "testing"

func TestReadValidImage(t *testing.T) {
	imagePath := "resources/image.jpg"
	image, status := readImage(imagePath)

	if status != nil {
		t.Error("Failed to read valid image with path = ", imagePath)
	}

	if len(image) != 113050 {
		t.Error("Wrong size of image. Shall be 113050, actual value is ", len(image))
	}
}

func TestReadInvalidImage(t *testing.T) {
	imagePath := "resources/noImage.jpg"
	_, status := readImage(imagePath)

	if status == nil {
		t.Error("Invalid image reading returned a true status")
	}
}

func TestReadEmptyImage(t *testing.T) {
	_, status := readImage("")

	if status == nil {
		t.Error("Invalid image reading returned a true status")
	}
}
