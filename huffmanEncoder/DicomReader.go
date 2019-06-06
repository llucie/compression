package main

import (
	"fmt"

	"github.com/suyashkumar/dicom"
)

type DicomReader struct{}

func (reader DicomReader) readFile(filePath string) []uint16 {
	fmt.Println("Reading file: ", filePath)

	p, err := dicom.NewParserFromFile(filePath, nil)
	opts := dicom.ParseOptions{DropPixelData: false}

	fmt.Println("Error in reading dicom file ", err)

	dataset, err := p.Parse(opts) // parse whole dicom
	fmt.Println("Error in parsing ", err)

	element, err := dataset.FindElementByName("PixelData")
	//image := element.Value[0]
	fmt.Println("Error in find element ", err, " value = ", element.Value[0])

	// i := 0

	// for i < 141 {
	// 	element := p.ParseNext(opts)
	// 	if element.Tag == "pouet" {
	// 		fmt.Println("Tag ", i, " : ", element.Tag, " Value = ")
	// 	}
	// 	i++
	// }

	// Create image with size = 3*3
	image := make([]uint16, 3*3)

	// // Fill image
	// for index, _ := range image {
	// 	image[index] = uint16(index)
	// }

	// for i := 0; i < 20; i++ {
	// 	image[i] = 0
	// }

	return image
}
